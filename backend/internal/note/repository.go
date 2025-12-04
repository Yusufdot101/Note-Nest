package note

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/filter"
	"github.com/Yusufdot101/note-nest/internal/project"
)

type Repository struct {
	DB            *sql.DB
	UpdateTimeout time.Duration
}

func NewRepository(db *sql.DB) (*Repository, error) {
	timeout, err := time.ParseDuration(os.Getenv("NOTE_UPDATE_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("invalid NOTE_UPDATE_TIMEOUT: %w", err)
	}
	return &Repository{
		DB:            db,
		UpdateTimeout: timeout,
	}, nil
}

func (r *Repository) insert(n *Note) error {
	insertQuery := `
		INSERT INTO notes
		(project_id, title, content, color, visibility)
		VALUES ($1, $2, $3, $4, $5);
	`
	updateProjectsQuery := `
		UPDATE projects
		SET entries_count = entries_count + 1
		WHERE id = $1;
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

	_, err := r.DB.ExecContext(ctx, insertQuery, values...)
	if err != nil {
		switch err.Error() {
		case `pq: insert or update on table "notes" violates foreign key constraint "notes_project_id_fkey"`:
			return customerrors.ErrNoRecord
		default:
			return err
		}
	}

	_, err = r.DB.ExecContext(ctx, updateProjectsQuery, n.ProjectID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) get(noteID int) (*Note, error) {
	query := `
		SELECT 
			id, project_id, created_at, updated_at, title, content, color, visibility, likes_count, 
			comments_count, saves_count
		FROM notes
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	note := &Note{}
	err := r.DB.QueryRowContext(ctx, query, noteID).Scan(
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
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, customerrors.ErrNoRecord
		default:
			return nil, err
		}
	}

	return note, nil
}

func (r *Repository) getMany(currentUserID, queryUserID, projectID int, title, visibility string, f *filter.Filter) ([]*Note, *filter.Metadata, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	baseQuery := `
		SELECT COUNT(*) OVER(),
			n.id, n.project_id, n.created_at, n.updated_at, n.title, n.content, n.color,
			n.visibility, n.likes_count, n.comments_count, n.saves_count
		FROM notes n
		JOIN projects p ON n.project_id = p.id
	`
	conds := []string{}
	args := []any{}
	idx := 1

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
	// CASE 2: ONLY projectID is provided
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

	var notes []*Note
	var totalResources int
	for rows.Next() {
		var note Note
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

func (r *Repository) delete(noteID, projectID int) error {
	deletQuery := `
		DELETE FROM notes
		WHERE id = $1
	`
	updateQuery := `
		UPDATE projects
		SET entries_count = entries_count - 1
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := r.DB.ExecContext(ctx, deletQuery, noteID)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return customerrors.ErrNoRecord
	}

	res, err = r.DB.ExecContext(ctx, updateQuery, projectID)
	if err != nil {
		return err
	}

	affectedRows, err = res.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return customerrors.ErrNoRecord
	}

	return nil
}

func (r *Repository) updateNoteTitleContent(userID, noteID int, title, content *string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Println(err)
		}
	}()

	// fetch note
	n := &Note{}
	fetchNoteQuery := `
		SELECT id, project_id, created_at, title, content
		FROM notes
		WHERE id = $1
		FOR UPDATE
	`

	err = tx.QueryRowContext(ctx, fetchNoteQuery, noteID).Scan(
		&n.ID,
		&n.ProjectID,
		&n.CreatedAt,
		&n.Title,
		&n.Content,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return customerrors.ErrNoRecord
		default:
			return err
		}
	}

	// fetch project
	fetchProjectQuery := `
		SELECT user_id
		FROM projects
		WHERE id = $1
	`

	p := &project.Project{}
	err = tx.QueryRowContext(ctx, fetchProjectQuery, n.ProjectID).Scan(
		&p.UserID,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return customerrors.ErrNoRecord
		default:
			return err
		}
	}

	// do checks
	if p.UserID != userID {
		return customerrors.ErrNoRecord
	}

	if time.Since(*n.CreatedAt) > r.UpdateTimeout {
		return customerrors.ErrUpdateTimeout
	}

	// update fields
	if title != nil {
		n.Title = *title
	}

	if content != nil {
		n.Content = *content
	}

	// save note
	updateQuery := `
		UPDATE notes
		SET title = $1,
			content = $2,
			updated_at = $3
		WHERE id = $4
	`

	values := []any{
		n.Title,
		n.Content,
		time.Now(),
		n.ID,
	}

	_, err = tx.ExecContext(ctx, updateQuery, values...)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) updateNoteVisibility(userID, noteID int, visibility string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Println(err)
		}
	}()

	// fetch note
	n := &Note{}
	fetchNoteQuery := `
		SELECT n.id
		FROM notes n
		INNER JOIN projects p
		ON n.project_id = p.id
		WHERE n.id = $1
			AND p.user_id = $2
		FOR UPDATE
	`

	err = tx.QueryRowContext(ctx, fetchNoteQuery, noteID, userID).Scan(
		&n.ID,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return customerrors.ErrNoRecord
		default:
			return err
		}
	}

	// save note
	updateQuery := `
		UPDATE notes
		SET visibility = $1,
			updated_at = $2
		WHERE id = $3
	`

	values := []any{
		visibility,
		time.Now(),
		n.ID,
	}

	_, err = tx.ExecContext(ctx, updateQuery, values...)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) updateNoteColor(userID, noteID int, color string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Println(err)
		}
	}()

	// fetch note
	n := &Note{}
	fetchNoteQuery := `
		SELECT id, project_id
		FROM notes
		WHERE id = $1
		FOR UPDATE
	`

	err = tx.QueryRowContext(ctx, fetchNoteQuery, noteID).Scan(
		&n.ID,
		&n.ProjectID,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return customerrors.ErrNoRecord
		default:
			return err
		}
	}

	// fetch project
	fetchProjectQuery := `
		SELECT user_id
		FROM projects
		WHERE id = $1
	`

	p := &project.Project{}
	err = tx.QueryRowContext(ctx, fetchProjectQuery, n.ProjectID).Scan(
		&p.UserID,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return customerrors.ErrNoRecord
		default:
			return err
		}
	}

	// do checks
	if p.UserID != userID {
		return customerrors.ErrNoRecord
	}

	// save note
	updateQuery := `
		UPDATE notes
		SET color = $1,
			updated_at = $2
		WHERE id = $3
	`

	values := []any{
		color,
		time.Now(),
		n.ID,
	}

	_, err = tx.ExecContext(ctx, updateQuery, values...)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) getOwner(userID, noteID int) (string, error) {
	query := `
		SELECT u.name
		FROM users u 
		INNER JOIN projects p
		ON p.user_id = u.id
		INNER JOIN notes n
		ON n.project_id = p.id
		WHERE n.id = $1
			AND ( n.visibility = 'public' OR p.user_id = $2 )
	`

	var name string
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, noteID, userID).Scan(
		&name,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", customerrors.ErrNoRecord
		}
		return "", err
	}

	return name, nil
}
