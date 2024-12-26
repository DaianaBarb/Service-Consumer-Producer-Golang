package aws

import (
	"project-golang/internal/adapters/cloud/aws"
    "project-golang/internal/adapters/cloud/aws/sqs"
	"project-golang/internal/consumer"
	"go.uber.org/fx"
	service "project-golang/internal/services"
)



func ModuleSqs() fx.Option {
	return fx.Options(
		NewModuleCloud(),
		fx.Provide(
			sqs.NewClient),
	)
}
func NewModuleCloud() fx.Option {
	return fx.Options(
		fx.Provide(aws.NewConfig),
	)
}

func ModuleAdapters() fx.Option {
	return fx.Options(
		fx.Provide(
		//repository.Newrepository
		),
	)
}

func NewModuleAdapters() fx.Option {
	return ModuleAdapters()
}


func ModuleService() fx.Option {
	return fx.Options(
		NewModuleAdapters(),
		fx.Provide(
			service.NewService,
		),
	)
}

func NewModuleService() fx.Option {
	return ModuleService()
}


func ModuleConsumer() fx.Option {
	return fx.Options(
		ModuleSqs(),
		fx.Provide(
			consumer.NewSqs,
		),
	)
}

func NewModuleConsumer() fx.Option {
	return ModuleConsumer()
}

