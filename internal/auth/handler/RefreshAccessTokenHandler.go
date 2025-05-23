package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/user-auth-system/internal/service"
	"github.com/whitestudios/user-auth-system/pkg/utils"
)

func RefreshTokenHandler(c *gin.Context) {
	req := struct {
		RefreshToken string `json:"refresh_token"`
	}{}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, "Invalid json, this request must be 'refresh_token' field", http.StatusBadRequest)
		return
	}

	token, err := service.GenerateAccessTokenByRefreshToken(req.RefreshToken)

	if err != nil {
		utils.SendError(c, err.Error(), http.StatusBadRequest)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "access token generated with success ",
		"access_token": token,
	})
}
