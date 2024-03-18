package util

import (
	"math/rand"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	number   = "1234567890"
)

func generateString(n int) string {
	var usr strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		usr.WriteByte(c)
	}
	return usr.String()
}

func generateNumber(n int) string {
	var usr strings.Builder
	k := len(number)
	for i := 0; i < n; i++ {
		c := number[rand.Intn(k)]
		usr.WriteByte(c)
	}
	return usr.String()
}

// RandomUsername generate a random username
func RandomString() string {
	return generateString(8)
}

// RandomEmail generate a random mail address
func RandomEmail() string {
	return generateString(8) + "@gmail.com"
}

// RandomPhoneNumber gerate a random phone number
func RandomPhoneNumber() string {
	return "+919" + generateNumber(9)
}

// RandomRole generate a random role for the user
func RandomRole() string {
	userType := []string{"owner", "employee"}
	return userType[rand.Intn(2)]
}
