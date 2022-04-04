package models

type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshtoken"`
}
type Refresh struct {
	RefreshToken string `json:"refresh"`
}
