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
	CreatedSimulation(ctx context.Context, simu *model.Simulation, tokenString string) error
	CreatedBorrower(tom *model.Borrower) error
	CreatedSetup(set *model.Setup) error
	FindByIdSimulation(simulationId string) (*model.Simulation, error)
	FindByIdSetup(setupId string) (*model.Setup, error)
	FindByIdBorrower(borrwerId string) (*model.Borrower, error)
	UpdateSetup(setupId string, newSetup *model.Setup) error
	UpdateSimulationStatus(simulationId string, status string) error
	SimulationResponseBorrower(response *model.SimulationResponseBorrower) error
}

type SimulationService struct {
	repository repository.IRepository
	sqsClient  sqsAws.Client
}

func NewSimulationService(repo repository.IRepository, sqs sqsAws.Client) ISimulationService {
	return &SimulationService{repository: repo,
		sqsClient: sqs}
}

// SimulationResponseBorrower implements ISimulationService.
func (s *SimulationService) SimulationResponseBorrower(response *model.SimulationResponseBorrower) error {
	// persistir no banco o status e enviar pra fila

	err := s.repository.UpdateSimulationStatus(response.SimulationId, response.Status)
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
func (s *SimulationService) CreatedBorrower(tom *model.Borrower) error {
	return s.repository.CreatedBorrower(model.ToBorrowerEntity(tom))
}

// CreatedSetup implements ISimulationService.
func (s *SimulationService) CreatedSetup(set *model.Setup) error {
	return s.repository.CreatedSetup(model.ToSetupEntity(set))
}

// CreatedSimulation implements ISimulationService.
func (s *SimulationService) CreatedSimulation(ctx context.Context, simu *model.Simulation, tokenString string) error {

	//verificar fraude
	// validar token e escopo
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

	_, err = s.validatingTokenJwtAndCheckingValidScope(tokenString, os.Getenv("EXPECTED_ESCOPE"))

	if err != nil {
		return err
	}
	setup, err := s.repository.FindByIdSetup(os.Getenv("SETUP_ID"))
	if err != nil {
		return err
	}

	juros, err := s.calculateInterest(model.ToSetupModel(setup), simu.NumberOfInstallments)
	simu.InterestRate = juros
	simu.Status = "CREATED"

	newSimu, err := s.repository.CreatedSimulation(model.ToSimulationEntity(simu))
	if err != nil {
		return err
	}
	err = s.sendMessageQueueNotification(ctx, &dto.QueuePublishPayload{
		SimulationId: newSimu.SimulationId,
	})
	if err != nil {
		return err
	}
	err = s.sendMessageQueueBorrowerAceptOrRejectedSimulation(ctx, &dto.QueuePublishPayload{
		SimulationId: newSimu.SimulationId,
	})
	if err != nil {
		return err
	}
	return nil
	// aplicar o try do pacote resiliense do GO pra caso der erro ao fazer um send na fila
}

// FindByIdBorrower implements ISimulationService.
func (s *SimulationService) FindByIdBorrower(borrwerId string) (*model.Borrower, error) {
	bo, err := s.repository.FindByIdBorrower(borrwerId)
	if err != nil {
		return nil, err
	}
	if (bo != &entity.Borrower{}) {

		return model.ToBorrowerdModel(bo), nil
	}
	return nil, errors.New("setup nao encontrado")
}

// FindByIdSetup implements ISimulationService.
func (s *SimulationService) FindByIdSetup(setupId string) (*model.Setup, error) {
	set, err := s.repository.FindByIdSetup(setupId)

	if err != nil {

	}
	if (set != &entity.Setup{}) {
		return model.ToSetupModel(set), nil
	}
	return nil, errors.New("simulation nao encontrada")
}

// FindByIdSimulation implements ISimulationService.
func (s *SimulationService) FindByIdSimulation(simulationId string) (*model.Simulation, error) {
	simula, err := s.repository.FindByIdSimulation(simulationId)
	if err != nil {
		return nil, err
	}
	if (simula != &entity.Simulation{}) {

		return model.ToSimulationModel(simula), nil
	}
	return nil, errors.New("simulation nao encontrada")

}

// UpdateSetup implements ISimulationService.
func (s *SimulationService) UpdateSetup(setupId string, newSetup *model.Setup) error {

	return s.repository.UpdateSetup(setupId, model.ToSetupEntity(newSetup))
}

// UpdateSimulationStatus implements ISimulationService.
func (s *SimulationService) UpdateSimulationStatus(simulationId string, status string) error {
	return s.repository.UpdateSimulationStatus(simulationId, status)
}

func (s *SimulationService) checkAntiFraude(request *dto.AntiFraudRequest) (*dto.AntiFraudResponse, error) {

	response, err := apiantifraude.CheckAntiFraud(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *SimulationService) validatingTokenJwtAndCheckingValidScope(tokenString string, expectedScope string) (*model.PayloadJWT, error) {

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

	// Extrair o payload do token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("falha ao extrair claims do token")
	}

	// Extrair os campos necessários do payload
	payload := &model.PayloadJWT{
		CredorID:  claims["credorId"].(string),
		Escopo:    claims["Escopo"].(string),
		Expiracao: int64(claims["Expiracao"].(float64)), // Timestamp
	}

	// Validar a expiração do token
	if time.Now().Unix() > payload.Expiracao {
		return nil, errors.New("token expirado")
	}

	// Validar o escopo
	if payload.Escopo != expectedScope {
		return nil, fmt.Errorf("escopo inválido: esperado '%s', encontrado '%s'", expectedScope, payload.Escopo)
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

func (s *SimulationService) GenerateJWT(payload model.PayloadJWT) (string, error) {
	expiracao := time.Now().Add(48 * time.Hour) //expira com 2 dias apartir da hora atual
	// Claims do token
	claims := jwt.MapClaims{
		"credorId":  payload.CredorID,
		"Escopo":    payload.Escopo,
		"Expiração": expiracao.Unix(),
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
