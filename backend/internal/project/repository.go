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
		(user_id, name, description, visibility, color)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`
	values := []any{
		p.UserID,
		p.Name,
		p.Description,
		p.Visibility,
		p.Color,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.DB.ExecContext(ctx, query, values...)
	return err
}

func (r *Repository) get(userID int, visibility string) ([]*Project, error) {
	query := `
		SELECT 
			id, created_at, updated_at, user_id, name, description, visibility, entries_count, likes_count, 
			comments_count, color
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
		return nil, err
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
			&p.EntriesCount,
			&p.LikesCount,
			&p.CommentsCount,
			&p.Color,
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

func (r *Repository) getOne(ID int) (*Project, error) {
	query := `
		SELECT 
			id, created_at, updated_at, user_id, name, description, visibility, entries_count, likes_count, 
			comments_count, color
		FROM projects
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	p := &Project{}
	err := r.DB.QueryRowContext(ctx, query, ID).Scan(
		&p.ID,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.UserID,
		&p.Name,
		&p.Description,
		&p.Visibility,
		&p.EntriesCount,
		&p.LikesCount,
		&p.CommentsCount,
		&p.Color,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, custom_errors.ErrNoRecord
		default:
			return nil, err
		}
	}
	return p, nil
}

func (r *Repository) delete(projectID int) error {
	query := `
		DELETE 
		FROM projects
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := r.DB.ExecContext(ctx, query, projectID)
	if err != nil {
		return nil
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return custom_errors.ErrNoRecord
	}

	return nil
}

// could not find a way around doing some logic here to avoid race conditions
func (r *Repository) update(userID, projectID int, name, description, visibility, color string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Println(err)
		}
	}()

	getQuery := `
		SELECT 
			id, updated_at, user_id, name, description, visibility, color
		FROM projects
		WHERE id = $1
		FOR UPDATE
	`
	p := &Project{}
	err = tx.QueryRowContext(ctx, getQuery, projectID).Scan(
		&p.ID,
		&p.UpdatedAt,
		&p.UserID,
		&p.Name,
		&p.Description,
		&p.Visibility,
		&p.Color,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return custom_errors.ErrNoRecord
		default:
			return err
		}
	}

	// user cannot update another's project
	if p.UserID != userID {
		return custom_errors.ErrNoRecord
	}

	if name != "" {
		p.Name = name
	}
	// description can be empty because its optional so we should not use if p.Description != "" {}
	p.Description = description
	if visibility != "" {
		p.Visibility = visibility
	}
	if color != "" {
		p.Color = color
	}
	now := time.Now()
	p.UpdatedAt = &now

	updateQuery := `
		UPDATE projects
		SET name = $1,
			description = $2,
			visibility = $3,
			color = $4,
			updated_at = $5
		WHERE id = $6
	`
	values := []any{
		p.Name,
		p.Description,
		p.Visibility,
		p.Color,
		p.UpdatedAt,
		p.ID,
	}

	_, err = tx.ExecContext(ctx, updateQuery, values...)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
