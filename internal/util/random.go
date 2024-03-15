package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func generateRandomUsername(n int) string {
	var usr strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		usr.WriteByte(c)
	}
	return usr.String()
}

// RandomUsername generate a random username
func RandomUsername() string {
	return generateRandomUsername(8)
}

// RandomRole generate a random role for the user
func RandomRole() string {
	userType := []string{"user", "admin"}
	return userType[rand.Intn(2)]
}
