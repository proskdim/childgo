package handler

import (
	model "childgo/app/models"
	"childgo/app/models/repo"
	"childgo/config"
	"childgo/utils"
	"childgo/utils/password"
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
	u := new(model.User)

	if err := utils.ParseBody(ctx, m); err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{"error": "invalid user data"})
	}

	if err := repo.FindUser(u, "email", m.Email).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	if err := password.Verify(u.Password, m.Password); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "incorrect password"})
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
	u := new(model.User)

	if err := utils.ParseBody(ctx, m); err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{"error": "invalid user data"})
	}

	if err := repo.FindUser(u, "email", m.Email).Error; err == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "user already exist"})
	}

	m.Password = password.Generate(m.Password)

	if err := repo.CreateUser(m).Error; err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user created error"})
	}

	return ctx.SendStatus(fiber.StatusOK)
}
