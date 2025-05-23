package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/user-auth-system/config"
	"github.com/whitestudios/user-auth-system/internal/user"
	"github.com/whitestudios/user-auth-system/pkg/utils"
)

var userRepository *user.UserRepository
var logger *config.Logger

func InitializeHandlers() {
	userRepository = user.NewUserRepository(config.GetSqlite())
	logger = config.NewLogger("handlers debug")
}

func LoginHandler(c *gin.Context) {
	req := UserReqDto{}
	if err := c.Bind(&req); err != nil {
		return
	}

	// validations
	if req.Email == "" && req.Password == "" {
		utils.SendError(c, "Invalid json, this request must be an email and password fields", http.StatusBadRequest)
		return
	}

	if !req.EmailIsValid() {
		utils.SendError(c, "Invalid email", http.StatusBadRequest)
		return
	}

	if !req.PasswordIsValid() {
		utils.SendError(c, "Invalid password", http.StatusBadRequest)
		return
	}

	var user user.User

	if err := userRepository.FindByEmail(&user, req.Email); err != nil {
		utils.SendError(c, err.Error(), http.StatusNotFound)
		return
	}

	if user.Password != req.Password {
		utils.SendError(c, "Invalid login, email or password is wrong", http.StatusBadRequest)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Correct login!",
	})
}

func RegisterHandler(c *gin.Context) {
	// parse to UserReqDto
	req := UserReqDto{}

	if err := c.Bind(&req); err != nil {
		return
	}

	// validations
	if req.Email == "" && req.Password == "" {
		utils.SendError(c, "Invalid json, this request must be an email and password fields", http.StatusBadRequest)
		return
	}

	if !req.EmailIsValid() {
		utils.SendError(c, "Invalid email", http.StatusBadRequest)
		return
	}

	if !req.PasswordIsValid() {
		utils.SendError(c, "Invalid password", http.StatusBadRequest)
		return
	}

	// insert on db
	user := user.User{
		Email:    req.Email,
		Password: req.Password,
	}

	if err := userRepository.Create(&user); err != nil {
		utils.SendError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User create with success",
		"data":    user,
	})
}
func TestListAllUsers(c *gin.Context) {
	users, err := userRepository.GetAll()

	if err != nil {
		logger.Errorf("error finding all in list all test route: %v", err.Error())
		utils.SendError(c, "error finding all in list all test route", http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"test-msg": "all users",
		"data":     users,
	})
}
