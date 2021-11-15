package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	_ "github.com/sachinsmc/housing-anywhere/docs"
	"github.com/sachinsmc/housing-anywhere/middlewares"
)

// Setup
// @title HA API
// @version 1.0
// @description This is a swagger doc for Housing Anywhere task
// @contact.name API Support
// @contact.email hey@sachinsmc.me
// @host localhost:3003
// @BasePath /
func Setup(app *fiber.App) {

	middlewares.CORS(app)

	app.Get("/", monitor.New())

	app.Use(etag.New())
	app.Use(logger.New())

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))

	api := app.Group("/api")

	v1 := api.Group("/v1")

	dns := v1.Group("/dns")

	dns.Post("/location", CalculateLocation)

}

type HTTPError struct {
	Status  string
	Message string
}
