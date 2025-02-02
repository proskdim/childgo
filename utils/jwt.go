package jwtUtil

import (
	"childgo/config"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrJwtTokenClaims = errors.New("wrong type of JWT token claims")
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

func Payload(token *jwt.Token) (jwt.MapClaims, error) {
	payload, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, ErrJwtTokenClaims
	}

	return payload, nil
}

func FindEmailFromToken(token *jwt.Token) (string, error) {
	var email string

	jwtPayload, err := Payload(token)

	if err != nil {
		return email, err
	}

	email = jwtPayload["sub"].(string)

	return email, err
}
