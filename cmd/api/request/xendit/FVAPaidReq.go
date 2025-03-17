package xendit

import "encoding/json"

type FVAPaidReq struct {
	Updated                  string `json:"updated"`
	Created                  string `json:"created"`
	PaymentID                string `json:"payment_id"`
	CallbackVirtualAccountID string `json:"callback_virtual_account_id"`
	OwnerID                  string `json:"owner_id"`
	ExternalID               string `json:"external_id"`
	AccountNumber            string `json:"account_number"`
	BankCode                 string `json:"bank_code"`
	Amount                   int64  `json:"amount"`
	TransactionTimestamp     string `json:"transaction_timestamp"`
	MerchantCode             string `json:"merchant_code"`
	ID                       string `json:"id"`
}

func UnmarshalFVAPaidReq(data []byte) (FVAPaidReq, error) {
	var r FVAPaidReq
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FVAPaidReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
