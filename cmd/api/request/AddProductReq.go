package request

import (
	_ "github.com/go-playground/validator/v10"
)

type AddProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int64  `json:"price" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
}
