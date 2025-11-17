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

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
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
			return custom_errors.ErrNoRecord
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
			id, project_id, created_at, title, content, color, visibility, likes_count, 
			comments_count
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
		&note.Title,
		&note.Content,
		&note.Color,
		&note.Visibility,
		&note.LikesCount,
		&note.CommentsCount,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, custom_errors.ErrNoRecord
		default:
			return nil, err
		}
	}

	return note, nil
}

func (r *Repository) getMany(currentUserID, queryUserID, projectID *int, visibility string) ([]*Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	uid := *currentUserID // guaranteed non-nil cause the endpoint has requireAccess middleware that needs access token which has the currentUserID

	baseQuery := `
		SELECT
			n.id, n.project_id, n.created_at, n.title, n.content, n.color,
			n.visibility, n.likes_count, n.comments_count
		FROM notes n
		JOIN projects p ON n.project_id = p.id
	`
	conds := []string{}
	args := []any{}
	idx := 1

	// =====================================================================
	// CASE 1: BOTH projectId AND userId are provided
	// =====================================================================
	if projectID != nil && queryUserID != nil {
		// fetch project owner
		var ownerID int
		err := r.DB.QueryRowContext(ctx,
			"SELECT user_id FROM projects WHERE id = $1",
			*projectID,
		).Scan(&ownerID)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, custom_errors.ErrNoRecord
			}
			return nil, err
		}

		// userId must match actual project owner
		if ownerID != *queryUserID {
			return nil, custom_errors.ErrNoRecord
		}

		// filter by project
		conds = append(conds, fmt.Sprintf("n.project_id = $%d", idx))
		args = append(args, *projectID)
		idx++

		// if the requester is the owner → full visibility rules apply
		if ownerID == uid {
			if visibility != "" {
				conds = append(conds, fmt.Sprintf("n.visibility = $%d", idx))
				args = append(args, visibility)
			}
		} else {
			// not owner → only public allowed
			conds = append(conds, "n.visibility = 'public'")
		}

		goto BUILD
	}

	// =====================================================================
	// CASE 2: ONLY projectId is provided
	// =====================================================================
	if projectID != nil {
		var ownerID int
		err := r.DB.QueryRowContext(ctx,
			"SELECT user_id FROM projects WHERE id = $1",
			*projectID,
		).Scan(&ownerID)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, custom_errors.ErrNoRecord
			}
			return nil, err
		}

		conds = append(conds, fmt.Sprintf("n.project_id = $%d", idx))
		args = append(args, *projectID)
		idx++

		if ownerID == uid {
			if visibility != "" {
				conds = append(conds, fmt.Sprintf("n.visibility = $%d", idx))
				args = append(args, visibility)
			}
		} else {
			conds = append(conds, "n.visibility = 'public'")
		}

		goto BUILD
	}

	// =====================================================================
	// CASE 3: ONLY userId is provided
	// =====================================================================
	if queryUserID != nil {
		conds = append(conds, fmt.Sprintf("p.user_id = $%d", idx))
		args = append(args, *queryUserID)
		idx++

		if *queryUserID == uid {
			if visibility != "" {
				conds = append(conds, fmt.Sprintf("n.visibility = $%d", idx))
				args = append(args, visibility)
			}
		} else {
			conds = append(conds, "n.visibility = 'public'")
		}

		goto BUILD
	}

	// =====================================================================
	// CASE 4: GLOBAL FEED (no projectId, no userId)
	// =====================================================================
	if visibility != "" {
		if visibility == "public" {
			conds = append(conds, "n.visibility = 'public'")
		} else {
			conds = append(conds, fmt.Sprintf("n.visibility = 'private' AND p.user_id = $%d", idx))
			args = append(args, uid)
		}
	} else {
		conds = append(conds, fmt.Sprintf("(n.visibility = 'public' OR p.user_id = $%d)", idx))
		args = append(args, uid)
	}

BUILD:
	if len(conds) == 0 {
		conds = []string{"1=1"}
	}

	query := baseQuery + " WHERE " + strings.Join(conds, " AND ") + " ORDER BY n.created_at DESC"

	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	var notes []*Note
	for rows.Next() {
		var note Note
		err := rows.Scan(
			&note.ID, &note.ProjectID, &note.CreatedAt, &note.Title, &note.Content,
			&note.Color, &note.Visibility, &note.LikesCount, &note.CommentsCount,
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

func (r *Repository) delete(noteID, projectID int) error {
	query := `
		DELETE FROM notes
		WHERE id = $1

		UPDATE projects
		SET entries_count = entries_count - 1
		WHERE id = $2
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := r.DB.ExecContext(ctx, query, noteID, projectID)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return custom_errors.ErrNoRecord
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
			return custom_errors.ErrNoRecord
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
			return custom_errors.ErrNoRecord
		default:
			return err
		}
	}

	// do checks
	if p.UserID != userID {
		return custom_errors.ErrNoRecord
	}

	if time.Since(*n.CreatedAt) > r.UpdateTimeout {
		return custom_errors.ErrUpdateTimeout
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
