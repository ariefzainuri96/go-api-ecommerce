package store

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/entity"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/request"
	"github.com/ariefzainuri96/go-api-ecommerce/internal/utils"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ProductStore struct {
	db     *sql.DB
	gormDb *gorm.DB
}

func (s *ProductStore) AddProduct(ctx context.Context, body *request.AddProductRequest) (entity.Product, error) {
	product := entity.Product{
		Name:        body.Name,
		Description: body.Description,
		Price:       decimal.NewFromFloat(body.Price),
		Quantity:    body.Quantity,
	}

	result := s.gormDb.WithContext(ctx).Create(&product)

	if result.Error != nil {
		return entity.Product{}, result.Error
	}

	return product, nil
}

func (s *ProductStore) GetProduct(ctx context.Context, req request.PaginationRequest) (utils.PaginateResult[entity.Product], error) {
	var products []entity.Product

	query := s.gormDb.Find(&products)

	var searchAllQuery string

	if req.SearchAll != "" {
		searchAllQuery = `
		products.name ILIKE ?
		OR products.description ILIKE ?
		OR CAST(products.quantity as TEXT) ILIKE ?
		OR CAST(products.price as TEXT) ILIKE ?
		`
	}

	result := utils.ApplyPagination[entity.Product](query, req, searchAllQuery)

	if result.Error != nil {
		return utils.PaginateResult[entity.Product]{}, result.Error
	}

	return result, nil
}

func (s *ProductStore) DeleteProduct(ctx context.Context, id uint) error {
	product := entity.Product{
		BaseEntity: entity.BaseEntity{
			ID: id,
		},
	}

	result := s.gormDb.Delete(&product)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no product found with id %d", id)
	}

	return nil
}

func (s *ProductStore) PatchProduct(ctx context.Context, id uint, patch map[string]any) (entity.Product, error) {

	product := entity.Product {
		BaseEntity: entity.BaseEntity{ ID: id },
	}

	result := s.gormDb.Model(&product).Updates(patch)

	if result.Error != nil {
		return entity.Product{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entity.Product{}, fmt.Errorf("no product found with id %v", id)
	}

	if err := s.gormDb.First(&product, id).Error; err != nil {
        return entity.Product{}, err
    }

	return product, nil
}
