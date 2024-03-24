package util

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"strings"
	"time"
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

func GetRandomImage4() (string, error) {
	// Read all files from the folder
	folderPath := "/home/layman/Downloads/instaloader-master/download"

	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return "", err
	}

	// Filter image files (you may adjust this to filter by specific file extensions)
	var imageFiles []string
	for _, file := range files {
		if !file.IsDir() {
			imageFiles = append(imageFiles, file.Name())
		}
	}

	// Select a random image file
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	randomIndex := rand.Intn(len(imageFiles))
	randomImage := imageFiles[randomIndex]

	// Construct the full file path
	// reactPath := "../../download/"
	// reactPath = "file:///home/layman/layman/Projects/App/Native/download/tree.png"
	// randomImage = "tree.png"
	httpserver := "http://192.168.29.71:8081/"
	imagePath := filepath.Join(httpserver, randomImage)

	return imagePath, nil
}

func GetRandomImage() (string, error) {
	// Read all files from the folder
	folderPath := "/home/layman/Downloads/instaloader-master/download"

	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return "", err
	}

	// Filter image files
	var imageFiles []string
	for _, file := range files {
		if !file.IsDir() && isImage(file.Name()) {
			imageFiles = append(imageFiles, file.Name())
		}
	}

	// Check if there are any image files
	if len(imageFiles) == 0 {
		return "", errors.New("no image files found in the folder")
	}

	// Select a random image file
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	randomIndex := rand.Intn(len(imageFiles))
	randomImage := imageFiles[randomIndex]

	// Construct the full file path
	httpserver := "http://192.168.29.71:8081/"
	imagePath := filepath.Join(httpserver, randomImage)

	return imagePath, nil
}

// Function to check if the file has an image extension
func isImage(filename string) bool {
	extension := strings.ToLower(filepath.Ext(filename))
	return extension == ".jpg" || extension == ".jpeg" || extension == ".png" || extension == ".gif" || extension == ".mp4"
}
