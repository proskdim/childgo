package jwtUtil

import (
	"childgo/config"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken   = errors.New("invalid jwt token")
)

func Get(payload jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(config.SecretKey)

	if err != nil {
		return "", ErrInvalidToken
	}

	return t, nil
}