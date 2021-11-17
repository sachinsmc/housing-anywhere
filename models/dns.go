package models

import _ "github.com/go-playground/validator"

type DNSRequest struct {
	X        string `json:"x" validate:"required"`
	Y        string `json:"y" validate:"required"`
	Z        string `json:"z" validate:"required"`
	Velocity string `json:"vel" validate:"required"`
}

type DNSResponse struct {
	Location string `json:"loc"`
}
