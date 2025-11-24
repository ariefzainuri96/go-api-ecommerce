package store

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/entity"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/request"
	"gorm.io/gorm"
)

type ProductStore struct {
	db     *sql.DB
	gormDb *gorm.DB
}

func (s *ProductStore) AddProduct(ctx context.Context, body *request.AddProductRequest) error {
	// query := `
	// 	INSERT INTO products (name, description, price, quantity)
	// 	VALUES ($1, $2, $3, $4);
	// `

	// _, err := s.db.ExecContext(ctx, query, body.Name, body.Description, body.Price, body.Quantity)

	product := entity.Product{
		Name:        body.Name,
		Description: body.Description,
		Price:       int64(body.Price),
		Quantity:    body.Quantity,
	}

	result := s.gormDb.WithContext(ctx).Create(&product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// func (s *ProductStore) GetById(ctx context.Context, id int64) (response.Blog, error) {
// 	var blog response.Blog

// 	query := `
// 		SELECT id, title, description, created_at
// 		FROM blogs
// 		WHERE id = $1;
// 	`

// 	err := s.db.
// 		QueryRowContext(ctx, query, id).
// 		Scan(&blog.ID, &blog.Title, &blog.Description, &blog.CreatedAt)

// 	if err != nil {
// 		return blog, err
// 	}

// 	return blog, nil
// }

func (s *ProductStore) GetAllProduct(ctx context.Context) ([]entity.Product, error) {
	var products []entity.Product

	query := `
		SELECT id, name, description, price, quantity, created_at
		FROM products;
	`

	rows, err := s.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.CreatedAt)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (s *ProductStore) DeleteProduct(ctx context.Context, id int64) error {
	query := `
		DELETE FROM products
		WHERE id = $1;
	`

	result, err := s.db.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	row, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if row == 0 {
		return fmt.Errorf("no product found with id %d", id)
	}

	return nil
}

func (s *ProductStore) PatchProduct(ctx context.Context, id int64, patch map[string]any) error {
	query := "UPDATE products SET "
	args := []any{}
	i := 1

	for key, value := range patch {
		query += fmt.Sprintf("%s = $%d,", key, i)
		args = append(args, value)
		i++
	}

	// Remove trailing comma and add WHERE clause
	query = strings.TrimSuffix(query, ",")
	query += fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, id)

	// Execute the query
	_, err := s.db.ExecContext(ctx, query, args...)

	if err != nil {
		return err
	}

	return nil
}

func (s *ProductStore) SearchProduct(ctx context.Context, search string) ([]entity.Product, error) {
	var products []entity.Product

	query := `
		SELECT * FROM products
		WHERE LOWER (name) ILIKE $1
		OR LOWER (description) ILIKE $1;
	`

	searchTerm := "%" + search + "%"

	rows, err := s.db.QueryContext(ctx, query, searchTerm)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.CreatedAt)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
