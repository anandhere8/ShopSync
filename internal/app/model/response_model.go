package model

type LoginResponse struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"after"`
	CreatedAt string `json:"createdAt"`
	Issuer    string `json:"issuer"`
}
