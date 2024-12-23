package consumidor

import (
	service "project-golang/internal/fx/server/processor/service"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return service.ModuleService()
}
