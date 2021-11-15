package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sachinsmc/housing-anywhere/controllers"
	"github.com/sachinsmc/housing-anywhere/models"
)

// CalculateLocation godoc
// @Summary Calculate Location
// @Description Calculate the location of the databank to upload gathered data
// @ID dns
// @Accept  json
// @Produce  json
// @Tags DNS
// @Param dns body models.DNSRequest true "DNS"
// @Success 200 {object} models.DNSResponse
// @Failure 500 {object} HTTPError
// @Router /api/v1/dns/location [post]
func CalculateLocation(c *fiber.Ctx) error {
	d := new(models.DNSRequest)
	if err := c.BodyParser(d); err != nil {
		fmt.Println("Failed to read body :", c.Body())
		return c.Status(500).JSON(&HTTPError{
			Status:  "Internal Server Error",
			Message: "Failed to read the body",
		})
	}
	return c.JSON(controllers.GetLocation(d))
}
