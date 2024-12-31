package handlres

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"project-golang/internal/domain/dto"
	service "project-golang/internal/services"
	"project-golang/internal/utils"
	"project-golang/internal/utils/errors"
	"strings"
	"time"
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
	HealthCheckHandler(w http.ResponseWriter, r *http.Request)
}

type SimulationHandler struct {
	service service.ISimulationService
}

// GenerateJWTw implements ISimulationHandler.
func (s *SimulationHandler) GenerateJWTw(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-User not exists"))
		return
	}
	request := &dto.JwtRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		//logar error
		utils.ErrorResponse(w, errors.UnprocessableEntityf("unprocessable entity error: %v", err))
		return
	}
	error := utils.ValidateStruct(&request)
	if error != nil {
		//logar error
		utils.ErrorResponse(w, errors.BadRequestf("bad request error: %v", err))
		return
	}

	token, err := s.service.GenerateJWT(*dto.ToPayloadJWTModel(*request))

	if error != nil {
		//logar error
		utils.ErrorResponse(w, errors.Internalf("request error: %v", err))
		return
	}

	utils.SuccessResponse(w, http.StatusCreated,  dto.JwtResponse{
		Token: token,
	})

}

func NewSimulationHandler(serv service.ISimulationService) ISimulationHandler {
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

// @Summary cria uma simulation
// @Description cria uma simulation de emprestimo
// @Tags simulation
// @Accept  json
// @Produce  json
// @Success 201 {simulation} Simulation
// @Router /v1/simulation [POST]
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

func (s *SimulationHandler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Testa a conexão com o banco
	err := s.service.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"status": "DOWN", "error": "%v"}`, err)))
		return
	}

	// Responde com sucesso
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"status": "UP", "response_time": "%v"}`, time.Since(start))))
}

func extractToken(r *http.Request) (string, error) {
	// Obtém o cabeçalho Authorization
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("cabeçalho Authorization ausente")
	}

	// O cabeçalho deve começar com "Bearer "
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("formato inválido do cabeçalho Authorization")
	}

	return parts[1], nil
}
