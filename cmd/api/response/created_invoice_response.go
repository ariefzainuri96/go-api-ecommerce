package response

import "encoding/json"

func UnmarshalCreatedInvoiceResp(data []byte) (CreatedInvoiceResp, error) {
	var r CreatedInvoiceResp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r CreatedInvoiceResp) Marshal() ([]byte, error) {
	marshal, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	return marshal, nil
}

func (r *CreatedInvoiceResp) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &r)
}

type CreatedInvoiceResp struct {
	ID                        string                  `json:"id"`
	ExternalID                string                  `json:"external_id"`
	UserID                    string                  `json:"user_id"`
	Status                    string                  `json:"status"`
	MerchantName              string                  `json:"merchant_name"`
	MerchantProfilePictureURL string                  `json:"merchant_profile_picture_url"`
	Amount                    int64                   `json:"amount"`
	PayerEmail                string                  `json:"payer_email"`
	Description               string                  `json:"description"`
	ExpiryDate                string                  `json:"expiry_date"`
	InvoiceURL                string                  `json:"invoice_url"`
	AvailableBanks            []AvailableBank         `json:"available_banks"`
	AvailableRetailOutlets    []AvailableRetailOutlet `json:"available_retail_outlets"`
	AvailableEwallets         []AvailableEwallet      `json:"available_ewallets"`
	AvailableQrCodes          []AvailableQrCode       `json:"available_qr_codes"`
	AvailableDirectDebits     []AvailableDirectDebit  `json:"available_direct_debits"`
	AvailablePaylaters        []AvailablePaylater     `json:"available_paylaters"`
	ShouldExcludeCreditCard   bool                    `json:"should_exclude_credit_card"`
	ShouldSendEmail           bool                    `json:"should_send_email"`
	Created                   string                  `json:"created"`
	Updated                   string                  `json:"updated"`
	Currency                  string                  `json:"currency"`
}

type AvailableBank struct {
	BankCode          string `json:"bank_code"`
	CollectionType    string `json:"collection_type"`
	TransferAmount    int64  `json:"transfer_amount"`
	BankBranch        string `json:"bank_branch"`
	AccountHolderName string `json:"account_holder_name"`
	IdentityAmount    int64  `json:"identity_amount"`
}

type AvailableDirectDebit struct {
	DirectDebitType string `json:"direct_debit_type"`
}

type AvailableEwallet struct {
	EwalletType string `json:"ewallet_type"`
}

type AvailablePaylater struct {
	PaylaterType string `json:"paylater_type"`
}

type AvailableQrCode struct {
	QrCodeType string `json:"qr_code_type"`
}

type AvailableRetailOutlet struct {
	RetailOutletName string `json:"retail_outlet_name"`
}
