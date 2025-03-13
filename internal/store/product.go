package store

import (
	"context"
	"database/sql"

	"github.com/ariefzainuri96/go-api-blogging/cmd/api/request"
	"github.com/ariefzainuri96/go-api-blogging/cmd/api/response"
)

type ProductStore struct {
	db *sql.DB
}

func (s *ProductStore) AddProduct(ctx context.Context, body *request.AddProductRequest) error {
	query := `
		INSERT INTO products (name, description, price, quantity)
		VALUES ($1, $2, $3, $4);
	`

	_, err := s.db.ExecContext(ctx, query, body.Name, body.Description, body.Price, body.Quantity)

	if err != nil {
		return err
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

func (s *ProductStore) GetAllProduct(ctx context.Context) ([]response.Product, error) {
	var products []response.Product

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
		var product response.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.CreatedAt)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

// func (s *ProductStore) DeleteById(ctx context.Context, id int64) error {
// 	query := `
// 		DELETE FROM blogs
// 		WHERE id = $1;
// 	`

// 	_, err := s.db.ExecContext(ctx, query, id)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
