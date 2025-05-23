package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/user-auth-system/internal/auth"
	"github.com/whitestudios/user-auth-system/internal/service"
	"github.com/whitestudios/user-auth-system/internal/user"
	"github.com/whitestudios/user-auth-system/pkg/utils"
)

func RegisterHandler(c *gin.Context) {
	// parse to UserReqDto
	req := auth.UserReqDto{}

	if err := c.Bind(&req); err != nil {
		return
	}

	if err := req.Validate(); err != nil {
		utils.SendError(c, err.Error(), http.StatusBadRequest)
	}

	hashPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		utils.SendError(c, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// insert on db
	user := user.User{
		Email:    req.Email,
		Password: hashPassword,
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

	if err := auth.UserRepository.Create(&user); err != nil {
		auth.Logger.Errorf("error creating user in user repository with register handler: %v", err.Error())
		utils.SendError(c, "a user with this email already exists", http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "User create with success",
		"data":          user,
		"token":         accessToken,
		"refresh-token": refreshToken,
	})
}
