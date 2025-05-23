package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/user-auth-system/internal/service"
	"github.com/whitestudios/user-auth-system/pkg/utils"
)

func ProtectedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			utils.AbortError(c, "missing authorization header", http.StatusUnauthorized)
			return
		}

		if len(tokenString) < 15 {
			utils.AbortError(c, "invalid token", http.StatusUnauthorized)
			return
		}

		tokenString = tokenString[len("Bearer "):]

		err := service.VerifyToken(tokenString)

		if err != nil {
			utils.AbortError(c, "invalid token", http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
