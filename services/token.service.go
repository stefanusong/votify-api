package services

import (
	"time"

	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
	"github.com/stefanusong/votify-api/config"
)

type Claims struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(id uuid.UUID, name, username string) (string, error) {
	jwtSecret := config.GetJWTSecretKey()
	expirationTime := time.Now().Add(time.Hour * 24) // one day expiration time

	claims := &Claims{
		ID:       id.String(),
		Name:     name,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	jwtSecret := config.GetJWTSecretKey()
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
