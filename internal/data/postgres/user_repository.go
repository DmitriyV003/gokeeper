package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"gokeeper/internal/applicationerrors"
	"gokeeper/internal/core"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: pool}
}

func (users *UserRepository) Create(ctx context.Context, user *core.User) error {
	sql := `INSERT INTO users (login, password, aes_secret, rsa_secret, created_at) VALUES ($1, $2, $3, $4, $5)`

	dbUser, err := users.GetByLogin(ctx, user.Login)
	if err != nil && !errors.Is(err, applicationerrors.ErrNotFound) {
		return applicationerrors.ErrInternalServer
	}

	if dbUser != nil {
		return applicationerrors.ErrConflict
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return applicationerrors.ErrNotFound
	}

	_, err = users.db.Exec(ctx, sql, user.Login, user.Password, user.AesSecret, user.RsaSecret, user.CreatedAt)

	if err != nil {
		return applicationerrors.ErrInternalServer
	}

	return nil
}

func (users *UserRepository) GetByLogin(ctx context.Context, login string) (*core.User, error) {
	sql := `SELECT id, login, password, aes_secret, rsa_secret, created_at FROM users WHERE login = $1`
	var user core.User

	row := users.db.QueryRow(ctx, sql, login)

	err := row.Scan(&user.ID, &user.Login, &user.Password, &user.AesSecret, &user.RsaSecret, &user.CreatedAt)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, applicationerrors.ErrNotFound
	}

	return &user, nil
}

func (users *UserRepository) GetByID(ctx context.Context, id int64) (*core.User, error) {
	sql := `SELECT id, login, created_at FROM users WHERE id = $1`
	var user core.User

	row := users.db.QueryRow(ctx, sql, id)

	err := row.Scan(&user.ID, &user.Login, &user.CreatedAt)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, applicationerrors.ErrNotFound
	}

	return &user, nil
}
