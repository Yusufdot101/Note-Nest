package user

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
)

var ErrDuplicateEmail = errors.New("duplicate email")

type Repository struct {
	DB *sql.DB
}

func (r *Repository) InsertUser(u *User) error {
	query := `
		INSERT INTO users
		(name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`
	values := []any{
		u.Name,
		u.Email,
		u.Password.hash,
	}

	// to prevent waiting indefinitely
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, values...).Scan(
		&u.ID,
		&u.CreatedAt,
	)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}

func (r *Repository) GetUserByEmail(email string) (*User, error) {
	query := `
		SELECT id, created_at, last_updated_at, name, email, password_hash
		FROM users
		WHERE email = $1
	`
	// to prevent waiting indefinitely
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	u := &User{}
	err := r.DB.QueryRowContext(ctx, query, email).Scan(
		&u.ID,
		&u.CreatedAt,
		&u.LastUpdatedAt,
		&u.Name,
		&u.Email,
		&u.Password.hash,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, custom_errors.ErrNoRecord
		default:
			return nil, err
		}
	}
	return u, nil
}
