package request

type UpdateCartRequest struct {
	Quantity int `json:"quantity" validate:"required"`
}
