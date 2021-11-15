package routes

import (
	"github.com/sachinsmc/housing-anywhere/config"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUsersRoute(t *testing.T) {

	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			description:  "get HTTP status 200",
			route:        "/api/v1/dns/location",
			expectedCode: 200,
		},
	}

	config.Init()

	app := fiber.New()

	Setup(app)

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)
		resp, _ := app.Test(req, 100)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
