package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitestudios/user-auth-system/config"
	"github.com/whitestudios/user-auth-system/internal/service"
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

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		utils.SendError(c, "Invalid login, email or password is wrong", http.StatusBadRequest)
		return
	}

	tokenString, err := service.GenerateJwt(req.Email)

	if err != nil {
		logger.Errorf("Error generating jwt on login handler: %v", err.Error())
		utils.SendError(c, "error generating jwt", http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":           "Correct login!",
		"token":         tokenString,
		"refresh-token": "refresh-token",
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

	if err := userRepository.Create(&user); err != nil {
		logger.Errorf("error creating user in user repository with register handler: %v", err.Error())
		utils.SendError(c, "a user with this email already exists", http.StatusBadRequest)
		return
	}

	token, err := service.GenerateJwt(req.Email)

	if err != nil {
		logger.Errorf("Error generating jwt on register handler: %v", err.Error())

		if err := userRepository.Delete(&user); err != nil {
			logger.Errorf("User created in database, but jwt was not created, go to login route: %v", err.Error())
			return
		}

		utils.SendError(c, "error generating jwt", http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User create with success",
		"data":    user,
		"token":   token,
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
