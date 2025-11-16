package comment

import (
	"context"
	"database/sql"
	"time"
)

type repository struct {
	DB *sql.DB
}

func (r *repository) insert(c *comment) error {
	query := `
		INSERT INTO comments
		(user_id, note_id, content)
		VALUES
		($1, $2, $3)

		UPDATE notes
		set comments_count = comments_count + 1
		WHERE id = $2
	`

	values := []any{
		c.UserID,
		c.NoteID,
		c.Content,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.DB.ExecContext(ctx, query, values...)
	return err
}
