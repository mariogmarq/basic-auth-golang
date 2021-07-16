package jwtUtils

import (
	"GoLandPruebas/internal/services"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type CustomClaim struct {
	Expiration time.Time
	UserId uint
}

func (c CustomClaim) Valid() error {
	if c.Expiration.Before(time.Now()) {
		return errors.New("expired token")
	}

	if _, err := services.UserService.GetUserById(c.UserId); err != nil {
		return errors.New("user does not exists")
	}

	return nil
}

func createClaim(t time.Duration, userId uint) CustomClaim {
	return CustomClaim{Expiration: time.Now().Add(t), UserId: userId}
}

var MyKey = []byte("hola mundo")

func CreateToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, createClaim(time.Hour, userId))
	signedString, err := token.SignedString(MyKey)
	if err != nil {
		return "", err
	}
	return signedString, err
}