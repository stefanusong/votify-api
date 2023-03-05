package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/stefanusong/votify-api/dto/request"
	"github.com/stefanusong/votify-api/dto/response"
	"github.com/stefanusong/votify-api/services"
)

type UserHandler interface {
	GetUserByUsername(c *gin.Context)
	Register(c *gin.Context)
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (handler *userHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := handler.userService.GetUserByUsername(username)
	if err != nil {
		resp := response.New(false, "Failed to get profile", nil, err.Error())
		c.JSON(http.StatusUnauthorized, resp)
		return
	}

	resp := response.New(true, "User Profile", user, nil)
	c.JSON(http.StatusOK, resp)
}

func (handler *userHandler) Register(c *gin.Context) {
	// Bind user
	var newUser request.RegisterUser
	if err := c.ShouldBindJSON(&newUser); err != nil {
		resp := response.New(false, "Failed to register", nil, err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	newUser.Username = strings.ToLower(newUser.Username)

	// Create user
	err := handler.userService.CreateUser(newUser)
	if err != nil {
		resp := response.New(false, "Failed to register", nil, err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := response.New(true, "User registered", nil, nil)
	c.JSON(http.StatusCreated, resp)
}
