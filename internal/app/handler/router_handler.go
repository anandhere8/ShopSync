package handler

import (
	"github.com/anandhere8/ShopSync/internal/app/auth"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine) {
	router.POST("/login", auth.LoginHandler)
}
