package handler

import (
	"childgo/database"
	"childgo/model"
	"childgo/model/user"
	jwtUtil "childgo/utils"
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
	db := database.DBConn

	userData := new(model.User)

	if err := ctx.BodyParser(&userData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user data"})
	}

	if _, err := user.Find(db, userData); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	payload := jwt.MapClaims{
		"sub": &userData.Email,
		"exp": ExpiriesTime,
	}

	token, err := jwtUtil.Get(payload)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "incorrect jwt token"})
	}

	return ctx.JSON(SignInResponse{
		JWTToken: token,
	})
}

func Signup(ctx *fiber.Ctx) error {
	db := database.DBConn

	userData := new(model.User)

	if err := ctx.BodyParser(&userData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user data"})
	}

	if _, err := user.FindByEmail(db, userData.Email); err == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "user already exist"})
	}

	db.Create(&userData)

	return ctx.SendStatus(fiber.StatusOK)
}
