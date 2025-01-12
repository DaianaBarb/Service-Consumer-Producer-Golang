package handlres

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"project-golang/internal/domain/dto"
	"project-golang/internal/domain/model"
	"project-golang/internal/services/mocks"
	"reflect"
	"strings"
	"testing"
	"time"

	"project-golang/internal/utils/errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

// Cria o token com método de assinatura HS256

func getToken() *jwt.Token {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"CredorId":  "TENANT_1",
		"Escope":    "escope",
		"Expiracao": float64(time.Now().Add(48 * time.Hour).Unix()),
	})
	return token

}

func TestSimulationHandler_GenerateJWTw(t *testing.T) {
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
		service *mocks.ISimulationService
	}
	tests := []struct {
		name               string
		fields             fields
		body               strings.Reader
		wantHTTPStatusCode int
		mock               func(service *mocks.ISimulationService)
	}{
		{name: "sucess",
			fields:             fields{service: new(mocks.ISimulationService)},
			wantHTTPStatusCode: 201,
			body: *strings.NewReader(`{
								"nome": "João Silva soares",
								"phone": "(11) 91234-5678",
								"email": "joao.silva2@email.com",
								"cpf": "123.456.789-01"
										}
`), mock: func(service *mocks.ISimulationService) {
				service.On("CreatedBorrower", mock.Anything, mock.Anything).Return(nil)
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			}},
		{
			name:               "validation error",
			fields:             fields{service: new(mocks.ISimulationService)},
			wantHTTPStatusCode: 422,
			body: *strings.NewReader(`{
								"nome": 0989,
								"phone": "",
								"email": "invalid-email",
								"cpf": ""
							}`),
			mock: func(service *mocks.ISimulationService) {
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
		{
			name:               "invalid token",
			fields:             fields{service: new(mocks.ISimulationService)},
			wantHTTPStatusCode: 401,
			body: *strings.NewReader(`{
								"nome": "João Silva",
								"phone": "(11) 91234-5678",
								"email": "joao.silva@email.com",
								"cpf": "123.456.789-01"
							}`),
			mock: func(service *mocks.ISimulationService) {
				service.On("TokenIsValid", mock.Anything).Return(nil, errors.New("invalid token"))
			},
		},
		{
			name:               "service error",
			fields:             fields{service: new(mocks.ISimulationService)},
			wantHTTPStatusCode: 500,
			body: *strings.NewReader(`{
								"nome": "João Silva",
								"phone": "(11) 91234-5678",
								"email": "joao.silva@email.com",
								"cpf": "123.456.789-01"
							}`),
			mock: func(service *mocks.ISimulationService) {
				service.On("CreatedBorrower", mock.Anything, mock.Anything).Return(errors.New("service error"))
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.service)
			s := &SimulationHandler{
				service: tt.fields.service,
			}
			r := httptest.NewRequest(http.MethodPost, "/v1/borrower", &tt.body)
			r.Header.Add("X-User", "daiana.soares")
			r.Header.Add("X-tenant-ID", "TENANT_1")
			r.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDcmVkb3JJZCI6IlRFTkFOVF8xIiwiRXNjb3BlIjoiZXNjb3BlIiwiRXhwaXJhY2FvIjoxNzM2MzUxMTk0fQ.L-bq2UEaGxvJm8v2sR77B6RsLOZJAfkvnI4NHcuzves")

			w := httptest.NewRecorder()

			s.CreatedBorrower(w, r)

			if !reflect.DeepEqual(tt.wantHTTPStatusCode, w.Code) {
				t.Errorf("Save() = %d, want %d", w.Code, tt.wantHTTPStatusCode)

			}

			tt.fields.service.AssertExpectations(t)

		})
	}
}

func TestSimulationHandler_CreatedSetup(t *testing.T) {
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
		mock   func(service *mocks.ISimulationService)
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
		service *mocks.ISimulationService
	}

	tests := []struct {
		name               string
		fields             fields
		body               strings.Reader
		wantHTTPStatusCode int
		mock               func(service *mocks.ISimulationService)
	}{
		{name: "sucess",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			body: *strings.NewReader(`{
										 "borrowerId": "7d08c5c6-6934-4a2e-882b-a7f572f4fb96",
										 "loanValue":  10.000,
										 "numberOfInstallments": 12,
										 "interestRate": 5.25
		                                }`),
			wantHTTPStatusCode: 201,
			mock: func(service *mocks.ISimulationService) {

				service.On("CreatedSimulation", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			}},

		{
			name: "service error",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			body: *strings.NewReader(`{
						"borrowerId": "7d08c5c6-6934-4a2e-882b-a7f572f4fb96",
						"loanValue": 10000,
						"numberOfInstallments": 12,
						"interestRate": 5.25
					}`),
			wantHTTPStatusCode: 500,
			mock: func(service *mocks.ISimulationService) {
				service.On("CreatedSimulation", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("service error"))
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
		{
			name: "invalid token",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			body: *strings.NewReader(`{
					"borrowerId": "7d08c5c6-6934-4a2e-882b-a7f572f4fb96",
					"loanValue": 10000,
					"numberOfInstallments": 12,
					"interestRate": 5.25
				}`),
			wantHTTPStatusCode: 401,
			mock: func(service *mocks.ISimulationService) {
				service.On("TokenIsValid", mock.Anything).Return(nil, errors.New("invalid token"))
			},
		},
		{
			name: "validation error",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			body: *strings.NewReader(`{
					"borrowerId": 45678,
					"loanValue": "-10000",
					"numberOfInstallments": 0,
					"interestRate": -5.25
				}`),
			wantHTTPStatusCode: 422,
			mock: func(service *mocks.ISimulationService) {
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.service)
			s := &SimulationHandler{
				service: tt.fields.service,
			}
			r := httptest.NewRequest(http.MethodPost, "/v1/simulation", &tt.body)
			r.Header.Add("X-User", "daiana.soares")
			r.Header.Add("X-tenant-ID", "TENANT_1")
			r.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDcmVkb3JJZCI6IlRFTkFOVF8xIiwiRXNjb3BlIjoiZXNjb3BlIiwiRXhwaXJhY2FvIjoxNzM2MzUxMTk0fQ.L-bq2UEaGxvJm8v2sR77B6RsLOZJAfkvnI4NHcuzves")

			w := httptest.NewRecorder()

			s.CreatedSimulation(w, r)

			if !reflect.DeepEqual(tt.wantHTTPStatusCode, w.Code) {
				t.Errorf("Save() = %d, want %d", w.Code, tt.wantHTTPStatusCode)

			}

			tt.fields.service.AssertExpectations(t)

		})
	}
}

func TestSimulationHandler_FindByIdBorrower(t *testing.T) {
	creatAt := time.Now()
	type fields struct {
		service *mocks.ISimulationService
	}
	type args struct {
		id string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantHTTPStatusCode int
		mock               func(service *mocks.ISimulationService)
	}{
		{
			name: "sucess",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "7d08c5c6-6934-4a2e",
			},
			wantHTTPStatusCode: 200,

			mock: func(service *mocks.ISimulationService) {

				service.On("FindByIdBorrower", mock.Anything, mock.Anything).Return(&model.Borrower{
					BorrowerId: "7d08c5c6-6934-4a2e",
					Name:       "daiana.soares",
					Phone:      "219897688",
					Email:      "daiana@hotmail",
					Cpf:        "12345355",
					CreatedAt:  &creatAt,
					UpdateAt:   &creatAt,
				}, nil)

				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)

			},
		},
		{
			name: "borrower not found",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "nonexistent-id",
			},
			wantHTTPStatusCode: 404,
			mock: func(service *mocks.ISimulationService) {
				service.On("FindByIdBorrower", mock.Anything, mock.Anything).Return(nil, errors.NotFoundf("not found"))
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
		{
			name: "service error",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "7d08c5c6-6934-4a2e",
			},
			wantHTTPStatusCode: 500,

			mock: func(service *mocks.ISimulationService) {
				service.On("FindByIdBorrower", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("service error"))
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
		{
			name: "invalid token",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "7d08c5c6-6934-4a2e",
			},
			wantHTTPStatusCode: 401,

			mock: func(service *mocks.ISimulationService) {
				service.On("TokenIsValid", mock.Anything).Return(nil, fmt.Errorf("invalid token"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.service)
			s := &SimulationHandler{
				service: tt.fields.service,
			}
			router := mux.NewRouter()
			router.HandleFunc("/v1/borrower/{id}", s.FindByIdBorrower)
			r := httptest.NewRequest(http.MethodGet, "/v1/borrower/"+tt.args.id, nil)
			r.Header.Add("X-User", "daiana.soares")
			r.Header.Add("X-tenant-ID", "TENANT_1")
			r.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDcmVkb3JJZCI6IlRFTkFOVF8xIiwiRXNjb3BlIjoiZXNjb3BlIiwiRXhwaXJhY2FvIjoxNzM2MzUxMTk0fQ.L-bq2UEaGxvJm8v2sR77B6RsLOZJAfkvnI4NHcuzves")

			w := httptest.NewRecorder()

			router.ServeHTTP(w, r)

			//	s.FindByIdBorrower(w, r)

			if !reflect.DeepEqual(tt.wantHTTPStatusCode, w.Code) {
				t.Errorf("Save() = %d, want %d", w.Code, tt.wantHTTPStatusCode)

			}

			tt.fields.service.AssertExpectations(t)
		})
	}
}

func TestSimulationHandler_FindByIdSetup(t *testing.T) {
	creatAt := time.Now()
	type fields struct {
		service *mocks.ISimulationService
	}
	type args struct {
		id string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantHTTPStatusCode int

		mock func(service *mocks.ISimulationService)
	}{
		{
			name: "success",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "setup-12345",
			},
			wantHTTPStatusCode: 200,
			mock: func(service *mocks.ISimulationService) {
				service.On("FindByIdSetup", mock.Anything, mock.Anything).Return(&model.Setup{
					SetupId:       "setup-12345",
					Capital:       100000,
					Fees:          5.25,
					InterestRate:  12.5,
					Escope:        "default",
					EscopeIsValid: true,
					CreatedAt:     &creatAt,
					UpdatedAt:     &creatAt,
				}, nil)
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
		{
			name: "setup not found",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "nonexistent-setup",
			},
			wantHTTPStatusCode: 404,
			mock: func(service *mocks.ISimulationService) {
				service.On("FindByIdSetup", mock.Anything, mock.Anything).Return(nil, errors.NotFoundf("not found"))
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
		{
			name: "service error",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "setup-12345",
			},
			wantHTTPStatusCode: 500,
			mock: func(service *mocks.ISimulationService) {
				service.On("FindByIdSetup", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("service error"))
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
		{
			name: "invalid token",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "setup-12345",
			},
			wantHTTPStatusCode: 401,
			mock: func(service *mocks.ISimulationService) {
				service.On("TokenIsValid", mock.Anything).Return(nil, fmt.Errorf("invalid token"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.service)
			s := &SimulationHandler{
				service: tt.fields.service,
			}
			router := mux.NewRouter()
			router.HandleFunc("/v1/setup/{id}", s.FindByIdSetup)
			r := httptest.NewRequest(http.MethodGet, "/v1/setup/"+tt.args.id, nil)
			r.Header.Add("X-User", "daiana.soares")
			r.Header.Add("X-tenant-ID", "TENANT_1")
			r.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDcmVkb3JJZCI6IlRFTkFOVF8xIiwiRXNjb3BlIjoiZXNjb3BlIiwiRXhwaXJhY2FvIjoxNzM2MzUxMTk0fQ.L-bq2UEaGxvJm8v2sR77B6RsLOZJAfkvnI4NHcuzves")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)

			if w.Code != tt.wantHTTPStatusCode {
				t.Errorf("expected status %d, got %d", tt.wantHTTPStatusCode, w.Code)
			}

			tt.fields.service.AssertExpectations(t)
		})
	}
}

func TestSimulationHandler_FindByIdSimulation(t *testing.T) {
	creatAt := time.Now()
	type fields struct {
		service *mocks.ISimulationService
	}

	type args struct {
		id string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantHTTPStatusCode int

		mock func(service *mocks.ISimulationService)
	}{
		{
			name: "success",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "simulation-12345",
			},
			wantHTTPStatusCode: 200,
			mock: func(service *mocks.ISimulationService) {
				service.On("FindByIdSimulation", mock.Anything, mock.Anything).Return(&model.Simulation{
					SimulationId:         "simulation-12345",
					BorrowerId:           "borrower-6789",
					LoanValue:            50000,
					NumberOfInstallments: 24,
					CreatedAt:            creatAt,
					UpdatedAt:            creatAt,
					Status:               "approved",
					InterestRate:         5.5,
				}, nil)
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
		{
			name: "simulation not found",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "nonexistent-simulation",
			},
			wantHTTPStatusCode: 404,
			mock: func(service *mocks.ISimulationService) {
				service.On("FindByIdSimulation", mock.Anything, mock.Anything).Return(nil, errors.NotFoundf("not found"))
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
		{
			name: "service error",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "simulation-12345",
			},
			wantHTTPStatusCode: 500,
			mock: func(service *mocks.ISimulationService) {
				service.On("FindByIdSimulation", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("service error"))
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
		{
			name: "invalid token",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "simulation-12345",
			},
			wantHTTPStatusCode: 401,
			mock: func(service *mocks.ISimulationService) {
				service.On("TokenIsValid", mock.Anything).Return(nil, fmt.Errorf("invalid token"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock(tt.fields.service)
		t.Run(tt.name, func(t *testing.T) {
			s := &SimulationHandler{
				service: tt.fields.service,
			}
			router := mux.NewRouter()
			router.HandleFunc("/v1/simulation/{id}", s.FindByIdSimulation)
			r := httptest.NewRequest(http.MethodGet, "/v1/simulation/"+tt.args.id, nil)
			r.Header.Add("X-User", "daiana.soares")
			r.Header.Add("X-tenant-ID", "TENANT_1")
			r.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDcmVkb3JJZCI6IlRFTkFOVF8xIiwiRXNjb3BlIjoiZXNjb3BlIiwiRXhwaXJhY2FvIjoxNzM2MzUxMTk0fQ.L-bq2UEaGxvJm8v2sR77B6RsLOZJAfkvnI4NHcuzves")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)

			if w.Code != tt.wantHTTPStatusCode {
				t.Errorf("expected status %d, got %d", tt.wantHTTPStatusCode, w.Code)
			}

			tt.fields.service.AssertExpectations(t)

		})
	}
}

func TestSimulationHandler_UpdateSetup(t *testing.T) {
	//createdAt := time.Now()
	//updatedAt := time.Now().Add(1 * time.Hour)
	type fields struct {
		service *mocks.ISimulationService
	}
	type args struct {
		id   string
		body string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantHTTPStatusCode int

		mock func(service *mocks.ISimulationService)
	}{
		{
			name: "success",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "setup-12345",
				body: `{
							"capital": 500.050,
							"fees": 150.50,
							"interestRate": 3.5,
							"escope": "escope"
		}`,
			},
			wantHTTPStatusCode: 200,
			mock: func(service *mocks.ISimulationService) {
				service.On("UpdateSetup", mock.Anything, mock.Anything).Return(nil)
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
		{
			name: "setup not found",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "nonexistent-setup",
				body: `{
					"capital": 500.050,
					"fees": 150.50,
					"interestRate": 3.5,
					"escope": "escope"
							}`,
			},
			wantHTTPStatusCode: 404,
			mock: func(service *mocks.ISimulationService) {
				service.On("UpdateSetup", mock.Anything, mock.Anything).Return(errors.NotFoundf("not found"))
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
		{
			name: " entitie error",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id:   "nonexistent-setup",
				body: `{"capital": 50000}`,
			},
			wantHTTPStatusCode: 400,
			mock: func(service *mocks.ISimulationService) {

				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
		{
			name: "service error",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "setup-12345",
				body: `{
					"capital": 500.050,
					"fees": 150.50,
					"interestRate": 3.5,
					"escope": "escope"
							}`,
			},
			wantHTTPStatusCode: 500,

			mock: func(service *mocks.ISimulationService) {
				service.On("UpdateSetup", mock.Anything, mock.Anything).Return(errors.Internalf("error"))
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
		{
			name: "invalid token",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id:   "setup-12345",
				body: `{"capital": 100000}`,
			},
			wantHTTPStatusCode: 401,
			mock: func(service *mocks.ISimulationService) {
				service.On("TokenIsValid", mock.Anything).Return(nil, fmt.Errorf("invalid token"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.service)
			s := &SimulationHandler{
				service: tt.fields.service,
			}
			router := mux.NewRouter()
			router.HandleFunc("/v1/setup/{id}", s.UpdateSetup).Methods(http.MethodPut)
			r := httptest.NewRequest(http.MethodPut, "/v1/setup/"+tt.args.id, strings.NewReader(tt.args.body))
			r.Header.Add("X-User", "daiana.soares")
			r.Header.Add("X-tenant-ID", "TENANT_1")
			r.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDcmVkb3JJZCI6IlRFTkFOVF8xIiwiRXNjb3BlIjoiZXNjb3BlIiwiRXhwaXJhY2FvIjoxNzM2MzUxMTk0fQ.L-bq2UEaGxvJm8v2sR77B6RsLOZJAfkvnI4NHcuzves")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, r)

			if w.Code != tt.wantHTTPStatusCode {
				t.Errorf("expected status %d, got %d", tt.wantHTTPStatusCode, w.Code)
			}

			tt.fields.service.AssertExpectations(t)
		})
	}
}

func TestSimulationHandler_UpdateSimulation(t *testing.T) {
	type fields struct {
		service *mocks.ISimulationService
	}
	type args struct {
		id   string
		body string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantHTTPStatusCode int
		mock               func(service *mocks.ISimulationService)
	}{
		{
			name: "success",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "123e4567-e89b-12d3-a456-426614174000",
				body: `{
							"borrowerId": "7d08c5c6-6934-4a2e-882b-a7f572f4fb96",
							"loanValue":  10.000,
							"numberOfInstallments": 12,
							"interestRate": 5.25
}`,
			},
			wantHTTPStatusCode: http.StatusOK,
			mock: func(service *mocks.ISimulationService) {
				service.On("UpdateSimulation", mock.Anything, mock.Anything).
					Return(nil)
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
		{
			name: "invalid ID",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "invalid-id",
				body: `{
					"loanValue": 20000
				}`,
			},
			wantHTTPStatusCode: http.StatusBadRequest,
			mock: func(service *mocks.ISimulationService) {
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
		{
			name: "validation error",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "123e4567-e89b-12d3-a456-426614174000",
				body: `{
					"loanValue": "-10000"
				}`,
			},
			wantHTTPStatusCode: http.StatusUnprocessableEntity,
			mock: func(service *mocks.ISimulationService) {
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
		{
			name: "service error",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				id: "123e4567-e89b-12d3-a456-426614174000",
				body: `{
					"borrowerId": "7d08c5c6-6934-4a2e",
					"loanValue": 20000,
					"numberOfInstallments": 24,
					"interestRate": 6.5,
					"status": "approved"
				}`,
			},
			wantHTTPStatusCode: http.StatusInternalServerError,
			mock: func(service *mocks.ISimulationService) {
				service.On("UpdateSimulation", mock.Anything, mock.Anything).
					Return(errors.New("service error"))
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.service)

			s := &SimulationHandler{
				service: tt.fields.service,
			}

			router := mux.NewRouter()
			router.HandleFunc("/v1/simulation/{id}", s.UpdateSimulation)

			r := httptest.NewRequest(http.MethodPut, "/v1/simulation/"+tt.args.id, strings.NewReader(tt.args.body))
			r.Header.Add("X-User", "daiana.soares")
			r.Header.Add("X-tenant-ID", "TENANT_1")
			r.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDcmVkb3JJZCI6IlRFTkFOVF8xIiwiRXNjb3BlIjoiZXNjb3BlIiwiRXhwaXJhY2FvIjoxNzM2MzUxMTk0fQ.L-bq2UEaGxvJm8v2sR77B6RsLOZJAfkvnI4NHcuzves")

			w := httptest.NewRecorder()

			router.ServeHTTP(w, r)

			if w.Code != tt.wantHTTPStatusCode {
				t.Errorf("UpdateSimulation() got = %d, want = %d", w.Code, tt.wantHTTPStatusCode)
			}

			tt.fields.service.AssertExpectations(t)
		})
	}
}

func TestSimulationHandler_BorrowerResponseToSimulation(t *testing.T) {
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
		mock   func(service *mocks.ISimulationService)
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
		mock   func(service *mocks.ISimulationService)
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

func TestSimulationHandler_FindSimulationsByParam(t *testing.T) {
	type fields struct {
		service *mocks.ISimulationService
	}
	type args struct {
		queryParams map[string]string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantHTTPStatusCode int
		wantBody           string
		mock               func(service *mocks.ISimulationService)
	}{
		{
			name: "success with parameters",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				queryParams: map[string]string{
					"borrowerId": "7d08c5c6-6934-4a2e",
					"status":     "approved",
					"page":       "1",
					"pageSize":   "2",
				},
			},
			wantHTTPStatusCode: http.StatusOK,
			mock: func(service *mocks.ISimulationService) {
				service.On("FindByParamSimulations", mock.Anything, mock.Anything).Return(&dto.SimulationPaginationResponse{
					Simulations: []dto.SimulationResponse{
						{SimulationId: "123e4567-e89b-12d3-a456-426614174001",
							BorrowerId:           "7d08c5c6-6934-4a2e",
							LoanValue:            20000,
							NumberOfInstallments: 24,
							InterestRate:         6.5,
							Status:               "approved"},
						{SimulationId: "123e4567-e89b-12d3-",
							BorrowerId:           "7d08c5c6-6934-4a2e",
							LoanValue:            40000,
							NumberOfInstallments: 24,
							InterestRate:         6.5,
							Status:               "approved"},
					},
				}, nil)
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)

			},
		},

		{
			name: "service error",
			fields: fields{
				service: new(mocks.ISimulationService),
			},
			args: args{
				queryParams: map[string]string{
					"borrowerId": "7d08c5c6-6934-4a2e",
				},
			},
			wantHTTPStatusCode: http.StatusInternalServerError,
			
			mock: func(service *mocks.ISimulationService) {
				service.On("FindByParamSimulations", mock.Anything, mock.Anything).Return(nil, errors.New("service error"))
				service.On("TokenIsValid", mock.Anything).Return(getToken(), nil)
			},
		},
	}
	for _, tt := range tests {
		tt.mock(tt.fields.service)
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.service)

			s := &SimulationHandler{
				service: tt.fields.service,
			}

			router := mux.NewRouter()
			router.HandleFunc("/v1/simulations", func(w http.ResponseWriter, r *http.Request) {
				q := r.URL.Query()
				for key, value := range tt.args.queryParams {
					q.Add(key, value)
				}
				r.URL.RawQuery = q.Encode()
				s.FindSimulationsByParam(w, r)
			})

			r := httptest.NewRequest(http.MethodGet, "/v1/simulations", nil)
			r.Header.Add("X-user-ID", "daiana.soares")
			r.Header.Add("X-tenant-ID", "TENANT_1")
			r.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDcmVkb3JJZCI6IlRFTkFOVF8xIiwiRXNjb3BlIjoiZXNjb3BlIiwiRXhwaXJhY2FvIjoxNzM2MzUxMTk0fQ.L-bq2UEaGxvJm8v2sR77B6RsLOZJAfkvnI4NHcuzves")

			w := httptest.NewRecorder()

			router.ServeHTTP(w, r)

			if w.Code != tt.wantHTTPStatusCode {
				t.Errorf("FindSimulationsByParam() got = %d, want = %d", w.Code, tt.wantHTTPStatusCode)
			}

			if tt.wantBody != "" && !reflect.DeepEqual(strings.TrimSpace(w.Body.String()), tt.wantBody) {
				t.Errorf("FindSimulationsByParam() body = %s, want = %s", w.Body.String(), tt.wantBody)
			}

			tt.fields.service.AssertExpectations(t)
		})
	}
}
