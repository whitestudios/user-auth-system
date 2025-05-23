package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/user-auth-system/internal/auth"
	"github.com/whitestudios/user-auth-system/pkg/utils"
)

func TestListAllUsers(c *gin.Context) {
	users, err := auth.UserRepository.GetAll()

	if err != nil {
		auth.Logger.Errorf("error finding all in list all test route: %v", err.Error())
		utils.SendError(c, "error finding all in list all test route", http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"test-msg": "all users",
		"data":     users,
	})
}
