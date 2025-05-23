package main

import (
	"github.com/whitestudios/user-auth-system/config"
	"github.com/whitestudios/user-auth-system/internal/handler"
)

func main() {
	config.Init()
	handler.Init()
}
