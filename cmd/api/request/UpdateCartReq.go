package request

type UpdateCartRequest struct {
	Quantity int64 `json:"quantity" validate:"required"`
}
