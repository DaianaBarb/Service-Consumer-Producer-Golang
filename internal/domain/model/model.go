package model

import "time"

type Model struct {
	Message string `json:"message"`
}

type Simulation struct {
	SimulationId         string
	BorrowerId           string
	LoanValue            string // valor emprestimo
	NumberOfInstallments float64    // quantidade de parcelas
	CreatedAt            *time.Time
	UpdateAt             *time.Time
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
	SetupId      string
	Capital      float64
	fees         float64 //juros
	InterestRate float64 // taxa de juros
}

type Contract struct {
	ContractId   string
	SimulationId string
	CreatedAt    *time.Time
	Status       string
	terms        string
}

type PayloadJWT struct {
	CredorID  string
	Escopo    string
	Expiracao int64
}
