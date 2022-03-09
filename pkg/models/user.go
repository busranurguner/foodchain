package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `json:"username"`
	Role     string             `json:"role"`
	Password string             `json:"password"`
	Refresh  string             `json:"refresh"`
}
