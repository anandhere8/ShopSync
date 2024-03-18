package repository

import (
	"context"

	db "github.com/anandhere8/ShopSync/db"
	sqlc "github.com/anandhere8/ShopSync/db/sqlc"
	"github.com/anandhere8/ShopSync/internal/app/model"
)

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

func RegisterUser(u model.User) (sqlc.User, error) {
	dbClient, err := db.GetDBInstance()
	if err != nil {
		return sqlc.User{}, err
	}
	args := sqlc.CreateUserParams{
		Firstname:    u.Firstname,
		Lastname:     u.Lastname,
		Username:     u.Username,
		Email:        u.Email,
		PhoneNumber:  u.PhoneNumber,
		PasswordHash: u.Password,
	}
	newUser, err := dbClient.CreateUser(context.Background(), args)
	if err != nil {
		return sqlc.User{}, err
	}
	return newUser, nil
}
