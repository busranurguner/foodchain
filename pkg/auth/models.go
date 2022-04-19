package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

type SignUpRequest struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `json:"username" validate:"required"`
	Email    string             `json:"email" validate:"required"`
	Role     string             `json:"role" validate:"required"`
	Password string             `json:"password" validate:"required"`
	Refresh  string             `json:"refresh"`
}
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type RefreshRequest struct {
	Refresh string `json:"refresh"`
}
