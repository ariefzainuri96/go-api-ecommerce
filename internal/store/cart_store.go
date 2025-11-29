package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/entity"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/request"
	"github.com/ariefzainuri96/go-api-ecommerce/internal/utils"
	"gorm.io/gorm"
)

type CartStore struct {
	db     *sql.DB
	gormDb *gorm.DB
}

func (s *CartStore) AddToCart(ctx context.Context, body request.AddToCartRequest, userId int) (entity.Cart, error) {
	// query := `
	// 	INSERT INTO carts (product_id, quantity, user_id)
	// 	VALUES ($1, $2, $3);
	// `

	cart := entity.Cart{
		ProductId: body.ProductID,
		UserId:    userId,
		Quantity:  body.Quantity,
	}

	result := s.gormDb.
		WithContext(ctx).
		Create(&cart)

	if result.Error != nil {
		return cart, result.Error
	}

	if err := s.gormDb.
		WithContext(ctx).
		Preload("Product", nil).
		First(&cart, cart.ID).Error; err != nil {
		return cart, err
	}

	return cart, nil
}

func (s *CartStore) DeleteFromCart(ctx context.Context, productID int) error {
	results := s.gormDb.
		WithContext(ctx).
		Delete(&entity.Cart{
			BaseEntity: entity.BaseEntity{
				ID: uint(productID),
			},
		})

	if results.Error != nil {
		return results.Error
	}

	if results.RowsAffected == 0 {
		return fmt.Errorf("no id found")
	}

	return nil
}

func (s *CartStore) GetCart(ctx context.Context, userID int, req request.PaginationRequest) (utils.PaginateResult[entity.Cart], error) {
	query := s.gormDb.
		WithContext(ctx).
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
		return utils.PaginateResult[entity.Cart]{}, result.Error
	}

	return result, nil
}

func (s *CartStore) UpdateQuantityCart(ctx context.Context, id int, data map[string]any) (entity.Cart, error) {
	cart := entity.Cart{
		BaseEntity: entity.BaseEntity{
			ID: uint(id),
		},
	}

	// update
	result := s.gormDb.
		WithContext(ctx).
		Model(&cart).
		Updates(data)

	// query := `
	// 	UPDATE shopping_carts
	// 	SET quantity = $1
	// 	WHERE id = $2;
	// `

	// _, err := s.db.ExecContext(ctx, query, quantity, id)

	if result.Error != nil {
		return entity.Cart{}, result.Error
	}

	if err := s.gormDb.
		WithContext(ctx).
		Preload("Product", nil).
		First(&cart, id).
		Error; err != nil {
		return entity.Cart{}, err
	}

	return cart, nil
}
