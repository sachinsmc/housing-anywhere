package routes

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sachinsmc/housing-anywhere/controllers"
	"github.com/sachinsmc/housing-anywhere/models"
	"reflect"
	"strconv"
)

var validate *validator.Validate

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
// @Failure 400 {object} HTTPError
// @Router /api/v1/dns/location [post]
func CalculateLocation(c *fiber.Ctx) error {
	d := new(models.DNSRequest)
	//fmt.Println("Body :", string(c.Body()))
	if err := c.BodyParser(d); err != nil {
		fmt.Println("Failed to read body :", err)
		return c.Status(500).JSON(&HTTPError{
			Status:  "Internal Server Error",
			Message: "Failed to read the body",
		})
	}
	// validate request
	validate = validator.New()
	err := validate.Struct(d)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}

		for _, err := range err.(validator.ValidationErrors) {
			return c.Status(400).JSON(&HTTPError{
				Status:  "Bad Request",
				Message: err.StructField() + " is required or invalid",
			})
		}
	}
	// check if all values are float type
	r := reflect.ValueOf(d).Elem()
	rt := r.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		rv := reflect.ValueOf(d)
		value := reflect.Indirect(rv).FieldByName(field.Name)
		_, err := strconv.ParseFloat(value.String(), 64)
		if err != nil {
			return c.Status(400).JSON(&HTTPError{
				Status:  "Bad Request",
				Message: field.Name + " must be floating number, all inputs must be floating type",
			})
		}
	}

	return c.JSON(controllers.GetLocation(d))
}
