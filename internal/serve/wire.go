//go:build wireinject
// +build wireinject

package serve

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

func InitServeApp(log *zap.SugaredLogger) (*JsonServer, error) {
	panic(wire.Build(
		NewJsonServer,
	))
}
