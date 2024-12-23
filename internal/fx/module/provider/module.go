package adapters_provider

import (
	adapters "project-golang/internal/fx/module/provider/adapters"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return adapters.ModuleAdapters()
}
