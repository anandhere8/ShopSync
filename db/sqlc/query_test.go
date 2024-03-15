package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/anandhere8/ShopSync/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	args := CreateUserParams{
		Username: util.RandomUsername(),
		Role:     util.RandomRole(),
	}
	account, err := testQueries.CreateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Username, account.Username)
	require.Equal(t, args.Role, account.Role)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	fmt.Printf("New user created with username - %v", args.Username)
	return account
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestUser(t *testing.T) {
	account1 := createRandomUser(t)
	account2, err := testQueries.GetUser(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Username, account2.Username)
	require.Equal(t, account1.Role, account2.Role)
	// require.Equal(t, account1.CreatedAt, account2.CreatedAt)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdate(t *testing.T) {
	account1 := createRandomUser(t)
	args := UpdateUserParams{
		Username: util.RandomUsername(),
		ID:       account1.ID,
		Role:     account1.Role,
	}
	fmt.Printf("Updating username from %s to %s", account1.Username, args.Username)
	account2, err := testQueries.UpdateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, args.Username, account2.Username)
	require.Equal(t, account1.Role, account2.Role)
	// require.Equal(t, account1.CreatedAt, account2.CreatedAt)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDelete(t *testing.T) {
	account1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), account1.ID)

	require.NoError(t, err)
	account2, err := testQueries.GetUser(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)

}

func TestListUsers(t *testing.T) {

	accounts, err := testQueries.ListUsers(context.Background())
	require.NoError(t, err)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
	log.Println(accounts)
}
