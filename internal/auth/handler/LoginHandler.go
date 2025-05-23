package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/user-auth-system/internal/auth"
	"github.com/whitestudios/user-auth-system/internal/service"
	"github.com/whitestudios/user-auth-system/internal/user"
	"github.com/whitestudios/user-auth-system/pkg/utils"
)

func LoginHandler(c *gin.Context) {
	req := auth.UserReqDto{}

	if err := c.Bind(&req); err != nil {
		return
	}

	if err := req.Validate(); err != nil {
		utils.SendError(c, err.Error(), http.StatusBadRequest)
	}

	var user user.User

	if err := auth.UserRepository.FindByEmail(&user, req.Email); err != nil {
		utils.SendError(c, err.Error(), http.StatusNotFound)
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		utils.SendError(c, "Invalid login, email or password is wrong", http.StatusBadRequest)
		return
	}

	accessToken, err := service.GenerateAccessJwt(req.Email)

	if err != nil {
		auth.Logger.Errorf("Error generating accessToken on register handler: %v", err.Error())

		utils.SendError(c, "error generating accessToken", http.StatusInternalServerError)
		return
	}

	refreshToken, err := service.GenerateRefreshJwt(req.Email)

	if err != nil {
		auth.Logger.Errorf("Error generating Refresh on register handler: %v", err.Error())

		utils.SendError(c, "error generating Refresh", http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":           "Correct login!",
		"access-token":  accessToken,
		"refresh-token": refreshToken,
	})
}
