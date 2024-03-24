package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func ValidateCredential(storedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		err = errors.New("invalid credential")
	}
	return err
}

func Encrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
