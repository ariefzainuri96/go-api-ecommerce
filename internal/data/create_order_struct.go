package data

type CreateOrderStruct struct {
	UserID         int64  `json:"user_id"`
	ProductID      int64  `json:"product_id"`
	Quantity       int    `json:"quantity"`
	TotalPrice     int64  `json:"total_price"`
	Status         string `json:"status"`
	InvoiceID      string `json:"invoice_id"`
	InvoiceURL     string `json:"invoice_url"`
	InvoiceExpDate string `json:"invoice_exp_date"`
}
