package api

import (
	"gin-template/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter (services service.Services)( *gin.Engine) {
	router := gin.Default()

	baseController := BaseController {
		services: services,
	}
	accountsRoute(router, baseController)

	return router
}

func accountsRoute(router *gin.Engine, baseController BaseController) {
	accountsController := AccountController{
		BaseController: baseController,
	}
	router.POST("/accounts", accountsController.createAccount)
}
