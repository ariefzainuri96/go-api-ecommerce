package store

import (
	"context"
	"database/sql"

	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/entity"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/request"
	response "github.com/ariefzainuri96/go-api-ecommerce/cmd/api/response"
	"github.com/ariefzainuri96/go-api-ecommerce/internal/data"
	"gorm.io/gorm"
)

type Storage struct {
	IProduct interface {
		GetAllProduct(context.Context) ([]entity.Product, error)
		AddProduct(context.Context, *request.AddProductRequest) error
		DeleteProduct(context.Context, int64) error
		PatchProduct(context.Context, int64, map[string]any) error
		SearchProduct(context.Context, string) ([]entity.Product, error)
	}
	IAuth interface {
		Login(context.Context, request.LoginRequest) (response.LoginData, error)
		Register(context.Context, request.RegisterReq) error
	}
	ICart interface {
		AddToCart(context.Context, request.AddToCartRequest) error
		GetCart(context.Context, int64) ([]entity.Cart, error)
		DeleteFromCart(context.Context, int64) error
		UpdateQuantityCart(context.Context, int64, int64) error
	}
	IOrder interface {
		CreateOrder(context.Context, data.CreateOrderStruct) error
		UpdateStatusOrder(context.Context, string, string) error
	}
	// create more interface here
}

func NewStorage(db *sql.DB, gormDb *gorm.DB) Storage {
	return Storage{
		IProduct: &ProductStore{db, gormDb},
		IAuth:    &AuthStore{db},
		ICart:    &CartStore{db, gormDb},
		IOrder:   &OrderStore{db},
	}
}
