package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	sqsAws "project-golang/internal/adapters/cloud/aws/sqs"

	"project-golang/internal/domain/model"
	service "project-golang/internal/services"

	//"github.com/americanas-go/config"
	"github.com/americanas-go/log"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type Consumer interface {
	Do()
}

const (
	root         = "worker.sqs"
	urlPath      = ".url"
	timeoutPath  = ".visibilityTimeout"
	sleepTimeout = ".sleepTimeout"
)

type worker struct {
	queueUrl          string
	visibilityTimeout int32
	sleepTimeout      time.Duration
	service         service.Processor
	client            sqsAws.Client
}

func (w *worker) Do() {
	logger := log.FromContext(context.TODO())
	go HandlerSigtermSignal()

	for {
		ctx := context.Background()

		result, err := w.client.ReceiveMessage(ctx, sqsAws.GenerateMessageInput(w.queueUrl, w.visibilityTimeout))
		if err != nil {
			logger.Debug(err.Error())
			return
		}
		if len(result.Messages) == 0 {
			time.Sleep(w.sleepTimeout)
			continue
		}
		var wg sync.WaitGroup
		for _, message := range result.Messages {
			wg.Add(1)
			go func(message types.Message) {
				defer wg.Done()
				w.processMessage(message, ctx)

			}(message)
		}

		wg.Wait()

	}

}

func HandlerSigtermSignal() {
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGTERM)
	<-termChan // Blocks here until interrupted

	log.Printf("Received SIGTERM and notify main state terminating\n")
	os.Exit(0)

}

func (w *worker) processMessage(message types.Message, ctx context.Context) {

	logger := log.FromContext(ctx)
	var model = &model.Model{}
	err := json.Unmarshal([]byte(*message.Body), model)
	if err != nil {
		logger.Debug(err.Error())
		return
	}
	fmt.Println(model.Message)
	// err = w.processor.Process(model)
	// if err != nil {
	// 	logger.Debug(err.Error())
	// 	return
	// }
	_, err = w.client.DeleteMessage(ctx, sqsAws.GenerateDeleteInput(w.queueUrl, message.ReceiptHandle))
	if err != nil {
		logger.Debug(err.Error())
		return
	}
	return

}

func NewSqs(sqsClient sqsAws.Client, service service.Processor) Consumer {
	return &worker{
		queueUrl:          os.Getenv("QUEUE_URL"),
	//	visibilityTimeout: int32(config.Int(os.Getenv("TIMEOUTPATH"))),
	//	sleepTimeout:      config.Duration(os.Getenv("SLEEPTIMEOUT")) * time.Millisecond,
		client:            sqsClient,
		service:         service,
	}

}
