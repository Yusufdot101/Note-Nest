package like

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
)

var ErrNoteAlreadyLike = errors.New("note already liked")

type repo struct {
	DB *sql.DB
}

func (r *repo) insert(l *like) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Println("rollback error: ", err)
		}
	}()

	insertQuery := `
		INSERT INTO likes (user_id, note_id)
		SELECT $1, $2
		FROM notes n
		JOIN projects p ON n.project_id = p.id
		WHERE n.id = $2
		  AND (
				-- public note: anyone allowed
				n.visibility = 'public'
				
				OR
				
				-- private note: only owner allowed
				(n.visibility = 'private' AND p.user_id = $1)
			  );
	`

	res, err := tx.ExecContext(ctx, insertQuery, l.userID, l.noteID)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "likes_pkey"`:
			return ErrNoteAlreadyLike
		case err.Error() == `pq: insert or update on table "likes" violates foreign key constraint "likes_note_id_fkey"`:
			return custom_errors.ErrNoRecord
		}
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return custom_errors.ErrNoRecord
	}

	updateNoteQuery := `
		UPDATE notes 
		SET likes_count = likes_count + 1
		WHERE id = $1
	`

	_, err = tx.ExecContext(ctx, updateNoteQuery, l.noteID)
	if err != nil {
		return err
	}

	updateProjectQuery := `
		UPDATE projects p
		SET likes_count = p.likes_count + 1
		FROM notes n
		WHERE n.id = $1
			AND n.project_id = p.id
	`

	_, err = tx.ExecContext(ctx, updateProjectQuery, l.noteID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *repo) delete(userID, noteID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Println("rollback error: ", err)
		}
	}()

	deleteQuery := `
		DELETE FROM likes l
		USING notes n
		JOIN projects p ON n.project_id = p.id
		WHERE l.note_id = n.id
			AND l.user_id = $1
			AND l.note_id = $2
			AND (
				n.visibility = 'public'
				OR p.user_id = $1
			  );
	`

	res, err := tx.ExecContext(ctx, deleteQuery, userID, noteID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return custom_errors.ErrNoRecord
	}

	updateNoteQuery := `
		UPDATE notes 
		SET likes_count = likes_count - 1
		WHERE id = $1
	`

	_, err = tx.ExecContext(ctx, updateNoteQuery, noteID)
	if err != nil {
		return err
	}

	updateProjectQuery := `
		UPDATE projects p
		SET likes_count = p.likes_count - 1
		FROM notes n
		WHERE n.id = $1
			AND n.project_id = p.id
	`

	_, err = tx.ExecContext(ctx, updateProjectQuery, noteID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *repo) isLiked(userID, noteID int) (bool, error) {
	query := `
		SELECT user_id
		FROM likes
		WHERE user_id = $1 AND note_id = $2
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id int
	err := r.DB.QueryRowContext(ctx, query, userID, noteID).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
