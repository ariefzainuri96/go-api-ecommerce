package store

import (
	"context"
	"database/sql"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/request"
	response "github.com/ariefzainuri96/go-api-ecommerce/cmd/api/response"
	entity "github.com/ariefzainuri96/go-api-ecommerce/cmd/api/entity"
	"github.com/ariefzainuri96/go-api-ecommerce/internal/data"
	"gorm.io/gorm"
)

type Storage struct {
	IProduct interface {
		GetProduct(context.Context, request.PaginationRequest) (response.ProductsResponse, error)
		AddProduct(context.Context, *request.AddProductRequest) (entity.Product, error)
		DeleteProduct(context.Context, uint) error
		PatchProduct(context.Context, uint, map[string]any) (entity.Product, error)		
	}
	IAuth interface {
		Login(context.Context, request.LoginRequest) (response.LoginData, error)
		Register(context.Context, request.RegisterReq) error
	}
	ICart interface {
		AddToCart(context.Context, request.AddToCartRequest, int) error
		GetCart(context.Context, int, request.PaginationRequest) (response.CartsResponse, error)
		DeleteFromCart(context.Context, int) error
		UpdateQuantityCart(context.Context, int, int) error		
	}
	IOrder interface {
		CreateOrder(context.Context, data.CreateOrderStruct) error
		UpdateStatusOrder(context.Context, string, string) error
		DeleteOrder(context.Context, int) error
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
