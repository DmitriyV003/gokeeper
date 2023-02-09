package sqlite

import (
	"context"
	"gokeeper/internal/core"
	"time"
)

type LoginSecretRepository struct {
	db *SQLite
}

func NewLoginSecretRepository(db *SQLite) *LoginSecretRepository {
	return &LoginSecretRepository{
		db: db,
	}
}

func (r *LoginSecretRepository) Create(ctx context.Context, loginSecret core.LoginSecret) error {
	query := `INSERT INTO logins (name, username, website, password, additional_data, created_at) VALUES ($1, $2, $3, $4, $5, $6);`

	if _, err := r.db.ExecContext(ctx, query, loginSecret.Name, loginSecret.Username, loginSecret.Website, loginSecret.Password, loginSecret.AdditionalData, time.Now()); err != nil {
		return err
	}

	return nil
}
