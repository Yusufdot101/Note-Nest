package like

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
)

var ErrResourceAlreadyLike = errors.New("resource already liked")

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
		if err := tx.Rollback(); err != nil && errors.Is(err, sql.ErrTxDone) {
			log.Println("rollback error: ", err)
		}
	}()

	var insertQuery, updateProjectQuery, updateNoteQuery, updateCommentQuery string
	switch l.resourceType {
	case noteResource:
		insertQuery = `
		INSERT INTO likes (user_id, note_id)
		SELECT $1, $2
		FROM notes n
		JOIN projects p ON n.project_id = p.id
		WHERE n.id = $2
		`

	case commentResource:
		insertQuery = `
		INSERT INTO likes (user_id, comment_id)
		SELECT $1, $2
		FROM comments c
		JOIN notes n ON c.note_id = n.id
		JOIN projects p ON n.project_id = p.id
		WHERE c.id = $2
		`
	default:
		return customerrors.ErrNoRecord
	}

	visibilityRules := `
		AND (
		-- public note: anyone allowed
		n.visibility = 'public'

		OR

		-- private note: only owner allowed
		(n.visibility = 'private' AND p.user_id = $1)
		);
	`

	values := []any{
		l.userID,
		l.resourceID,
	}

	insertQuery += visibilityRules
	res, err := tx.ExecContext(ctx, insertQuery, values...)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "likes_note_unique"`:
			return ErrResourceAlreadyLike
		case err.Error() == `pq: duplicate key value violates unique constraint "likes_comment_unique"`:
			return ErrResourceAlreadyLike
		case err.Error() == `pq: insert or update on table "likes" violates foreign key constraint "likes_note_id_fkey"`:
			return customerrors.ErrNoRecord
		}
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return customerrors.ErrNoRecord
	}

	if l.resourceType == commentResource {
		updateCommentQuery = `
			UPDATE comments
			SET likes_count = likes_count + 1
			WHERE id = $1
		`

		_, err = tx.ExecContext(ctx, updateCommentQuery, l.resourceID)
		if err != nil {
			return err
		}
		goto COMMIT
	}

	updateNoteQuery = `
		UPDATE notes
		SET likes_count = likes_count + 1
		WHERE id = $1
	`

	_, err = tx.ExecContext(ctx, updateNoteQuery, l.resourceID)
	if err != nil {
		return err
	}

	updateProjectQuery = `
		UPDATE projects p
		SET likes_count = p.likes_count + 1
		FROM notes n
		WHERE n.id = $1
			AND n.project_id = p.id
	`

	_, err = tx.ExecContext(ctx, updateProjectQuery, l.resourceID)
	if err != nil {
		return err
	}

COMMIT:
	if err := tx.Commit(); err != nil && !errors.Is(err, sql.ErrTxDone) {
		return err
	}

	return nil
}

func (r *repo) delete(l *like) error {
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

	var deleteQuery, updateProjectQuery, updateNoteQuery string
	switch l.resourceType {
	case noteResource:
		deleteQuery = `
			DELETE FROM likes l
			USING notes n
			JOIN projects p ON n.project_id = p.id
			WHERE l.note_id = n.id
				AND l.user_id = $1
				AND l.note_id = $2
		`

	case commentResource:
		deleteQuery = `
			DELETE FROM likes l
			USING comments c
			JOIN notes n ON c.note_id = n.id
			JOIN projects p ON n.project_id = p.id
			WHERE l.comment_id = n.id
				AND l.user_id = $1
				AND l.comment_id = $2
		`

	default:
		return customerrors.ErrNoRecord
	}

	visibilityRules := `
		AND (
		-- public note: anyone allowed
		n.visibility = 'public'

		OR

		-- private note: only owner allowed
		(n.visibility = 'private' AND p.user_id = $1)
		);
	`
	deleteQuery += visibilityRules

	values := []any{
		l.userID,
		l.resourceID,
	}

	res, err := tx.ExecContext(ctx, deleteQuery, values...)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return customerrors.ErrNoRecord
	}

	if l.resourceType == commentResource {
		updateNoteQuery = `
			UPDATE comments
			SET likes_count = likes_count - 1
			WHERE id = $1
		`

		_, err = tx.ExecContext(ctx, updateNoteQuery, l.resourceID)
		if err != nil {
			return err
		}
		goto COMMIT
	}

	updateNoteQuery = `
		UPDATE notes
		SET likes_count = likes_count - 1
		WHERE id = $1
	`

	_, err = tx.ExecContext(ctx, updateNoteQuery, l.resourceID)
	if err != nil {
		return err
	}

	updateProjectQuery = `
			UPDATE projects p
			SET likes_count = p.likes_count - 1
			FROM notes n
			WHERE n.id = $1
				AND n.project_id = p.id
		`

	_, err = tx.ExecContext(ctx, updateProjectQuery, l.resourceID)
	if err != nil {
		return err
	}

COMMIT:
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *repo) isLiked(l *like) (bool, error) {
	query := `
		SELECT user_id
		FROM likes
		WHERE user_id = $1
	`

	switch l.resourceType {
	case noteResource:
		query += "AND note_id = $2"
	case commentResource:
		query += "AND comment_id = $2"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	values := []any{
		l.userID,
		l.resourceID,
	}
	var id int
	err := r.DB.QueryRowContext(ctx, query, values...).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
