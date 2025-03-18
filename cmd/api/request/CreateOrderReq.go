package request

type CreateInvoiceRequest struct {
	ExternalID  string `json:"external_id" validate:"required"`
	Amount      int64  `json:"amount" validate:"required"`
	PayerEmail  string `json:"payer_email" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateOrderRequest struct {
	ProductId int64 `json:"product_id" validate:"required"`
	Quantity  int   `json:"quantity" validate:"required"`
	CreateInvoiceRequest
}
