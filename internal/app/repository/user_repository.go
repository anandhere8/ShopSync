package repository

import (
	"context"

	db "github.com/anandhere8/ShopSync/db"
	sqlc "github.com/anandhere8/ShopSync/db/sqlc"
	"github.com/anandhere8/ShopSync/internal/app/model"
	"github.com/anandhere8/ShopSync/internal/app/service"
)

func GetUserByUsername(username string) (sqlc.User, error) {
	dbClient, err := db.GetDBInstance()
	if err != nil {
		return sqlc.User{}, err
	}
	newUser, err := dbClient.GetUserByUsername(context.Background(), username)
	if err != nil {
		return sqlc.User{}, err
	}
	return newUser, nil
}

func RegisterUser(u model.User) (sqlc.User, error) {
	dbClient, err := db.GetDBInstance()
	if err != nil {
		return sqlc.User{}, err
	}
	passwordHash, err := service.Encrypt(u.Password)
	if err != nil {
		return sqlc.User{}, err
	}
	args := sqlc.CreateUserParams{
		Firstname:    u.Firstname,
		Lastname:     u.Lastname,
		Username:     u.Username,
		Email:        u.Email,
		PhoneNumber:  u.PhoneNumber,
		PasswordHash: passwordHash,
	}
	newUser, err := dbClient.CreateUser(context.Background(), args)
	if err != nil {
		return sqlc.User{}, err
	}
	return newUser, nil
}
