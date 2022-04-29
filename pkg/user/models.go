package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetAllRequest struct {
	Username string `json:"username" `
	Email    string `json:"email" `
	Role     string `json:"role" `
	Password string `json:"password" `
	Refresh  string `json:"refresh"`
}

type CreateRequest struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `json:"username" validate:"required"`
	Email    string             `json:"email" validate:"required"`
	Role     string             `json:"role" validate:"required"`
	Password string             `json:"password" validate:"required"`
	Refresh  string             `json:"refresh"`
}
type UpdateRequest struct {
	ID       string `bson:"_id"`
	Password string `json:"password"`
}

type DeleteRequest struct {
	ID string `bson:"_id"`
}
