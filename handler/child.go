package handler

import (
	"childgo/config"
	"childgo/config/database"
	"childgo/model"
	"childgo/model/user"
	"childgo/utils/pagination"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	badRequest = fiber.StatusBadRequest
	notFound = fiber.StatusNotFound
	statusOk = fiber.StatusOK
)

// get all childs for current user
func Childs(ctx *fiber.Ctx) error {
	db := database.DBConn
	page, err := strconv.Atoi(ctx.Query("page", "1"))

	if err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "cannot parse page param"})
	}

	fetchedUser := ctx.Locals(config.ContextJwtUser).(*model.User)

	childs, err := user.FindAllChilds(db, fetchedUser)

	pagy := pagination.Paginate(&pagination.Option{
		DB:      db,
		Page:    page,
		Limit:   10,
		ShowSQL: true,
	}, &childs)

	if err != nil {
		return ctx.Status(notFound).JSON(fiber.Map{"error": "error to fetch all childs"})
	}

	return ctx.JSON(pagy)
}

// add new child
func NewChild(ctx *fiber.Ctx) error {
	fetchedUser := ctx.Locals(config.ContextJwtUser).(*model.User)

	child := new(model.Child)

	if err := ctx.BodyParser(child); err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "invalid child data"})
	}

	dbChild, err := user.AddChild(database.DBConn, fetchedUser, child)

	if err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "error to add child"})
	}

	return ctx.JSON(dbChild)
}

// get child by id
func GetChild(ctx *fiber.Ctx) error {
	fetchedUser := ctx.Locals(config.ContextJwtUser).(*model.User)

	childId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "failed to convert id param"})
	}

	child, err := user.FindChild(database.DBConn, fetchedUser, childId)

	if err != nil {
		return ctx.Status(notFound).JSON(fiber.Map{"error": "child not found"})
	}

	return ctx.JSON(child)
}

// delete child for current user
func DeleteChild(ctx *fiber.Ctx) error {
	db := database.DBConn
	fetchedUser := ctx.Locals(config.ContextJwtUser).(*model.User)

	childId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "failed to convert id param"})
	}

	child, err := user.FindChild(db, fetchedUser, childId)

	if err != nil {
		return ctx.Status(notFound).JSON(fiber.Map{"error": "child by id not found"})
	}

	if _, err := user.DeleteChild(db, child); err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "error to delete child"})
	}

	return ctx.SendStatus(statusOk)
}

func UpdateChild(ctx *fiber.Ctx) error {
	db := database.DBConn
	fetchedUser := ctx.Locals(config.ContextJwtUser).(*model.User)

	childId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "failed to convert id param"})
	}

	// parse
	newChild := new(model.Child)

	if err := ctx.BodyParser(newChild); err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "invalid child data"})
	}

	sourceChild, err := user.FindChild(db, fetchedUser, childId)

	if err != nil {
		return ctx.Status(notFound).JSON(fiber.Map{"error": "child by id not found"})
	}

	result, err := user.UpdateChild(db, sourceChild, newChild)

	if err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "error to update child"})
	}

	return ctx.Status(statusOk).JSON(result)
}
