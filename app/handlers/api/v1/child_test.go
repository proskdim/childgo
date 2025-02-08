package handler_test

import (
	"bytes"
	"childgo/app"
	model "childgo/app/models"
	"childgo/app/models/repo"
	"childgo/app/types"
	storage "childgo/config/database"
	"childgo/utils/password"
	"childgo/utils/uuidv7"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	UserEmail    = "user@mail.ru"
	UserPassword = "qwerty123"
)

func TestGetChild(t *testing.T) {
	s := storage.Storage

	clear(s.DB)

	app := app.StartupTest(s)

	user, err := createUser()

	if err != nil {
		t.Fatal(err)
	}

	req := &types.SigninRequest{Email: UserEmail, Password: UserPassword}

	token, err := fetchToken(app, req)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEqual(t, "", token)

	uuid, _ := uuidv7.Generate()

	child := &model.Child{
		ID:       *uuid,
		UserID:   user.ID,
		Name:     "test",
		Age:      10,
		Birthday: time.Now(),
	}

	repo.CreateChild(child)

	tests := []struct {
		desc       string
		childId    string
		token      string
		statusCode int
	}{
		{
			desc:       "return child (200)",
			childId:    child.ID.String(),
			token:      token,
			statusCode: 200,
		},
		{
			desc:       "return BadRequest code",
			childId:    "otherid",
			token:      token,
			statusCode: 400,
		},
		{
			desc:       "return Unauthorized code",
			childId:    child.ID.String(),
			token:      "token",
			statusCode: 401,
		},
	}

	for _, test := range tests {
		url := fmt.Sprintf("/api/v1/child/%s", test.childId)

		req := httptest.NewRequest("GET", url, nil)

		req.Header.Set("authorization", fmt.Sprintf("Bearer %v", test.token))
		req.Header.Set("content-type", "application/json")

		resp, err := app.Test(req, -1)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, test.statusCode, resp.StatusCode)
	}
}

func TestNewChild(t *testing.T) {
	s := storage.Storage

	clear(s.DB)

	app := app.StartupTest(s)

	_, err := createUser()

	if err != nil {
		t.Fatal(err)
	}

	req := &types.SigninRequest{Email: UserEmail, Password: UserPassword}

	token, err := fetchToken(app, req)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotEqual(t, "", token)

	tests := []struct {
		desc       string
		child      *types.ChildRequest
		token      string
		statusCode int
	}{
		{
			desc:       "return new child (200)",
			child:      &types.ChildRequest{Name: "test", Age: 10, Birthday: time.Now()},
			token:      token,
			statusCode: 200,
		},
		{
			desc:       "failed request (UnprocessableEntity)",
			child:      &types.ChildRequest{Name: "", Age: 0, Birthday: time.Now()},
			token:      token,
			statusCode: 422,
		},
		{
			desc:       "failed request (UnprocessableEntity)",
			child:      &types.ChildRequest{Name: "test", Age: 10},
			token:      token,
			statusCode: 422,
		},
		{
			desc:       "failed request without token (BadRequest)",
			child:      &types.ChildRequest{Name: "test", Age: 10, Birthday: time.Now()},
			token:      "",
			statusCode: 400,
		},
	}

	for _, test := range tests {
		buf := new(bytes.Buffer)

		json.NewEncoder(buf).Encode(test.child)

		req := httptest.NewRequest("POST", "/api/v1/child", buf)

		req.Header.Set("authorization", fmt.Sprintf("Bearer %v", test.token))
		req.Header.Set("content-type", "application/json")

		resp, err := app.Test(req, -1)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, test.statusCode, resp.StatusCode)

		if resp.StatusCode == 200 {
			pr := new(types.ChildResponse)

			json.NewDecoder(resp.Body).Decode(pr)

			assert.Equal(t, test.child.Name, pr.Name)
			assert.Equal(t, test.child.Age, pr.Age)
			assert.Equal(t, test.child.Birthday.GoString(), pr.Birthday.GoString())
		}
	}
}

func fetchToken(app *fiber.App, req *types.SigninRequest) (string, error) {
	resp, err := signin(app, req)

	var token string

	if err != nil {
		return token, err
	}

	if resp.StatusCode != 200 {
		return token, err
	}

	res := new(types.SigninResponse)

	json.NewDecoder(resp.Body).Decode(res)

	return res.JWTToken, nil
}

func createUser() (*model.User, error) {
	uuid, _ := uuidv7.Generate()

	m := &model.User{
		ID:       *uuid,
		Email:    UserEmail,
		Password: password.Generate(UserPassword),
	}

	if err := repo.CreateUser(m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func clear(db *gorm.DB) {
	repo.DeleteUsers(db)
	repo.DeleteChilds(db)
}
