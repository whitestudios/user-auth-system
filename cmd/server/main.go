package main

import (
	"github.com/whitestudios/user-auth-system/config"
	"github.com/whitestudios/user-auth-system/internal/handler"
	"github.com/whitestudios/user-auth-system/internal/service"
)

func main() {
	service.InitializeJwtService("ultra-plus-max-secret-key")
	config.Init()
	handler.Init()
}
