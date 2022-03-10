package models

type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshtoken"`
	Status       string `json:"status"`
}
type Refresh struct {
	RefreshToken string `json:"refresh"`
}
