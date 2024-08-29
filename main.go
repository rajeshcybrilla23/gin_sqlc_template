package main

import (
	"gin-template/api_v2/router"
	"gin-template/config"
)

func main() {
	init := config.Init()
	app := router.Init(init)
	app.Run(init.EnvVaraibles.HTTPServerAddress)
}
