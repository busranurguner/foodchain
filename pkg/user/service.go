package user

import (
	"github.com/busranurguner/foodchain/pkg/models"
)

type UserService interface {
	GetAll(req GetAllRequest) ([]models.User, error)
	GetByID(id string) (*models.User, error)
	Update(req UpdateRequest) error
	Delete(req DeleteRequest) error
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
func (u userService) GetByID(id string) (*models.User, error) {
	return u.repo.GetByID(id)
}

func (u userService) Update(req UpdateRequest) error {
	user, err := u.repo.GetByID(req.ID)
	if err != nil {
		return err
	}
	user.Password = req.Password
	return u.repo.Update(user)
}

func (u userService) Delete(req DeleteRequest) error {
	return u.repo.Delete(req.ID)
}
