package store

import (
	"context"
	"database/sql"

	"github.com/ariefzainuri96/go-api-blogging/cmd/api/middleware"
	"github.com/ariefzainuri96/go-api-blogging/cmd/api/request"
	"github.com/ariefzainuri96/go-api-blogging/cmd/api/response"
)

type AuthStore struct {
	db *sql.DB
}

func (store *AuthStore) Login(ctx context.Context, body request.LoginRequest) (response.LoginData, error) {
	query := `SELECT id, email, created_at FROM users_login WHERE email = $1 AND password = $2 `

	row := store.db.QueryRowContext(ctx, query, body.Email, body.Password)

	var login response.LoginData

	err := row.Scan(&login.ID, &login.Email, &login.CreatedAt)

	if err != nil {
		return login, err
	}

	token, err := middleware.GenerateToken(body.Email, login.ID)

	if err != nil {
		return login, err
	}

	login.Token = token

	return login, nil
}

func (store *AuthStore) Register(ctx context.Context, body request.LoginRequest) error {
	query := `INSERT INTO users_login (email, password) VALUES ($1, $2)`

	_, err := store.db.ExecContext(ctx, query, body.Email, body.Password)

	if err != nil {
		return err
	}

	return nil
}
