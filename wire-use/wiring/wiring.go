//go:build wireinject
// +build wireinject

package wiring

import (
	"github.com/google/wire"
	"wire-use/app"
)

func InitializeServices() (*app.UserService, error) {
	wire.Build(app.NewUserService, app.NewDatabase)
	return &app.UserService{}, nil
}
