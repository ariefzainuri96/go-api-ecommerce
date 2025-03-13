package store

import (
	"context"
	"database/sql"

	"github.com/ariefzainuri96/go-api-blogging/cmd/api/request"
	response "github.com/ariefzainuri96/go-api-blogging/cmd/api/response"
)

type Storage struct {
	Product interface {
		GetAllProduct(context.Context) ([]response.Product, error)
		AddProduct(context.Context, *request.AddProductRequest) error
		// DeleteProduct(context.Context, int64) (response.Blog, error)
		// UpdateProduct(context.Context, int64) error
	}
	Auth interface {
		Login(context.Context, request.LoginRequest) (response.LoginData, error)
		Register(context.Context, request.LoginRequest) error
	}
	Cart interface {
		AddToCart(context.Context, request.AddToCartRequest) error
	}
	// create more interface here
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Product: &ProductStore{db},
		Auth:    &AuthStore{db},
		Cart:    &CartStore{db},
	}
}
