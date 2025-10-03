package response

import (
	"encoding/json"

	en "github.com/ariefzainuri96/go-api-ecommerce/cmd/api/entity"
)

// @Model
type ProductsResponse struct {
	BaseResponse
	Products []en.Product `json:"products"`
}

func (r ProductsResponse) MarshalProductsResponse() ([]byte, error) {
	marshal, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	return marshal, nil
}

func (r *ProductsResponse) UnmarshalProductsResponse(data []byte) error {
	return json.Unmarshal(data, &r)
}
