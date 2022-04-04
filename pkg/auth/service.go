package auth

import (
	"github.com/busranurguner/foodchain/pkg/models"
)

type AuthService interface {
	SignUp(req SignUpRequest) error
	Login(req LoginRequest) (string, string, error)
	Refresh(req RefreshRequest) (string, string, error)
}

type authService struct {
	repo AuthRepository
}

var _ AuthService = authService{}

func NewService(repo AuthRepository) AuthService {
	return authService{repo: repo}
}

func (a authService) SignUp(req SignUpRequest) error {
	users := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
		Refresh:  req.Refresh,
	}
	return a.repo.SignUp(users)
}

func (a authService) Login(req LoginRequest) (string, string, error) {
	return a.repo.Login(req)
}

func (a authService) Refresh(req RefreshRequest) (string, string, error) {
	return a.repo.Refresh(req)
}
