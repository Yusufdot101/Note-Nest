package user

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/token"
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
			return nil, customerrors.ErrNoRecord
		default:
			return nil, err
		}
	}
	return u, nil
}

func (r *Repository) UpdatePasswordUsingToken(tokenString, newPassword string) error {
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

	fetchQuery := `
		SELECT u.id, u.password_hash, last_updated_at
		FROM users u
		INNER JOIN tokens t
			ON t.user_id = u.id
		WHERE t.token_string = $1 
			AND t.expires > Now()
			AND t.use = $2
		FOR UPDATE
	`
	u := &User{}
	err = tx.QueryRowContext(ctx, fetchQuery, tokenString, token.RESET).Scan(
		&u.ID,
		&u.Password.hash,
		&u.LastUpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return customerrors.ErrNoRecord
		}
		return err
	}

	err = u.Password.Set(newPassword)
	if err != nil {
		return err
	}

	updateQuery := `
		UPDATE users
		set password_hash = $1,
			last_updated_at = $2
		WHERE id = $3
	`
	_, err = tx.ExecContext(ctx, updateQuery, u.Password.hash, time.Now(), u.ID)
	if err != nil {
		return err
	}

	deleteQurey := `
		DELETE FROM tokens
		WHERE token_string = $1
	`
	res, err := tx.ExecContext(ctx, deleteQurey, tokenString)
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

	if err := tx.Commit(); err != nil && !errors.Is(err, sql.ErrTxDone) {
		return err
	}

	return nil
}

func (r *Repository) GetUserByProvider(providerName, userSub string) (*User, error) {
	query := `
		SELECT u.id, u.created_at, u.last_updated_at, u.name, u.email
		FROM users u
		INNER JOIN providers p
		ON u.id = p.user_id
		WHERE p.provider_name = $1 AND p.sub = $2
	`
	// to prevent waiting indefinitely
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	u := &User{}
	err := r.DB.QueryRowContext(ctx, query, providerName, userSub).Scan(
		&u.ID,
		&u.CreatedAt,
		&u.LastUpdatedAt,
		&u.Name,
		&u.Email,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, customerrors.ErrNoRecord
		default:
			return nil, err
		}
	}
	return u, nil
}
