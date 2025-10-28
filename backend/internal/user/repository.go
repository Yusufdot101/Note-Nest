package user

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var ErrDuplicateEmail = errors.New("duplicate email")

type Repository struct {
	DB *sql.DB
}

func (r *Repository) insertUser(u *User) error {
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
