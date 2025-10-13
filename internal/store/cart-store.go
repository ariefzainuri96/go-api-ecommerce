package store

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"strings"

	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/entity"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/request"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/response"
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
	var resp response.CartsResponse

	offset := (req.Page - 1) * req.PageSize

	query := gorm.G[entity.Cart](s.gormDb).
		Where(entity.Cart{UserId: userID}).
		Preload("Product", nil).
		Offset(offset).
		Limit(req.PageSize)

	// Optional ordering
	if req.OrderBy != "" {
		sortDirection := "ASC"
		if strings.ToUpper(req.Sort) == "DESC" {
			sortDirection = "DESC"
		}
		query = query.Order(fmt.Sprintf("%s %s", req.OrderBy, sortDirection))
	}

	// Optional search filtering
	if req.SearchField != "" && req.SearchValue != "" {
		query = query.Where(fmt.Sprintf("%s ILIKE ?", req.SearchField), "%"+req.SearchValue+"%")
	} else if req.SearchAll != "" {
		search := "%" + req.SearchAll + "%"
		query = query.Where("product_name ILIKE ? OR description ILIKE ?", search, search)
	}

	// Fetch paginated data
	carts, err := query.Find(ctx)
	if err != nil {
		return resp, fmt.Errorf("failed to fetch carts: %w", err)
	}

	// Count total rows (without offset/limit)
	var total int64
	countQuery := s.gormDb.Model(&entity.Cart{}).Where("user_id = ?", userID)
	if req.SearchField != "" && req.SearchValue != "" {
		countQuery = countQuery.Where(fmt.Sprintf("%s ILIKE ?", req.SearchField), "%"+req.SearchValue+"%")
	} else if req.SearchAll != "" {
		search := "%" + req.SearchAll + "%"
		countQuery = countQuery.Where("product_name ILIKE ? OR description ILIKE ?", search, search)
	}

	if err := countQuery.Count(&total).Error; err != nil {
		return resp, fmt.Errorf("failed to count carts: %w", err)
	}

	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	// Build response
	resp = response.CartsResponse{
		BaseResponse: response.BaseResponse{
			Message: "success",
			Status:  http.StatusOK,
		},
		Carts: carts,
		Pagination: response.PaginationMetadata{
			Page:      req.Page,
			PageSize:  req.PageSize,
			TotalData: total,
			TotalPage: totalPages,
		},
	}

	return resp, nil
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
