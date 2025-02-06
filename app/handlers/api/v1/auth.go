package handler

import (
	model "childgo/app/models"
	"childgo/app/models/repo"
	"childgo/config"
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
)

type (
	SignInResponse struct {
		JWTToken string `json:"jwt_token"`
	}
)

var ExpiriesTime = time.Now().Add(time.Hour * 24).Unix()

func Signin(ctx *fiber.Ctx) error {
	m := new(model.User)

	if err := ctx.BodyParser(m); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user data"})
	}

	u := &model.User{}

	if err := repo.FindUser(u, m).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	payload := jwt.MapClaims{
		"sub": u.Email,
		"exp": ExpiriesTime,
	}

	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := claim.SignedString([]byte(config.SecretKey))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "incorrect jwt token"})
	}

	return ctx.JSON(SignInResponse{
		JWTToken: token,
	})
}

func Signup(ctx *fiber.Ctx) error {
	m := new(model.User)

	if err := ctx.BodyParser(m); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user data"})
	}

	u := &model.User{}

	if err := repo.FindUser(u, "email", m.Email).Error; err == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "user already exist"})
	}

	if err := repo.CreateUser(m).Error; err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user created error"})
	}

	return ctx.SendStatus(fiber.StatusOK)
}
