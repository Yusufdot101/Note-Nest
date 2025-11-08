package project

import (
	"context"
	"database/sql"
	"time"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) insert(p *Project) error {
	query := `
		INSERT INTO projects
		(user_id, name, description, visibility)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`
	values := []any{
		p.UserID,
		p.Name,
		p.Description,
		p.Visibility,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, values...).Scan(
		&p.UserID,
		&p.CreatedAt,
	)
	return err
}
