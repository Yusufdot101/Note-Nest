package save

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
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
		return custom_errors.ErrNoRecord
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
