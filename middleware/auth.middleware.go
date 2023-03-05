package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/stefanusong/votify-api/dto/response"
	"github.com/stefanusong/votify-api/services"
)

func AuthMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		resp := response.New(false, "Not Authorized", nil, "No token provided")
		c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}

	// get token value without bearer
	token = strings.Split(token, "Bearer ")[1]

	claims, err := services.ParseToken(token)
	if err != nil {
		resp := response.New(false, "Not Authorized", nil, err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}

	c.Set("userid", claims.ID)
	c.Set("username", claims.Username)
	c.Next()
}
