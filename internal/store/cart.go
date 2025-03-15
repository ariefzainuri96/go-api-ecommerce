package store

import (
	"context"
	"database/sql"

	"github.com/ariefzainuri96/go-api-blogging/cmd/api/request"
	"github.com/ariefzainuri96/go-api-blogging/cmd/api/response"
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

func (s *CartStore) GetCart(ctx context.Context, userID int64) ([]response.Cart, error) {
	query := `
		SELECT shopping_carts.id, products.name, products.price, users.name, shopping_carts.quantity, products.price * shopping_carts.quantity as total
		FROM shopping_carts
		INNER JOIN products ON shopping_carts.product_id = products.id 
		INNER JOIN users ON shopping_carts.user_id = users.id
		WHERE shopping_carts.user_id = $1;
	`

	rows, err := s.db.QueryContext(ctx, query, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var carts []response.Cart

	for rows.Next() {
		var cart response.Cart
		err := rows.Scan(&cart.ID, &cart.ProductName, &cart.ProductPrice, &cart.FullName, &cart.Quantity, &cart.TotalAmount)
		if err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}

	return carts, nil
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
