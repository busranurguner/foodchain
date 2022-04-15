package auth

import (
	"context"

	"github.com/busranurguner/foodchain/pkg/logger"
	"github.com/busranurguner/foodchain/pkg/models"
	"github.com/busranurguner/foodchain/pkg/token"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionName string = "User"
)

type AuthRepository interface {
	SignUp(user *models.User) error
	Login(req LoginRequest) (string, string, error)
	Refresh(req RefreshRequest) (string, string, error)
}

type authRepository struct {
	DB *mongo.Collection
}

var _ AuthRepository = authRepository{}

func NewRepository(DB *mongo.Database) AuthRepository {
	return authRepository{DB: DB.Collection(collectionName)}
}

//SignUp implements AuthRepository
func (a authRepository) SignUp(user *models.User) error {
	user.ID = primitive.NewObjectID()
	_, err := a.DB.InsertOne(context.TODO(), user)
	if err != nil {
		logger.L.Error("There was an error sign up the user")
		return err
	}
	return nil

}

//Login implements AuthRepository
func (a authRepository) Login(req LoginRequest) (string, string, error) {
	var foundUser models.User
	err := a.DB.FindOne(context.TODO(), bson.M{"username": req.Username, "password": req.Password}).Decode(&foundUser)
	if err != nil {
		logger.L.Error("User not found")
		return "", "", err
	}
	atoken, rtoken, err := token.Token(foundUser.Username, foundUser.Password, foundUser.Role)
	if err != nil {
		return "", "", err
	}
	//refresh token add
	update := bson.M{"refresh": rtoken}
	_, err = a.DB.UpdateOne(context.TODO(), bson.M{"_id": foundUser.ID}, bson.M{"$set": update})
	if err != nil {
		return "", "", err
	}

	return atoken, rtoken, nil
}

//Refresh implements AuthRepository
func (a authRepository) Refresh(req RefreshRequest) (string, string, error) {
	var foundUser models.User
	err := a.DB.FindOne(context.TODO(), bson.M{"refresh": req.Refresh}).Decode(&foundUser)
	if err != nil {
		return "", "", nil
	}
	atoken, rtoken, err := token.Token(foundUser.Username, foundUser.Password, foundUser.Role)
	if err != nil {
		return "", "", nil
	}
	update := bson.M{"refresh": rtoken}
	_, err = a.DB.UpdateOne(context.TODO(), bson.M{"_id": foundUser.ID}, bson.M{"$set": update})
	if err != nil {
		return "", "", nil
	}

	return atoken, rtoken, nil
}
