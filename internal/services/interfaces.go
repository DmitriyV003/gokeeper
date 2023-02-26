package services

import (
	"context"
	"gokeeper/internal/core"
)

type UserRepo interface {
	GetByLogin(ctx context.Context, login string) (*core.User, error)
	Create(ctx context.Context, user *core.User) error
}

type KeyService interface {
	GenerateKeys() (string, string, error)
}
