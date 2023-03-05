package services

import (
	"errors"

	"github.com/stefanusong/votify-api/dto/request"
	"github.com/stefanusong/votify-api/helpers"
	"github.com/stefanusong/votify-api/repositories"
)

type AuthService interface {
	Login(user request.LoginUser) (any, error)
}

type authService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (svc *authService) Login(loginUser request.LoginUser) (any, error) {
	user := svc.repo.GetUserByUsername(loginUser.Username)

	if user == nil || !helpers.CompareEncryption(user.Password, loginUser.Password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := GenerateToken(user.ID, user.Name, user.Username)
	if err != nil {
		return nil, err
	}

	resp := map[string]any{"token": token}
	return resp, nil
}
