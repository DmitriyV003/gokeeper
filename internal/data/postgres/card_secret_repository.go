package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"gokeeper/internal/core"
	"time"
)

type CardSecretRepository struct {
	db *pgxpool.Pool
}

func NewCardSecretRepository(db *pgxpool.Pool) *CardSecretRepository {
	return &CardSecretRepository{
		db: db,
	}
}

func (l *CardSecretRepository) Create(ctx context.Context, request core.CardSecret) (int64, error) {
	sql := `INSERT INTO cards (cardholder_name, type, expire_date, valid_from, number, secret_code, additional_data, created_at, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	var id int64
	err := l.db.QueryRow(ctx, sql, request.CardholderName, request.Type, request.ExpireDate, request.ValidFrom, request.Number, request.SecretCode, request.AdditionalData, time.Now(), request.UserID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("unable to insert cards to db: %w", err)
	}

	return id, nil
}

func (l *CardSecretRepository) UpdateByID(ctx context.Context, request core.CardSecret) error {
	sql := `UPDATE cards SET cardholder_name = $1, type = $2, expire_date = $3, valid_from = $4, number = $5, secret_code = $6, additional_data = $7, updated_at = $8`

	_, err := l.db.Exec(ctx, sql, request.CardholderName, request.Type, request.ExpireDate, request.ValidFrom, request.Number, request.SecretCode, request.AdditionalData, time.Now())
	if err != nil {
		return fmt.Errorf("unable to update cards in db: %w", err)
	}

	return nil
}

func (l *CardSecretRepository) DeleteById(ctx context.Context, id int64) error {
	sql := `DELETE cards WHERE id = $1`

	_, err := l.db.Exec(ctx, sql, id)
	if err != nil {
		return fmt.Errorf("unable to update cards in db: %w", err)
	}

	return nil
}
