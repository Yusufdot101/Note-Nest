package comment

import (
	"context"
	"database/sql"
	"log"
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

func (r *repository) get(userID, noteID int) ([]*comment, error) {
	query := `
		SELECT c.id, c.created_at, c.user_id, c.note_id, c.content, c.likes_count 
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
