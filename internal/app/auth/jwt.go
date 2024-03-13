package auth

import (
	"fmt"
	"time"

	"github.com/anandhere8/ShopSync/internal/app/model"
	"github.com/anandhere8/ShopSync/key"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID, username string) (string, error) {

	privateKey, err := key.LoadPrivetKey()
	if err != nil {
		return "Failed to load the private key", err
	}
	// fmt.Println(privateKey)
	claims := model.CustomClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Audience:  []string{"user"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		fmt.Println("Failed to create token")
		fmt.Println(err)
		return "", err
	}
	return tokenString, nil
}
