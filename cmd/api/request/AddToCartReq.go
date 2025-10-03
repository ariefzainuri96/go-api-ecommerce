package request

type AddToCartRequest struct {
	ProductID int64 `json:"product_id" validate:"required"`
	Quantity  int   `json:"quantity" validate:"required"`
	UserID    int64 `json:"user_id" validate:"required"`
}
