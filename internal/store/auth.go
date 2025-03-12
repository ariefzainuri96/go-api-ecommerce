package store

import (
	"context"
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/ariefzainuri96/go-api-blogging/cmd/api/middleware"
	"github.com/ariefzainuri96/go-api-blogging/cmd/api/request"
	"github.com/ariefzainuri96/go-api-blogging/cmd/api/response"
)

type AuthStore struct {
	db *sql.DB
}

func (store *AuthStore) Login(ctx context.Context, body request.LoginRequest) (response.LoginData, error) {
	query := `SELECT name, email, password FROM users WHERE email = $1;`

	row := store.db.QueryRowContext(ctx, query, body.Email)

	var login response.LoginData
	var password string

	err := row.Scan(&login.Name, &login.Email, &password)

	if err != nil {
		return login, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(body.Password))

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
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2)`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	result, err := store.db.ExecContext(ctx, query, body.Name, body.Email, string(hashedPassword))

	if err != nil {
		return err
	}

	row, _ := result.RowsAffected()

	log.Println(row)

	return nil
}
