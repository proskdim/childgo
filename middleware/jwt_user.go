package middleware

import (
	"childgo/config"
	"childgo/database"
	"childgo/model"
	"childgo/model/user"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// fetch user before api handlers
func JwtUserMiddleware(ctx *fiber.Ctx) error {
	token, ok := ctx.Locals(config.CONTEXT_KEY_USER).(*jwt.Token)
	db := database.DBConn

	if !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "wrong type of JWT token in context"})
	}

	user, err := fetchUser(db, token)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	ctx.Locals(config.CONTEXT_JWT_USER, user)

	return ctx.Next()
}

// fetch user from db by jwt token
func fetchUser(db *gorm.DB, token *jwt.Token) (*model.User, error) {
	payload, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("failed to extraction token")
	}

	email := payload["sub"].(string)

	dbUser, err := user.FindByEmail(db, email)

	if err != nil {
		return nil, errors.New("user by jwt token not found")
	}

	return dbUser, nil
}
