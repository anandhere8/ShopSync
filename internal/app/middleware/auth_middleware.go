package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	// c.JSON(http.StatusUnauthorized, gin.H{
	// 	"error": "Invalid user",
	// })
	c.Abort()
}
