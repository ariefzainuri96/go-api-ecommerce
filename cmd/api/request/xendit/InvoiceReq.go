package xendit

import (
	"encoding/json"
)

func UnmarshalInvoiceReq(data []byte) (InvoiceReq, error) {
	var r InvoiceReq
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InvoiceReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type InvoiceReq struct {
	ID                     string `json:"id"`
	ExternalID             string `json:"external_id"`
	UserID                 string `json:"user_id"`
	IsHigh                 bool   `json:"is_high"`
	PaymentMethod          string `json:"payment_method"`
	Status                 string `json:"status"`
	MerchantName           string `json:"merchant_name"`
	Amount                 int64  `json:"amount"`
	PaidAmount             int64  `json:"paid_amount"`
	BankCode               string `json:"bank_code"`
	PaidAt                 string `json:"paid_at"`
	PayerEmail             string `json:"payer_email"`
	Description            string `json:"description"`
	AdjustedReceivedAmount int64  `json:"adjusted_received_amount"`
	FeesPaidAmount         int64  `json:"fees_paid_amount"`
	Updated                string `json:"updated"`
	Created                string `json:"created"`
	Currency               string `json:"currency"`
	PaymentChannel         string `json:"payment_channel"`
	PaymentDestination     string `json:"payment_destination"`
}
