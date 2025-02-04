package handler

import (
	"childgo/database"
	"childgo/model"
	jwtUtil "childgo/utils"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type (
	SignInResponse struct {
		JWTToken string `json:"jwt_token"`
	}
)

var ExpiriesTime = time.Now().Add(time.Hour * 24).Unix()

func Signin(ctx *fiber.Ctx) error {
	user := new(model.User)

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user data"})
	}

	dbUser := database.DBConn.Where("email = ? AND password = ?", &user.Email, &user.Password).First(&user)

	if errors.Is(dbUser.Error, gorm.ErrRecordNotFound) {
		return ctx.SendStatus(fiber.StatusUnprocessableEntity)
	}

	payload := jwt.MapClaims{
		"sub": &user.Email,
		"exp": ExpiriesTime,
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
	user := new(model.User)

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user data"})
	}

	if err := database.DBConn.First(&user, "email = ?", &user.Email).Error; err == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "user already exist"})
	}

	database.DBConn.Create(&user)

	return ctx.SendStatus(fiber.StatusOK)
}
