package store

import (
	"context"
	"database/sql"

	"github.com/ariefzainuri96/go-api-blogging/cmd/api/request"
)

type CartStore struct {
	db *sql.DB
}

func (s *CartStore) AddToCart(ctx context.Context, body request.AddToCartRequest) error {
	query := `
		INSERT INTO shopping_carts (product_id, quantity, user_id)
		VALUES ($1, $2, $3);
	`

	_, err := s.db.ExecContext(ctx, query, body.ProductID, body.Quantity, body.UserID)

	if err != nil {
		return err
	}

	return nil
}
