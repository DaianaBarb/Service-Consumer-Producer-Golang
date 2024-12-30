package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"project-golang/internal/domain/entity"
	"project-golang/internal/logger"
	"sync"
)

var tenantConfigCache = sync.Map{}

type IRepository interface {
	CreatedSimulation(simu *entity.Simulation) (*entity.Simulation, error)
	CreatedBorrower(tom *entity.Borrower) error
	CreatedSetup(set *entity.Setup) error
	FindByIdSimulation(simulationId string) (*entity.Simulation, error)
	FindByIdSetup(setupId string) (*entity.Setup, error)
	FindByIdBorrower(borrwerId string) (*entity.Borrower, error)
	UpdateSetup(setupId string, newSetup *entity.Setup) error
	UpdateSimulationStatus(simulationId string, status string) error
}

type Repository struct {
	db     *sql.DB
	logger logger.ILogCloudWatch
}

func NewRepository(db *sql.DB, log logger.ILogCloudWatch) IRepository {

	return &Repository{
		db:     db,
		logger: log,
	}
}

func (r *Repository) CreatedSimulation(simu *entity.Simulation) (*entity.Simulation, error) {
	simula := &entity.Simulation{}

	query := `
		INSERT INTO simulations (borrower_id, loan_value, number_installments, interest_rate, status,created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
	`

	// _, err := r.db.Exec(query, simu.BorrowerId, simu.LoanValue, simu.NumberOfInstallments, simu.InterestRate, simu.Status)
	// if err != nil {
	// 	return err
	// }

	err := r.db.QueryRow(
		query,
		simu.BorrowerId,
		simu.LoanValue,
		simu.NumberOfInstallments,
		simu.InterestRate,
		simu.Status,
	).Scan(
		simula.SimulationId,
		simula.BorrowerId,
		simula.LoanValue,
		simula.NumberOfInstallments,
		simula.InterestRate,
		simula.Status,
		simula.CreatedAt,
		simula.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return simu, nil

}

func (r *Repository) CreatedBorrower(tom *entity.Borrower) error {
	query := `
	INSERT INTO borrower ( name, phone, email, cpf,created_at, updated_at)
	VALUES ($1, $2, $3, $4, NOW(), NOW())
`

	_, err := r.db.Exec(query, tom.Name, tom.Phone, tom.Email, tom.Cpf)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) CreatedSetup(set *entity.Setup) error {

	query := `
	INSERT INTO setup ( setup_id, capital, fees, interest_rate,created_at, updated_at)
	VALUES ($1, $2, $3,$4, NOW(), NOW())
`

	_, err := r.db.Exec(query, os.Getenv("SETUP_ID"), set.Capital, set.Fees, set.InterestRate)
	if err != nil {
		return err
	}

	return nil

}

func (r *Repository) UpdateSimulationStatus(simulationId string, status string) error {
	query := `
		UPDATE simulations 
		SET status = $1, updated_at = NOW() 
		WHERE id = $2
	`

	_, err := r.db.Exec(query, status, simulationId)
	return err
}

func (r *Repository) UpdateSetup(setupId string, newSetup *entity.Setup) error {

	query := `
	UPDATE setup 
	SET capital = $1, fees = $2, interest_rate = $3, updated_at = NOW() 
	WHERE id = $4
`

	_, err := r.db.Exec(query, newSetup.Capital, newSetup.Fees, newSetup.InterestRate, setupId)

	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindByIdSimulation(simulationId string) (*entity.Simulation, error) {
	simulation := &entity.Simulation{}

	query := `
	SELECT id, simulation_id, borrower_id, loan_value, number_installments, interest_rate, status, created_at, updated_at
	FROM simulations
	WHERE id = $1
`
	row := r.db.QueryRow(query, simulationId)

	err := row.Scan(
		simulation.SimulationId,
		simulation.BorrowerId,
		simulation.LoanValue,
		simulation.NumberOfInstallments,
		simulation.InterestRate,
		simulation.Status,
		simulation.CreatedAt,
		simulation.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("simulação com ID %s não encontrada", simulationId)
		}
		return nil, err
	}
	return simulation, nil

}
func (r *Repository) FindByIdSetup(setupId string) (*entity.Setup, error) {

	if config, ok := tenantConfigCache.Load(setupId); ok {
		return config.(*entity.Setup), nil
	}

	setup := &entity.Setup{}

	query := `
	SELECT setup_id, capital, fees, interest_rate, created_at, updated_at
	FROM setup
	WHERE id = $1
`
	row := r.db.QueryRow(query, setupId)

	err := row.Scan(
		setup.SetupId,
		setup.Capital,
		setup.Fees,
		setup.InterestRate,
		setup.CreatedAt,
		setup.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("setup com ID %s não encontrada", setupId)
		}

		tenantConfigCache.Store(setupId, setup)
		return nil, err
	}

	return setup, nil
}
func (r *Repository) FindByIdBorrower(borrwerId string) (*entity.Borrower, error) {

	borrwer := &entity.Borrower{}

	query := `
	SELECT borrwer_id, name, phone, email, cpf, created_at, updated_at
	FROM borrwer
	WHERE id = $1
`
	row := r.db.QueryRow(query, borrwerId)

	err := row.Scan(
		borrwer.BorrowerId,
		borrwer.Name,
		borrwer.Phone,
		borrwer.Email,
		borrwer.Cpf,
		borrwer.CreatedAt,
		borrwer.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("borrwer  com ID %s não encontrada", borrwerId)
		}
		return nil, err
	}

	return borrwer, nil

}
