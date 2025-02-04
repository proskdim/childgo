package handler

import (
	"childgo/config"
	"childgo/database"
	"childgo/model"
	"childgo/model/user"
	jwtUtil "childgo/utils"
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

func fetchUserFromPayload(ctx *fiber.Ctx) (*model.User, error) {
	token, err := fetchToken(ctx)

	if err != nil {
		return nil, err
	}

	user, err := fetchUser(token)

	if err != nil {
		return nil, err
	}

	return user, err
}

func fetchUser(token *jwt.Token) (*model.User, error) {
	db := database.DBConn
	email, err := jwtUtil.FindEmailFromToken(token)

	if err != nil {
		return nil, err
	}

	dbUser, err := user.FindByEmail(db, email)

	if err != nil {
		return nil, err
	}

	return dbUser, nil
}

func fetchToken(ctx *fiber.Ctx) (token *jwt.Token, err error) {
	token, err = checkJwt(ctx)

	if err != nil {
		return nil, err
	}

	return token, err
}

// get all childs for current user
func Childs(ctx *fiber.Ctx) error {
	fetchedUser, err := fetchUserFromPayload(ctx)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	childs, err := user.FindAllChilds(database.DBConn, fetchedUser)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return ctx.JSON(childs)
}

// add new child
func NewChild(ctx *fiber.Ctx) error {
	fetchedUser, err := fetchUserFromPayload(ctx)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	child := new(model.Child)

	if err := ctx.BodyParser(child); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	dbChild, err := user.AddChild(database.DBConn, fetchedUser, child)

	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(dbChild)
}

// get child by id
func GetChild(ctx *fiber.Ctx) error {
	fetchedUser, err := fetchUserFromPayload(ctx)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	childId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	child, err := user.FindChild(database.DBConn, fetchedUser, childId)

	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return ctx.JSON(child)
}

// delete child for current user
func DeleteChild(ctx *fiber.Ctx) error {
	fetchedUser, err := fetchUserFromPayload(ctx)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	childId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	child, err := user.FindChild(database.DBConn, fetchedUser, childId)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "child by id not found"})
	}

	if _, err := user.DeleteChild(database.DBConn, child); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func UpdateChild(ctx *fiber.Ctx) error {
	fetchedUser, err := fetchUserFromPayload(ctx)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	childId, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
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
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
