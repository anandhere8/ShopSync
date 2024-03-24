package handler

import (
	"net/http"

	"github.com/anandhere8/ShopSync/internal/util"
	"github.com/gin-gonic/gin"
)

func testHandler(ctx *gin.Context) {
	img, _ := util.GetRandomImage()
	ctx.JSON(http.StatusOK, gin.H{
		"url": img,
	})
}
