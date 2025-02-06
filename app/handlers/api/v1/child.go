package handler

import (
	model "childgo/app/models"
	"childgo/app/models/repo"
	"childgo/app/types"
	"childgo/config/database"
	"childgo/utils"
	"childgo/utils/pagination"
	"childgo/utils/uuidv7"
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
		return ctx.Status(badRequest).JSON(utils.MapErr("cannot parse page param"))
	}

	u := utils.GetUser(ctx)

	chs := []model.Child{}

	if err := repo.FindChildrensByUser(&chs, u.ID).Error; err != nil {
		return ctx.Status(notFound).JSON(utils.MapErr("error to fetch all childs"))
	}

	pagy := pagination.Paginate(&pagination.Option{
		DB:      database.DBConn,
		Page:    page,
		Limit:   10,
		ShowSQL: true,
		Conds:   []any{"user_id", u.ID},
	}, &chs)

	return ctx.JSON(pagy)
}

// add new child
func NewChild(ctx *fiber.Ctx) error {
	c := new(model.Child)
	u := utils.GetUser(ctx)

	if err := utils.ParseBody(ctx, c); err != nil {
		return ctx.Status(err.Code).JSON(utils.MapErr("invalid child data"))
	}

	uuid, err := uuidv7.Generate()

	if err != nil {
		return ctx.Status(err.Code).JSON(utils.MapErr("failed generate uuid"))
	}

	c.ID = *uuid
	c.UserID = u.ID

	if err := repo.CreateChild(c).Error; err != nil {
		return ctx.Status(badRequest).JSON(utils.MapErr("error to add child"))
	}

	return ctx.JSON(c)
}

// get child by id
func GetChild(ctx *fiber.Ctx) error {
	c := new(model.Child)
	u := utils.GetUser(ctx)

	id, err := utils.ParseID(ctx)

	if err != nil {
		return ctx.Status(err.Code).JSON(utils.MapErr("failed read id param"))
	}

	if err := repo.FindChildByUser(c, id, u.ID).Error; err != nil {
		return ctx.Status(notFound).JSON(utils.MapErr("child not found"))
	}

	return ctx.JSON(c)
}

// delete child for current user
func DeleteChild(ctx *fiber.Ctx) error {
	u := utils.GetUser(ctx)

	id, err := utils.ParseID(ctx)

	if err != nil {
		return ctx.Status(err.Code).JSON(utils.MapErr("failed read id param"))
	}

	res := repo.DeleteChild(id, u.ID)

	if res.RowsAffected == 0 {
		return ctx.Status(conflict).JSON(utils.MapErr("unable to delete child"))
	}

	if res.Error != nil {
		return ctx.Status(badRequest).JSON(utils.MapErr("error do delete child"))
	}

	return ctx.SendStatus(statusOk)
}

// update child for current user
func UpdateChild(ctx *fiber.Ctx) error {
	c := new(model.Child)
	u := utils.GetUser(ctx)

	id, err := utils.ParseID(ctx)

	if err != nil {
		return ctx.Status(err.Code).JSON(utils.MapErr("failed read id param"))
	}

	if err := utils.ParseBody(ctx, c); err != nil {
		return ctx.Status(err.Code).JSON(utils.MapErr("invalid child data"))
	}

	res := repo.UpdateChild(id, u.ID, c)

	if res.Error != nil {
		return ctx.Status(badRequest).JSON(utils.MapErr("error to update child"))
	}

	if res.RowsAffected == 0 {
		return ctx.Status(notFound).JSON(utils.MapErr("child not found"))
	}

	return ctx.JSON(types.MsgResp{
		Message: "successfull updated",
	})
}
