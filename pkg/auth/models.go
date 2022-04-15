package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

type SignUpRequest struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `json:"username" `
	Email    string             `json:"email" `
	Role     string             `json:"role" `
	Password string             `json:"password" `
	Refresh  string             `json:"refresh"`
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RefreshRequest struct {
	Refresh string `json:"refresh"`
}
