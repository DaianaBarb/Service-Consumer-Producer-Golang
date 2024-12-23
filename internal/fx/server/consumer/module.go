package consumer

import (
	worker "project-golang/internal/fx/server/consumer/worker"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return worker.Module()
}
