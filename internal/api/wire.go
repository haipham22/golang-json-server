//go:build wireinject
// +build wireinject

package api

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

func InitApp(
	isDebugMode bool,
	db string,
	log *zap.SugaredLogger,
) (*Handler, func(), error) {
	panic(wire.Build(
		NewApiBiz,
		//postgres.NewGormDB,
		//wire.NewSet(storage.NewStorage),
	))

	return &Handler{}, nil, nil
}
