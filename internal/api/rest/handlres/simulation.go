package handlres

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"project-golang/internal/domain/dto"
	"project-golang/internal/domain/model"
	service "project-golang/internal/services"
	"project-golang/internal/utils"
	"project-golang/internal/utils/errors"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type ISimulationHandler interface {
	CreatedSimulation(w http.ResponseWriter, r *http.Request)
	CreatedBorrower(w http.ResponseWriter, r *http.Request)
	CreatedSetup(w http.ResponseWriter, r *http.Request)
	FindByIdSimulation(w http.ResponseWriter, r *http.Request)
	FindByIdSetup(w http.ResponseWriter, r *http.Request)
	FindByIdBorrower(w http.ResponseWriter, r *http.Request)
	UpdateSetup(w http.ResponseWriter, r *http.Request)
	UpdateSimulation(w http.ResponseWriter, r *http.Request)
	BorrowerResponseToSimulation(w http.ResponseWriter, r *http.Request)
	GenerateJWTw(w http.ResponseWriter, r *http.Request)
	HealthCheckHandler(w http.ResponseWriter, r *http.Request)
	FindSimulationsByParam(w http.ResponseWriter, r *http.Request)
}

type SimulationHandler struct {
	service service.ISimulationService
}

func NewSimulationHandler(serv service.ISimulationService) ISimulationHandler {
	return &SimulationHandler{service: serv}
}

// @Summary Find Simulations By Param
// @Description Find Simulations By param, with all the parameters of the simulation object with the page and pageSize as desired ex: status="acepted", borrowerId=12345
// @Tags simulations
// @Accept  json
// @Produce  json
//
//	@Success 201 {object} dto.SimulationPaginationResponse
//
// exemple:{ "simulations": [
//
//	    {
//	      "simulationId": "sim123",
//	      "borrowerId": "bor123",
//	      "loanValue": 50000.00,
//	      "numberOfInstallments": 24,
//	      "createdAt": "2024-12-25T15:04:05Z",
//	      "updatedAt": "2024-12-25T15:04:05Z",
//	      "satus": "APPROVED",
//	      "interestRate": 3.5
//	    },
//	    {
//	      "simulationId": "sim124",
//	      "borrowerId": "bor124",
//	      "loanValue": 30000.00,
//	      "numberOfInstallments": 12,
//	      "createdAt": "2024-12-25T16:10:15Z",
//	      "updatedAt": "2024-12-25T16:10:15Z",
//	      "satus": "PENDING",
//	      "interestRate": 2.8
//	    }
//	  ],
//	  "page": "1",
//	  "pageSize": "10"
//	} JwtRequest
//
// @Router /v1/simulation [GET]
// FindSimulationsByParam implements ISimulationHandler.
func (s *SimulationHandler) FindSimulationsByParam(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-user-ID")
	tenant := r.Header.Get("X-tenant-ID")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-user-ID not exists"))
		return
	}

	if len(tenant) == 0 || tenant == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-tenant-ID not exists"))
		return
	}

	token, err := extractToken(r)
	if err != nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil && tokenJwt == nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}

	layout := "2006-01-02T15:04:05"

	updatedAt, _ := time.Parse(layout, r.URL.Query().Get("updatedAt"))

	createdAt, _ := time.Parse(layout, r.URL.Query().Get("createdAt"))

	simulationId := r.URL.Query().Get("simulationId")
	borrowerId := r.URL.Query().Get("borrowerId ")
	loanValue, _ := strconv.ParseFloat(r.URL.Query().Get("loanValue"), 64)
	numberOfInstallments, _ := strconv.ParseFloat(r.URL.Query().Get("numberOfInstallments"), 64)

	status := r.URL.Query().Get("status")
	interestRate, _ := strconv.ParseFloat(r.URL.Query().Get("interestRate"), 64)
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	simulations, err := s.service.FindByParamSimulations(&model.Params{
		Simu: &model.SimulationParam{
			SimulationId:         &simulationId,
			BorrowerId:           &borrowerId,
			LoanValue:            &loanValue,
			NumberOfInstallments: &numberOfInstallments,
			Status:               &status,
			InterestRate:         &interestRate,
			CreatedAt:            &createdAt,
			UpdatedAt:            &updatedAt,
		},
		Page:     page,
		PageSize: pageSize,
	}, tenant)
	if err != nil {
		utils.ErrorResponse(w, errors.Internalf("request error: %v", err))
		return
	}

	utils.SuccessResponse(w, http.StatusOK, simulations)

}

// @Summary GenerateJWTw
// @Description generate token jwt
// @Tags GenerateJWTw
// @Accept  json
// @Produce  json
// @Success 201 {object} dto.JwtResponse
// {"token": "12233hdjfj474748wkdmms"} JwtRequest
// @Router /v1/generateJwt [POST]

func (s *SimulationHandler) GenerateJWTw(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User")

	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-user-ID not exists"))
		return
	}

	request := dto.JwtRequest{}

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

	token, err := s.service.GenerateJWT(*dto.ToPayloadJWTModel(request))

	if err != nil {
		//logar error
		utils.ErrorResponse(w, errors.Internalf("request error: %v", err))
		return
	}

	utils.SuccessResponse(w, http.StatusCreated, dto.JwtResponse{
		Token: token,
	})

}

// @Summary Created Borrower
// @Description Created Borrower
// @Tags Borrower
// @Accept  json
// @Produce  json
// @Success 201
// @Router /v1/Borrower [POST]

func (s *SimulationHandler) CreatedBorrower(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User-ID")
	tenant := r.Header.Get("X-tenant-ID")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-user-ID not exists"))
		return
	}

	if len(tenant) == 0 || tenant == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-tenant-ID not exists"))
		return
	}
	token, err := extractToken(r)
	if err != nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil && tokenJwt == nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}

	request := dto.BorrowerRequest{}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		//logar error
		utils.ErrorResponse(w, errors.UnprocessableEntityf("unprocessable entity error: %v", err))
		return
	}
	error := utils.ValidateStruct(request)
	if error != nil {
		//logar error
		utils.ErrorResponse(w, errors.BadRequestf("bad request error: %v", err))
		return
	}

	error = utils.Validate(request)
	if error != nil {
		//logar error
		utils.ErrorResponse(w, errors.BadRequestf("bad request error: %v", err))
		return
	}

	err = s.service.CreatedBorrower(dto.ToBorrowerModel(&request), tenant)

	if err != nil {
		//logar error
		utils.ErrorResponse(w, errors.Internalf("request error: %v", err))
		return
	}

	utils.SuccessResponse(w, http.StatusCreated, nil)

}

// @Summary Created Setup
// @Description Created Setup
// @Tags Borrower
// @Accept  json
// @Produce  json
// @Success 201
// @Router /v1/setup [POST]
// CreatedSetup implements ISimulationHandler.
func (s *SimulationHandler) CreatedSetup(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User-ID")
	tenant := r.Header.Get("X-tenant-ID")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-user-ID not exists"))
		return
	}

	if len(tenant) == 0 || tenant == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-tenant-ID not exists"))
		return
	}

	token, err := extractToken(r)
	if err != nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil || tokenJwt == nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}

	request := dto.SetupRequest{}

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
	error = utils.Validate(request)
	if error != nil {
		//logar error
		utils.ErrorResponse(w, errors.BadRequestf("bad request error: %v", err))
		return
	}

	err = s.service.CreatedSetup(dto.ToSetupModel(&request), tenant)

	if err != nil {
		//logar error
		utils.ErrorResponse(w, errors.Internalf("request error: %v", err))
		return
	}

	utils.SuccessResponse(w, http.StatusCreated, nil)

}

// @Summary Created Simulation
// @Description cCreated Simulation
// @Tags Setup
// @Accept  json
// @Produce  json
// @Success 201
// @Router /v1/simulation [POST]
func (s *SimulationHandler) CreatedSimulation(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User-ID")
	tenant := r.Header.Get("X-tenant-ID")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-user-ID not exists"))
		return
	}

	if len(tenant) == 0 || tenant == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-tenant-ID not exists"))
		return
	}
	token, err := extractToken(r)
	if err != nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil || tokenJwt == nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}

	request := dto.SimulationRequest{}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		//logar error
		utils.ErrorResponse(w, errors.UnprocessableEntityf("unprocessable entity error: %v", err))
		return
	}

	error := utils.ValidateStruct(request)
	if error != nil {
		//logar error
		utils.ErrorResponse(w, errors.BadRequestf("bad request error: %v", err))
		return
	}
	error = utils.Validate(request)
	if error != nil {
		//logar error
		utils.ErrorResponse(w, errors.BadRequestf("bad request error: %v", err))
		return
	}

	err = s.service.CreatedSimulation(ctx, dto.ToSimulationModel(&request), tokenJwt, tenant)

	if err != nil {
		//logar error
		utils.ErrorResponse(w, errors.Internalf("request error: %v", err))
		return
	}

	utils.SuccessResponse(w, http.StatusCreated, nil)

}

// @Summary Find By Id Borrower
// @Description Find By Id Borrower
// @Tags simulation
// @Accept  json
// @Produce  json
//
//	@Success 200 {object} dto.BorrowerResponse
//
//	"example:{
//		  "borrowerId": "12345",
//		  "name": "João Silva",
//		  "phone": "123456789",
//		  "email": "joao.silva@email.com",
//		  "cpf": "123.456.789-00",
//		  "createdAt": "2024-12-25T15:04:05Z",
//		  "updatedAt": "2024-12-25T15:04:05Z"
//		} "
//
// @Router /v1/borrower/{id} [GET]
// FindByIdBorrower implements ISimulationHandler.
func (s *SimulationHandler) FindByIdBorrower(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User-ID")
	tenant := r.Header.Get("X-tenant-ID")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-user-ID not exists"))
		return
	}

	if len(tenant) == 0 || tenant == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-tenant-ID not exists"))
		return
	}
	token, err := extractToken(r)
	if err != nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil && tokenJwt == nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {

		utils.ErrorResponse(w, errors.BadRequestf("error to parser query params: %v", id))
		return
	}

	borrower, err := s.service.FindByIdBorrower(id, tenant)
	if err != nil {

		utils.ErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.SuccessResponse(w, http.StatusCreated, dto.BorrowerResponse{
		BorrowerId: borrower.BorrowerId,
		Name:       borrower.Name,
		Phone:      borrower.Phone,
		Email:      borrower.Email,
		Cpf:        borrower.Cpf,
		CreatedAt:  borrower.CreatedAt,
		UpdateAt:   borrower.UpdateAt,
	})

}

// @Summary Find By Id Setup
// @Description Find By Id Setup
// @Tags Borrower
// @Accept  json
// @Produce  json
//
//	@Success 200 {object} dto.SetupResponse
//
//	"example:{
//		  "setupId": "abc123",
//		  "capital": 10000.00,
//		  "fees": 500.00,
//		  "interestRate": 5.5,
//		  "escope": "Investment",
//		  "escopeIdValid": true,
//		  "createdAt": "2024-12-25T15:04:05Z",
//		  "updatedAt": "2024-12-25T15:04:05Z"
//		}"
//
// @Router /v1/setup/{id} [GET]
// FindByIdSetup implements ISimulationHandler.
func (s *SimulationHandler) FindByIdSetup(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User-ID")
	tenant := r.Header.Get("X-tenant-ID")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-user-ID not exists"))
		return
	}

	if len(tenant) == 0 || tenant == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-tenant-ID not exists"))
		return
	}

	token, err := extractToken(r)
	if err != nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil && tokenJwt == nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {

		utils.ErrorResponse(w, errors.BadRequestf("error to parser query params: %v", id))
		return
	}

	setup, err := s.service.FindByIdSetup(id, tenant)
	if err != nil {

		utils.ErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.SuccessResponse(w, http.StatusCreated, dto.SetupResponse{
		SetupId:       setup.SetupId,
		Capital:       setup.Capital,
		Fees:          setup.Fees,
		InterestRate:  setup.InterestRate,
		Escope:        setup.Escope,
		EscopeIdValid: setup.EscopeIsValid,
		CreatedAt:     setup.CreatedAt,
		UpdatedAt:     setup.UpdatedAt,
	})
}

// @Summary Find By Id Simulation
// @Description Find By Id Simulation
// @Tags simulation
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.SimulationResponse
// example:{
//   "simulationId": "sim123",
//   "borrowerId": "bor123",
//   "loanValue": 50000.00,
//   "numberOfInstallments": 24,
//   "createdAt": "2024-12-25T15:04:05Z",
//   "updatedAt": "2024-12-25T15:04:05Z",
//   "satus": "APPROVED",
//   "interestRate": 3.5
// } FindByIdSimulation

// @Router /v1/simulation/{id} [GET]
// FindByIdSimulation implements ISimulationHandler.
func (s *SimulationHandler) FindByIdSimulation(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User-ID")
	tenant := r.Header.Get("X-tenant-ID")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-user-ID not exists"))
		return
	}

	if len(tenant) == 0 || tenant == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-tenant-ID not exists"))
		return
	}
	token, err := extractToken(r)
	if err != nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", token))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil && tokenJwt == nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {

		utils.ErrorResponse(w, errors.BadRequestf("error to parser query params: %v", id))
		return
	}

	simu, err := s.service.FindByIdSimulation(id, tenant)
	if err != nil {

		utils.ErrorResponse(w, err)
		return
	}

	utils.SuccessResponse(w, http.StatusOK, dto.SimulationResponse{
		SimulationId:         simu.SimulationId,
		InterestRate:         simu.InterestRate,
		CreatedAt:            simu.CreatedAt,
		UpdatedAt:            simu.UpdatedAt,
		BorrowerId:           simu.BorrowerId,
		LoanValue:            simu.LoanValue,
		NumberOfInstallments: simu.NumberOfInstallments,
		Status:               simu.Status,
	})
}

// @Summary Update Setup
// @Description Update Setup
// @Tags Setup
// @Accept  json
// @Produce  json
// @Success 200
// @Router /v1/setup [PUT]
// UpdateSetup implements ISimulationHandler.
func (s *SimulationHandler) UpdateSetup(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User-ID")
	tenant := r.Header.Get("X-tenant-ID")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-user-ID not exists"))
		return
	}

	if len(tenant) == 0 || tenant == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-tenant-ID not exists"))
		return
	}

	token, err := extractToken(r)
	if err != nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil && tokenJwt == nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}

	newSetup := dto.SetupRequest{}

	err = json.NewDecoder(r.Body).Decode(&newSetup)
	if err != nil {
		//logar error
		utils.ErrorResponse(w, errors.UnprocessableEntityf("unprocessable entity error: %v", err))
		return
	}
	error := utils.ValidateStruct(newSetup)
	if error != nil {
		//logar error
		utils.ErrorResponse(w, errors.BadRequestf("bad request error: %v", err))
		return
	}

	error = utils.Validate(newSetup)
	if error != nil {
		//logar error
		utils.ErrorResponse(w, errors.BadRequestf("bad request error: %v", err))
		return
	}
	err = s.service.UpdateSetup(&model.Setup{
		Capital:      newSetup.Capital,
		Fees:         newSetup.Fees,
		InterestRate: newSetup.InterestRate,
		Escope:       newSetup.Escope,
	}, tenant)
	if err != nil {

		utils.ErrorResponse(w,  err)
		return
	}

	utils.SuccessResponse(w, http.StatusOK, nil)

}

// @Summary Update Simulation Status
// @Description Update Simulation Status
// @Tags simulation
// @Accept  json
// @Produce  json
// @Success 200
// @Router /v1/simulation/{id} [PUT]
// UpdateSimulationStatus implements ISimulationHandler.
func (s *SimulationHandler) UpdateSimulation(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User-ID")
	tenant := r.Header.Get("X-tenant-ID")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-user-ID not exists"))
		return
	}

	if len(tenant) == 0 || tenant == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-tenant-ID not exists"))
		return
	}
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {

		utils.ErrorResponse(w, errors.BadRequestf("error to parser query params: %v", id))
		return
	}
	token, err := extractToken(r)
	if err != nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil && tokenJwt == nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}

	newSimu := dto.SimulationRequest{}

	err = json.NewDecoder(r.Body).Decode(&newSimu)
	if err != nil {
		//logar error
		utils.ErrorResponse(w, errors.UnprocessableEntityf("unprocessable entity error: %v", err))
		return
	}
	error := utils.ValidateStruct(newSimu)
	if error != nil {
		//logar error
		utils.ErrorResponse(w, errors.BadRequestf("bad request error: %v", err))
		return
	}

	error = utils.Validate(newSimu)
	if error != nil {
		//logar error
		utils.ErrorResponse(w, errors.BadRequestf("bad request error: %v", err))
		return
	}
	vars = mux.Vars(r)
	id, ok = vars["id"]
	if !ok {

		utils.ErrorResponse(w, errors.BadRequestf("error to parser query params: %v", id))
		return
	}

	err = s.service.UpdateSimulation(&model.Simulation{
		SimulationId:         id,
		BorrowerId:           newSimu.BorrowerId,
		LoanValue:            newSimu.LoanValue,
		NumberOfInstallments: newSimu.NumberOfInstallments,
		InterestRate:         newSimu.InterestRate,
	}, tenant)
	if err != nil {

		utils.ErrorResponse(w, errors.Internalf("error to parser query params: %v", err))
		return
	}

	utils.SuccessResponse(w, http.StatusOK, nil)

}

// @Summary Borrower Response To Simulation
// @Description borrower's response to the loan simulation
// @Tags simulation
// @Accept  json
// @Produce  json
// @Success 200
// @Router /v1/simulation/{id} [PATCH]
func (s *SimulationHandler) BorrowerResponseToSimulation(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "generatJWT", r.Header.Get("X-User"))
	user := r.Header.Get("X-User-ID")
	tenant := r.Header.Get("X-tenant-ID")
	w.Header().Set("Content-Type", "application/json")

	if len(user) == 0 || user == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-user-ID not exists"))
		return
	}

	if len(tenant) == 0 || tenant == "" {

		// fazer log user invalid
		utils.ErrorResponse(w, errors.BadRequestf("header X-tenant-ID not exists"))
		return
	}
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {

		utils.ErrorResponse(w, errors.BadRequestf("error to parser query params: %v", id))
		return
	}

	token, err := extractToken(r)
	if err != nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	tokenJwt, err := s.service.TokenIsValid(token)

	if err != nil && tokenJwt == nil {

		utils.ErrorResponse(w, errors.Unauthorizedf("Unauthorized error: %v", err))
		return
	}
	request := &dto.BorrowerResponseTosimulationRequest{}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		//logar error

		utils.ErrorResponse(w, errors.UnprocessableEntityf("unprocessable entity error: %v", err))
		return
	}
	if request.Status != "accepted" && request.Status != "rejected" {

		utils.ErrorResponse(w, errors.UnprocessableEntityf("unprocessable entity error: %v", err))
		return

	}
	error := utils.ValidateStruct(&request)
	if error != nil {
		//logar error

		utils.ErrorResponse(w, errors.BadRequestf("bad request error: %v", err))
		return
	}

	err = s.service.SimulationResponseBorrower(id, &model.SimulationResponseBorrower{
		Status: request.Status}, tenant)

	if err != nil {
		//logar error

		utils.ErrorResponse(w, errors.Internalf("request error: %v", err))
		return
	}

	utils.SuccessResponse(w, http.StatusOK, nil)

}

// @Summary HealthCheckHandler
// @Description cria uma simulation de emprestimo
// @Tags simulation
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.HelfCheckResponse

// @Router /v1/health-handler [GET]
func (s *SimulationHandler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Testa a conexão com o banco
	err := s.service.Ping()
	if err != nil {

		_, err = w.Write([]byte(fmt.Sprintf(`{"status": "DOWN", "error": "%v"}`, err)))
		if err != nil {
			utils.ErrorResponse(w, err)
		}
		return
	}

	// Responde com sucesso
	utils.SuccessResponse(w, http.StatusOK, dto.HelfCheckResponse{
		Status:        "UP",
		Response_time: &start,
	})
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
