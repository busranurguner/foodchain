package user

import (
	"context"

	"github.com/busranurguner/foodchain/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionName string = "User"
)

type UserRepository interface {
	GetAll(req GetAllRequest) ([]models.User, error)
	GetByID(id string) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id string) error
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

func (u userRepository) GetByID(id string) (*models.User, error) {
	var user models.User
	bid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = u.DB.FindOne(context.TODO(), bson.M{"_id": bid}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u userRepository) Create(user *models.User) error {
	user.ID = primitive.NewObjectID()
	_, err := u.DB.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (u userRepository) Update(user *models.User) error {
	update := bson.M{"password": user.Password}
	_, err := u.DB.UpdateOne(context.TODO(), bson.M{"_id": user.ID}, bson.M{"$set": update})
	if err != nil {
		return err
	}
	return nil
}

func (u userRepository) Delete(id string) error {
	bid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = u.DB.DeleteOne(context.TODO(), bson.M{"_id": bid})
	if err != nil {
		return err
	}
	return nil
}
