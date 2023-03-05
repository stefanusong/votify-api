package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/stefanusong/votify-api/dto/request"
	"github.com/stefanusong/votify-api/dto/response"
	"github.com/stefanusong/votify-api/services"
)

type AuthHandler interface {
	Login(c *gin.Context)
	AuthCheck(c *gin.Context)
}

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

func (handler *authHandler) Login(c *gin.Context) {
	var user request.LoginUser
	if err := c.ShouldBindJSON(&user); err != nil {
		resp := response.New(false, "Failed to login", nil, "Username and password is required")
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	user.Username = strings.ToLower(user.Username)

	data, err := handler.authService.Login(user)
	if err != nil {
		resp := response.New(false, "Failed to login", nil, err.Error())
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, resp)
			return
		}

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := response.New(true, "User logged in", data, nil)
	c.JSON(http.StatusOK, resp)
}

func (handler *authHandler) AuthCheck(c *gin.Context) {
	// if this route is not returning 403 status code, we know the user is authenticated
	// (auth middleware does the job 👌)
	c.JSON(http.StatusOK, gin.H{})
}
