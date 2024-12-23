package sqs

import (
	"project-golang/internal/fx/module/cloud/aws"
	"project-golang/internal/provider/adapters/cloud/aws/sqs"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		aws.NewModule(),
		fx.Provide(
			sqs.NewClient),
	)
}
