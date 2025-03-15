package store

import (
	"context"
	"database/sql"
	"errors"
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
	query := `SELECT id, name, email, password, is_admin FROM users WHERE email = $1;`

	row := store.db.QueryRowContext(ctx, query, body.Email)

	var login response.LoginData
	var password string
	var isAdmin bool

	err := row.Scan(&login.ID, &login.Name, &login.Email, &password, &isAdmin)

	if err != nil {
		return login, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(body.Password))

	if err != nil {
		return login, errors.New("Invalid email or password!")
	}

	log.Println("userid", login.ID)

	token, err := middleware.GenerateToken(body.Email, isAdmin, login.ID)

	if err != nil {
		return login, err
	}

	login.Token = token

	return login, nil
}

func (store *AuthStore) Register(ctx context.Context, body request.LoginRequest) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`

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
