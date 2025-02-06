package handler

import (
	model "childgo/app/models"
	"childgo/app/models/repo"
	"childgo/config/database"
	"childgo/utils"
	"childgo/utils/pagination"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	badRequest = fiber.StatusBadRequest
	notFound   = fiber.StatusNotFound
	statusOk   = fiber.StatusOK
	conflict   = fiber.StatusConflict
)

// get all childs for current user
func Childs(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Query("page", "1"))

	if err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "cannot parse page param"})
	}

	u := utils.GetUser(ctx)

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
	c := new(model.Child)
	u := utils.GetUser(ctx)

	if err := utils.ParseBody(ctx, c); err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{"error": "invalid child data"})
	}

	c.UserID = u.ID

	if err := repo.CreateChild(c).Error; err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "error to add child"})
	}

	return ctx.JSON(c)
}

// get child by id
func GetChild(ctx *fiber.Ctx) error {
	c := new(model.Child)
	u := utils.GetUser(ctx)

	id, err := utils.ParseID(ctx)

	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{"error": "failed read id param"})
	}

	if err := repo.FindChildByUser(c, id, u.ID).Error; err != nil {
		return ctx.Status(notFound).JSON(fiber.Map{"error": "child not found"})
	}

	return ctx.JSON(c)
}

// delete child for current user
func DeleteChild(ctx *fiber.Ctx) error {
	u := utils.GetUser(ctx)

	id, err := utils.ParseID(ctx)

	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{"error": "failed read id param"})
	}

	res := repo.DeleteChild(id, u.ID)

	if res.RowsAffected == 0 {
		return ctx.Status(conflict).JSON(fiber.Map{"error": "unable to delete child"})
	}

	if res.Error != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "error do delete child"})
	}

	return ctx.SendStatus(statusOk)
}

// update child for current user
func UpdateChild(ctx *fiber.Ctx) error {
	c := new(model.Child)
	u := utils.GetUser(ctx)

	id, err := utils.ParseID(ctx)

	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{"error": "failed read id param"})
	}

	if err := utils.ParseBody(ctx, c); err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{"error": "invalid child data"})
	}

	if err := repo.UpdateChild(id, u.ID, c).Error; err != nil {
		return ctx.Status(badRequest).JSON(fiber.Map{"error": "error to update child"})
	}

	return ctx.Status(statusOk).JSON(c)
}
