package service

import (
	"context"
	sqsAws "project-golang/internal/adapters/cloud/aws/sqs"
	"project-golang/internal/domain/model"
	"project-golang/internal/repository"
	"reflect"
	"testing"

	"github.com/golang-jwt/jwt/v4"
)

func TestSimulationService_SimulationResponseBorrower(t *testing.T) {
	type fields struct {
		repository repository.IRepository
		sqsClient  sqsAws.Client
	}
	type args struct {
		id       string
		response *model.SimulationResponseBorrower
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
			s := &SimulationService{
				repository: tt.fields.repository,
				sqsClient:  tt.fields.sqsClient,
			}
			if err := s.SimulationResponseBorrower(tt.args.id, tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("SimulationService.SimulationResponseBorrower() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSimulationService_CreatedBorrower(t *testing.T) {
	type fields struct {
		repository repository.IRepository
		sqsClient  sqsAws.Client
	}
	type args struct {
		tom *model.Borrower
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
			s := &SimulationService{
				repository: tt.fields.repository,
				sqsClient:  tt.fields.sqsClient,
			}
			if err := s.CreatedBorrower(tt.args.tom); (err != nil) != tt.wantErr {
				t.Errorf("SimulationService.CreatedBorrower() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSimulationService_CreatedSetup(t *testing.T) {
	type fields struct {
		repository repository.IRepository
		sqsClient  sqsAws.Client
	}
	type args struct {
		set *model.Setup
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
			s := &SimulationService{
				repository: tt.fields.repository,
				sqsClient:  tt.fields.sqsClient,
			}
			if err := s.CreatedSetup(tt.args.set); (err != nil) != tt.wantErr {
				t.Errorf("SimulationService.CreatedSetup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSimulationService_CreatedSimulation(t *testing.T) {
	type fields struct {
		repository repository.IRepository
		sqsClient  sqsAws.Client
	}
	type args struct {
		ctx   context.Context
		simu  *model.Simulation
		token *jwt.Token
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
			s := &SimulationService{
				repository: tt.fields.repository,
				sqsClient:  tt.fields.sqsClient,
			}
			if err := s.CreatedSimulation(tt.args.ctx, tt.args.simu, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("SimulationService.CreatedSimulation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSimulationService_FindByIdBorrower(t *testing.T) {
	type fields struct {
		repository repository.IRepository
		sqsClient  sqsAws.Client
	}
	type args struct {
		borrwerId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Borrower
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationService{
				repository: tt.fields.repository,
				sqsClient:  tt.fields.sqsClient,
			}
			got, err := s.FindByIdBorrower(tt.args.borrwerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("SimulationService.FindByIdBorrower() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimulationService.FindByIdBorrower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimulationService_FindByIdSetup(t *testing.T) {
	type fields struct {
		repository repository.IRepository
		sqsClient  sqsAws.Client
	}
	type args struct {
		setupId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Setup
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationService{
				repository: tt.fields.repository,
				sqsClient:  tt.fields.sqsClient,
			}
			got, err := s.FindByIdSetup(tt.args.setupId)
			if (err != nil) != tt.wantErr {
				t.Errorf("SimulationService.FindByIdSetup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimulationService.FindByIdSetup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimulationService_FindByIdSimulation(t *testing.T) {
	type fields struct {
		repository repository.IRepository
		sqsClient  sqsAws.Client
	}
	type args struct {
		simulationId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Simulation
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationService{
				repository: tt.fields.repository,
				sqsClient:  tt.fields.sqsClient,
			}
			got, err := s.FindByIdSimulation(tt.args.simulationId)
			if (err != nil) != tt.wantErr {
				t.Errorf("SimulationService.FindByIdSimulation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimulationService.FindByIdSimulation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimulationService_UpdateSetup(t *testing.T) {
	type fields struct {
		repository repository.IRepository
		sqsClient  sqsAws.Client
	}
	type args struct {
		setupId  string
		newSetup *model.Setup
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
			s := &SimulationService{
				repository: tt.fields.repository,
				sqsClient:  tt.fields.sqsClient,
			}
			if err := s.UpdateSetup(tt.args.newSetup); (err != nil) != tt.wantErr {
				t.Errorf("SimulationService.UpdateSetup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSimulationService_UpdateSimulationStatus(t *testing.T) {
	type fields struct {
		repository repository.IRepository
		sqsClient  sqsAws.Client
	}
	type args struct {
		simu *model.Simulation
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
			s := &SimulationService{
				repository: tt.fields.repository,
				sqsClient:  tt.fields.sqsClient,
			}
			if err := s.UpdateSimulation(tt.args.simu); (err != nil) != tt.wantErr {
				t.Errorf("SimulationService.UpdateSimulationStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSimulationService_TokenIsValid(t *testing.T) {
	type fields struct {
		repository repository.IRepository
		sqsClient  sqsAws.Client
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *jwt.Token
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationService{
				repository: tt.fields.repository,
				sqsClient:  tt.fields.sqsClient,
			}
			got, err := s.TokenIsValid(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("SimulationService.TokenIsValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimulationService.TokenIsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimulationService_GenerateJWT(t *testing.T) {
	type fields struct {
		repository repository.IRepository
		sqsClient  sqsAws.Client
	}
	type args struct {
		payload model.PayloadJWT
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationService{
				repository: tt.fields.repository,
				sqsClient:  tt.fields.sqsClient,
			}
			got, err := s.GenerateJWT(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("SimulationService.GenerateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SimulationService.GenerateJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimulationService_Ping(t *testing.T) {
	type fields struct {
		repository repository.IRepository
		sqsClient  sqsAws.Client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationService{
				repository: tt.fields.repository,
				sqsClient:  tt.fields.sqsClient,
			}
			if err := s.Ping(); (err != nil) != tt.wantErr {
				t.Errorf("SimulationService.Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
