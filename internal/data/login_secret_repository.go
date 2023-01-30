package data

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"gokeeper/internal/proto"
	"time"
)

type LoginSecretRepository struct {
	db *pgxpool.Pool
}

func (l *LoginSecretRepository) Create(ctx context.Context, request proto.CreateLoginSecretRequest) error {
	sql := `INSERT INTO logins (name, username, website, password, additional_data, created_at) VALUES ($1, $2, $3, $4, $5, $6)`

	err := l.db.QueryRow(ctx, sql, request.Name, request.Username, request.Website, request.Password, request.AdditionalData, time.Now())
	if err != nil {
		return fmt.Errorf("unable to insert login to db: %w", err)
	}

	return nil
}

func (l *LoginSecretRepository) UpdateByID(ctx context.Context, request proto.UpdateLoginSecretRequest) error {
	sql := `UPDATE logins SET name = $1, username = $2, website = $3, password = $4, additional_data = $5, updated_at = $6 WHERE id = $7`

	_, err := l.db.Exec(ctx, sql, request.Name, request.Username, request.Website, request.Password, request.AdditionalData, time.Now(), request.ID)
	if err != nil {
		return fmt.Errorf("unable to update login in db: %w", err)
	}

	return nil
}
