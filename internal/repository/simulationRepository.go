package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"project-golang/internal/domain/entity"
	"project-golang/internal/domain/model"
	"project-golang/internal/logger"
	"sync"

	"github.com/google/uuid"
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
	UpdateSimulation(simu *entity.Simulation) error
	GetSimulations(param *model.Params) ([]entity.Simulation, error)
	Ping() error
}

type Repository struct {
	db     *sql.DB
	logger logger.ILogCloudWatch
}

// Ping implements IRepository.
func (r *Repository) Ping() error {
	err := r.db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func NewRepository(db *sql.DB, log logger.ILogCloudWatch) IRepository {

	return &Repository{
		db:     db,
		logger: log,
	}
}

func (r *Repository) CreatedSimulation(simu *entity.Simulation) (*entity.Simulation, error) {
	var simula entity.Simulation

	var SimulationId uuid.UUID
	query := `
		INSERT INTO simulation (borrower_id, loan_value, number_of_installments, status, interest_rate, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING simulation_id, borrower_id, loan_value, number_of_installments, status, interest_rate, created_at, updated_at;
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
		simu.Status,
		simu.InterestRate,
	).Scan(
		&SimulationId,
		&simula.BorrowerId,
		&simula.LoanValue,
		&simula.NumberOfInstallments,
		&simula.Status,
		&simula.InterestRate,
		&simula.CreatedAt,
		&simula.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	simu.SimulationId = SimulationId.String()

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
	set.Escope = "escope"
	set.EscopeIsValid = true
	query := `
	INSERT INTO setup ( setup_id, capital, fees, interest_rate, escope, escope_is_valid, created_at, updated_at)
	VALUES ($1, $2, $3,$4, $5, $6, NOW(), NOW())
`

	_, err := r.db.Exec(query, os.Getenv("SETUP_ID"), set.Capital, set.Fees, set.InterestRate, set.Escope, set.EscopeIsValid)
	if err != nil {
		return err
	}

	return nil

}

func (r *Repository) UpdateSimulation(simu *entity.Simulation) error {

	query := `
		UPDATE simulations 
		SET status = $1, borrower_id = $2, loan_value = $3, number_installments = $4, interest_rate = $5, updated_at = NOW() 
		WHERE id = $6
	`

	_, err := r.db.Exec(query, simu.Status, simu.BorrowerId, simu.LoanValue, simu.NumberOfInstallments, simu.InterestRate, simu.SimulationId)
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
	simulation := entity.Simulation{}

	query := `
	SELECT  simulation_id, borrower_id, loan_value, number_of_installments, interest_rate, status, created_at, updated_at
	FROM simulation
	WHERE simulation_id = $1
`
	row := r.db.QueryRow(query, simulationId)

	err := row.Scan(
		&simulation.SimulationId,
		&simulation.BorrowerId,
		&simulation.LoanValue,
		&simulation.NumberOfInstallments,
		&simulation.InterestRate,
		&simulation.Status,
		&simulation.CreatedAt,
		&simulation.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("simulação com ID %s não encontrada", simulationId)
		}
		return nil, err
	}
	return &simulation, nil

}
func (r *Repository) FindByIdSetup(setupId string) (*entity.Setup, error) {

	if config, ok := tenantConfigCache.Load(setupId); ok {
		return config.(*entity.Setup), nil
	}

	setup := entity.Setup{}

	query := `
	SELECT setup_id, capital, fees, interest_rate, escope, escope_is_valid, created_at, updated_at
	FROM setup
	WHERE setup_id = $1
`
	row := r.db.QueryRow(query, setupId)

	err := row.Scan(
		&setup.SetupId,
		&setup.Capital,
		&setup.Fees,
		&setup.InterestRate,
		&setup.Escope,
		&setup.EscopeIsValid,
		&setup.CreatedAt,
		&setup.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("setup com ID %s não encontrada", setupId)
		}

		tenantConfigCache.Store(setupId, setup)
		return nil, err
	}

	return &setup, nil
}
func (r *Repository) FindByIdBorrower(borrwerId string) (*entity.Borrower, error) {

	borrwer := entity.Borrower{}

	query := `
	SELECT borrower_id, name, phone, email, cpf, created_at, updated_at
	FROM borrower
	WHERE borrower_id = $1
`
	row := r.db.QueryRow(query, borrwerId)

	err := row.Scan(
		&borrwer.BorrowerId,
		&borrwer.Name,
		&borrwer.Phone,
		&borrwer.Email,
		&borrwer.Cpf,
		&borrwer.CreatedAt,
		&borrwer.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("borrwer  com ID %s não encontrada", borrwerId)
		}
		return nil, err
	}

	return &borrwer, nil

}

func (r *Repository) GetSimulations(param *model.Params) ([]entity.Simulation, error) {

	offset := (param.Page - 1) * param.PageSize
	// Base query
	query := `
		SELECT id, borrower_id, interest_rate, status, loan_value,number_installments created_at,updated_at, simulation_id
		FROM simulation
		WHERE 1=1
	`

	// Dynamic filters
	args := []interface{}{}
	if param.Simu.BorrowerId != nil {
		query += " AND borrower_id = $1"
		args = append(args, *param.Simu.BorrowerId)
	}
	if param.Simu.InterestRate != nil {
		query += fmt.Sprintf(" AND interest_rate = $%d", len(args)+1)
		args = append(args, *param.Simu.InterestRate)
	}
	if param.Simu.Status != nil {
		query += fmt.Sprintf(" AND status = $%d", len(args)+1)
		args = append(args, *param.Simu.Status)
	}
	if param.Simu.LoanValue != nil {
		query += fmt.Sprintf(" AND loan_value = $%d", len(args)+1)
		args = append(args, *param.Simu.LoanValue)
	}
	if param.Simu.NumberOfInstallments != nil {
		query += fmt.Sprintf(" AND number_installments = $%d", len(args)+1)
		args = append(args, *param.Simu.NumberOfInstallments)
	}
	if param.Simu.CreatedAt != nil {
		query += fmt.Sprintf(" AND created_at = $%d", len(args)+1)
		args = append(args, *param.Simu.CreatedAt)
	}
	if param.Simu.UpdatedAt != nil {
		query += fmt.Sprintf(" AND updated_at = $%d", len(args)+1)
		args = append(args, *param.Simu.UpdatedAt)
	}
	if param.Simu.SimulationId != nil {
		query += fmt.Sprintf(" AND simulation_id = $%d", len(args)+1)
		args = append(args, *param.Simu.SimulationId)
	}

	// Adiciona paginação
	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, param.PageSize, offset)

	// Executa a consulta
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse dos resultados
	var simulations []entity.Simulation
	for rows.Next() {
		var sim entity.Simulation
		if err := rows.Scan(&sim.SimulationId, &sim.BorrowerId, &sim.LoanValue, &sim.NumberOfInstallments, &sim.InterestRate, &sim.Status, &sim.CreatedAt, &sim.UpdatedAt); err != nil {
			return nil, err
		}
		simulations = append(simulations, sim)
	}

	return simulations, nil
}
