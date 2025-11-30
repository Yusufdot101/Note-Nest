package token

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) InsertToken(token *Token) error {
	query := `
		INSERT INTO tokens
		(use, user_id, token_string, expires)
		VALUES ($1, $2, $3, $4)
	`
	values := []any{
		token.Use,
		token.UserID,
		token.TokenString,
		token.Expires,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.DB.ExecContext(ctx, query, values...)
	return err
}

func (r *Repository) GetByTokenStringAndUse(tokenString string, tokenUse TokenUse) (*Token, error) {
	query := `
		SELECT user_id, token_string FROM tokens
		WHERE token_string = $1 AND use = $2
			AND expires > NOW()
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	token := &Token{}
	err := r.DB.QueryRowContext(ctx, query, tokenString, tokenUse).Scan(
		&token.UserID,
		&token.TokenString,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, customerrors.ErrNoRecord
		default:
			return nil, err
		}
	}

	return token, nil
}

func (r *Repository) DeleteByTokenStringAndUse(tokenString string, tokenUse TokenUse) error {
	query := `
		DELETE FROM tokens
		WHERE token_string = $1 AND use = $2
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.DB.ExecContext(ctx, query, tokenString, tokenUse)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return customerrors.ErrNoRecord
	}

	return nil
}
