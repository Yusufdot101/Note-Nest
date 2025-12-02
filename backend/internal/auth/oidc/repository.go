package oidc

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
)

type repository struct {
	DB *sql.DB
}

func (r *repository) insert(ui *userInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Println("Rollback error: ", err)
		}
	}()

	insertUsersQuery := `
		INSERT INTO users
		(name, email)
		VALUES ($1, $2)
		RETURNING id
	`

	err = tx.QueryRowContext(ctx, insertUsersQuery, ui.Name, ui.Email).Scan(
		&ui.UserID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return customerrors.ErrNoRecord
		}
		return err
	}

	insertProvidersQuery := `
		INSERT INTO providers
		(provider_name, user_id, sub)
		VALUES ($1, $2, $3)
	`

	_, err = tx.ExecContext(ctx, insertProvidersQuery, ui.ProviderName, ui.UserID, ui.Sub)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
