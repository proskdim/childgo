package handler

import (
	"childgo/config"
	"childgo/config/database"
	"childgo/model"
	"childgo/model/user"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// get all childs for current user
func Childs(ctx *fiber.Ctx) error {
	fetchedUser := ctx.Locals(config.ContextJwtUser).(*model.User)

	childs, err := user.FindAllChilds(database.DBConn, fetchedUser)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "error to fetch all childs"})
	}

	return ctx.JSON(childs)
}

// add new child
func NewChild(ctx *fiber.Ctx) error {
	fetchedUser := ctx.Locals(config.ContextJwtUser).(*model.User)

	child := new(model.Child)

	if err := ctx.BodyParser(child); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid child data"})
	}

	dbChild, err := user.AddChild(database.DBConn, fetchedUser, child)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error to add child"})
	}

	return ctx.JSON(dbChild)
}

// get child by id
func GetChild(ctx *fiber.Ctx) error {
	fetchedUser := ctx.Locals(config.ContextJwtUser).(*model.User)

	childId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to convert id param"})
	}

	child, err := user.FindChild(database.DBConn, fetchedUser, childId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "child not found"})
	}

	return ctx.JSON(child)
}

// delete child for current user
func DeleteChild(ctx *fiber.Ctx) error {
	fetchedUser := ctx.Locals(config.ContextJwtUser).(*model.User)

	childId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to convert id param"})
	}

	child, err := user.FindChild(database.DBConn, fetchedUser, childId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "child by id not found"})
	}

	if _, err := user.DeleteChild(database.DBConn, child); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error to delete child"})
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func UpdateChild(ctx *fiber.Ctx) error {
	fetchedUser := ctx.Locals(config.ContextJwtUser).(*model.User)

	childId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to convert id param"})
	}

	// parse
	newChild := new(model.Child)

	if err := ctx.BodyParser(newChild); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid child data"})
	}

	sourceChild, err := user.FindChild(database.DBConn, fetchedUser, childId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "child by id not found"})
	}

	result, err := user.UpdateChild(database.DBConn, sourceChild, newChild)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error to update child"})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
