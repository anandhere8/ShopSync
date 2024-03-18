package handler

import (
	"log"
	"net/http"

	"github.com/anandhere8/ShopSync/internal/app/model"
	"github.com/anandhere8/ShopSync/internal/app/repository"
	"github.com/gin-gonic/gin"
)

func registerHandler(ctx *gin.Context) {
	var usr model.User
	if err := ctx.BindJSON(&usr); err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to parse the Payload",
		})

		return
	}
	repository.RegisterUser(usr)
	ctx.JSON(http.StatusOK, gin.H{
		"message ": "User created successfully",
	})
}
