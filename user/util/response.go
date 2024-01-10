package util

import (
	"github.com/gin-gonic/gin"
)

func ResponseError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"code":    status,
		"success": false,
		"message": message,
	})
}

func ResponseJson(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{
		"code":    status,
		"success": true,
		"data":    data,
	})
}
