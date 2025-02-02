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

type ProfileResponse struct {
	Email string `json:"email"`
}

var (
	ErrUserNotFound = errors.New("user not found")
	ErrJwtContext   = errors.New("wrong type of JWT token in context")
)

func Profile(ctx *fiber.Ctx) error {
	token, ok := ctx.Locals(config.ContextKeyUser).(*jwt.Token)

	if !ok {
		return ctx.Status(fiber.StatusBadRequest).SendString(ErrJwtContext.Error())
	}

	email, err := jwtUtil.FindEmailFromToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	_, err = FindByEmail(email)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return ctx.JSON(ProfileResponse{
		Email: email,
	})
}

func FindByEmail(email string) (*gorm.DB, error) {
	user := new(model.User)

	dbUser := database.DBConn.Where("email = ?", email).First(&user)

	if errors.Is(dbUser.Error, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}

	return dbUser, nil
}
