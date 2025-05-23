package auth

import (
	"github.com/whitestudios/user-auth-system/config"
	"github.com/whitestudios/user-auth-system/internal/user"
)

var UserRepository *user.UserRepository
var Logger *config.Logger

func InitializeHandlers() {
	UserRepository = user.NewUserRepository(config.GetSqlite())
	Logger = config.NewLogger("handlers debug")
}
