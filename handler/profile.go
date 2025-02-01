package handler

import (
	"childgo/config"
	"childgo/database"
	"childgo/model"
	jwtUtil "childgo/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrJwtContext   = errors.New("wrong type of JWT token in context")
)

func Profile(ctx *fiber.Ctx) error {
	token, ok := ctx.Context().Value(config.ContextKeyUser).(*jwt.Token)

	if !ok {
		return ctx.Status(fiber.StatusBadRequest).SendString(ErrJwtContext.Error())
	}

	jwtPayload, err := jwtUtil.Payload(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	email := jwtPayload["sub"].(string)

	user := new(model.User)

	dbUser := database.DBConn.Where("email = ?", email).First(&user)

	if errors.Is(dbUser.Error, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusConflict).SendString(ErrUserNotFound.Error())
	}

	return ctx.JSON(ProfileResponse{
		Email: email,
	})
}
