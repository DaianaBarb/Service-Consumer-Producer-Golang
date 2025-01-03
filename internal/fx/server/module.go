package server

import (
	"project-golang/internal/api/rest/routes"
	modules "project-golang/internal/fx/modules"

	worker "project-golang/internal/consumer"

	"go.uber.org/fx"
)

func Start() {

	module := fx.Options(
		modules.NewModuleService(),
		modules.NewModuleConsumer(),
		modules.ModuleRouterSimulation(),
	)

	fx.New(
		module,
		fx.Provide(
			fx.Annotate(
				modules.CloseDB,
			),
		),
		fx.Invoke(
			func(job worker.Consumer, job2 routes.IRoutes) {
				job2.RegisterRoutes()
				job.Do()
			},
		),
	).Run()

}
