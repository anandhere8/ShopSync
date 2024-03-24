package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/anandhere8/ShopSync/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	args := CreateUserParams{
		Firstname:    util.RandomString(),
		Lastname:     util.RandomString(),
		Email:        util.RandomEmail(),
		PhoneNumber:  util.RandomPhoneNumber(),
		PasswordHash: util.RandomString(),
		Username:     util.RandomString(),
	}

	account, err := testQueries.CreateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, args.Firstname, account.Firstname)
	require.Equal(t, args.Lastname, account.Lastname)
	require.Equal(t, args.Username, account.Username)
	require.Equal(t, args.PhoneNumber, account.PhoneNumber)
	require.Equal(t, args.Email, account.Email)
	require.Equal(t, args.PasswordHash, account.PasswordHash)

	return account
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestUpdateUser(t *testing.T) {
	oldUser := createRandomUser(t)
	newFirstname := "Anand"
	usr, _ := json.MarshalIndent(oldUser, "", "")
	var newUser UpdateUserParams
	_ = json.Unmarshal(usr, &newUser)
	newUser.Firstname = newFirstname
	account, err := testQueries.UpdateUser(context.Background(), newUser)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, newUser.Firstname, account.Firstname)
}

func TestGetUserByID(t *testing.T) {
	oldUser := createRandomUser(t)
	newUser, err := testQueries.GetUserByID(context.Background(), oldUser.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, newUser)
	require.Equal(t, oldUser.Firstname, newUser.Firstname)
	require.Equal(t, oldUser.Lastname, newUser.Lastname)
	require.Equal(t, oldUser.Username, newUser.Username)
	require.Equal(t, oldUser.PasswordHash, newUser.PasswordHash)
	require.Equal(t, oldUser.CreatedAt, newUser.CreatedAt)
	require.Equal(t, oldUser.UserID, newUser.UserID)
	require.Equal(t, oldUser.Email, newUser.Email)
	require.Equal(t, oldUser.PhoneNumber, newUser.PhoneNumber)
}

func TestDeleteUser(t *testing.T) {
	oldUser := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), oldUser.UserID)
	require.NoError(t, err)
	newUser, err := testQueries.GetUserByID(context.Background(), oldUser.UserID)
	require.Error(t, err, sql.ErrNoRows.Error())
	require.Empty(t, newUser)
}

func TestListUser(t *testing.T) {
	allUsr, err := testQueries.ListUsers(context.Background())
	require.NoError(t, err)
	for _, usr := range allUsr {
		require.NotEmpty(t, usr)
		testQueries.DeleteUser(context.Background(), usr.UserID)
	}
}
