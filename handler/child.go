package handler

import (
	"childgo/config"
	"childgo/database"
	"childgo/model"
	"childgo/model/child"
	"strconv"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
)

func checkJwt(ctx *fiber.Ctx) (*jwt.Token, error) {
	token, ok := ctx.Locals(config.ContextKeyUser).(*jwt.Token)

	if !ok {
		return nil, ErrJwtContext
	}

	return token, nil
}

func Childs(ctx *fiber.Ctx) error {
	if _, err := checkJwt(ctx); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	childs := make([]model.Child, 0)

	database.DBConn.Find(&childs)

	return ctx.JSON(childs)
}

func NewChild(ctx *fiber.Ctx) error {
	if _, err := checkJwt(ctx); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	child := new(model.Child)

	if err := ctx.BodyParser(child); err != nil {
		return ctx.Status(503).SendString(err.Error())
	}

	database.DBConn.Create(&child)

	return ctx.JSON(child)
}

func GetChild(ctx *fiber.Ctx) error {
	if _, err := checkJwt(ctx); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	childId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	child, err := child.FindById(database.DBConn, childId)

	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return ctx.JSON(child)
}

func DeleteChild(ctx *fiber.Ctx) error {
	if _, err := checkJwt(ctx); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	childId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	child, err := child.FindById(database.DBConn, childId)

	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	database.DBConn.Delete(&child)

	return ctx.SendStatus(fiber.StatusOK)
}
