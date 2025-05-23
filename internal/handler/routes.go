package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/user-auth-system/internal/auth"
)

func initializeRoutes(router *gin.Engine) {
	basePath := "/api"

	api := router.Group(basePath)
	{
		api.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "Hello world"})
		})
	}

	user := router.Group(basePath + "/user")

	{
		user.POST("/login", auth.LoginHandler)
		user.POST("/register", auth.RegisterHandler)
	}
}
