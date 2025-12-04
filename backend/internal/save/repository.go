package save

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/filter"
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

func (r *repository) getSavedNotes(currentUserID, queryUserID, projectID int, title, visibility string, f *filter.Filter) ([]*note.Note, *filter.Metadata, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	baseQuery := `
		SELECT COUNT(*) OVER(),
			n.id, n.project_id, n.created_at, n.updated_at, n.title, n.content, n.color,
			n.visibility, n.likes_count, n.comments_count, n.saves_count
		FROM notes n
		JOIN saves s ON n.id = s.note_id
		JOIN projects p ON n.project_id = p.id
	`
	conds := []string{
		"s.user_id = $1",
	}
	args := []any{
		currentUserID,
	}
	idx := 2

	// =====================================================================
	// CASE 1: BOTH projectID AND userID are provided
	// =====================================================================
	if queryUserID != -1 && projectID != -1 {
		var owner int
		err := r.DB.QueryRowContext(ctx, "SELECT user_id from projects where id = $1", projectID).Scan(&owner)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil, customerrors.ErrNoRecord
			}
			return nil, nil, err
		}

		// userId must match actual project owner
		if owner != queryUserID {
			return nil, nil, customerrors.ErrNoRecord
		}

		conds = append(conds, fmt.Sprintf("n.project_id = $%d", idx))
		args = append(args, projectID)
		idx++
		if visibility != "" {
			if queryUserID != currentUserID {
				if visibility != "public" {
					return nil, nil, customerrors.ErrNoRecord
				}
				conds = append(conds, "n.visibility = 'public'")
			} else {
				conds = append(conds, fmt.Sprintf("n.visibility = $%d", idx))
				args = append(args, visibility)
				idx++
			}
		} else {
			conds = append(conds, fmt.Sprintf("( n.visibility = 'public' OR p.user_id = $%d )", idx))
			args = append(args, currentUserID)
			idx++
		}

		goto BUILD
	}

	// =====================================================================
	// CASE 2: ONLY userID is provided
	// =====================================================================
	if queryUserID != -1 {
		conds = append(conds, fmt.Sprintf("p.user_id = $%d", idx))
		args = append(args, queryUserID)
		idx++
		if visibility != "" {
			if queryUserID != currentUserID && visibility == "private" {
				return nil, nil, customerrors.ErrNoRecord
			}
			conds = append(conds, "n.visibility = 'public'")
		} else {
			conds = append(conds, fmt.Sprintf("( n.visibility = 'public' OR p.user_id = $%d )", idx))
			args = append(args, currentUserID)
			idx++
		}
		goto BUILD
	}

	// =====================================================================
	// CASE 3: ONLY projectID is provided
	// =====================================================================
	if projectID != -1 {
		var owner int
		err := r.DB.QueryRowContext(ctx, "SELECT user_id from projects where id = $1", projectID).Scan(&owner)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil, customerrors.ErrNoRecord
			}
			return nil, nil, err
		}

		conds = append(conds, fmt.Sprintf("n.project_id = $%d", idx))
		args = append(args, projectID)
		idx++
		if visibility != "" {
			if owner != currentUserID {
				if visibility != "public" {
					return nil, nil, customerrors.ErrNoRecord
				} else {
					conds = append(conds, "n.visibility = 'public'")
				}
			} else {
				conds = append(conds, fmt.Sprintf("n.visibility = $%d", idx))
				args = append(args, visibility)
				idx++
			}
		} else {
			conds = append(conds, fmt.Sprintf("(n.visibility = 'public' OR p.user_id = $%d)", idx))
			args = append(args, currentUserID)
			idx++
		}

		goto BUILD
	}

	if visibility != "" {
		if visibility == "public" {
			conds = append(conds, "n.visibility = 'public'")
		} else {
			conds = append(conds, fmt.Sprintf("n.visibility = 'private' AND p.user_id = $%d", idx))
			args = append(args, currentUserID)
			idx++
		}
	} else {
		conds = append(conds, fmt.Sprintf("(n.visibility = 'public' OR p.user_id = $%d)", idx))
		args = append(args, currentUserID)
		idx++
	}

BUILD:
	if len(conds) == 0 {
		conds = []string{"1=1"}
	}

	query := baseQuery + " WHERE " + strings.Join(conds, " AND ")
	query += fmt.Sprintf(`
			AND (
				$%d = ''
				OR to_tsvector('simple', n.title) @@ to_tsquery('simple', $%d)
			)
			ORDER BY %s %s, id ASC
			LIMIT $%d
			OFFSET $%d
		`, idx, idx, f.SortColumn(), f.SortDirection(), idx+1, idx+2)

	// to_tsquery wont work directly if you pass spaces, like "go is great" because spaces are treated as operators so you need to convert to into "go&is&great"
	// we add ':*' before the '&' so that partial word search is possible
	words := strings.Fields(title) // splits by spaces
	for i := range words {
		words[i] += ":*" // add prefix operator
	}
	formattedTitle := strings.Join(words, " & ") // join with &
	args = append(args, formattedTitle)
	args = append(args, f.Limit())
	args = append(args, f.Offset())

	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	var notes []*note.Note
	var totalResources int
	for rows.Next() {
		var note note.Note
		err := rows.Scan(
			&totalResources,
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
		)
		if err != nil {
			return nil, nil, err
		}
		notes = append(notes, &note)
	}

	if err = rows.Err(); err != nil {
		return nil, nil, err
	}

	metadata := filter.GenerateMetadata(f.Page, f.PageSize, totalResources)
	return notes, metadata, nil
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
