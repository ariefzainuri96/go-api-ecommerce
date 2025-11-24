package store

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/entity"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/request"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/response"
	"github.com/ariefzainuri96/go-api-ecommerce/internal/utils"
	"gorm.io/gorm"
)

type CartStore struct {
	db     *sql.DB
	gormDb *gorm.DB
}

func (s *CartStore) AddToCart(ctx context.Context, body request.AddToCartRequest, userId int) error {
	// query := `
	// 	INSERT INTO carts (product_id, quantity, user_id)
	// 	VALUES ($1, $2, $3);
	// `

	result := s.gormDb.WithContext(ctx).Create(&entity.Cart{
		ProductId: body.ProductID,
		UserId:    userId,
		Quantity:  body.Quantity,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *CartStore) DeleteFromCart(ctx context.Context, productID int) error {
	rows, err := gorm.G[entity.Cart](s.gormDb).
		Where(&entity.Cart{ProductId: productID}).
		Delete(ctx)

	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

func (s *CartStore) GetCart(ctx context.Context, userID int, req request.PaginationRequest) (response.CartsResponse, error) {
	query := s.gormDb.WithContext(ctx).
		Model(&entity.Cart{}).
		Where(entity.Cart{UserId: userID}).
		Preload("Product", nil).
		// if you want to perform, like search or filtering using field from the related table,
		// you should make Joins first
		Joins("INNER JOIN products ON products.id = carts.product_id")

	var searchAllQuery string

	if req.SearchAll != "" {
		searchAllQuery = "products.name ILIKE ? OR CAST(carts.quantity as TEXT) ILIKE ?"
	}

	result := utils.ApplyPagination[entity.Cart](query, req, searchAllQuery)

	if result.Error != nil {
		return response.CartsResponse{}, result.Error
	}

	return response.CartsResponse{
		BaseResponse: response.BaseResponse{
			Message: "Success",
			Status:  http.StatusOK,
		},
		Carts:      result.Data,
		Pagination: result.Pagination,
	}, nil
}

func (s *CartStore) UpdateQuantityCart(ctx context.Context, id int, quantity int) error {
	result := s.gormDb.WithContext(ctx).
		Model(&entity.Cart{}).
		UpdateColumn("Quantity", quantity)

	// query := `
	// 	UPDATE shopping_carts
	// 	SET quantity = $1
	// 	WHERE id = $2;
	// `

	// _, err := s.db.ExecContext(ctx, query, quantity, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
