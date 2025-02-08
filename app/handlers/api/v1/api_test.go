package handler_test

import (
	"childgo/app"
	"childgo/app/types"
	storage "childgo/config/database"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	app := app.StartupTest(storage.Storage)

	req := httptest.NewRequest("GET", "/api/v1", nil)

	resp, _ := app.Test(req, -1)

	assert.Equal(t, 200, resp.StatusCode)

	h := new(types.HealthResp)

	json.NewDecoder(resp.Body).Decode(h)

	assert.Equal(t, "ok", h.Status)
}