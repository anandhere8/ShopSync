package service

import (
	"testing"

	"github.com/anandhere8/ShopSync/internal/util"
	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	password := util.RandomString()
	password_hash, err := Encrypt(password)
	require.NoError(t, err)
	require.NotEmpty(t, password_hash)
}
