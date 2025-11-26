package request

import (
	_ "github.com/go-playground/validator/v10"
)

type AddProductRequest struct {
	Name        string  `json:"name" validate:"required,max=255"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0,lte=9999999999"`
	Quantity    int     `json:"quantity" validate:"required,gt=0,lte=9999"`
}
