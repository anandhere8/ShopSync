package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func homehandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the home page",
	})
	fmt.Println(c.GetHeader("user-agent"))
}
