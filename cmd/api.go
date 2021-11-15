package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sachinsmc/housing-anywhere/config"
	"github.com/sachinsmc/housing-anywhere/routes"
	"github.com/spf13/viper"
)

func Run() {
	config.Init()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":" + viper.GetString("server.port"))
}
