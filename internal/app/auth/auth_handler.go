package auth

import (
	"fmt"
	"net/http"

	"github.com/anandhere8/ShopSync/internal/app/model"
	"github.com/anandhere8/ShopSync/internal/app/repository"
	"github.com/anandhere8/ShopSync/internal/app/service"
	"github.com/gin-gonic/gin"
)

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
	if !service.ValidateCredential(username, password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}
	userID, _ := repository.GetUserID(username)
	token, err := GenerateJWT(userID, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
