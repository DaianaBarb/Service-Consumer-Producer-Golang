package consumidor_service

import (
	adapters "project-golang/internal/fx/module/provider"

	service "project-golang/internal/services"

	"go.uber.org/fx"
)

func ModuleService() fx.Option {
	return fx.Options(
		adapters.NewModule(),
		fx.Provide(
			service.NewService,
		),
	)
}

func NewModule() fx.Option {
	return ModuleService()
}
