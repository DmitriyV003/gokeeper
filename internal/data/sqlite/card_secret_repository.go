package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gokeeper/internal/core"
	"time"
)

type CardSecretRepository struct {
	db *SQLite
}

func NewCardSecretRepository(db *SQLite) *CardSecretRepository {
	return &CardSecretRepository{
		db: db,
	}
}

func (r *CardSecretRepository) Create(ctx context.Context, request core.CardSecret) error {
	sql := `INSERT INTO cards (cardholder_name, type, expire_date, valid_from, number, secret_code, additional_data, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`

	if _, err := r.db.ExecContext(ctx, sql, request.CardholderName, request.Type, request.ExpireDate, request.ValidFrom, request.Number, request.SecretCode, request.AdditionalData, time.Now()); err != nil {
		return err
	}

	return nil
}

func (r *CardSecretRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM cards WHERE id = $1;`

	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("deleting card from SQLite: %w", err)
	}

	return nil
}

func (r *CardSecretRepository) GetById(ctx context.Context, id int64) (*core.CardSecret, error) {
	query := `SELECT id, website, login, enc_password, additional_data, user_id FROM cards WHERE id = $1;`

	secret := new(core.CardSecret)
	if err := r.db.QueryRowContext(ctx, query, id).Scan(
		&secret.ID,
		&secret.CardholderName,
		&secret.SecretCode,
		&secret.Number,
		&secret.ValidFrom,
		&secret.ExpireDate,
		&secret.AdditionalData,
		&secret.UserID,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return secret, nil
}

func (r *CardSecretRepository) GetAll(ctx context.Context, userID int) ([]*core.CardSecret, error) {
	query := `SELECT id, website, login, enc_password, additional_data, user_id
		FROM cards
		WHERE user_id = $1
		ORDER BY website, login;`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	list := make([]*core.CardSecret, 0)

	for rows.Next() {
		secret := new(core.CardSecret)
		if err = rows.Scan(
			&secret.ID,
			&secret.CardholderName,
			&secret.SecretCode,
			&secret.Number,
			&secret.ValidFrom,
			&secret.ExpireDate,
			&secret.AdditionalData,
			&secret.UserID,
		); err != nil {
			return nil, err
		}

		list = append(list, secret)
	}

	return list, nil
}
