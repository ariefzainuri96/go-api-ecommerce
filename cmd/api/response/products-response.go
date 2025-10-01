package response

import (
	"encoding/json"

	"github.com/ariefzainuri96/go-api-ecommerce/internal/data"
)

type ProductsResponse struct {
	BaseResponse
	Products []data.Product `json:"products"`
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
