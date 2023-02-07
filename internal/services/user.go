package services

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"gokeeper/internal/core"
	"gokeeper/internal/core/services"
	"gokeeper/internal/data/postgres"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService struct {
	keysService *services.KeysService
	repo        *postgres.UserRepository
}

func NewUserService(keysService *services.KeysService, repo *postgres.UserRepository) *UserService {
	return &UserService{
		repo:        repo,
		keysService: keysService,
	}
}

func (s *UserService) Create(ctx context.Context, login, password string) (*core.User, error) {
	user, err := s.repo.GetByLogin(ctx, login)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, services.ErrLoginIsTaken
	}

	hashedPassword, err := s.hashPassword(password)
	if err != nil {
		return nil, err
	}

	aesSecret, privateKey, err := s.keysService.GenerateKeys()
	if err != nil {
		log.Debug().Msg("1234")
		return nil, fmt.Errorf("generating keys for new user: %w", err)
	}
	u := core.User{
		ID:        0,
		Login:     login,
		Password:  hashedPassword,
		AesSecret: aesSecret,
		RsaSecret: privateKey,
		CreatedAt: time.Now(),
	}

	if err = s.repo.Create(ctx, &u); err != nil {
		return nil, err
	}

	return s.repo.GetByLogin(ctx, login)
}

func (s *UserService) FindByLogin(ctx context.Context, login string) (*core.User, error) {
	return s.repo.GetByLogin(ctx, login)
}

func (s *UserService) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
