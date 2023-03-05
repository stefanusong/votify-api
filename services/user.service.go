package services

import (
	"errors"
	"regexp"

	"github.com/stefanusong/votify-api/dto/request"
	"github.com/stefanusong/votify-api/helpers"
	"github.com/stefanusong/votify-api/models"
	"github.com/stefanusong/votify-api/repositories"
)

type UserService interface {
	CreateUser(newUser request.RegisterUser) error
	GetUserByUsername(username string) (any, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (svc *userService) CreateUser(newUser request.RegisterUser) error {
	hashedPassword, err := helpers.Encrypt(newUser.Password)
	if err != nil {
		return err
	}

	isUsernameAlphaNumeric := regexp.MustCompile(`^[a-z0-9]*$`).MatchString(newUser.Username)
	if !isUsernameAlphaNumeric {
		return errors.New("username must be alphanumeric")
	}

	if svc.isUsernameTaken(newUser.Username) {
		return errors.New("username already taken")
	}

	createdUser := models.NewUser(newUser.Name, newUser.Username, hashedPassword)
	errInsert := svc.repo.InsertUser(createdUser)
	if errInsert != nil {
		return err
	}

	return nil
}

func (svc *userService) isUsernameTaken(username string) bool {
	user := svc.repo.GetUserByUsername(username)
	return user != nil
}

func (svc *userService) GetUserByUsername(username string) (any, error) {
	if username == "" {
		return nil, errors.New("invalid username")
	}

	user := svc.repo.GetUserByUsername(username)

	if user == nil {
		return nil, errors.New("user not found")
	}
	res := map[string]interface{}{
		"name":     user.Name,
		"username": user.Username,
	}

	return res, nil
}
