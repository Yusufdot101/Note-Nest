package project

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/filter"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) insert(p *Project) error {
	query := `
		INSERT INTO projects
		(user_id, title, description, visibility, color)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`
	values := []any{
		p.UserID,
		p.Title,
		p.Description,
		p.Visibility,
		p.Color,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.DB.ExecContext(ctx, query, values...)
	return err
}

func (r *Repository) get(currentUserID, queryUserID int, title, visibility string, filter filter.Filter) ([]*Project, error) {
	baseQuery := `
		SELECT 
			id, created_at, updated_at, user_id, title, description, visibility, entries_count, likes_count, 
			comments_count, color
		FROM projects
	`

	conds := []string{}
	args := []any{}
	idx := 1

	// =====================================================================
	// CASE 1: userID is provided
	// =====================================================================
	if queryUserID != -1 {
		conds = append(conds, fmt.Sprintf("user_id = $%d", idx))
		args = append(args, queryUserID)
		idx++

		if queryUserID == currentUserID {
			if visibility != "" {
				conds = append(conds, fmt.Sprintf("visibility = $%d", idx))
				args = append(args, visibility)
			}
		} else {
			// cant access only public projects of others
			if visibility != "public" && visibility != "" {
				return nil, nil
			}
			conds = append(conds, "visibility = 'public'")
		}
		goto BUILD
	}

	// =====================================================================
	// CASE 2: userID not provided
	// =====================================================================
	if visibility != "" {
		switch visibility {
		case "public":
			conds = append(conds, "visibility = 'public'")
		case "private":
			conds = append(conds, fmt.Sprintf("visibility = 'private' AND user_id = $%d", idx))
			args = append(args, currentUserID)
			idx++
		default:
			return nil, nil
		}
	} else {
		conds = append(conds, fmt.Sprintf("visibility = 'public' OR user_id = $%d", idx))
		args = append(args, currentUserID)
		idx++
	}

BUILD:
	if len(conds) == 0 {
		conds = []string{"1=1"}
	}

	query := baseQuery + " WHERE " + strings.Join(conds, " AND ")
	query += fmt.Sprintf(`
			AND (to_tsvector('simple', title) @@ plainto_tsquery('simple', $%d) OR $%d = '')
			ORDER BY %s %s, id ASC
			LIMIT $%d
			OFFSET $%d
		`, idx, idx, filter.SortColumn(), filter.SortDirection(), idx+1, idx+2)

	args = append(args, title)
	args = append(args, filter.Limit())
	args = append(args, filter.Offest())

	log.Println(query)
	log.Println(conds)
	log.Println(args)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	projects := []*Project{}
	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("rows close error:", err)
		}
	}()
	for rows.Next() {
		p := &Project{}
		err = rows.Scan(
			&p.ID,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.UserID,
			&p.Title,
			&p.Description,
			&p.Visibility,
			&p.EntriesCount,
			&p.LikesCount,
			&p.CommentsCount,
			&p.Color,
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil

	/*
		queryUserID := currentUserID
		// -1 means it was not given
		if userID != -1 {
			queryUserID = userID
			visibility = "public"
		}

		query := `
			SELECT
				id, created_at, updated_at, user_id, title, description, visibility, entries_count, likes_count,
				comments_count, color
			FROM projects
			WHERE user_id = $1
		`
		args := []any{userID}
		if visibility != "" {
			query += " AND visibility = $2"
			args = append(args, visibility)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		projects := []*Project{}
		rows, err := r.DB.QueryContext(ctx, query, args...)
		if err != nil {
			return nil, err
		}
		defer func() {
			if err := rows.Close(); err != nil {
				log.Println("rows close error:", err)
			}
		}()
		for rows.Next() {
			p := &Project{}
			err = rows.Scan(
				&p.ID,
				&p.CreatedAt,
				&p.UpdatedAt,
				&p.UserID,
				&p.Title,
				&p.Description,
				&p.Visibility,
				&p.EntriesCount,
				&p.LikesCount,
				&p.CommentsCount,
				&p.Color,
			)
			if err != nil {
				return nil, err
			}
			projects = append(projects, p)
		}
		if err = rows.Err(); err != nil {
			return nil, err
		}

		return projects, nil
	*/
}

func (r *Repository) getOne(ID int) (*Project, error) {
	query := `
		SELECT 
			id, created_at, updated_at, user_id, title, description, visibility, entries_count, likes_count, 
			comments_count, color
		FROM projects
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	p := &Project{}
	err := r.DB.QueryRowContext(ctx, query, ID).Scan(
		&p.ID,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.UserID,
		&p.Title,
		&p.Description,
		&p.Visibility,
		&p.EntriesCount,
		&p.LikesCount,
		&p.CommentsCount,
		&p.Color,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, custom_errors.ErrNoRecord
		default:
			return nil, err
		}
	}
	return p, nil
}

func (r *Repository) delete(projectID int) error {
	query := `
		DELETE 
		FROM projects
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := r.DB.ExecContext(ctx, query, projectID)
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

	return nil
}

// could not find a way around doing some logic here to avoid race conditions
func (r *Repository) update(userID, projectID int, title, description, visibility, color *string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Println(err)
		}
	}()

	getQuery := `
		SELECT 
			id, updated_at, user_id, title, description, visibility, color
		FROM projects
		WHERE id = $1
		FOR UPDATE
	`
	p := &Project{}
	err = tx.QueryRowContext(ctx, getQuery, projectID).Scan(
		&p.ID,
		&p.UpdatedAt,
		&p.UserID,
		&p.Title,
		&p.Description,
		&p.Visibility,
		&p.Color,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return custom_errors.ErrNoRecord
		default:
			return err
		}
	}

	// user cannot update another's project
	if p.UserID != userID {
		return custom_errors.ErrNoRecord
	}

	if title != nil {
		p.Title = *title
	}

	if description != nil {
		p.Description = *description
	}

	if visibility != nil {
		p.Visibility = *visibility
	}

	if color != nil {
		p.Color = *color
	}

	updateQuery := `
		UPDATE projects
		SET title = $1,
			description = $2,
			visibility = $3,
			color = $4,
			updated_at = $5
		WHERE id = $6
	`
	now := time.Now()
	values := []any{
		p.Title,
		p.Description,
		p.Visibility,
		p.Color,
		now,
		p.ID,
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
