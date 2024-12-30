package handlres

import (
	"net/http"
	service "project-golang/internal/services"
)

type ISimulationHandler interface {
	CreatedSimulation(w http.ResponseWriter, r *http.Request)
	CreatedBorrower(w http.ResponseWriter, r *http.Request)
	CreatedSetup(w http.ResponseWriter, r *http.Request)
	FindByIdSimulation(w http.ResponseWriter, r *http.Request)
	FindByIdSetup(w http.ResponseWriter, r *http.Request)
	FindByIdBorrower(w http.ResponseWriter, r *http.Request)
	UpdateSetup(w http.ResponseWriter, r *http.Request)
	UpdateSimulationStatus(w http.ResponseWriter, r *http.Request)
	SimulationResponseBorrower(w http.ResponseWriter, r *http.Request)
	GenerateJWTw(w http.ResponseWriter, r *http.Request)
}

type SimulationHandler struct {
	service *service.ISimulationService
}

// GenerateJWTw implements ISimulationHandler.
func (s *SimulationHandler) GenerateJWTw(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func NewSimulationHandler(serv *service.ISimulationService) ISimulationHandler {
	return &SimulationHandler{service: serv}
}

// CreatedBorrower implements ISimulationHandler.
func (s *SimulationHandler) CreatedBorrower(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// CreatedSetup implements ISimulationHandler.
func (s *SimulationHandler) CreatedSetup(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// CreatedSimulation implements ISimulationHandler.
func (s *SimulationHandler) CreatedSimulation(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// FindByIdBorrower implements ISimulationHandler.
func (s *SimulationHandler) FindByIdBorrower(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// FindByIdSetup implements ISimulationHandler.
func (s *SimulationHandler) FindByIdSetup(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// FindByIdSimulation implements ISimulationHandler.
func (s *SimulationHandler) FindByIdSimulation(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// UpdateSetup implements ISimulationHandler.
func (s *SimulationHandler) UpdateSetup(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// UpdateSimulationStatus implements ISimulationHandler.
func (s *SimulationHandler) UpdateSimulationStatus(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (s *SimulationHandler) SimulationResponseBorrower(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
