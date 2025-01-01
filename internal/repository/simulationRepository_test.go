package repository

import (
	"database/sql"
	"project-golang/internal/domain/entity"
	"project-golang/internal/domain/model"
	"project-golang/internal/logger"
	"reflect"
	"testing"
)

func TestRepository_CreatedSimulation(t *testing.T) {
	type fields struct {
		db     *sql.DB
		logger logger.ILogCloudWatch
	}
	type args struct {
		simu *entity.Simulation
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Simulation
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			got, err := r.CreatedSimulation(tt.args.simu)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.CreatedSimulation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.CreatedSimulation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_CreatedBorrower(t *testing.T) {
	type fields struct {
		db     *sql.DB
		logger logger.ILogCloudWatch
	}
	type args struct {
		tom *entity.Borrower
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			if err := r.CreatedBorrower(tt.args.tom); (err != nil) != tt.wantErr {
				t.Errorf("Repository.CreatedBorrower() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_CreatedSetup(t *testing.T) {
	type fields struct {
		db     *sql.DB
		logger logger.ILogCloudWatch
	}
	type args struct {
		set *entity.Setup
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			if err := r.CreatedSetup(tt.args.set); (err != nil) != tt.wantErr {
				t.Errorf("Repository.CreatedSetup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_UpdateSimulationStatus(t *testing.T) {
	type fields struct {
		db     *sql.DB
		logger logger.ILogCloudWatch
	}
	type args struct {
		simulationId string
		status       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			if err := r.UpdateSimulationStatus(tt.args.simulationId, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("Repository.UpdateSimulationStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_UpdateSetup(t *testing.T) {
	type fields struct {
		db     *sql.DB
		logger logger.ILogCloudWatch
	}
	type args struct {
		setupId  string
		newSetup *entity.Setup
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			if err := r.UpdateSetup(tt.args.setupId, tt.args.newSetup); (err != nil) != tt.wantErr {
				t.Errorf("Repository.UpdateSetup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_FindByIdSimulation(t *testing.T) {
	type fields struct {
		db     *sql.DB
		logger logger.ILogCloudWatch
	}
	type args struct {
		simulationId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Simulation
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			got, err := r.FindByIdSimulation(tt.args.simulationId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.FindByIdSimulation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.FindByIdSimulation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_FindByIdSetup(t *testing.T) {
	type fields struct {
		db     *sql.DB
		logger logger.ILogCloudWatch
	}
	type args struct {
		setupId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Setup
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			got, err := r.FindByIdSetup(tt.args.setupId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.FindByIdSetup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.FindByIdSetup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_FindByIdBorrower(t *testing.T) {
	type fields struct {
		db     *sql.DB
		logger logger.ILogCloudWatch
	}
	type args struct {
		borrwerId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Borrower
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			got, err := r.FindByIdBorrower(tt.args.borrwerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.FindByIdBorrower() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.FindByIdBorrower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetSimulations(t *testing.T) {
	type fields struct {
		db     *sql.DB
		logger logger.ILogCloudWatch
	}
	type args struct {
		param *model.Params
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Simulation
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			got, err := r.GetSimulations(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetSimulations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetSimulations() = %v, want %v", got, tt.want)
			}
		})
	}
}
