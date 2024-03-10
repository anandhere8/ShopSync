package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/anandhere8/ShopSync/internal/app/model"
	"github.com/gin-gonic/gin"
)

const DefaultTokenExpiration = time.Hour * 1

func isValid(username, password string) bool {
	if username == "anand" && password == "123" {
		return true
	}
	return false
}

func getUserIDByUsername(string) uint {
	return uint(123)
}

func LoginHandler(c *gin.Context) {
	fmt.Println("Content-type ; ", c.ContentType())
	var loginRequst model.LoginRequest
	err := c.Bind(&loginRequst)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid format",
		})
		return
	}
	username := loginRequst.Username
	password := loginRequst.Password
	if !isValid(username, password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	userID := getUserIDByUsername(username)
	token, err := GenerateJWT(userID, username, "abc", DefaultTokenExpiration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
