package config

import (
	"gin-template/api_v2/controller"
	"gin-template/db/util"
)

type Initialization struct {
	EnvVaraibles util.Config
	AccCtrl      controller.AccountControllerImpl
}

func NewInitialization(envVaraibles util.Config,
	accCtrl controller.AccountControllerImpl) *Initialization {
	return &Initialization{
		EnvVaraibles: envVaraibles,
		AccCtrl:      accCtrl,
	}
}
