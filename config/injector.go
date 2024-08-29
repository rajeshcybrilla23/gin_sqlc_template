// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"gin-template/api_v2/controller"
	"gin-template/api_v2/repository"
	"gin-template/api_v2/service"

	"github.com/google/wire"
)

var envVariables = wire.NewSet(GetEnvVariables)

var store = wire.NewSet(ConnectToDB)

var accountRepo = wire.NewSet(repository.AccountRepositoryImpInit)

var accountSvc = wire.NewSet(service.AccountServiceImplInit)

var accountCtrl = wire.NewSet(controller.AccountControllerImplInit)

func Init() *Initialization {
	wire.Build(NewInitialization, store, envVariables, accountCtrl, accountSvc, accountRepo)
	return nil
}
