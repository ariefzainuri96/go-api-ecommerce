package response

import (
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/entity"
)

type CartsResponse struct {
	BaseResponse
	Carts      []entity.Cart      `json:"carts"`
	Pagination PaginationMetadata `json:"pagination"`
}

type CartResponse struct {
	BaseResponse
	Cart entity.Cart `json:"cart"`
}
