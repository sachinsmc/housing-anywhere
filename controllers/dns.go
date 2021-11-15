package controllers

import (
	"fmt"
	"github.com/sachinsmc/housing-anywhere/models"
	"github.com/spf13/viper"
	"strconv"
)

func GetLocation(d *models.DNSRequest) *models.DNSResponse {
	sectorId := viper.GetFloat64("sector.id")
	x, _ := strconv.ParseFloat(d.X, 64)
	y, _ := strconv.ParseFloat(d.Y, 64)
	z, _ := strconv.ParseFloat(d.Z, 64)
	vel, _ := strconv.ParseFloat(d.Velocity, 64)
	loc := x*sectorId + y*sectorId + z*sectorId + vel
	return &models.DNSResponse{Location: fmt.Sprintf("%.2f", loc)}
}
