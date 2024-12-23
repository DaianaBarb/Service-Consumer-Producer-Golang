package consumer

import (
	sqsClient "project-golang/internal/fx/module/cloud/aws/sqs"
	"project-golang/internal/server/consumer"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		sqsClient.Module(),
		fx.Provide(
			consumer.NewSqs,
		),
	)
}


