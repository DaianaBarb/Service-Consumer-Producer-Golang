package model

import (
	"project-golang/internal/domain/entity"
	"time"
)

type Model struct {
	Message string `json:"message"`
}

type Simulation struct {
	SimulationId         string
	BorrowerId           string
	LoanValue            float64 // valor emprestimo
	NumberOfInstallments float64 // quantidade de parcelas
	CreatedAt            *time.Time
	UpdatedAt            *time.Time
	Status               string
	InterestRate         float64
}

type Borrower struct {
	BorrowerId string
	Name       string
	Phone      string
	Email      string
	Cpf        string
	CreatedAt  *time.Time
	UpdateAt   *time.Time
}

type Setup struct {
	SetupId       string
	Capital       float64
	Fees          float64 //juros
	InterestRate  float64 // taxa de juros
	Escope        string
	EscopeIsValid bool
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
type LogEntry struct {
	Level     string `json:"level"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// type Contract struct {
// 	ContractId   string
// 	SimulationId string
// 	CreatedAt    *time.Time
// 	Status       string
// 	terms        string
// }

type SimulationResponseBorrower struct {
	SimulationId string
	Status       string
}

type PayloadJWT struct {
	CredorID  string
	Escopo    string
	Expiracao int64
}

func ToSetupEntity(setup *Setup) *entity.Setup {
	return &entity.Setup{
		SetupId:      setup.SetupId,
		Capital:      setup.Capital,
		Fees:         setup.Fees,
		InterestRate: setup.InterestRate,
		CreatedAt:    setup.CreatedAt,
		UpdatedAt:    setup.UpdatedAt,
	}

}

func ToSimulationEntity(simu *Simulation) *entity.Simulation {

	return &entity.Simulation{
		SimulationId:         simu.SimulationId,
		BorrowerId:           simu.BorrowerId,
		LoanValue:            simu.LoanValue,
		NumberOfInstallments: simu.NumberOfInstallments,
		CreatedAt:            simu.CreatedAt,
		UpdatedAt:            simu.UpdatedAt,
		Status:               simu.Status,
		InterestRate:         simu.InterestRate,
	}

}

func ToSimulationModel(simu *entity.Simulation) *Simulation {

	return &Simulation{
		SimulationId:         simu.SimulationId,
		BorrowerId:           simu.BorrowerId,
		LoanValue:            simu.LoanValue,
		NumberOfInstallments: simu.NumberOfInstallments,
		CreatedAt:            simu.CreatedAt,
		UpdatedAt:            simu.UpdatedAt,
		Status:               simu.Status,
		InterestRate:         simu.InterestRate,
	}

}

func ToSetupModel(setup *entity.Setup) *Setup {
	return &Setup{
		SetupId:      setup.SetupId,
		Capital:      setup.Capital,
		Fees:         setup.Fees,
		InterestRate: setup.InterestRate,
		CreatedAt:    setup.CreatedAt,
		UpdatedAt:    setup.UpdatedAt,
	}

}

func ToBorrowerdModel(bo *entity.Borrower) *Borrower {
	return &Borrower{
		BorrowerId: bo.BorrowerId,
		Name:       bo.BorrowerId,
		Phone:      bo.Phone,
		Email:      bo.Email,
		Cpf:        bo.Cpf,
		CreatedAt:  bo.CreatedAt,
		UpdateAt:   bo.UpdatedAt}

}
func ToBorrowerEntity(bo *Borrower) *entity.Borrower {

	return &entity.Borrower{
		BorrowerId: bo.BorrowerId,
		Name:       bo.BorrowerId,
		Phone:      bo.Phone,
		Email:      bo.Email,
		Cpf:        bo.Cpf,
		CreatedAt:  bo.CreatedAt,
		UpdatedAt:  bo.UpdateAt,
	}

}

type SimulationParam struct {
	SimulationId         *string
	BorrowerId           *string
	LoanValue            *float64 // valor emprestimo
	NumberOfInstallments *float64 // quantidade de parcelas
	CreatedAt            *time.Time
	UpdatedAt            *time.Time
	Status               *string
	InterestRate         *float64
}

type Params struct {
	Simu     *SimulationParam
	Page     int
	PageSize int
}
