package entity

import (
	"time"
)

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
