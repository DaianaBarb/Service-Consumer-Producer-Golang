package server

import (
	consumidor "project-golang/internal/fx/server/consumer"
	processor "project-golang/internal/fx/server/processor"

	worker "project-golang/internal/server/consumer"

	"github.com/americanas-go/config"
	gifx "github.com/americanas-go/ignite/go.uber.org/fx.v1"
	"go.uber.org/fx"
)

func Start() {
	config.Load()

	module := fx.Options(
		processor.NewModule(),
		consumidor.NewModule(),
	)
	gifx.NewApp(
		module,
		fx.Invoke(
			func(job worker.Job) {
				job.Do()
			},
		),
	).Run()

}
