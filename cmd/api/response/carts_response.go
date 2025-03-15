package response

import (
	"encoding/json"
)

type CartsResponse struct {
	BaseResponse
	Carts []Cart `json:"carts"`
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

type Cart struct {
	ID           int64  `json:"id"`
	ProductName  string `json:"product_name"`
	ProductPrice int64  `json:"product_price"`
	FullName     string `json:"full_name"`
	Quantity     int    `json:"quantity"`
	TotalAmount  int64  `json:"total_amount"`
}
