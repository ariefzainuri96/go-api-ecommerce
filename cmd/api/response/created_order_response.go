package response

import "encoding/json"

type CreatedOrderResp struct {
	BaseResponse
	Data CreatedOrderData `json:"data"`
}

type CreatedOrderData struct {
	InvoiceUrl string `json:"invoice_url"`
	InvoiceID  string `json:"id"`
	Status     string `json:"status"`
	ExpiryDate string `json:"expiry_date"`
}

func (r CreatedOrderResp) Marshal() ([]byte, error) {
	marshal, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	return marshal, nil
}

func (r *CreatedOrderResp) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &r)
}
