package handler

import (
	model "childgo/app/models"
	"childgo/app/models/repo"
	"childgo/config"
	"childgo/config/database"
	"childgo/utils/pagination"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	badRequest = fiber.StatusBadRequest
	notFound   = fiber.StatusNotFound
	statusOk   = fiber.StatusOK
)

// get all childs for current user
func Childs(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Query("page", "1"))

	if err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "cannot parse page param"})
	}

	u := ctx.Locals(config.ContextJwtUser).(*model.User)

	chs := []model.Child{}

	if err := repo.FindChildrensByUser(&chs, u.ID).Error; err != nil {
		return ctx.Status(notFound).JSON(fiber.Map{"error": "error to fetch all childs"})
	}

	pagy := pagination.Paginate(&pagination.Option{
		DB:      database.DBConn,
		Page:    page,
		Limit:   10,
		ShowSQL: true,
		Cond:    fmt.Sprintf("user_id = %v", u.ID),
	}, &chs)

	return ctx.JSON(pagy)
}

// add new child
func NewChild(ctx *fiber.Ctx) error {
	u := ctx.Locals(config.ContextJwtUser).(*model.User)

	c := &model.Child{}

	if err := ctx.BodyParser(c); err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "invalid child data"})
	}

	c.UserID = u.ID

	if err := repo.CreateChild(c).Error; err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "error to add child"})
	}

	return ctx.JSON(c)
}

// get child by id
func GetChild(ctx *fiber.Ctx) error {
	u := ctx.Locals(config.ContextJwtUser).(*model.User)

	childId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "failed to convert id param"})
	}

	c := &model.Child{}

	if err := repo.FindChildByUser(c, childId, u.ID).Error; err != nil {
		return ctx.Status(notFound).JSON(fiber.Map{"error": "child not found"})
	}

	return ctx.JSON(c)
}

// delete child for current user
func DeleteChild(ctx *fiber.Ctx) error {
	u := ctx.Locals(config.ContextJwtUser).(*model.User)

	childId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "failed to convert id param"})
	}

	res := repo.DeleteChild(childId, u.ID)

	if res.RowsAffected == 0 {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "unable to delete child"})
	}

	if res.Error != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "error do delete child"})
	}

	return ctx.SendStatus(statusOk)
}

// update child for current user
func UpdateChild(ctx *fiber.Ctx) error {
	u := ctx.Locals(config.ContextJwtUser).(*model.User)

	childId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "failed to convert id param"})
	}

	c := &model.Child{}

	if err := ctx.BodyParser(c); err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "invalid child data"})
	}

	if err := repo.UpdateChild(childId, u.ID, c).Error; err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "error to update child"})
	}

	return ctx.Status(statusOk).JSON(c)
}
