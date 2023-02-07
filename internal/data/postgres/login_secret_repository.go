package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"gokeeper/internal/core"
	"time"
)

type LoginSecretRepository struct {
	db *pgxpool.Pool
}

func NewLoginSecretRepository(db *pgxpool.Pool) *LoginSecretRepository {
	return &LoginSecretRepository{
		db: db,
	}
}

func (l *LoginSecretRepository) Create(ctx context.Context, request core.LoginSecret) error {
	sql := `INSERT INTO logins (name, username, website, password, additional_data, created_at, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	err := l.db.QueryRow(ctx, sql, request.Name, request.Username, request.Website, request.Password, request.AdditionalData, time.Now(), request.UserID)
	if err != nil {
		return fmt.Errorf("unable to insert login to db: %w", err)
	}

	return nil
}

func (l *LoginSecretRepository) UpdateByID(ctx context.Context, request core.LoginSecret) error {
	sql := `UPDATE logins SET name = $1, username = $2, website = $3, password = $4, additional_data = $5, updated_at = $6 WHERE id = $7`

	_, err := l.db.Exec(ctx, sql, request.Name, request.Username, request.Website, request.Password, request.AdditionalData, time.Now(), request.ID)
	if err != nil {
		return fmt.Errorf("unable to update login in db: %w", err)
	}

	return nil
}

func (l *LoginSecretRepository) DeleteById(ctx context.Context, id int64) error {
	sql := `DELETE logins WHERE id = $1`

	_, err := l.db.Exec(ctx, sql, id)
	if err != nil {
		return fmt.Errorf("unable to update login in db: %w", err)
	}

	return nil
}
