package handler

import (
	model "childgo/app/models"
	"childgo/app/models/repo"
	"childgo/app/types"
	storage "childgo/config/database"
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

	childs := []types.Child{}

	pagy := pagination.Paginate(&pagination.Option{
		DB:      storage.Storage.DB,
		Model:   &model.Child{},
		Page:    page,
		Limit:   10,
		ShowSQL: true,
		Conds:   []any{"user_id", u.ID},
	}, &childs)

	return ctx.JSON(pagy)
}

// add new child
func NewChild(ctx *fiber.Ctx) error {
	c := new(model.Child)
	u := utils.GetUser(ctx)

	if err := utils.ParseBodyValidator(ctx, c); err != nil {
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

	address := types.AddressResponse{
		City:      c.Address.City,
		Street:    c.Address.Street,
		House:     c.Address.House,
		Apartment: c.Address.Apartment,
	}

	return ctx.JSON(&types.ChildCreateResponse{
		Child: &types.ChildResponse{
			ID:       c.ID,
			Name:     c.Name,
			Age:      c.Age,
			Birthday: c.Birthday,
			Address:  address,
		}})
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

	address := types.AddressResponse{
		City:      c.Address.City,
		Street:    c.Address.Street,
		House:     c.Address.House,
		Apartment: c.Address.Apartment,
	}

	return ctx.JSON(&types.ChildCreateResponse{
		Child: &types.ChildResponse{
			ID:       c.ID,
			Name:     c.Name,
			Age:      c.Age,
			Birthday: c.Birthday,
			Address:  address,
		}})
}

// delete child for current user
func DeleteChild(ctx *fiber.Ctx) error {
	u := utils.GetUser(ctx)

	id, err := utils.ParseID(ctx)

	if err != nil {
		return ctx.Status(err.Code).JSON(utils.MapErr("failed read id param"))
	}

	if err := repo.DeleteChild(id, u.ID); err != nil {
		return ctx.Status(badRequest).JSON(utils.MapErr(err.Error()))
	}

	return ctx.JSON(&types.MsgResp{
		Message: "Child successfully deleted",
	})
}

// update child for current user
func UpdateChild(ctx *fiber.Ctx) error {
	c := new(model.Child)
	u := utils.GetUser(ctx)

	id, err := utils.ParseID(ctx)

	if err != nil {
		return ctx.Status(err.Code).JSON(utils.MapErr("failed read id param"))
	}

	if err := utils.ParseBodyValidator(ctx, c); err != nil {
		return ctx.Status(err.Code).JSON(utils.MapErr("invalid child data"))
	}

	if err := repo.UpdateChild(id, u.ID, c); err != nil {
		return ctx.Status(badRequest).JSON(utils.MapErr(err.Error()))
	}

	return ctx.JSON(types.MsgResp{
		Message: "Child successfull updated",
	})
}
