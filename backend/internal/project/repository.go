package project

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
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

	_, err := r.DB.ExecContext(ctx, query, values...)
	return err
}

func (r *Repository) get(userID int, visibility string) ([]*Project, error) {
	query := `
		SELECT id, created_at, updated_at, user_id, name, description, visibility
		FROM projects
		WHERE user_id = $1
	`
	args := []any{userID}
	if visibility != "" {
		query += " AND visibility = $2"
		args = append(args, visibility)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	projects := []*Project{}
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, custom_errors.ErrNoRecord
		default:
			return nil, err
		}
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("rows close error:", err)
		}
	}()
	for rows.Next() {
		p := &Project{}
		err = rows.Scan(
			&p.ID,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.UserID,
			&p.Name,
			&p.Description,
			&p.Visibility,
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
