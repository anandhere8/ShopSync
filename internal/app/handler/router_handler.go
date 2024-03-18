package handler

import (
	"github.com/anandhere8/ShopSync/internal/app/auth"
	"github.com/anandhere8/ShopSync/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine) {
	router.POST("/login", auth.LoginHandler)
	router.POST("/register", registerHandler)
	router.GET("/test", testHandler)
	// router.Use(middleware.AuthMiddleware())
	authRouter := router.Group("/").Use(middleware.AuthMiddleware())
	authRouter.GET("/home", homehandler)
}
