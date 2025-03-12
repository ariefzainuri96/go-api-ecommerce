package store

import (
	"context"
	"database/sql"

	response "github.com/ariefzainuri96/go-api-blogging/cmd/api/response"
)

type BlogsStore struct {
	db *sql.DB
}

func (s *BlogsStore) CreateWithDB(ctx context.Context, body *response.Blog) error {
	query := `
		INSERT INTO blogs (title, description)
		VALUES ($1, $2) RETURNING id, created_at;
	`

	err := s.db.
		QueryRowContext(ctx, query, body.Title, body.Description).
		Scan(&body.ID, &body.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (s *BlogsStore) GetById(ctx context.Context, id int64) (response.Blog, error) {
	var blog response.Blog

	query := `
		SELECT id, title, description, created_at
		FROM blogs
		WHERE id = $1;
	`

	err := s.db.
		QueryRowContext(ctx, query, id).
		Scan(&blog.ID, &blog.Title, &blog.Description, &blog.CreatedAt)

	if err != nil {
		return blog, err
	}

	return blog, nil
}

func (s *BlogsStore) GetAll(ctx context.Context) ([]response.Blog, error) {
	var blogs []response.Blog

	query := `
		SELECT id, title, description, created_at
		FROM blogs;
	`

	rows, err := s.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var blog response.Blog
		err := rows.Scan(&blog.ID, &blog.Title, &blog.Description, &blog.CreatedAt)

		if err != nil {
			return nil, err
		}

		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func (s *BlogsStore) DeleteById(ctx context.Context, id int64) error {
	query := `
		DELETE FROM blogs
		WHERE id = $1;
	`

	_, err := s.db.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}
