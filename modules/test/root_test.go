package test

import (
	"bytes"
	"io"
	"net/http/httptest"
	"os"
	"testing"

	"contacts_api/modules/database"
	"contacts_api/modules/rest"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func initServer() *fiber.App {
	db, err := database.Init("test.db")
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	rest.Setup(app, db)
	return app
}

func cleanDB() {
	os.Remove("test.db")
}

func TestSetup(t *testing.T) {
	app := initServer()
	defer cleanDB()

	{
		req := httptest.NewRequest("GET", "/v1/users", nil)
		resp, err := app.Test(req, -1)
		b, _ := io.ReadAll(resp.Body)

		assert.Equalf(t, nil, err, "error should be nil")
		assert.Equalf(t, 200, resp.StatusCode, "status should be 200")
		assert.Equalf(t, `{"error":null,"success":{"data":{"users":[]}}}`, string(b), "no users should be returned")
	}
	{
		req := httptest.NewRequest("GET", "/v1/skills", nil)
		resp, err := app.Test(req, -1)
		b, _ := io.ReadAll(resp.Body)

		assert.Equalf(t, nil, err, "error should be nil")
		assert.Equalf(t, 200, resp.StatusCode, "status should be 200")
		assert.Equalf(t, `{"error":null,"success":{"data":{"skills":[]}}}`, string(b), "no skills should be returned")
	}
	{
		req := httptest.NewRequest("DELETE", "/v1/users/1", nil)
		resp, err := app.Test(req, -1)

		assert.Equalf(t, nil, err, "error should be nil")
		assert.Equalf(t, 403, resp.StatusCode, "not allowed to use this enpoint without token")
	}
	{
		req := httptest.NewRequest("POST", "/v1/skills", bytes.NewBuffer([]byte(`[{"name":"skill1"},{"name":"go"},{"name":"python"}]`)))
		resp, err := app.Test(req, -1)

		assert.Equalf(t, nil, err, "error should be nil")
		assert.Equalf(t, 200, resp.StatusCode, "status should be 200")
	}
	{
		req := httptest.NewRequest("GET", "/v1/skills", nil)
		resp, err := app.Test(req, -1)
		b, _ := io.ReadAll(resp.Body)

		assert.Equalf(t, nil, err, "error should be nil")
		assert.Equalf(t, 200, resp.StatusCode, "status should be 200")
		assert.Equalf(t, `{"error":null,"success":{"data":{"skills":[{"name":"skill1"},{"name":"go"},{"name":"python"}]}}}`, string(b), "list of created skills should be returned")
	}
	{
		req := httptest.NewRequest("POST", "/v1/users", bytes.NewBuffer([]byte(`{"name": "mr","surname": "me","address": "my adress 44","email": "test2@test.test","phone": "+41749999991"}`)))
		resp, err := app.Test(req, -1)

		assert.Equalf(t, nil, err, "error should be nil")
		assert.Equalf(t, 200, resp.StatusCode, "status should be 200")
	}
	{
		req := httptest.NewRequest("GET", "/v1/users", nil)
		resp, err := app.Test(req, -1)
		b, _ := io.ReadAll(resp.Body)

		assert.Equalf(t, nil, err, "error should be nil")
		assert.Equalf(t, 200, resp.StatusCode, "status should be 200")
		assert.Equalf(t, `{"error":null,"success":{"data":{"users":[{"id":1,"name":"mr","surname":"me","address":"my adress 44","email":"test2@test.test","phone":"+41749999991","skills":null}]}}}`, string(b), "one created user should be returned")
	}
}
