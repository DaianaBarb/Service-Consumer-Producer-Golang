package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Client interface {
	ReceiveMessage(context context.Context, input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error)
	DeleteMessage(context context.Context, message *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error)
	SendMessage(ctx context.Context, input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error)
}

type client struct {
	awsClient *sqs.Client
}

func (c *client) ReceiveMessage(context context.Context, input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {

	return c.awsClient.ReceiveMessage(context, input)
}

func (c *client) DeleteMessage(context context.Context, message *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	return c.awsClient.DeleteMessage(context, message)
}

func NewClient(awsConfig aws.Config) Client {

	awsClient := sqs.NewFromConfig(awsConfig)

	return &client{awsClient}
}

func (c *client) SendMessage(ctx context.Context, input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {

	output, err := c.awsClient.SendMessage(ctx, input, nil)

	if err != nil {
		return output, err
	}
	return output, nil

}
