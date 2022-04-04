package user

import (
	"context"

	"github.com/busranurguner/foodchain/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionName string = "User"
)

type UserRepository interface {
	GetAll(req GetAllRequest) ([]models.User, error)
}

type userRepository struct {
	DB *mongo.Collection
}

var _ UserRepository = userRepository{}

func NewRepository(DB *mongo.Database) UserRepository {
	return userRepository{DB: DB.Collection(collectionName)}
}

// GetAll implements UserRepository
func (u userRepository) GetAll(req GetAllRequest) ([]models.User, error) {
	var users []models.User
	cursor, err := u.DB.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user models.User
		if err = cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, cursor.Err()
}
