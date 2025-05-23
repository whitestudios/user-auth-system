package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/user-auth-system/internal/auth"
	"github.com/whitestudios/user-auth-system/internal/middleware"
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

	protected := router.Group(basePath + "/protected")

	{
		protected.Use(middleware.ProtectedHandler())
		protected.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "welcome to this protected area!",
			})
		})
	}
}
