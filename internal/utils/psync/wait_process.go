package psync

import (
	"log"
	"os"
	"os/signal"
	"sort"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"project-golang/internal/utils"
)

const TIME_WAIT_MILIS = (5 * time.Millisecond)

const SCALING_FACTOR_WORKLOAD int32 = 25

type WaitProcessScaling struct {
	countProcess       int32
	maxProcess         int32
	autoScalingProcess int32
	factorWorkload     int32
	useAutoScaling     bool
}

type WaitProcessWorkload struct {
	mutexWorkload     sync.Mutex
	pendingWorkload   int32
	walkingWorkload   int32
	pendingWorkloadAt *time.Time
	finishedWorkload  int32
	averageWorkload   int32
	averageWorkloadAt *time.Time
}

// A WaitProcess waits for a goroutine within a group to finish first before releasing another goroutine.
type WaitProcess struct {
	mutexChilds sync.Mutex
	state       int32
	scaling     WaitProcessScaling
	workload    WaitProcessWorkload
	name        string
	startAt     time.Time
	childs      []*WaitProcess
}

type StatusWaitProcessRunning struct {
	UseAutoScaling       bool  `json:"useAutoScaling"`
	FactorWorkload       int32 `json:"factorWorkload"`
	AutoScalingProcesses int32 `json:"autoScalingProcesses"`
	MaximumProcesses     int32 `json:"maximumProcesses"`
	RunningProcesses     int32 `json:"runningProcesses"`
}

type StatusWaitProcessWorkload struct {
	PendingWorkload   int32      `json:"pendingWorkload"`
	WalkingWorkload   int32      `json:"walkingWorkload"`
	PendingWorkloadAt *time.Time `json:"pendingWorkloadAt"`
	AverageWorkload   int32      `json:"averageWorkload"`
	AverageWorkloadAt *time.Time `json:"averageWorkloadAt"`
}

type StatusWaitProcess struct {
	GroupName string                    `json:"groupName"`
	State     string                    `json:"state"`
	StartAt   time.Time                 `json:"startAt"`
	Processes StatusWaitProcessRunning  `json:"processes"`
	Workload  StatusWaitProcessWorkload `json:"workload"`
	Childs    []*StatusWaitProcess      `json:"childs,omitempty"`
}

// NewWaitProcess create and initialising a WaitProcess.
func NewWaitProcess(maximumProcess int32, nameGroup string) *WaitProcess {
	timeStart := time.Now()
	newWp := &WaitProcess{
		state: 0,
		scaling: WaitProcessScaling{
			countProcess:       0,
			maxProcess:         maximumProcess,
			autoScalingProcess: maximumProcess,
			useAutoScaling:     true,
			factorWorkload:     SCALING_FACTOR_WORKLOAD,
		},
		workload: WaitProcessWorkload{
			pendingWorkload:   0,
			walkingWorkload:   0,
			pendingWorkloadAt: nil,
			finishedWorkload:  0,
			averageWorkload:   0,
			averageWorkloadAt: nil,
		},
		startAt: timeStart,
		name:    nameGroup,
	}
	return newWp
}

// NewChildWaitProcess create and initialising a WaitProcess as child of a main WaitProcess.
func (wp *WaitProcess) NewChildWaitProcess(maximumProcess int32, nameGroup string) *WaitProcess {
	newWp := NewWaitProcess(maximumProcess, nameGroup)

	wp.mutexChilds.Lock()
	defer wp.mutexChilds.Unlock()

	wp.childs = append(wp.childs, newWp)

	if len(wp.childs) > 1 {
		sort.SliceStable(wp.childs, func(i, j int) bool {
			return wp.childs[i].name < wp.childs[j].name
		})
	}

	return newWp
}

func (wp *WaitProcess) IsActive() bool {
	return atomic.LoadInt32(&wp.state) == 0
}

func (wp *WaitProcess) IsTerminating() bool {
	return atomic.LoadInt32(&wp.state) == 1
}

func (wp *WaitProcess) IsPaused() bool {
	return atomic.LoadInt32(&wp.state) == 2
}

func (wp *WaitProcess) GetStateName() string {
	actualState := atomic.LoadInt32(&wp.state)
	if actualState == 1 {
		return "Terminating"
	} else if actualState == 2 {
		return "Paused"
	}

	return "Active"
}

// HandleSigtermSignal await termChan signal and notify state terminating
func (wp *WaitProcess) HandleSigtermSignal() {
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	<-termChan // Blocks here until interrupted

	log.Printf("Received SIGTERM and notify main state terminating\n")
	wp.SetTerminating()
}

func (wp *WaitProcess) GetStatusWaitProcess() *StatusWaitProcess {

	wp.workload.mutexWorkload.Lock()
	workload := StatusWaitProcessWorkload{
		PendingWorkload:   wp.workload.pendingWorkload,
		WalkingWorkload:   wp.workload.walkingWorkload,
		PendingWorkloadAt: wp.workload.pendingWorkloadAt,
		AverageWorkload:   wp.workload.averageWorkload,
		AverageWorkloadAt: wp.workload.averageWorkloadAt,
	}
	wp.workload.mutexWorkload.Unlock()

	status := &StatusWaitProcess{
		GroupName: wp.name,
		State:     wp.GetStateName(),
		StartAt:   wp.startAt,
		Processes: StatusWaitProcessRunning{
			UseAutoScaling:       wp.scaling.useAutoScaling,
			FactorWorkload:       atomic.LoadInt32(&wp.scaling.factorWorkload),
			AutoScalingProcesses: atomic.LoadInt32(&wp.scaling.autoScalingProcess),
			MaximumProcesses:     atomic.LoadInt32(&wp.scaling.maxProcess),
			RunningProcesses:     atomic.LoadInt32(&wp.scaling.countProcess),
		},
		Workload: workload,
	}

	for _, wpChild := range wp.childs {
		statusChild := wpChild.GetStatusWaitProcess()
		status.Childs = append(status.Childs, statusChild)
	}

	return status
}

func (wp *WaitProcess) SetPendingWorkload(pendingWorkload int32, walkingWorkload int32, timeAt *time.Time) {

	if wp.scaling.useAutoScaling && wp.scaling.autoScalingProcess > 0 {
		newMaxProcess := pendingWorkload / wp.scaling.factorWorkload
		if newMaxProcess < 1 {
			newMaxProcess = 1
		} else {
			if newMaxProcess > wp.scaling.autoScalingProcess {
				newMaxProcess = wp.scaling.autoScalingProcess
			}
		}

		atomic.StoreInt32(&wp.scaling.maxProcess, newMaxProcess)
	}

	wp.workload.mutexWorkload.Lock()
	defer wp.workload.mutexWorkload.Unlock()
	wp.workload.pendingWorkload = pendingWorkload
	wp.workload.walkingWorkload = walkingWorkload
	wp.workload.pendingWorkloadAt = timeAt
}

func (wp *WaitProcess) AddFinishedWorkload(countFinished int32, timeAt *time.Time) {
	wp.workload.mutexWorkload.Lock()
	defer wp.workload.mutexWorkload.Unlock()

	wp.workload.finishedWorkload = wp.workload.finishedWorkload + countFinished

	if wp.workload.averageWorkloadAt == nil {
		wp.workload.averageWorkloadAt = utils.GetTimeNow()

	} else {

		elapsed := timeAt.Sub(*wp.workload.averageWorkloadAt)
		if elapsed >= time.Minute {
			var calculateNewAverage float32 = float32(wp.workload.finishedWorkload) / float32(elapsed.Seconds())

			wp.workload.averageWorkload = int32(calculateNewAverage * 60)
			if wp.workload.averageWorkload < 1 {
				wp.workload.averageWorkload = 1
			}
			wp.workload.averageWorkloadAt = timeAt
			wp.workload.finishedWorkload = 0
		}
	}
}

func (wp *WaitProcess) SetScalingFactorWorkload(factor int32) {
	atomic.StoreInt32(&wp.scaling.factorWorkload, factor)
}

func (wp *WaitProcess) SetMaximumProcess(maximumProcess int32) {
	atomic.StoreInt32(&wp.scaling.maxProcess, maximumProcess)
	atomic.StoreInt32(&wp.scaling.autoScalingProcess, maximumProcess)
}

func (wp *WaitProcess) SetMaximumProcessChild(nameChild string, maximumProcess int32) bool {
	for _, wpChild := range wp.childs {
		if wpChild.name == nameChild {
			wpChild.SetMaximumProcess(maximumProcess)
			return true
		}
	}
	return false
}

func (wp *WaitProcess) SetMaximumProcessChildForAll(maximumProcess int32) bool {
	for _, wpChild := range wp.childs {
		wpChild.SetMaximumProcess(maximumProcess)
	}
	return true
}

func (wp *WaitProcess) SetActive() {
	atomic.StoreInt32(&wp.state, 0)
}

func (wp *WaitProcess) SetTerminating() {
	atomic.StoreInt32(&wp.state, 1)
	for _, child := range wp.childs {
		child.SetTerminating()
	}
}

func (wp *WaitProcess) SetPaused() {
	atomic.StoreInt32(&wp.state, 2)
}

// Add increases the number of processes running in the group.
func (wp *WaitProcess) Add() {
	atomic.AddInt32(&wp.scaling.countProcess, 1)
}

// Done decreases the number of processes running in the group.
func (wp *WaitProcess) Done(nameProcessChild string) {
	atomic.AddInt32(&wp.scaling.countProcess, -1)
	if len(nameProcessChild) > 0 {
		log.Printf("Done process child [%s] - group [%s]\n", nameProcessChild, wp.name)
	}
}

// Wait waits for a goroutine to finish before releasing another process for execution in the group, or return
// when the state is changed for terminating or paused.
func (wp *WaitProcess) Wait() {
	for atomic.LoadInt32(&wp.scaling.countProcess) >= wp.scaling.maxProcess {
		if wp.IsTerminating() || wp.IsPaused() {
			return
		}
		//runtime.Gosched()
		time.Sleep(TIME_WAIT_MILIS)
	}
}

// WaitAll waits for all goroutines to finish for execution in the group.
func (wp *WaitProcess) WaitAll() {
	for atomic.LoadInt32(&wp.scaling.countProcess) > 0 {
		//runtime.Gosched()
		time.Sleep(TIME_WAIT_MILIS)
	}
}
