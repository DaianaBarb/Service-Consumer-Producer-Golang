package handlres

import (
	"net/http"
	service "project-golang/internal/services"
	"project-golang/internal/services/mocks"
	"testing"
)

func TestSimulationHandler_FindSimulationsByParam(t *testing.T) {
	type fields struct {
		service *mocks.ISimulationService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mock    func(*mocks.ISimulationService)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationHandler{
				service: tt.fields.service,
			}
			s.FindSimulationsByParam(tt.args.w, tt.args.r)
		})
	}
}

func TestSimulationHandler_GenerateJWTw(t *testing.T) {
	type fields struct {
		service service.ISimulationService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationHandler{
				service: tt.fields.service,
			}
			s.GenerateJWTw(tt.args.w, tt.args.r)
		})
	}
}

func TestSimulationHandler_CreatedBorrower(t *testing.T) {
	type fields struct {
		service service.ISimulationService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationHandler{
				service: tt.fields.service,
			}
			s.CreatedBorrower(tt.args.w, tt.args.r)
		})
	}
}

func TestSimulationHandler_CreatedSetup(t *testing.T) {
	type fields struct {
		service service.ISimulationService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationHandler{
				service: tt.fields.service,
			}
			s.CreatedSetup(tt.args.w, tt.args.r)
		})
	}
}

func TestSimulationHandler_CreatedSimulation(t *testing.T) {
	type fields struct {
		service service.ISimulationService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationHandler{
				service: tt.fields.service,
			}
			s.CreatedSimulation(tt.args.w, tt.args.r)
		})
	}
}

func TestSimulationHandler_FindByIdBorrower(t *testing.T) {
	type fields struct {
		service service.ISimulationService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationHandler{
				service: tt.fields.service,
			}
			s.FindByIdBorrower(tt.args.w, tt.args.r)
		})
	}
}

func TestSimulationHandler_FindByIdSetup(t *testing.T) {
	type fields struct {
		service service.ISimulationService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationHandler{
				service: tt.fields.service,
			}
			s.FindByIdSetup(tt.args.w, tt.args.r)
		})
	}
}

func TestSimulationHandler_FindByIdSimulation(t *testing.T) {
	type fields struct {
		service service.ISimulationService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationHandler{
				service: tt.fields.service,
			}
			s.FindByIdSimulation(tt.args.w, tt.args.r)
		})
	}
}

func TestSimulationHandler_UpdateSetup(t *testing.T) {
	type fields struct {
		service service.ISimulationService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationHandler{
				service: tt.fields.service,
			}
			s.UpdateSetup(tt.args.w, tt.args.r)
		})
	}
}

func TestSimulationHandler_UpdateSimulation(t *testing.T) {
	type fields struct {
		service service.ISimulationService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationHandler{
				service: tt.fields.service,
			}
			s.UpdateSimulation(tt.args.w, tt.args.r)
		})
	}
}

func TestSimulationHandler_BorrowerResponseToSimulation(t *testing.T) {
	type fields struct {
		service service.ISimulationService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationHandler{
				service: tt.fields.service,
			}
			s.BorrowerResponseToSimulation(tt.args.w, tt.args.r)
		})
	}
}

func TestSimulationHandler_HealthCheckHandler(t *testing.T) {
	type fields struct {
		service service.ISimulationService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationHandler{
				service: tt.fields.service,
			}
			s.HealthCheckHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_extractToken(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractToken(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
