package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// Generate ECDSA key pair
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println("Error generating ECDSA key pair:", err)
		return
	}

	// Generate JWT token
	tokenString, err := generateToken(privateKey)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}
	fmt.Println("Generated Token:", tokenString)

	// Verify token
	err = verifyToken(tokenString, &privateKey.PublicKey)
	if err != nil {
		fmt.Println("Token verification failed:", err)
		return
	}
	fmt.Println("Token verification successful")
}

func generateToken(privateKey *ecdsa.PrivateKey) (string, error) {
	// Create JWT token claims
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
		Issuer:    "example.com",
	}

	// Create and sign JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func verifyToken(tokenString string, publicKey *ecdsa.PublicKey) error {
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
		return fmt.Errorf("token verification failed: %v", err)
	}
	return nil
}
