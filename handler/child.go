package handler

import (
	"childgo/config"
	"childgo/database"
	"childgo/model"
	"errors"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func Childs(ctx *fiber.Ctx) error {
	_, ok := ctx.Locals(config.ContextKeyUser).(*jwt.Token)

	if !ok {
		return ctx.Status(fiber.StatusBadRequest).SendString(ErrJwtContext.Error())
	}

	childs := make([]model.Child, 0)

	database.DBConn.Find(&childs)

	return ctx.JSON(childs)
}

func NewChild(ctx *fiber.Ctx) error {
	_, ok := ctx.Locals(config.ContextKeyUser).(*jwt.Token)

	if !ok {
		return ctx.Status(fiber.StatusBadRequest).SendString(ErrJwtContext.Error())
	}

	child := new(model.Child)

	if err := ctx.BodyParser(child); err != nil {
		return ctx.Status(503).SendString(err.Error())
	}

	database.DBConn.Create(&child)

	return ctx.JSON(child)
}

func GetChild(ctx *fiber.Ctx) error {
	_, ok := ctx.Locals(config.ContextKeyUser).(*jwt.Token)

	if !ok {
		return ctx.Status(fiber.StatusBadRequest).SendString(ErrJwtContext.Error())
	}

	id := ctx.Params("id")

	child := new(model.Child)

	dbChild := database.DBConn.First(&child, id)

	if errors.Is(dbChild.Error, gorm.ErrRecordNotFound) {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return ctx.JSON(child)
}

func DeleteChild(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	child := new(model.Child)

	dbChild := database.DBConn.First(&child, id)

	if errors.Is(dbChild.Error, gorm.ErrRecordNotFound) {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	database.DBConn.Delete(&child)

	return ctx.SendStatus(fiber.StatusOK)
}
