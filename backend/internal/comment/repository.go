package comment

import (
	"context"
	"database/sql"
	"time"
)

type repository struct {
	DB *sql.DB
}

func (r *repository) insert(c *comment, projectID int) error {
	insertQuery := `
		INSERT INTO comments
		(user_id, note_id, content)
		VALUES
		($1, $2, $3);
	`

	updateNotesQuery := `
		UPDATE notes
		set comments_count = comments_count + 1
		WHERE id = $1;
	`

	updateProjectsQuery := `
		UPDATE projects
		set comments_count = comments_count + 1
		WHERE id = $1;
	`

	values := []any{
		c.UserID,
		c.NoteID,
		c.Content,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.DB.ExecContext(ctx, insertQuery, values...)
	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, updateNotesQuery, c.NoteID)
	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, updateProjectsQuery, projectID)
	if err != nil {
		return err
	}

	return nil
}
