package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func homehandler(c *gin.Context) {
	username, _ := c.Get("username")
	msg := fmt.Sprintf("Welcome to the home page %s", username)
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
	// fmt.Println(c.GetHeader("user-agent"))
}
