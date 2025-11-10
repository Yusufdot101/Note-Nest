package note

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) insert(n *Note) error {
	query := `
		INSERT INTO notes
		(project_id, title, content, color, visibility)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, likes_count, comments_count
	`
	values := []any{
		n.ProjectID,
		n.Title,
		n.Content,
		n.Color,
		n.Visibility,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, values...).Scan(
		&n.ID,
		&n.CreatedAt,
		&n.LikesCount,
		&n.CommentsCount,
	)
	if err != nil {
		switch {
		case err.Error() == `pq: insert or update on table "notes" violates foreign key constraint "notes_project_id_fkey"`:
			return custom_errors.ErrNoRecord
		default:
			return err
		}
	}
	return nil
}

func (r *Repository) get(ProjectID, noteID int, visibility string) ([]*Note, error) {
	query := `
		SELECT 
			id, project_id, created_at, title, content, color, visibility, likes_count, 
			comments_count
		FROM notes
		WHERE project_id = $1
	`
	args := []any{
		ProjectID,
	}
	argNum := 2
	if noteID != 0 {
		query += fmt.Sprintf(" AND id = $%d", argNum)
		args = append(args, noteID)
		argNum++
	}
	if visibility != "" {
		query += fmt.Sprintf(" AND visibility = $%d", argNum)
		args = append(args, visibility)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("rows close error:", err)
		}
	}()
	notes := []*Note{}
	for rows.Next() {
		n := &Note{}
		err = rows.Scan(
			&n.ID,
			&n.ProjectID,
			&n.CreatedAt,
			&n.Title,
			&n.Color,
			&n.Color,
			&n.Visibility,
			&n.LikesCount,
			&n.CommentsCount,
		)
		if err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}
