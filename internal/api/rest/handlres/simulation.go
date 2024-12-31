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

func NewSimulationHandler(serv service.ISimulationService) ISimulationHandler {
	return &SimulationHandler{service: serv}
}

// @Summary GenerateJWTw
// @Description gera um token jwt pra solicitação de emprestimo
// @Tags GenerateJWTw
// @Accept  json
// @Produce  json
// @Success 201 {jwtRequest} JwtRequest
// @Router /v1/simulation/generateJwt [POST]

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

	if err != nil {
		//logar error
		utils.ErrorResponse(w, errors.Internalf("request error: %v", err))
		return
	}

	utils.SuccessResponse(w, http.StatusCreated, dto.JwtResponse{
		Token: token,
	})

}

// @Summary CreatedBorrower
// @Description cria um Borrower
// @Tags Borrower
// @Accept  json
// @Produce  json
// @Success 201 {borrower} Borrower
// @Router /v1/Borrower [POST]

func (s *SimulationHandler) CreatedBorrower(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-User not exists"))
		return
	}

	token, err := extractToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil && tokenJwt == nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}

	request := &dto.BorrowerRequest{}

	err = json.NewDecoder(r.Body).Decode(&request)
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

	err = s.service.CreatedBorrower(dto.ToBorrowerModel(request))

	if err != nil {
		//logar error
		utils.ErrorResponse(w, errors.Internalf("request error: %v", err))
		return
	}

	utils.SuccessResponse(w, http.StatusCreated, nil)

}

// @Summary CreatedSetup
// @Description cria um Setup
// @Tags Setup
// @Accept  json
// @Produce  json
// @Success 201 {setup} Setup
// @Router /v1/setup [POST]
// CreatedSetup implements ISimulationHandler.
func (s *SimulationHandler) CreatedSetup(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-User not exists"))
		return
	}

	token, err := extractToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil && tokenJwt == nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}

	request := &dto.SetupRequest{}

	err = json.NewDecoder(r.Body).Decode(&request)
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

	err = s.service.CreatedSetup(dto.ToSetupModel(request))

	if err != nil {
		//logar error
		utils.ErrorResponse(w, errors.Internalf("request error: %v", err))
		return
	}

	utils.SuccessResponse(w, http.StatusCreated, nil)

}

// @Summary cria uma simulation
// @Description cria uma simulation de emprestimo
// @Tags simulation
// @Accept  json
// @Produce  json
// @Success 201 {simulation} Simulation
// @Router /v1/simulation [POST]
func (s *SimulationHandler) CreatedSimulation(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-User not exists"))
		return
	}

	token, err := extractToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil && tokenJwt == nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}

	request := &dto.SimulationRequest{}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		//logar error
		utils.ErrorResponse(w, errors.UnprocessableEntityf("unprocessable entity error: %v", err))
		return
	}

	err = s.service.CreatedSimulation(ctx, dto.ToSimulationModel(request), tokenJwt)

	if err != nil {
		//logar error
		utils.ErrorResponse(w, errors.Internalf("request error: %v", err))
		return
	}

	utils.SuccessResponse(w, http.StatusCreated, nil)

}

// FindByIdBorrower implements ISimulationHandler.
func (s *SimulationHandler) FindByIdBorrower(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-User not exists"))
		return
	}

	token, err := extractToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil && tokenJwt == nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
}

// FindByIdSetup implements ISimulationHandler.
func (s *SimulationHandler) FindByIdSetup(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-User not exists"))
		return
	}

	token, err := extractToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil && tokenJwt == nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
}

// FindByIdSimulation implements ISimulationHandler.
func (s *SimulationHandler) FindByIdSimulation(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-User not exists"))
		return
	}

	token, err := extractToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil && tokenJwt == nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
}

// UpdateSetup implements ISimulationHandler.
func (s *SimulationHandler) UpdateSetup(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-User not exists"))
		return
	}

	token, err := extractToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil && tokenJwt == nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
}

// UpdateSimulationStatus implements ISimulationHandler.
func (s *SimulationHandler) UpdateSimulationStatus(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-User not exists"))
		return
	}

	token, err := extractToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil && tokenJwt == nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
}

func (s *SimulationHandler) SimulationResponseBorrower(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-User not exists"))
		return
	}

	token, err := extractToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil && tokenJwt == nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
}

// @Summary HealthCheckHandler
// @Description cria uma simulation de emprestimo
// @Tags simulation
// @Accept  json
// @Produce  json
// @Success 200 {healthCheckHandler} healthCheckHandler
// @Router /v1/health-handler [GET]
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
