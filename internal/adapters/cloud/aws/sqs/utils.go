package sqs

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	sqsAws "github.com/aws/aws-sdk-go-v2/service/sqs"
)

func GenerateMessageInput(url string, timeout int32) *sqsAws.ReceiveMessageInput {
	sqsAws.EndpointResolverFromURL("http://localhost:4566")
	return &sqsAws.ReceiveMessageInput{
		QueueUrl:            aws.String(url),
		MaxNumberOfMessages: 10,
		VisibilityTimeout:   timeout,
		WaitTimeSeconds:     20,
	}
}

func GenerateDeleteInput(url string, id *string) *sqsAws.DeleteMessageInput {
	return &sqsAws.DeleteMessageInput{
		QueueUrl:      aws.String(url),
		ReceiptHandle: id,
	}
}
