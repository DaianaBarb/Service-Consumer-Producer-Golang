package logger

import (
	"context"
	"encoding/json"
	"log"
	"project-golang/internal/domain/model"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

// Estrutura para o log

var (
	logGroupName  = "my-log-group"  // Nome do grupo de logs no CloudWatch
	logStreamName = "my-log-stream" // Nome do stream de logs
)

type ILogCloudWatch interface {
	SendLog(level, message string)
}

type CloudWatch struct {
	cli *cloudwatchlogs.Client
}

func NewCloudWatch(cli *cloudwatchlogs.Client) ILogCloudWatch {
	return &CloudWatch{
		cli: cli,
	}

}

// Função para enviar logs ao CloudWatch
func (c *CloudWatch) SendLog(level, message string) {
	// Formata o log como JSON
	logEntry := model.LogEntry{
		Level:     level,
		Message:   message,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	logData, err := json.Marshal(logEntry)
	if err != nil {
		log.Fatalf("Erro ao formatar log como JSON: %v", err)
	}

	// Obtém o token de sequência do stream de logs
	resp, err := c.cli.DescribeLogStreams(context.TODO(), &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName:        &logGroupName,
		LogStreamNamePrefix: &logStreamName,
	})
	if err != nil {
		log.Fatalf("Erro ao descrever streams de logs: %v", err)
	}

	var sequenceToken *string
	for _, stream := range resp.LogStreams {
		if *stream.LogStreamName == logStreamName && stream.UploadSequenceToken != nil {
			sequenceToken = stream.UploadSequenceToken
			break
		}
	}

	// Envia o log
	_, err = c.cli.PutLogEvents(context.TODO(), &cloudwatchlogs.PutLogEventsInput{
		LogGroupName:  &logGroupName,
		LogStreamName: &logStreamName,
		LogEvents: []types.InputLogEvent{
			{
				Message:   aws.String(string(logData)),
				Timestamp: aws.Int64(time.Now().UnixMilli()),
			},
		},
		SequenceToken: sequenceToken,
	})
	if err != nil {
		log.Fatalf("Erro ao enviar log para o CloudWatch: %v", err)
	}
}
