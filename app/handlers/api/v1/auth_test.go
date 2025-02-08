package handler_test

import (
	"bytes"
	"childgo/app"
	"childgo/app/models/repo"
	"childgo/app/types"
	storage "childgo/config/database"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var (
	TestEmail    = "user@mail.ru"
	TestPassword = "qwerty123"
)

type signUpRequest = types.SignupRequest
type signInRequest = types.SigninRequest

func TestSignup(t *testing.T) {
	s := storage.Storage
	app := app.StartupTest(s)

	repo.DeleteUsers(s.DB)

	validUser := &signUpRequest{Email: TestEmail, Password: TestPassword}

	tests := []struct {
		desc         string
		user         *signUpRequest
		statusCode   int
		expectedBody string
	}{
		{
			desc:         "success result (200)",
			user:         validUser,
			statusCode:   200,
			expectedBody: validUser.Email,
		},
		{
			desc:         "failed result (UnprocessableEntity)",
			user:         &signUpRequest{Email: "", Password: ""},
			statusCode:   422,
			expectedBody: "",
		},
		{
			desc:         "failed result (UnprocessableEntity)",
			user:         &signUpRequest{Email: TestEmail, Password: ""},
			statusCode:   422,
			expectedBody: "",
		},
		{
			desc:         "failed result (UnprocessableEntity)",
			user:         &signUpRequest{Email: "", Password: TestPassword},
			statusCode:   422,
			expectedBody: "",
		},
	}

	for _, test := range tests {
		resp, err := signup(app, test.user)
		res := new(types.SignupResponse)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, test.statusCode, resp.StatusCode)

		json.NewDecoder(resp.Body).Decode(res)

		assert.Equal(t, test.expectedBody, res.Email)
	}
}

func TestSignin(t *testing.T) {
	s := storage.Storage
	app := app.StartupTest(s)

	repo.DeleteUsers(s.DB)

	tests := []struct {
		desc        string
		user        *signInRequest
		signupUser  *signUpRequest
		statusCode  int
		createUser  bool
		expectedJWT bool
	}{
		{
			desc:        "success result (200)",
			user:        &signInRequest{Email: TestEmail, Password: TestPassword},
			signupUser:  &signUpRequest{Email: TestEmail, Password: TestPassword},
			statusCode:  200,
			createUser:  true,
			expectedJWT: true,
		},
		{
			desc:        "failed result (NotFound)",
			user:        &signInRequest{Email: "other@mail.ru", Password: "password"},
			signupUser:  &signUpRequest{Email: TestEmail, Password: TestPassword},
			statusCode:  404,
			createUser:  true,
			expectedJWT: false,
		},
		{
			desc:        "failed result (Unauthorized)",
			user:        &signInRequest{Email: TestEmail, Password: "other password"},
			signupUser:  &signUpRequest{Email: TestEmail, Password: TestPassword},
			statusCode:  401,
			createUser:  true,
			expectedJWT: false,
		},
	}

	for _, test := range tests {
		if test.createUser {
			resp, err := signup(app, test.signupUser)

			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != 200 {
				t.Fatal(resp.StatusCode)
			}
		}

		resp, err := signin(app, test.user)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, test.statusCode, resp.StatusCode)

		res := new(types.SigninResponse)

		if test.expectedJWT {
			json.NewDecoder(resp.Body).Decode(res)
			assert.NotEqual(t, "", res.JWTToken)
		}

		repo.DeleteUsers(s.DB)
	}
}

// signup is helper test function
func signup(app *fiber.App, user *types.SignupRequest) (*http.Response, error) {
	buf := new(bytes.Buffer)

	json.NewEncoder(buf).Encode(user)

	req := jsonReq("/api/v1/signup", buf)

	return app.Test(req, -1)
}

// signin is helper test function
func signin(app *fiber.App, user *types.SigninRequest) (*http.Response, error) {
	buf := new(bytes.Buffer)

	json.NewEncoder(buf).Encode(user)

	req := jsonReq("/api/v1/signin", buf)

	return app.Test(req, -1)
}

// jsonReq is helper function
func jsonReq(url string, buf *bytes.Buffer) *http.Request {
	req := httptest.NewRequest("POST", url, buf)

	req.Header.Set("content-type", "application/json")

	return req
}
