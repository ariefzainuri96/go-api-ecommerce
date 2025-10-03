package response

import (
	"encoding/json"

	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/entity"
)

type CartsResponse struct {
	BaseResponse
	Carts []entity.Cart `json:"carts"`
}

func (r CartsResponse) MarshalResponse() ([]byte, error) {
	marshal, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	return marshal, nil
}

func (r *CartsResponse) UnmarshalResponse(data []byte) error {
	return json.Unmarshal(data, &r)
}

type CartResponse struct {
	entity.BaseEntity
	Product  entity.Product `json:"product"`
	Quantity int            `json:"quantity"`
}
