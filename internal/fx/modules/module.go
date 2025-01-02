package aws

import (
	"project-golang/internal/adapters/cloud/aws"
	"project-golang/internal/adapters/cloud/aws/sqs"
	"project-golang/internal/adapters/db/postegres"
	apiantifraude "project-golang/internal/adapters/integrations/apiAntiFraude"
	"project-golang/internal/api/rest/handlres"
	"project-golang/internal/api/rest/routes"
	"project-golang/internal/consumer"
	"project-golang/internal/logger"
	"project-golang/internal/repository"
	service "project-golang/internal/services"

	"go.uber.org/fx"
)

// injetando dependencias da API
//--------------------------------------------------

func ModuleRepository() fx.Option {
	return fx.Options(
		NewModuleAtributesRepository(),
		fx.Provide(
			repository.NewRepository),
	)
}
func NewModuleAtributesRepository() fx.Option {
	return fx.Options(
		fx.Provide(logger.InitCloudWatch),
		fx.Provide(postegres.Connect),
		fx.Provide(logger.NewCloudWatch),
		fx.Provide(apiantifraude.NewApiAntifraude),
	)
}

func ModuleServiceSimulation() fx.Option {
	return fx.Options(
		NewModuleAtributesServiceSimulation(),
		fx.Provide(
			service.NewSimulationService),
	)
}
func NewModuleAtributesServiceSimulation() fx.Option {
	return ModuleRepository()
}

func ModuleHandlerSimulation() fx.Option {
	return fx.Options(
		NewModuleAtributesHandlerSimulation(),
		fx.Provide(
			handlres.NewSimulationHandler),
	)
}
func NewModuleAtributesHandlerSimulation() fx.Option {
	return ModuleServiceSimulation()
}

func ModuleRouterSimulation() fx.Option {
	return fx.Options(
		NewModuleAtributesRouterSimulation(),
		fx.Provide(
			routes.NewRoutes),
	)
}

func NewModuleAtributesRouterSimulation() fx.Option {
	return ModuleHandlerSimulation()
}

//----------------------------------------------------

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
