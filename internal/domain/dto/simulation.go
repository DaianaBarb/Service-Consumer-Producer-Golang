package dto

import (
	"project-golang/internal/domain/model"
	"time"
)

type AntiFraudRequest struct {
	BorrowerId string  `json:"borrowerId"`
	LoanValue  float64 `json:"valor"`
}

type AntiFraudResponse struct {
	BorrowerId string `json:"borrowerId"`
}

type QueuePublishPayload struct {
	SimulationId string `json:"simulationId "`
}

type JwtRequest struct {
	CredorId string `json:"credorId "`
	Escopo   string `json:"escopoId "`
}

type SimulationRequest struct {
	BorrowerId           string  `json:"borrowerId"`
	LoanValue            float64 `json:"loanValue"`
	NumberOfInstallments float64 `json:"numberOfInstallments"`
	InterestRate         float64 `json:"interestRate"`
}

type HelfCheckResponse struct {
	Status        string     `json:"status"`
	Response_time *time.Time `json:"response_time"`
}

type SimulationResponse struct {
	SimulationId         string     `json:"simulationId "`
	BorrowerId           string     `json:"borrowerId "`
	LoanValue            float64    `json:"loanValue"`
	NumberOfInstallments float64    `json:"numberOfInstallments"`
	CreatedAt            *time.Time `json:"createdAt "`
	UpdatedAt            *time.Time `json:"updatedAt "`
	Status               string     `json:"satus"`
	InterestRate         float64    `json:"interestRate "`
}

type SimulationPaginationResponse struct {
	Simulations []SimulationResponse `json:"simulations "`
	Page        string               `json:"page"`
	PageSize    string               `json:"pageSize"`
}

type SimulationResponseBorrowerRequest struct {
	SimulationId string `json:"simulationId "`
	Response     bool   `json:"response "`
}

type JwtResponse struct {
	Token string `json:"token"`
}

func ToPayloadJWTModel(r JwtRequest) *model.PayloadJWT {
	return &model.PayloadJWT{
		CredorID: r.CredorId,
		Escopo:   r.CredorId,
	}

}

type BorrowerRequest struct {
	Name  string `json:"nome"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
}

type BorrowerResponse struct {
	BorrowerId string     `json:"borrewerId"`
	Name       string     `json:"name"`
	Phone      string     `json:"phone"`
	Email      string     `json:"email"`
	Cpf        string     `json:"cpf"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdateAt   *time.Time `json:"updatedAt"`
}

type BorrowerResponseTosimulationRequest struct {
	Status       string `json:"status"`
}

type SetupRequest struct {
	Capital      float64 `json:"capital"`
	Fees         float64 `json:"fees"`         //juros
	InterestRate float64 `json:"interestRate"` // taxa de juros
	Escope       string  `json:"escope"`
}
type SetupResponse struct {
	SetupId       string     `json:"setupId"`
	Capital       float64    `json:"capital"`
	Fees          float64    `json:"fees"`         //juros
	InterestRate  float64    `json:"interestRate"` // taxa de juros
	Escope        string     `json:"escope"`
	EscopeIdValid bool       `json:"escopeIdValid"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
}

func ToBorrowerModel(r *BorrowerRequest) *model.Borrower {
	return &model.Borrower{
		Name:  r.Name,
		Phone: r.Phone,
		Email: r.Email,
		Cpf:   r.Cpf,
	}

}

func ToSetupModel(s *SetupRequest) *model.Setup {
	return &model.Setup{
		Capital:      s.Capital,
		Fees:         s.Fees,
		InterestRate: s.InterestRate,
		Escope:       s.Escope,
	}

}

func ToSimulationModel(s *SimulationRequest) *model.Simulation {
	return &model.Simulation{
		BorrowerId:           s.BorrowerId,
		LoanValue:            s.LoanValue,
		NumberOfInstallments: s.NumberOfInstallments,
		InterestRate:         s.InterestRate,
	}

}
