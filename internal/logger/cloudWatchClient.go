package logger

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

// Inicializa a configuração da AWS e o cliente do CloudWatch
func InitCloudWatch() *cloudwatchlogs.Client {

	client := &cloudwatchlogs.Client{}
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Erro ao carregar a configuração da AWS: %v", err)
	}

	client = cloudwatchlogs.NewFromConfig(cfg)

	// Criação do grupo de logs, se não existir
	_, err = client.CreateLogGroup(context.TODO(), &cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: &logGroupName,
	})
	if err != nil && !isResourceAlreadyExistsError(err) {
		fmt.Printf("Grupo de logs já existe ou houve erro: %v\n", err)
	}

	// Criação do stream de logs, se não existir
	_, err = client.CreateLogStream(context.TODO(), &cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  &logGroupName,
		LogStreamName: &logStreamName,
	})
	if err != nil && !isResourceAlreadyExistsError(err) {
		fmt.Printf("Stream de logs já existe ou houve erro: %v\n", err)
	}
	return client
}

// Verifica se o erro é de recurso já existente
func isResourceAlreadyExistsError(err error) bool {
	var opErr *types.ResourceAlreadyExistsException
	return err != nil && errors.As(err, &opErr)
}
