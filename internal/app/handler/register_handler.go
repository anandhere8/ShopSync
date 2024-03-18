package handler

import (
	"net/http"

	"github.com/anandhere8/ShopSync/internal/app/model"
	"github.com/anandhere8/ShopSync/internal/app/repository"
	"github.com/anandhere8/ShopSync/internal/util"
	"github.com/gin-gonic/gin"
)

func registerHandler(ctx *gin.Context) {
	var usr model.User
	if err := ctx.BindJSON(&usr); err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))

		return
	}

	user, err := repository.RegisterUser(usr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))

		return
	}

	ctx.JSON(http.StatusOK, user)
}
