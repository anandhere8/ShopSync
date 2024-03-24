package handler

import (
	"net/http"
	"path/filepath"

	"github.com/anandhere8/ShopSync/internal/app/repository"
	"github.com/anandhere8/ShopSync/internal/util"
	"github.com/gin-gonic/gin"
)

func userProfile(c *gin.Context) {
	usernameValue, _ := c.Get("username")
	username := usernameValue.(string)
	// msg := fmt.Sprintf("Your username is %s", username)
	user, err := repository.GetUserByUsername(username)
	user.PasswordHash = ""
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
	}
	c.JSON(http.StatusOK, user)
}

func uploadProfilePic(c *gin.Context) {
	// userValue, _ := c.Get("user_id")
	// userId, _ := strconv.ParseInt(userValue.(string), 10, 64)
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	filePath := filepath.Join("img", file.Filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
