package handler

import "github.com/gin-gonic/gin"

var router *gin.Engine

func Init() {
	router = gin.Default()
	initializeRoutes(router)
	router.Run()
}
