package api

import "github.com/gin-gonic/gin"

func SetupRouter ()( *gin.Engine) {
	router := gin.Default()

	//router.POST("/accounts", createAccount)

	return router
}
