package server

import (
	modules "project-golang/internal/fx/modules"


	worker "project-golang/internal/consumer"

	"go.uber.org/fx"
)







func Start() {

	module := fx.Options(
		modules.NewModuleService(),
		modules.NewModuleConsumer(),
	)
	fx.New(
		module,
		fx.Invoke(
			func(job worker.Consumer) {
				job.Do()
			},
		),
	).Run()

}
