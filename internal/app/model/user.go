package model

type User struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
