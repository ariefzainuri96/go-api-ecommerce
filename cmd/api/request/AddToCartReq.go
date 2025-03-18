package request

type AddToCartRequest struct {
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
	UserID    int `json:"user_id" validate:"required"`
}
