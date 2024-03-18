package handler

import (
	"github.com/anandhere8/ShopSync/internal/app/auth"
	"github.com/anandhere8/ShopSync/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine) {
	router.POST("/login", auth.LoginHandler)
	router.POST("/register", registerHandler)
	router.Use(middleware.AuthMiddleware())
	router.GET("/home", homehandler)
}
