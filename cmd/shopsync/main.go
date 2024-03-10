package main

import (
	routing "github.com/anandhere8/ShopSync/internal/app/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routing.ConfigureRoutes(router)
	router.Run(":8080")
}
