package store

import (
	"context"
	"database/sql"
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

func (s *CartStore) AddToCart(ctx context.Context, body request.AddToCartRequest, userId int64) error {
	// query := `
	// 	INSERT INTO carts (product_id, quantity, user_id)
	// 	VALUES ($1, $2, $3);
	// `

	result := s.gormDb.WithContext(ctx).Create(&entity.Cart{
		ProductId: body.ProductID,
		UserId:    userId,
		Quantity:  body.Quantity,
	})

	// _, err := s.db.ExecContext(ctx, query, body.ProductID, body.Quantity, body.UserID)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *CartStore) DeleteFromCart(ctx context.Context, productID int64) error {
	query := `
		DELETE FROM shopping_carts
		WHERE id = $1;
	`

	_, err := s.db.ExecContext(ctx, query, productID)

	if err != nil {
		return err
	}

	return nil
}

func (s *CartStore) GetCart(ctx context.Context, userID int64, req request.PaginationRequest) (response.CartsResponse, error) {
	query := s.gormDb.WithContext(ctx).
		Model(&entity.Cart{}).
		Where(entity.Cart{UserId: userID}).
		Preload("Product", nil).
		Joins("INNER JOIN products ON carts.product_id = products.id")

	var searchAllQuery string

	if req.SearchAll != "" {
		searchAllQuery = "products.name ILIKE ? OR carts.quantity ILIKE ?"
	}

	result := utils.ApplyPagination[entity.Cart](query, req, searchAllQuery)

	return response.CartsResponse{
		BaseResponse: response.BaseResponse{
			Message: "success",
			Status:  http.StatusOK,
		},
		Carts:      result.Data,
		Pagination: result.Pagination,
	}, nil
}

func (s *CartStore) UpdateQuantityCart(ctx context.Context, id int64, quantity int64) error {
	query := `
		UPDATE shopping_carts
		SET quantity = $1
		WHERE id = $2;
	`

	_, err := s.db.ExecContext(ctx, query, quantity, id)

	if err != nil {
		return err
	}

	return nil
}
