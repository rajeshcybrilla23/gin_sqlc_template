package router

import (
	"gin-template/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	// router.Use(gin.Logger())
	// router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		user := api.Group("/account")
		//user.GET("", init.UserCtrl.GetAllUserData)
		user.POST("", init.AccCtrl.CreateAccount)
		// user.GET("/:userID", init.UserCtrl.GetUserById)
		// user.PUT("/:userID", init.UserCtrl.UpdateUserData)
		// user.DELETE("/:userID", init.UserCtrl.DeleteUser)
	}

	return router
}
