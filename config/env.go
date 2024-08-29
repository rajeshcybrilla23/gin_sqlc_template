package config

import (
	"gin-template/db/util"
	"log"
)

func GetEnvVariables() util.Config {
	config, err := util.LoadConfig(".") // . mean current folder
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	return config
}
