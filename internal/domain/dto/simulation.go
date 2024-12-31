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

type SimulationResponseBorrowerRequest struct {
	SimulationId string `json:"simulationId "`
	Response     bool   `json:"response "`
}

type JwtResponse struct {
	Token string `json:"token "`
}

func ToPayloadJWTModel(r JwtRequest) *model.PayloadJWT {
	return &model.PayloadJWT{
		CredorID: r.CredorId,
		Escopo:   r.CredorId,
	}

}
