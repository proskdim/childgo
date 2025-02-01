package handler

import (
	"childgo/config"
	jwtUtil "childgo/utils"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type (
	SignUpRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	SignInRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	SignInResponse struct {
		JWTToken string `json:"jwt_token"`
	}

	ProfileResponse struct {
		Email string `json:"email"`
	}

	User struct {
		Email    string
		password string
	}
)

var users = map[string]User{}

func Signin(ctx *fiber.Ctx) error {
	regReq := SignInRequest{}

	if err := ctx.BodyParser(&regReq); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	user, exists := users[regReq.Email]

	if !exists {
		return ctx.SendStatus(fiber.StatusUnprocessableEntity)
	}

	if regReq.Password != user.password {
		return ctx.SendStatus(fiber.StatusUnprocessableEntity)
	}

	payload := jwt.MapClaims{
		"sub": regReq.Email,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token, err := jwtUtil.Get(payload)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return ctx.JSON(SignInResponse{
		JWTToken: token,
	})
}

func Signup(ctx *fiber.Ctx) error {
	regReq := SignUpRequest{}

	if err := ctx.BodyParser(&regReq); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	if _, exists := users[regReq.Email]; exists {
		return ctx.SendStatus(fiber.StatusConflict)
	}

	users[regReq.Email] = User{
		Email:    regReq.Email,
		password: regReq.Password,
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func Profile(ctx *fiber.Ctx) error {
	token, ok := ctx.Context().Value(config.ContextKeyUser).(*jwt.Token)

	if !ok {
		logrus.WithFields(logrus.Fields{
			"jwt_token_context_value": ctx.Context().Value(config.ContextKeyUser),
		}).Error("wrong type of JWT token in context")

		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	jwtPayload, err := jwtUtil.Payload(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	userInfo, ok := users[jwtPayload["sub"].(string)]

	if !ok {
		return errors.New("user not found")
	}

	return ctx.JSON(ProfileResponse{
		Email: userInfo.Email,
	})
}
