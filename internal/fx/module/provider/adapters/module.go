package adapters

import (
	"go.uber.org/fx"
)

func ModuleAdapters() fx.Option {
	return fx.Options(
		fx.Provide(
		//repository.Newrepository
		),
	)
}
