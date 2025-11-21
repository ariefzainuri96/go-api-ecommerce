package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ariefzainuri96/go-api-ecommerce/internal/data"
)

type OrderStore struct {
	db *sql.DB
}

func (s *OrderStore) DeleteOrder(ctx context.Context, invoiceID int) error {
	query := `
		DELETE FROM orders
		WHERE invoice_id = $1;
	`

	_, err := s.db.ExecContext(ctx, query, invoiceID)

	if err != nil {
		return err
	}

	return nil
}

func (s *OrderStore) CreateOrder(ctx context.Context, body data.CreateOrderStruct) error {
	query := `
		INSERT INTO orders (user_id, total_price, status, product_id, quantity, invoice_id, invoice_url, invoice_exp_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
	`

	_, err := s.db.ExecContext(ctx, query, body.UserID, body.TotalPrice, body.Status, body.ProductID, body.Quantity, body.InvoiceID, body.InvoiceURL, body.InvoiceExpDate)

	if err != nil {
		return err
	}

	return nil
}

func (s *OrderStore) UpdateStatusOrder(ctx context.Context, invoiceID string, status string) error {
	query := `
		UPDATE orders
		SET status = $1, updated_at = NOW()
		WHERE invoice_id = $2;
	`

	res, err := s.db.ExecContext(ctx, query, status, invoiceID)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("No order found with invoice id %s", invoiceID)
	}

	return nil
}
