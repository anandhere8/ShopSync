package service

import (
	"github.com/anandhere8/ShopSync/internal/app/repository"
)

func ValidateCredential(username, password string) bool {
	storedPassword, err := repository.GetUserPassword(username)
	if err != nil {
		return false
	}
	// fmt.Println(storedPassword)
	// fmt.Println(password)
	return password == storedPassword
}
