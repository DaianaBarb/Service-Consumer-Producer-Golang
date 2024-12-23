package aws

import (
	"project-golang/internal/provider/adapters/cloud/aws"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(aws.NewConfig),
	)
}
