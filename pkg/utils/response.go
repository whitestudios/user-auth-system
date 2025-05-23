package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSuccess(c *gin.Context, op string, data any) {
	c.Header("Content-type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": op,
		"data":    data,
	})
}

func SendError(c *gin.Context, op string, code int) {
	c.JSON(code, gin.H{
		"message":   op,
		"errorCode": code,
	})
}
