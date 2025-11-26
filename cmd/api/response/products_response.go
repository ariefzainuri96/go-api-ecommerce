package response

import (
	en "github.com/ariefzainuri96/go-api-ecommerce/cmd/api/entity"
)

// @Model
type ProductsResponse struct {
	BaseResponse
	Products []en.Product `json:"products"`
	Pagination PaginationMetadata `json:"pagination"`
}

type ProductResponse struct {
	BaseResponse
	Product en.Product `json:"product"`
}