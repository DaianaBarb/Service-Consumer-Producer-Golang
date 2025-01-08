package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	sqsAws "project-golang/internal/adapters/cloud/aws/sqs"
	apiantifraude "project-golang/internal/adapters/integrations/apiAntiFraude"
	"project-golang/internal/domain/dto"
	"project-golang/internal/domain/entity"
	"project-golang/internal/domain/model"
	"project-golang/internal/repository"
	"time"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/golang-jwt/jwt/v4"
)

// SecretKey é a chave usada para assinar o token.
var SecretKey = []byte("chave_secreta_")

type ISimulationService interface {
	CreatedSimulation(ctx context.Context, simu *model.Simulation, token *jwt.Token, schema string) error
	CreatedBorrower(tom *model.Borrower, schema string) error
	CreatedSetup(set *model.Setup, schema string) error
	FindByIdSimulation(simulationId string, schema string) (*model.Simulation, error)
	FindByIdSetup(setupId string, schema string) (*model.Setup, error)
	FindByIdBorrower(borrwerId string, schema string) (*model.Borrower, error)
	UpdateSetup(newSetup *model.Setup, schema string) error
	UpdateSimulation(m *model.Simulation, schema string) error
	SimulationResponseBorrower(id string, response *model.SimulationResponseBorrower, schema string) error
	GenerateJWT(payload model.PayloadJWT) (string, error)
	TokenIsValid(tokenString string) (*jwt.Token, error)
	FindByParamSimulations(param *model.Params, schema string) (*dto.SimulationPaginationResponse, error)

	Ping() error
}

type SimulationService struct {
	repository    repository.IRepository
	sqsClient     sqsAws.Client
	apiAntifraude apiantifraude.IApiAntifraude
}

// FindByParamSimulations implements ISimulationService.
func (s *SimulationService) FindByParamSimulations(param *model.Params, schema string) (*dto.SimulationPaginationResponse, error) {
	simulations, err := s.repository.GetSimulations(param, schema)
	if err != nil {
		return nil, err
	}

	return dto.ToSimulationPaginationResponse(param.Page, param.PageSize, simulations), nil

}

func NewSimulationService(repo repository.IRepository, sqs sqsAws.Client, anti apiantifraude.IApiAntifraude) ISimulationService {
	return &SimulationService{repository: repo,
		sqsClient:     sqs,
		apiAntifraude: anti}
}

// SimulationResponseBorrower implements ISimulationService.
func (s *SimulationService) SimulationResponseBorrower(id string, response *model.SimulationResponseBorrower, schema string) error {

	newSimu := &entity.Simulation{SimulationId: response.SimulationId,
		Status: response.Status}
	// persistir no banco o status e enviar pra fila

	err := s.repository.UpdateSimulation(newSimu, schema)
	if err != nil {
		return err
	}
	err = s.sendMessageQueueNotification(context.Background(), &dto.QueuePublishPayload{SimulationId: response.SimulationId})

	if err != nil {
		return err
	}
	return nil
}

// CreatedBorrower implements ISimulationService.
func (s *SimulationService) CreatedBorrower(tom *model.Borrower, schema string) error {
	return s.repository.CreatedBorrower(model.ToBorrowerEntity(tom), schema)
}

// CreatedSetup implements ISimulationService.
func (s *SimulationService) CreatedSetup(set *model.Setup, schema string) error {
	return s.repository.CreatedSetup(model.ToSetupEntity(set), schema)
}

// CreatedSimulation implements ISimulationService.
func (s *SimulationService) CreatedSimulation(ctx context.Context, simu *model.Simulation, token *jwt.Token, schema string) error {

	//verificar fraude
	// o token ja e validado no handler
	// validated escopo
	// buscar setup
	// calcular juros
	// salvar simulação no banco
	//enviar pra fila de notificações
	//enviar pra fila pro tomador aceitar a siulação
	_, err := s.checkAntiFraude(&dto.AntiFraudRequest{BorrowerId: simu.BorrowerId,
		LoanValue: simu.LoanValue})

	if err != nil {
		return err
	}

	setup, err := s.repository.FindByIdSetup(os.Getenv("SETUP_ID"), schema)
	if err != nil {
		return err
	}
	if setup.EscopeIsValid {
		_, err := s.validateScope(token, setup.Escope)
		if err != nil {
			return err
		}
	} else {
		return errors.New("escope in setup invalid")
	}

	juros, err := s.calculateInterest(model.ToSetupModel(setup), simu.NumberOfInstallments)
	simu.InterestRate = juros
	simu.Status = "CREATED"

	newSimu, err := s.repository.CreatedSimulation(model.ToSimulationEntity(simu), schema)
	if err != nil {
		return err
	}
	err = s.sendMessageQueueNotification(ctx, &dto.QueuePublishPayload{
		SimulationId: newSimu.SimulationId,
	})
	if err != nil {
		return err
	}
	return nil
	// aplicar o try do pacote resiliense do GO pra caso der erro ao fazer um send na fila
}

// FindByIdBorrower implements ISimulationService.
func (s *SimulationService) FindByIdBorrower(borrwerId string, schema string) (*model.Borrower, error) {
	bo, err := s.repository.FindByIdBorrower(borrwerId, schema)
	if err != nil {
		return nil, err
	}
	if (bo != &entity.Borrower{}) {

		return model.ToBorrowerdModel(bo), nil
	}
	return nil, errors.New("setup nao encontrado")
}

// FindByIdSetup implements ISimulationService.
func (s *SimulationService) FindByIdSetup(setupId string, schema string) (*model.Setup, error) {
	set, err := s.repository.FindByIdSetup(setupId, schema)

	if err != nil {
		return nil,err

	}
	if (set != &entity.Setup{}) {
		return model.ToSetupModel(set), nil
	}
	return nil, errors.New("simulation nao encontrada")
}

// FindByIdSimulation implements ISimulationService.
func (s *SimulationService) FindByIdSimulation(simulationId string, schema string) (*model.Simulation, error) {
	simula, err := s.repository.FindByIdSimulation(simulationId, schema)
	if err != nil {
		return nil, err
	}
	if (simula != &entity.Simulation{}) {

		return model.ToSimulationModel(simula), nil
	}
	return nil, errors.New("simulation nao encontrada")

}

// UpdateSetup implements ISimulationService.
func (s *SimulationService) UpdateSetup(newSetup *model.Setup, schema string) error {

	escope, err := s.theScopeIsValid(newSetup.Escope)
	if err != nil {
		return err
	}
	if escope {
		return s.repository.UpdateSetup(os.Getenv("SETUP_ID"), model.ToSetupEntity(newSetup), schema)

	}
	return errors.New("escopo invalido")

}

// UpdateSimulationStatus implements ISimulationService.
func (s *SimulationService) UpdateSimulation(m *model.Simulation, schema string) error {
	simu := &entity.Simulation{SimulationId: m.SimulationId,
		Status:               m.Status,
		BorrowerId:           m.BorrowerId,
		LoanValue:            m.LoanValue,
		NumberOfInstallments: m.NumberOfInstallments,
		InterestRate:         m.InterestRate}
	return s.repository.UpdateSimulation(simu, schema)
}

func (s *SimulationService) checkAntiFraude(request *dto.AntiFraudRequest) (*dto.AntiFraudResponse, error) {

	response, err := s.apiAntifraude.CheckAntiFraud(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *SimulationService) TokenIsValid(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// alg e o algoritimo de assinatura
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inválido: %v", token.Header["alg"])
		}
		return SecretKey, nil
	})

	// Verificar se houve erro no parse
	if err != nil {
		return nil, fmt.Errorf("falha ao validar o token: %w", err)
	}

	// Garantir que o token seja válido
	if !token.Valid {
		return nil, errors.New("token inválido")
	}

	return token, nil

}

func (s *SimulationService) validateScope(token *jwt.Token, expectedScope string) (*model.PayloadJWT, error) {
	// Extrair o payload do token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("falha ao extrair claims do token")
	}

	// Extrair os campos necessários do payload
	payload := &model.PayloadJWT{
		CredorID:  claims["CredorId"].(string),
		Escope:    claims["Escope"].(string),
		Expiracao: int64(claims["Expiracao"].(float64)), // Timestamp
	}

	// Validar a expiração do token
	if time.Now().Unix() > payload.Expiracao {
		return nil, errors.New("token expirado")
	}

	// Validar o escopo
	if payload.Escope != expectedScope {
		return nil, fmt.Errorf("escopo inválido: esperado '%s', encontrado '%s'", expectedScope, payload.Escope)
	}
	return payload, nil

}

func (s *SimulationService) calculateInterest(setup *model.Setup, numberOfInstallments float64) (float64, error) {
	// Calcular os juros simples: J = C x i x t
	juros := setup.Capital * setup.InterestRate * numberOfInstallments
	return juros, nil
}

func (s *SimulationService) sendMessageQueueNotification(ctx context.Context, payload *dto.QueuePublishPayload) error {

	format, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	url := os.Getenv("QUEUE_NOTIFICATION")
	_, err = s.sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: aws.String(string(format)),
		QueueUrl:    &url,
	})

	if err != nil {
		return err
	}
	return nil

}

func (s *SimulationService) sendMessageQueueBorrowerAceptOrRejectedSimulation(ctx context.Context, payload *dto.QueuePublishPayload) error {
	format, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	url := os.Getenv("QUEUE_BORROWER")
	_, err = s.sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: aws.String(string(format)),
		QueueUrl:    &url,
	})

	if err != nil {
		return err
	}
	return nil
}

//bloquear a geração de token frequente, sempre que a criação de token for chamada verificar se o tokan ja expirou

func (s *SimulationService) GenerateJWT(payload model.PayloadJWT) (string, error) {
	expiracao := time.Now().Add(48 * time.Hour) //expira com 2 dias apartir da hora atual
	// Claims do token
	claims := jwt.MapClaims{
		"CredorId":  os.Getenv("CREDOR_ID"),
		"Escope":    payload.Escope,
		"Expiracao": expiracao.Unix(),
	}

	// Cria o token com método de assinatura HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assina o token com a chave secreta
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", fmt.Errorf("erro ao assinar o token: %w", err)
	}

	return tokenString, nil
}

func (s *SimulationService) Ping() error {

	err := s.repository.Ping()
	if err != nil {

		return err
	}
	return nil
}

func (s *SimulationService) theScopeIsValid(escope string) (bool, error) {
	return true, nil
}
