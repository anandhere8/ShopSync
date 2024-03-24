package auth

import (
	"fmt"
	"time"

	"github.com/anandhere8/ShopSync/internal/app/model"
	"github.com/anandhere8/ShopSync/key"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID, username string) (model.Token, error) {

	privateKey, err := key.LoadPrivetKey()
	if err != nil {
		return model.Token{}, err
	}
	loc, _ := time.LoadLocation("Asia/Kolkata")
	now := time.Now().In(loc)
	claims := model.CustomClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			Audience:  []string{"user"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		fmt.Println("Failed to create token")
		fmt.Println(err)
		return model.Token{}, err
	}

	curtoken := model.Token{
		Access_token: tokenString,
		Issuer:       username,
		Auth_type:    "bearer",
		ExpiresAt:    claims.RegisteredClaims.ExpiresAt.String(),
		CreatedAt:    claims.RegisteredClaims.IssuedAt.String(),
	}

	return curtoken, nil
}
