package entity

import (
	"time"
)

type Simulation struct {
	SimulationId         string    `json:"simulation_id"`
	BorrowerId           string    `json:"borrower_id"`
	LoanValue            float64   `json:"loan_value"`             // valor emprestimo
	NumberOfInstallments float64   `json:"number_of_installments"` // quantidade de parcelas
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	Status               string    `json:"status"`
	InterestRate         float64   `json:"interest_rate"`
}

type Borrower struct {
	BorrowerId string
	Name       string
	Phone      string
	Email      string
	Cpf        string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
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

// type Contract struct {
// 	ContractId   string
// 	SimulationId string
// 	CreatedAt    *time.Time
// 	Status       string
// 	terms        string
// }
