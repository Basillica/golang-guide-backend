package utils

import (
	"errors"
	"time"

	"github.com/Basillica/golang-guide/models"
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	User models.User `json:"user"`
	jwt.StandardClaims
}

func GenerateJWT(user models.User, secret string) (string, error) {
	claims := CustomClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "github.com/Basillica/golang-guide",
			Audience:  "github.com/Basillica/golang-guide",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateJWT(tokenString, secret string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token string provided")
	}
	return claims, nil
}
