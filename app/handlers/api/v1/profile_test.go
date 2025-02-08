package handler_test

import (
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

	"github.com/stretchr/testify/assert"
)

var (
	ProfileEmail    = "profile@mail.ru"
	ProfilePassword = "qwerty12345"
)

func TestProfile(t *testing.T) {
	s := storage.Storage

	repo.DeleteUsers(s.DB)

	app := app.StartupTest(s)
	uuid, _ := uuidv7.Generate()

	m := &model.User{
		ID:       *uuid,
		Email:    ProfileEmail,
		Password: password.Generate(ProfilePassword),
	}

	repo.CreateUser(m)

	r := &types.SigninRequest{Email: ProfileEmail, Password: ProfilePassword}

	resp, err := signin(app, r)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatal(resp.StatusCode)
	}

	res := new(types.SigninResponse)

	json.NewDecoder(resp.Body).Decode(res)

	tests := []struct {
		desc         string
		token        string
		statusCode   int
		expectedBody string
	}{
		{
			desc:         "return user email (200)",
			token:        res.JWTToken,
			statusCode:   200,
			expectedBody: r.Email,
		},
		{
			desc: "return unauthorized (401)",
			token: "other token",
			statusCode: 401,
			expectedBody: "",
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", "/api/v1/profile", nil)

		req.Header.Set("authorization", fmt.Sprintf("Bearer %v", test.token))
		req.Header.Set("content-type", "application/json")

		resp, err := app.Test(req, -1)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, test.statusCode, resp.StatusCode)

		pr := new(types.ProfileResponse)

		json.NewDecoder(resp.Body).Decode(pr)

		assert.Equal(t, test.expectedBody, pr.Email)
	}
}
