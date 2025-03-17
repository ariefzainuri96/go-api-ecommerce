package xendit

import (
	"encoding/json"
)

type FVACreatedUpdatedReq struct {
	IsClosed       bool   `json:"is_closed"`
	Status         string `json:"status"`
	Currency       string `json:"currency"`
	OwnerId        string `json:"owner_id"`
	ExternalId     string `json:"external_id"`
	BankCode       string `json:"bank_code"`
	MerchantCode   string `json:"merchant_code"`
	Name           string `json:"name"`
	AccountNumber  string `json:"account_number"`
	IsSingleUse    bool   `json:"is_single_use"`
	ExpirationDate string `json:"expiration_date"`
	Id             string `json:"id"`
}

func (r FVACreatedUpdatedReq) Marshal() ([]byte, error) {
	marshal, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	return marshal, nil
}

func (r *FVACreatedUpdatedReq) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &r)
}
