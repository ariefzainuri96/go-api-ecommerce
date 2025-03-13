package request

type AddToCartRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
	UserID    int `json:"user_id"`
}
