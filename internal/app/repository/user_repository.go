package repository

import "github.com/anandhere8/ShopSync/internal/app/model"

func GetUserPassword(username string) (string, error) {
	if username == "anand" {
		return "123", nil
	}
	return "kkk", nil
}

func GetUserID(username string) (string, error) {
	if username == "anand" {
		return "ID123", nil
	}
	return "asfsaf", nil
}

func RegisterUser(model.User) error {
	return nil
}
