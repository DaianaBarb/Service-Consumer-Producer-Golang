package consumer

import (
	"context"
	"fmt"
	"sync"
	"time"

	sqsAws "project-golang/internal/provider/adapters/cloud/aws/sqs"

	service "project-golang/internal/services"

	"github.com/americanas-go/config"
	"github.com/americanas-go/log"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type Job interface {
	Do()
}

const (
	root         = "worker.sqs"
	urlPath      = ".url"
	timeoutPath  = ".visibilityTimeout"
	sleepTimeout = ".sleepTimeout"
)

type sqs struct {
	queueUrl          string
	visibilityTimeout int32
	sleepTimeout      time.Duration
	processor         service.Processor
	client            sqsAws.Client
}

func (w *sqs) Do() {
	logger := log.FromContext(context.TODO())
	for {
		ctx := context.Background()

		result, err := w.client.ReceiveMessage(ctx, sqsAws.GenerateMessageInput(w.queueUrl, w.visibilityTimeout))
		if err != nil {
			logger.Debug(err.Error())
			return
		}
		if len(result.Messages) == 0 {
			time.Sleep(w.sleepTimeout)
			continue
		}
		var wg sync.WaitGroup
		for _, message := range result.Messages {
			wg.Add(1)
			go func(message types.Message) {
				defer wg.Done()
				w.processMessage(message, ctx)
			}(message)
		}

		wg.Wait()

	}

}

func (w *sqs) processMessage(message types.Message, ctx context.Context) {
	logger := log.FromContext(ctx)
	//var model = ""
	// err := json.Unmarshal([]byte(*message.Body), &model)
	// if err != nil {
	// 	logger.Debug(err.Error())
	// 	return
	// }
	fmt.Printf("Body: %s\n", *&message.Body)
	// err = w.processor.Process(model)
	// if err != nil {
	// 	logger.Debug(err.Error())
	// 	return
	// }
	_, err = w.client.DeleteMessage(ctx, sqsAws.GenerateDeleteInput(w.queueUrl, message.ReceiptHandle))
	if err != nil {
		logger.Debug(err.Error())
		return
	}

}

func NewSqs(sqsClient sqsAws.Client, processor service.Processor) Job {
	return &sqs{
		queueUrl:          config.String(root + urlPath),
		visibilityTimeout: int32(config.Int(root + timeoutPath)),
		sleepTimeout:      config.Duration(root+sleepTimeout) * time.Millisecond,
		client:            sqsClient,
		processor:         processor,
	}

}
