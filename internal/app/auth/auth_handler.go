package auth

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/anandhere8/ShopSync/internal/app/model"
	"github.com/anandhere8/ShopSync/internal/app/repository"
	"github.com/anandhere8/ShopSync/internal/app/service"
	"github.com/anandhere8/ShopSync/internal/util"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {

	fmt.Println("Content-type ; ", c.ContentType())
	var loginRequst model.LoginRequest
	err := c.Bind(&loginRequst)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	username := loginRequst.Username
	password := loginRequst.Password
	user, err := repository.GetUserByUsername(username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	if !service.ValidateCredential(user.PasswordHash, password) {
		c.JSON(http.StatusUnauthorized, util.ErrorResponse(err))
		return
	}

	token, err := GenerateJWT(strconv.FormatInt(user.UserID, 10), username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, token)
}
