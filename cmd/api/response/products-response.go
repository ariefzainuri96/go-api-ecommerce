package response

import "encoding/json"

type ProductsResponse struct {
	BaseResponse
	Products []Product `json:"products"`
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

type Product struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
	CreatedAt   string `json:"created_at"`
}
