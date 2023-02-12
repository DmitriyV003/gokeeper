package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	query := `INSERT INTO login_secrets (username, website, password, additional_data, created_at) VALUES ($1, $2, $3, $4, $5);`

	if _, err := r.db.ExecContext(ctx, query, loginSecret.Username, loginSecret.Website, loginSecret.Password, loginSecret.AdditionalData, time.Now()); err != nil {
		return err
	}

	return nil
}

func (r *LoginSecretRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM login_secrets WHERE id = $1;`

	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("deleting login from SQLite: %w", err)
	}

	return nil
}

func (r *LoginSecretRepository) GetById(ctx context.Context, id int64) (*core.LoginSecret, error) {
	query := `SELECT id, website, login, enc_password, additional_data, user_id FROM login_secrets WHERE id = $1;`

	login := new(core.LoginSecret)
	if err := r.db.QueryRowContext(ctx, query, id).Scan(
		&login.ID,
		&login.Website,
		&login.Username,
		&login.Password,
		&login.AdditionalData,
		&login.UserID,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return login, nil
}

func (r *LoginSecretRepository) GetAll(ctx context.Context, userID int) ([]*core.LoginSecret, error) {
	query := `SELECT id, website, login, enc_password, additional_data, user_id
		FROM login_secrets
		WHERE user_id = $1
		ORDER BY website, login;`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	list := make([]*core.LoginSecret, 0)

	for rows.Next() {
		secret := new(core.LoginSecret)
		if err = rows.Scan(
			&secret.ID,
			&secret.Website,
			&secret.Username,
			&secret.Password,
			&secret.AdditionalData,
			&secret.UserID,
		); err != nil {
			return nil, err
		}

		list = append(list, secret)
	}

	return list, nil
}
