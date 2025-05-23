package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/user-auth-system/internal/auth"
	"github.com/whitestudios/user-auth-system/internal/auth/handler"
	"github.com/whitestudios/user-auth-system/internal/middleware"
)

func initializeRoutes(router *gin.Engine) {
	auth.InitializeHandlers()

	basePath := "/api"

	api := router.Group(basePath)
	{
		api.GET("/allUsers", handler.TestListAllUsers)
	}
	// public routes
	user := router.Group(basePath + "/public")

	{
		user.POST("/refresh", handler.RefreshTokenHandler)
		user.POST("/login", handler.LoginHandler)
		user.POST("/register", handler.RegisterHandler)
	}

	protected := router.Group(basePath + "/protected")

	{
		// protected, just can access with a valid access token
		protected.Use(middleware.ProtectedHandler())

		protected.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "welcome to this protected area!",
			})
		})
	}
}
