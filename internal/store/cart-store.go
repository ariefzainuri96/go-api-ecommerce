package store

import (
	"context"
	"database/sql"

	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/entity"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/request"
	"gorm.io/gorm"
)

type CartStore struct {
	db     *sql.DB
	gormDb *gorm.DB
}

func (s *CartStore) AddToCart(ctx context.Context, body request.AddToCartRequest) error {
	// query := `
	// 	INSERT INTO carts (product_id, quantity, user_id)
	// 	VALUES ($1, $2, $3);
	// `

	result := s.gormDb.WithContext(ctx).Create(&entity.Cart{
		ProductId: body.ProductID,
		UserId:    body.UserID,
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

func (s *CartStore) GetCart(ctx context.Context, userID int64) ([]entity.Cart, error) {
	// rows, err := s.gormDb.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
	carts, err := gorm.G[entity.Cart](s.gormDb).
		Where(entity.Cart{UserId: userID}).
		Preload("Product", nil).
		Find(ctx)

	if err != nil {
		return nil, err
	}

	// query := `
	// 	SELECT shopping_carts.id, products.name, products.price, users.name, shopping_carts.quantity, products.price * shopping_carts.quantity as total
	// 	FROM shopping_carts
	// 	INNER JOIN products ON shopping_carts.product_id = products.id
	// 	INNER JOIN users ON shopping_carts.user_id = users.id
	// 	WHERE shopping_carts.user_id = $1;
	// `

	// rows, err := s.db.QueryContext(ctx, query, userID)

	// if err != nil {
	// 	return nil, err
	// }

	// defer rows.Close()

	// var carts []entity.Cart

	// for rows.Next() {
	// 	var cart entity.Cart
	// 	err := rows.Scan(&cart.ID, &cart.ProductName, &cart.ProductPrice, &cart.FullName, &cart.Quantity, &cart.TotalAmount)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	carts = append(carts, cart)
	// }

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
