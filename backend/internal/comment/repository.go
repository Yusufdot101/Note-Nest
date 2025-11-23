package comment

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
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

func (r *repository) get(userID, noteID int) ([]*comment, error) {
	query := `
		SELECT c.id, c.created_at, c.is_edited, c.user_id, c.note_id, c.content, c.likes_count 
		FROM comments c
		INNER JOIN notes n
		ON c.note_id = n.id
		INNER JOIN projects p
		ON n.project_id = p.id
		WHERE n.id = $1
			AND ( n.visibility = 'public' OR p.user_id = $2 )
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.DB.QueryContext(ctx, query, noteID, userID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	comments := []*comment{}

	for rows.Next() {
		c := &comment{}
		err = rows.Scan(
			&c.ID,
			&c.CreatedAt,
			&c.Edited,
			&c.UserID,
			&c.NoteID,
			&c.Content,
			&c.LikesCount,
		)
		if err != nil {
			return nil, err
		}

		comments = append(comments, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *repository) update(userID, commentID int, newContent string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Println("rollback error: ", err)
		}
	}()

	fetchQuery := `
		SELECT c.content
		FROM comments c
		INNER JOIN notes n
		ON c.note_id = n.id
		INNER JOIN projects p
		ON n.project_id = p.id
		WHERE c.id = $1 AND c.user_id = $2
			AND ( n.visibility = 'public' OR p.user_id = $2 )
		FOR UPDATE
	`

	values := []any{
		commentID,
		userID,
	}

	var content string
	err = tx.QueryRowContext(ctx, fetchQuery, values...).Scan(&content)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return custom_errors.ErrNoRecord
		}
		return err
	}

	updateQuery := `
		UPDATE comments
		SET content = $1,
			is_edited = true
		WHERE id = $2
	`
	_, err = tx.ExecContext(ctx, updateQuery, newContent, commentID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
