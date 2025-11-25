package save

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/note"
)

var ErrNoteAlreadySaved = errors.New("note already saved")

type repository struct {
	DB *sql.DB
}

func (r *repository) insert(s *save) error {
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

	insertQuery := `
		INSERT INTO saves 
			( user_id, note_id )
		SELECT $1, $2
		FROM notes n
		JOIN projects p
		ON n.project_id = p.id
		WHERE n.id = $2
			AND ( n.visibility = 'public' OR p.user_id = $1 )
	`

	values := []any{
		s.userID,
		s.noteID,
	}

	res, err := tx.ExecContext(ctx, insertQuery, values...)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "saves_pkey"`:
			return ErrNoteAlreadySaved
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

	updateNotesQuery := `
		UPDATE notes
		SET saves_count = saves_count + 1
		WHERE id = $1
	`

	_, err = tx.ExecContext(ctx, updateNotesQuery, s.noteID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *repository) isSaved(userID, noteID int) (bool, error) {
	query := `
		SELECT user_id
		FROM saves
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

func (r *repository) getSavedNotes(userID int) ([]*note.Note, error) {
	query := `
		SELECT 
			n.id, n.project_id, n.created_at, n.updated_at, n.title, n.content, n.color, n.visibility, n.likes_count, 
			n.comments_count, n.saves_count, n.shares_count
		FROM notes n
		INNER JOIN saves s
		ON n.id = s.note_id
		INNER JOIN projects p
		ON n.project_id = p.id
		WHERE s.user_id = $1
			AND ( n.visibility = 'public' OR p.user_id = $1 )
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	var notes []*note.Note
	for rows.Next() {
		var note note.Note
		err := rows.Scan(
			&note.ID,
			&note.ProjectID,
			&note.CreatedAt,
			&note.UpdatedAt,
			&note.Title,
			&note.Content,
			&note.Color,
			&note.Visibility,
			&note.LikesCount,
			&note.CommentsCount,
			&note.SavesCount,
			&note.SharesCount,
		)
		if err != nil {
			return nil, err
		}
		notes = append(notes, &note)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *repository) delete(userID, noteID int) error {
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

	deleteQuery := `
		DELETE FROM saves
		WHERE user_id = $1 AND note_id = $2
	`

	values := []any{
		userID,
		noteID,
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

	updateNotesQuery := `
		UPDATE notes
		SET saves_count = saves_count - 1
		WHERE id = $1
	`

	_, err = tx.ExecContext(ctx, updateNotesQuery, noteID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
