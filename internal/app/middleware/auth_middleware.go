package middleware

import (
	"crypto/ecdsa"
	"fmt"
	"net/http"
	"strings"

	"github.com/anandhere8/ShopSync/key"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	authorizationHeader = "Authorization"
	authorizationType   = "Bearer"
)

func verifyToken(tokenString string, publicKey *ecdsa.PublicKey) (*jwt.Token, error) {
	// Parse and verify JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check token signing method
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return ECDSA public key for verification
		return publicKey, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("token verification failed: %v", err)
	}
	return token, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader(authorizationHeader)
		fields := strings.Fields(tokenString)
		if len(fields) < 2 {
			c.Abort()
		}
		accessToken := fields[1]
		fmt.Println(accessToken)
		publicKey, err := key.LoadPublicKey()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to load the public key"})
			c.Abort()
		}
		token, err := verifyToken(accessToken, publicKey)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		// fmt.Println(token)
		// fmt.Println(reflect.TypeOf(token.Claims))

		for key, val := range token.Claims.(jwt.MapClaims) {
			c.Set(key, val)
		}
		// c.Set("claims", )
		fmt.Println(c)
		c.Next()
	}
}
