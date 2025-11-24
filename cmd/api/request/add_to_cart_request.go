package request

type AddToCartRequest struct {
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int   `json:"quantity" validate:"required,gt=0"`
}
