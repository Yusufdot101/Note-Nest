package note

import (
	"context"
	"database/sql"
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
