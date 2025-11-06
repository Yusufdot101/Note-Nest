package token

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) InsertToken(token *Token) error {
	query := `
		INSERT INTO refresh_tokens
		(user_id, token_string, expires)
		VALUES ($1, $2, $3)
	`
	values := []any{
		token.UserID,
		token.TokenString,
		token.Expires,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.DB.ExecContext(ctx, query, values...)
	return err
}

func (r *Repository) GetByTokenString(tokenString string) (*Token, error) {
	query := `
		SELECT user_id, token_string FROM refresh_tokens
		WHERE token_string = $1 AND expires > NOW()
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	token := &Token{}
	err := r.DB.QueryRowContext(ctx, query, tokenString).Scan(
		&token.UserID,
		&token.TokenString,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, custom_errors.ErrNoRecord
		default:
			return nil, err
		}
	}

	return token, nil
}

func (r *Repository) DeleteByTokenString(tokenString string) error {
	query := `
		DELETE FROM refresh_tokens
		WHERE token_string = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.DB.ExecContext(ctx, query, tokenString)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return custom_errors.ErrNoRecord
	}

	return nil
}
