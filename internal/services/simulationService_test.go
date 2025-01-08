package service

import (
	"context"
	"errors"
	"project-golang/internal/adapters/cloud/aws/sqs/mocks"
	mockApi "project-golang/internal/adapters/integrations/apiAntiFraude/mocks"
	"project-golang/internal/domain/dto"
	"project-golang/internal/domain/entity"
	"project-golang/internal/domain/model"
	mockRepo "project-golang/internal/repository/mocks"
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/mock"
)

func TestSimulationService_FindByParamSimulations(t *testing.T) {

	status := "REJECTED"
	borro := "672c85c7-ea1e-43c8-ae7a"
	creatAt := time.Now()

	updateAt := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
	type fields struct {
		repository *mockRepo.IRepository
	}
	type args struct {
		param  *model.Params
		schema string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.SimulationPaginationResponse
		wantErr bool
		mock    func(repo *mockRepo.IRepository)
	}{
		{
			name: "success - filter by status",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				param: &model.Params{
					Simu: &model.SimulationParam{
						BorrowerId: &borro,
						Status:     &status,
					},
					Page:     1,
					PageSize: 10,
				},
				schema: "TENANT_1",
			},
			want: &dto.SimulationPaginationResponse{
				Simulations: []dto.SimulationResponse{
					{
						SimulationId:         "8c2f-1357b59725a2",
						BorrowerId:           "672c85c7-ea1e-43c8-ae7a",
						Status:               "REJECTED",
						LoanValue:            5.000,
						NumberOfInstallments: 12,
						InterestRate:         100,
						CreatedAt:            creatAt,
						UpdatedAt:            updateAt,
					},
					{
						SimulationId:         "ae7a-0d9f5537fcd3",
						BorrowerId:           "672c85c7-ea1e-43c8-ae7a",
						Status:               "REJECTED",
						LoanValue:            3.00,
						NumberOfInstallments: 10,
						InterestRate:         100,
						CreatedAt:            creatAt,
						UpdatedAt:            updateAt,
					},
				},
				Page:     "1",
				PageSize: "10",
			},
			wantErr: false,
			mock: func(repo *mockRepo.IRepository) {

				repo.On("GetSimulations", mock.Anything, mock.Anything).Return([]entity.Simulation{
					{SimulationId: "8c2f-1357b59725a2",
						BorrowerId:           "672c85c7-ea1e-43c8-ae7a",
						Status:               "REJECTED",
						LoanValue:            5.000,
						NumberOfInstallments: 12,
						InterestRate:         100,
						CreatedAt:            creatAt,
						UpdatedAt:            updateAt},
					{SimulationId: "ae7a-0d9f5537fcd3",
						BorrowerId:           "672c85c7-ea1e-43c8-ae7a",
						Status:               "REJECTED",
						LoanValue:            3.00,
						NumberOfInstallments: 10,
						InterestRate:         100,
						CreatedAt:            creatAt,
						UpdatedAt:            updateAt},
				}, nil)

			}},
		{
			name: "success - filter by created_at",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				param: &model.Params{
					Simu: &model.SimulationParam{
						CreatedAt: &creatAt,
					},
					Page:     1,
					PageSize: 10,
				},
				schema: "TENANT_1",
			},
			want: &dto.SimulationPaginationResponse{
				Simulations: []dto.SimulationResponse{
					{
						SimulationId:         "ae7a-0d9f5537fcd3",
						BorrowerId:           "672c85c7-ea1e-43c8-ae7a",
						Status:               "REJECTED",
						LoanValue:            3.00,
						NumberOfInstallments: 10,
						InterestRate:         100,
						CreatedAt:            creatAt,
						UpdatedAt:            updateAt,
					},
				},
				Page:     "1",
				PageSize: "10",
			},
			wantErr: false,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("GetSimulations", mock.Anything, mock.Anything).Return([]entity.Simulation{
					{
						SimulationId:         "ae7a-0d9f5537fcd3",
						BorrowerId:           "672c85c7-ea1e-43c8-ae7a",
						Status:               "REJECTED",
						LoanValue:            3.00,
						NumberOfInstallments: 10,
						InterestRate:         100,
						CreatedAt:            creatAt,
						UpdatedAt:            updateAt,
					},
				}, nil)
			},
		},
		{
			name: "success - no results",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				param: &model.Params{
					Simu: &model.SimulationParam{
						Status: &status,
					},
					Page:     1,
					PageSize: 10,
				},
				schema: "TENANT_1",
			},
			want: &dto.SimulationPaginationResponse{
				Simulations: []dto.SimulationResponse{},
				Page:        "1",
				PageSize:    "10",
			},
			wantErr: false,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("GetSimulations", mock.Anything, mock.Anything).Return([]entity.Simulation{}, nil)
			},
		},
		{
			name: "error - repository failure",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				param: &model.Params{
					Simu: &model.SimulationParam{
						Status: &status,
					},
					Page:     1,
					PageSize: 10,
				},
				schema: "TENANT_1",
			},
			want:    nil,
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("GetSimulations", mock.Anything, mock.Anything).Return(nil, errors.New("database error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.repository)
			s := &SimulationService{
				repository: tt.fields.repository,
			}
			got, err := s.FindByParamSimulations(tt.args.param, tt.args.schema)
			if (err != nil) != tt.wantErr {
				t.Errorf("SimulationService.FindByParamSimulations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimulationService.FindByParamSimulations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimulationService_SimulationResponseBorrower(t *testing.T) {
	type fields struct {
		repository *mockRepo.IRepository
		sqsClient  *mocks.Client
	}
	type args struct {
		id       string
		response *model.SimulationResponseBorrower
		schema   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(repo *mockRepo.IRepository, sqsClient *mocks.Client)
	}{
		{name: "sucess valid response",
			fields: fields{repository: new(mockRepo.IRepository),
				sqsClient: new(mocks.Client)},
			args: args{
				id: "ae7a-0d9f5537fcd3",
				response: &model.SimulationResponseBorrower{
					SimulationId: "ae7a-0d9f5537fcd3",
					Status:       "ACCEPTED",
				},
				schema: "TENANT_1",
			},
			wantErr: false,
			mock: func(repo *mockRepo.IRepository, sqsClient *mocks.Client) {
				repo.On("UpdateSimulation", mock.Anything, mock.Anything).Return(nil)
				sqsClient.On("SendMessage", mock.Anything, mock.Anything).Return(&sqs.SendMessageOutput{}, nil)

			}},
		{
			name: "error - invalid ID",
			fields: fields{
				repository: new(mockRepo.IRepository),
				sqsClient:  new(mocks.Client),
			},
			args: args{
				id: "",
				response: &model.SimulationResponseBorrower{
					SimulationId: "ae7a-0d9f5537fcd3",
					Status:       "ACCEPTED",
				},
				schema: "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository, sqsClient *mocks.Client) {
				repo.On("UpdateSimulation", mock.Anything, mock.Anything).Return(errors.New("invalid ID"))
			},
		},
		// Erro - Resposta nula
		{
			name: "error - nil response",
			fields: fields{
				repository: new(mockRepo.IRepository),
				sqsClient:  new(mocks.Client),
			},
			args: args{
				id:       "ae7a-0d9f5537fcd3",
				response: &model.SimulationResponseBorrower{},
				schema:   "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository, sqsClient *mocks.Client) {
				repo.On("UpdateSimulation", mock.Anything, mock.Anything).Return(errors.New("response is nil"))
			},
		},
		// Erro - Schema inválido
		{
			name: "error - invalid schema",
			fields: fields{
				repository: new(mockRepo.IRepository),
				sqsClient:  new(mocks.Client),
			},
			args: args{
				id: "ae7a-0d9f5537fcd3",
				response: &model.SimulationResponseBorrower{
					SimulationId: "ae7a-0d9f5537fcd3",
					Status:       "ACCEPTED",
				},
				schema: "",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository, sqsClient *mocks.Client) {
				repo.On("UpdateSimulation", mock.Anything, mock.Anything).Return(errors.New("invalid schema"))
			},
		},
		// Erro - Status inválido
		{
			name: "error - invalid status",
			fields: fields{
				repository: new(mockRepo.IRepository),
				sqsClient:  new(mocks.Client),
			},
			args: args{
				id: "ae7a-0d9f5537fcd3",
				response: &model.SimulationResponseBorrower{
					SimulationId: "ae7a-0d9f5537fcd3",
					Status:       "INVALID_STATUS",
				},
				schema: "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository, sqsClient *mocks.Client) {
				repo.On("UpdateSimulation", mock.Anything, mock.Anything).Return(errors.New("invalid status"))
			},
		},
		// Erro - Falha ao atualizar simulação
		{
			name: "error - failed to update simulation",
			fields: fields{
				repository: new(mockRepo.IRepository),
				sqsClient:  new(mocks.Client),
			},
			args: args{
				id: "ae7a-0d9f5537fcd3",
				response: &model.SimulationResponseBorrower{
					SimulationId: "ae7a-0d9f5537fcd3",
					Status:       "ACCEPTED",
				},
				schema: "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository, sqsClient *mocks.Client) {
				repo.On("UpdateSimulation", mock.Anything, mock.Anything).Return(errors.New("failed to update"))
			},
		},
		// Erro - Falha no envio da mensagem SQS
		{
			name: "error - failed to send SQS message",
			fields: fields{
				repository: new(mockRepo.IRepository),
				sqsClient:  new(mocks.Client),
			},
			args: args{
				id: "ae7a-0d9f5537fcd3",
				response: &model.SimulationResponseBorrower{
					SimulationId: "ae7a-0d9f5537fcd3",
					Status:       "ACCEPTED",
				},
				schema: "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository, sqsClient *mocks.Client) {
				repo.On("UpdateSimulation", mock.Anything, mock.Anything).Return(nil)
				sqsClient.On("SendMessage", mock.Anything, mock.Anything).Return(nil, errors.New("failed to send message"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.repository, tt.fields.sqsClient)
			s := &SimulationService{
				repository: tt.fields.repository,
				sqsClient:  tt.fields.sqsClient,
			}
			if err := s.SimulationResponseBorrower(tt.args.id, tt.args.response, tt.args.schema); (err != nil) != tt.wantErr {
				t.Errorf("SimulationService.SimulationResponseBorrower() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSimulationService_CreatedBorrower(t *testing.T) {
	creatAt := time.Now()
	type fields struct {
		repository *mockRepo.IRepository
	}
	type args struct {
		tom    *model.Borrower
		schema string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(repo *mockRepo.IRepository)
	}{
		{
			name:   "sucess valid borrower",
			fields: fields{repository: new(mockRepo.IRepository)},
			args: args{
				tom: &model.Borrower{
					Name:      "daiana",
					Phone:     "219908767",
					Email:     "daiana.soares@",
					Cpf:       "12356789",
					CreatedAt: &creatAt,
				},
				schema: "TENANT_1",
			},
			wantErr: false,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("CreatedBorrower", mock.Anything, mock.Anything).Return(nil)

			},
		},
		{
			name:   "error - empty schema",
			fields: fields{repository: new(mockRepo.IRepository)},
			args: args{
				tom: &model.Borrower{
					Name:      "joao",
					Phone:     "219123456",
					Email:     "joao.silva@domain.com",
					Cpf:       "98765432100",
					CreatedAt: &creatAt,
				},
				schema: "",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("CreatedBorrower", mock.Anything, mock.Anything).Return(errors.New("schema is empty"))
			},
		},
		// Erro - Dados inválidos (email mal formatado)
		{
			name:   "error - invalid email",
			fields: fields{repository: new(mockRepo.IRepository)},
			args: args{
				tom: &model.Borrower{
					Name:      "carlos",
					Phone:     "219555555",
					Email:     "invalid-email",
					Cpf:       "11122233344",
					CreatedAt: &creatAt,
				},
				schema: "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("CreatedBorrower", mock.Anything, mock.Anything).Return(errors.New("invalid email format"))
			},
		},
		// Erro - CPF inválido
		{
			name:   "error - invalid CPF",
			fields: fields{repository: new(mockRepo.IRepository)},
			args: args{
				tom: &model.Borrower{
					Name:      "ana",
					Phone:     "219333444",
					Email:     "ana.maria@domain.com",
					Cpf:       "invalid-cpf",
					CreatedAt: &creatAt,
				},
				schema: "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("CreatedBorrower", mock.Anything, mock.Anything).Return(errors.New("invalid CPF"))
			},
		},
		// Erro - Falha no repositório
		{
			name:   "error - repository failure",
			fields: fields{repository: new(mockRepo.IRepository)},
			args: args{
				tom: &model.Borrower{
					Name:      "marcos",
					Phone:     "219876543",
					Email:     "marcos.rocha@domain.com",
					Cpf:       "12345678901",
					CreatedAt: &creatAt,
				},
				schema: "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("CreatedBorrower", mock.Anything, mock.Anything).Return(errors.New("database error"))
			},
		},
		// Erro - Borrower nulo
		{
			name:   "error - nil borrower",
			fields: fields{repository: new(mockRepo.IRepository)},
			args: args{
				tom:    &model.Borrower{},
				schema: "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("CreatedBorrower", mock.Anything, mock.Anything).Return(errors.New("borrower is nil"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.repository)
			s := &SimulationService{
				repository: tt.fields.repository,
			}
			if err := s.CreatedBorrower(tt.args.tom, tt.args.schema); (err != nil) != tt.wantErr {
				t.Errorf("SimulationService.CreatedBorrower() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSimulationService_CreatedSetup(t *testing.T) {
	creatAt := time.Now()

	type fields struct {
		repository *mockRepo.IRepository
	}
	type args struct {
		set    *model.Setup
		schema string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(repo *mockRepo.IRepository)
	}{
		{
			name: "success - valid setup",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				set: &model.Setup{
					Capital:       50.000,
					Fees:          100.00,
					InterestRate:  500.00,
					Escope:        "escope",
					EscopeIsValid: true,
					CreatedAt:     &creatAt,
				},
				schema: "TENANT_1",
			},
			wantErr: false,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("CreatedSetup", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "error - invalid schema",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				set: &model.Setup{
					Capital:       50.000,
					Fees:          100.00,
					InterestRate:  500.00,
					Escope:        "escope",
					EscopeIsValid: true,
					CreatedAt:     &creatAt,
				},
				schema: "",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("CreatedSetup", mock.Anything, mock.Anything).Return(errors.New("invalid schema"))
			},
		},
		{
			name: "error - invalid setup",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				set: &model.Setup{
					Capital:       0, // Capital inválido
					Fees:          -10.00,
					InterestRate:  -5.00,
					Escope:        "",
					EscopeIsValid: false,
					CreatedAt:     &creatAt,
				},
				schema: "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("CreatedSetup", mock.Anything, mock.Anything).Return(errors.New("invalid setup data"))
			},
		},
		{
			name: "error - repository failure",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				set: &model.Setup{
					Capital:       50.000,
					Fees:          100.00,
					InterestRate:  500.00,
					Escope:        "escope",
					EscopeIsValid: true,
					CreatedAt:     &creatAt,
				},
				schema: "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("CreatedSetup", mock.Anything, mock.Anything).Return(errors.New("database error"))
			},
		},
		{
			name: "error - nil setup",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				set:    &model.Setup{},
				schema: "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("CreatedSetup", mock.Anything, mock.Anything).Return(errors.New("nil setup"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.repository)
			s := &SimulationService{
				repository: tt.fields.repository,
			}
			if err := s.CreatedSetup(tt.args.set, tt.args.schema); (err != nil) != tt.wantErr {
				t.Errorf("SimulationService.CreatedSetup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSimulationService_CreatedSimulation(t *testing.T) {
	expiracao := time.Now().Add(48 * time.Hour) //expira com 2 dias apartir da hora atual
	// Criando as claims (informações do token)
	claims := jwt.MapClaims{
		"CredorId":  "TENANT_1",
		"Escope":    "escope",
		"Expiracao": float64(expiracao.Unix()),
	}

	// Cria o token com método de assinatura HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	creatAt := time.Now()
	type fields struct {
		repository    *mockRepo.IRepository
		sqsClient     *mocks.Client
		apiAntifraude *mockApi.IApiAntifraude
	}
	type args struct {
		ctx    context.Context
		simu   *model.Simulation
		token  *jwt.Token
		schema string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(repo *mockRepo.IRepository, sqsClient *mocks.Client, apiAntifraude *mockApi.IApiAntifraude)
	}{
		{
			name: "sucess valid simulation",
			fields: fields{
				repository:    new(mockRepo.IRepository),
				sqsClient:     new(mocks.Client),
				apiAntifraude: new(mockApi.IApiAntifraude),
			},
			args: args{
				ctx: context.Background(),
				simu: &model.Simulation{
					BorrowerId:           "672c85c7-ea1e-",
					LoanValue:            50.000,
					InterestRate:         500.00,
					NumberOfInstallments: 12,
				},
				token:  token,
				schema: "TENANT_1",
			},
			wantErr: false,
			mock: func(repo *mockRepo.IRepository, sqsClient *mocks.Client, apiAntifraude *mockApi.IApiAntifraude) {

				repo.On("FindByIdSetup", mock.Anything, mock.Anything).Return(&entity.Setup{
					SetupId:       "TENANT_1",
					Capital:       50.000,
					Fees:          10.000,
					InterestRate:  500.00,
					Escope:        "escope",
					EscopeIsValid: true,
					CreatedAt:     &creatAt,
				}, nil)
				apiAntifraude.On("CheckAntiFraud", mock.Anything).Return(&dto.AntiFraudResponse{
					BorrowerId: "672c85c7-ea1e-",
				}, nil)
				repo.On("CreatedSimulation", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&entity.Simulation{SimulationId: "672c85c7-ea1e-43c8",
					BorrowerId:           "672c85c7-ea1e-",
					LoanValue:            50.00,
					InterestRate:         500.00,
					NumberOfInstallments: 12,
					Status:               "CREATED",
					CreatedAt:            creatAt}, nil)
				sqsClient.On("SendMessage", mock.Anything, mock.Anything).Return(&sqs.SendMessageOutput{}, nil)

			},
		},
		{
			name: "error invalid setup",
			fields: fields{
				repository:    new(mockRepo.IRepository),
				sqsClient:     new(mocks.Client),
				apiAntifraude: new(mockApi.IApiAntifraude),
			},
			args: args{
				ctx: context.Background(),
				simu: &model.Simulation{
					BorrowerId:           "672c85c7-ea1e-",
					LoanValue:            50.000,
					InterestRate:         500.00,
					NumberOfInstallments: 12,
				},
				token:  token,
				schema: "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository, sqsClient *mocks.Client, apiAntifraude *mockApi.IApiAntifraude) {

				repo.On("FindByIdSetup", mock.Anything, mock.Anything).Return(nil, errors.New("setup not found"))
				apiAntifraude.On("CheckAntiFraud", mock.Anything).Return(&dto.AntiFraudResponse{
					BorrowerId: "672c85c7-ea1e-",
				}, nil)
			},
		},
		{
			name: "error antifraud validation",
			fields: fields{
				repository:    new(mockRepo.IRepository),
				sqsClient:     new(mocks.Client),
				apiAntifraude: new(mockApi.IApiAntifraude),
			},
			args: args{
				ctx: context.Background(),
				simu: &model.Simulation{
					BorrowerId:           "672c85c7-ea1e-",
					LoanValue:            50.000,
					InterestRate:         500.00,
					NumberOfInstallments: 12,
				},
				token:  token,
				schema: "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository, sqsClient *mocks.Client, apiAntifraude *mockApi.IApiAntifraude) {

				repo.On("FindByIdSetup", mock.Anything, mock.Anything).Return(&entity.Setup{
					SetupId:       "TENANT_1",
					Capital:       50.000,
					Fees:          10.000,
					InterestRate:  500.00,
					Escope:        "escope",
					EscopeIsValid: true,
					CreatedAt:     &creatAt,
				}, nil)

				apiAntifraude.On("CheckAntiFraud", mock.Anything).Return(nil, errors.New("antifraud failed"))
			},
		},
		{
			name: "error creating simulation",
			fields: fields{
				repository:    new(mockRepo.IRepository),
				sqsClient:     new(mocks.Client),
				apiAntifraude: new(mockApi.IApiAntifraude),
			},
			args: args{
				ctx: context.Background(),
				simu: &model.Simulation{
					BorrowerId:           "672c85c7-ea1e-",
					LoanValue:            50.000,
					InterestRate:         500.00,
					NumberOfInstallments: 12,
				},
				token:  token,
				schema: "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository, sqsClient *mocks.Client, apiAntifraude *mockApi.IApiAntifraude) {

				repo.On("FindByIdSetup", mock.Anything, mock.Anything).Return(&entity.Setup{
					SetupId:       "TENANT_1",
					Capital:       50.000,
					Fees:          10.000,
					InterestRate:  500.00,
					Escope:        "escope",
					EscopeIsValid: true,
					CreatedAt:     &creatAt,
				}, nil)

				apiAntifraude.On("CheckAntiFraud", mock.Anything).Return(&dto.AntiFraudResponse{
					BorrowerId: "672c85c7-ea1e-",
				}, nil)

				repo.On("CreatedSimulation", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("database error"))
			},
		},
		{
			name: "error sending message to SQS",
			fields: fields{
				repository:    new(mockRepo.IRepository),
				sqsClient:     new(mocks.Client),
				apiAntifraude: new(mockApi.IApiAntifraude),
			},
			args: args{
				ctx: context.Background(),
				simu: &model.Simulation{
					BorrowerId:           "672c85c7-ea1e-",
					LoanValue:            50.000,
					InterestRate:         500.00,
					NumberOfInstallments: 12,
				},
				token:  token,
				schema: "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository, sqsClient *mocks.Client, apiAntifraude *mockApi.IApiAntifraude) {

				repo.On("FindByIdSetup", mock.Anything, mock.Anything).Return(&entity.Setup{
					SetupId:       "TENANT_1",
					Capital:       50.000,
					Fees:          10.000,
					InterestRate:  500.00,
					Escope:        "escope",
					EscopeIsValid: true,
					CreatedAt:     &creatAt,
				}, nil)

				apiAntifraude.On("CheckAntiFraud", mock.Anything).Return(&dto.AntiFraudResponse{
					BorrowerId: "672c85c7-ea1e-",
				}, nil)

				repo.On("CreatedSimulation", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&entity.Simulation{
					SimulationId:         "672c85c7-ea1e-43c8",
					BorrowerId:           "672c85c7-ea1e-",
					LoanValue:            50.00,
					InterestRate:         500.00,
					NumberOfInstallments: 12,
					Status:               "CREATED",
					CreatedAt:            creatAt,
				}, nil)

				sqsClient.On("SendMessage", mock.Anything, mock.Anything).Return(nil, errors.New("sqs error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.repository, tt.fields.sqsClient, tt.fields.apiAntifraude)
			s := &SimulationService{
				repository:    tt.fields.repository,
				sqsClient:     tt.fields.sqsClient,
				apiAntifraude: tt.fields.apiAntifraude,
			}
			if err := s.CreatedSimulation(tt.args.ctx, tt.args.simu, tt.args.token, tt.args.schema); (err != nil) != tt.wantErr {
				t.Errorf("SimulationService.CreatedSimulation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSimulationService_FindByIdBorrower(t *testing.T) {
	creatAt := time.Now()
	type fields struct {
		repository *mockRepo.IRepository
	}
	type args struct {
		borrwerId string
		schema    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Borrower
		wantErr bool
		mock    func(repo *mockRepo.IRepository)
	}{
		{
			name: "success - valid borrower",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				borrwerId: "123456",
				schema:    "TENANT_1",
			},
			want: &model.Borrower{
				BorrowerId: "123456",
				Name:       "John Doe",
				Phone:      "123456789",
				Email:      "johndoe@example.com",
				Cpf:        "12345678901",
				CreatedAt:  &creatAt,
			},
			wantErr: false,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("FindByIdBorrower", mock.Anything, mock.Anything).Return(&entity.Borrower{
					BorrowerId: "123456",
					Name:       "John Doe",
					Phone:      "123456789",
					Email:      "johndoe@example.com",
					Cpf:        "12345678901",
					CreatedAt:  &creatAt,
				}, nil)
			},
		},
		{
			name: "error - borrower not found",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				borrwerId: "999999",
				schema:    "TENANT_1",
			},
			want:    nil,
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("FindByIdBorrower", mock.Anything, mock.Anything).Return(nil, errors.New("borrower not found"))
			},
		},
		{
			name: "error - invalid schema",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				borrwerId: "123456",
				schema:    "",
			},
			want:    nil,
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("FindByIdBorrower", mock.Anything, mock.Anything).Return(nil, errors.New("invalid schema"))
			},
		},
		{
			name: "error - repository failure",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				borrwerId: "123456",
				schema:    "TENANT_1",
			},
			want:    nil,
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("FindByIdBorrower", mock.Anything, mock.Anything).Return(nil, errors.New("database error"))
			},
		},
		{
			name: "error - empty borrower ID",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				borrwerId: "",
				schema:    "TENANT_1",
			},
			want:    nil,
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("FindByIdBorrower", mock.Anything, mock.Anything).Return(nil, errors.New("borrower ID cannot be empty"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.repository)
			s := &SimulationService{
				repository: tt.fields.repository,
			}
			got, err := s.FindByIdBorrower(tt.args.borrwerId, tt.args.schema)
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
	creatAt := time.Now()
	type fields struct {
		repository *mockRepo.IRepository
	}
	type args struct {
		setupId string
		schema  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Setup
		wantErr bool
		mock    func(repo *mockRepo.IRepository)
	}{
		{name: "sucess find valid",
			fields: fields{repository: new(mockRepo.IRepository)},
			args: args{
				setupId: "TENANT_1",
				schema:  "TENANT_1",
			},
			want: &model.Setup{Capital: 50.000,
				SetupId:       "TENANT_1",
				Fees:          100.000,
				InterestRate:  10.000,
				Escope:        "escope",
				EscopeIsValid: true,
				CreatedAt:     &creatAt},
			mock: func(repo *mockRepo.IRepository) {
				repo.On("FindByIdSetup", mock.Anything, mock.Anything).Return(&entity.Setup{Capital: 50.000,
					SetupId:       "TENANT_1",
					Fees:          100.000,
					InterestRate:  10.000,
					Escope:        "escope",
					EscopeIsValid: true,
					CreatedAt:     &creatAt}, nil)
			},
			wantErr: false},
		{
			name:   "error setup not found",
			fields: fields{repository: new(mockRepo.IRepository)},
			args: args{
				setupId: "TENANT_2",
				schema:  "TENANT_2",
			},
			want:    nil,
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("FindByIdSetup", mock.Anything, mock.Anything).Return(nil, errors.New("setup not found"))
			},
		},
		{
			name:   "error invalid schema",
			fields: fields{repository: new(mockRepo.IRepository)},
			args: args{
				setupId: "TENANT_1",
				schema:  "",
			},
			want:    nil,
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("FindByIdSetup", mock.Anything, mock.Anything).Return(nil, errors.New("invalid schema"))
			},
		},
		{
			name:   "error database connection",
			fields: fields{repository: new(mockRepo.IRepository)},
			args: args{
				setupId: "TENANT_1",
				schema:  "TENANT_1",
			},
			want:    nil,
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("FindByIdSetup", mock.Anything, mock.Anything).Return(nil, errors.New("database connection error"))
			},
		},

		{
			name:   "error with empty setupId",
			fields: fields{repository: new(mockRepo.IRepository)},
			args: args{
				setupId: "",
				schema:  "TENANT_1",
			},
			want:    nil,
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("FindByIdSetup", mock.Anything, mock.Anything).Return(nil, errors.New("empty setupId"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.repository)
			s := &SimulationService{
				repository: tt.fields.repository,
			}
			got, err := s.FindByIdSetup(tt.args.setupId, tt.args.schema)
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

func TestSimulationService_UpdateSetup(t *testing.T) {
	creatAt := time.Now()
	type fields struct {
		repository *mockRepo.IRepository
	}
	type args struct {
		newSetup *model.Setup
		schema   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(repo *mockRepo.IRepository)
	}{
		{
			name: "success - valid setup update",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				newSetup: &model.Setup{
					SetupId:       "TENANT_1",
					Capital:       100000.00,
					Fees:          2000.00,
					InterestRate:  15.00,
					Escope:        "new_escope",
					EscopeIsValid: true,
					CreatedAt:     &creatAt,
				},
				schema: "TENANT_1",
			},
			wantErr: false,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("UpdateSetup", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "error - invalid schema",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				newSetup: &model.Setup{
					SetupId:       "TENANT_1",
					Capital:       100000.00,
					Fees:          2000.00,
					InterestRate:  15.00,
					Escope:        "new_escope",
					EscopeIsValid: true,
					CreatedAt:     &creatAt,
				},
				schema: "",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("UpdateSetup", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("invalid schema"))
			},
		},
		{
			name: "error - repository failure",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				newSetup: &model.Setup{
					SetupId:       "TENANT_1",
					Capital:       100000.00,
					Fees:          2000.00,
					InterestRate:  15.00,
					Escope:        "new_escope",
					EscopeIsValid: true,
					CreatedAt:     &creatAt,
				},
				schema: "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("UpdateSetup", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("database error"))
			},
		},
		{
			name: "error - invalid setup data",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				newSetup: &model.Setup{
					SetupId:       "",
					Capital:       0,
					Fees:          0,
					InterestRate:  0,
					Escope:        "",
					EscopeIsValid: false,
					CreatedAt:     nil,
				},
				schema: "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("UpdateSetup", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("invalid setup data"))
			},
		},
		{
			name: "error - setup not found",
			fields: fields{
				repository: new(mockRepo.IRepository),
			},
			args: args{
				newSetup: &model.Setup{
					SetupId:       "INVALID_ID",
					Capital:       100000.00,
					Fees:          2000.00,
					InterestRate:  15.00,
					Escope:        "escope",
					EscopeIsValid: true,
					CreatedAt:     &creatAt,
				},
				schema: "TENANT_1",
			},
			wantErr: true,
			mock: func(repo *mockRepo.IRepository) {
				repo.On("UpdateSetup", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("setup not found"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.repository)
			s := &SimulationService{
				repository: tt.fields.repository,
			}
			if err := s.UpdateSetup(tt.args.newSetup, tt.args.schema); (err != nil) != tt.wantErr {
				t.Errorf("SimulationService.UpdateSetup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
