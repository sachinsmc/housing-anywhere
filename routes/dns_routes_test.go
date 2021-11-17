package routes

import (
	"bytes"
	"github.com/sachinsmc/housing-anywhere/config"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {

	tests := []struct {
		description  string
		route        string
		expectedCode int
		body         []byte
		location     []byte
	}{
		{
			description:  "get HTTP status 200",
			route:        "/api/v1/dns/location",
			expectedCode: 200,
			body:         []byte(`{"vel":"20.0","x":"123.12","y":"456.56","z":"789.89"}`),
		},
		{
			description:  "get HTTP status 500",
			route:        "/api/v1/dns/location",
			expectedCode: 500,
			body:         []byte(`"vel":"20.0","x":"123.12","y":"456.56","z":"789.89"`),
		},
		{
			description:  "get HTTP status 400",
			route:        "/api/v1/dns/location",
			expectedCode: 400,
			body:         []byte(`{"vel":"string","x":"123.12","y":"456.56","z":"789.89"}`),
		},
		{
			description:  "get HTTP status 400",
			route:        "/api/v1/dns/location",
			expectedCode: 400,
			body:         []byte(`{"x":"123.12","y":"456.56","z":"789.89"}`),
		},
	}

	config.Init()

	app := fiber.New()

	Setup(app)
	for _, test := range tests {
		req := httptest.NewRequest("POST", test.route, bytes.NewBuffer(test.body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req, 500)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
