package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/whitestudios/user-auth-system/internal/auth"
)

func initializeRoutes(router *gin.Engine) {
	auth.InitializeHandlers()

	basePath := "/api"

	api := router.Group(basePath)
	{
		api.GET("/allUsers", auth.TestListAllUsers)
	}

	user := router.Group(basePath + "/user")

	{
		user.POST("/login", auth.LoginHandler)
		user.POST("/register", auth.RegisterHandler)
	}
}
