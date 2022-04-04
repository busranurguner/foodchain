package user

import (
	"github.com/busranurguner/foodchain/pkg/models"
)

type UserService interface {
	GetAll(req GetAllRequest) ([]models.User, error)
}

type userService struct {
	repo UserRepository
}

var _ UserService = userService{}

func NewService(repo UserRepository) UserService {
	return userService{repo: repo}
}

func (u userService) GetAll(req GetAllRequest) ([]models.User, error) {
	model, err := u.repo.GetAll(req)
	if err != nil {
		return nil, err
	}
	return model, err
}
