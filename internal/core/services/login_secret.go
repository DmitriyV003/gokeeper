package services

import (
	"context"
	"errors"
	"fmt"
	"gokeeper/internal/data/postgres"
	"gokeeper/internal/data/sqlite"
	"gokeeper/internal/proto"
)

type LoginSecretService struct {
	authService    *AuthService
	client         proto.LoginSecretServiceClient
	settingsRepo   *sqlite.SettingsRepository
	masterPassword string
	keysService    *KeysService
	usersRepo      *postgres.UserRepository
}

func NewLoginSecretService(authService *AuthService, client proto.LoginSecretServiceClient, settingsRepo *sqlite.SettingsRepository, masterPassword string, keysService *KeysService, userRepo *postgres.UserRepository) *LoginSecretService {
	return &LoginSecretService{
		authService:    authService,
		client:         client,
		settingsRepo:   settingsRepo,
		masterPassword: masterPassword,
		keysService:    keysService,
		usersRepo:      userRepo,
	}
}

func (l *LoginSecretService) Create(ctx context.Context, name, username, website, password, additionalData string) error {
	jwt, _, _ := l.settingsRepo.Get(ctx, "jwt")
	req := proto.CreateLoginSecretRequest{
		Name:           name,
		Username:       username,
		Website:        website,
		Password:       "",
		AdditionalData: additionalData,
	}

	token := Token{
		Value:  jwt,
		Claims: map[string]interface{}{},
	}

	l.authService.ParseTokenWithClaims(&token)

	login := token.Claims["sub"]
	strLogin := login.(string)
	if strLogin == "" {
		return errors.New("login empty")
	}
	user, err := l.usersRepo.GetByLogin(ctx, strLogin)
	if err != nil {
		return fmt.Errorf("error to get user: %w", err)
	}

	securityService := NewSecurityService(user.AesSecret, user.RsaSecret, l.masterPassword, l.keysService)
	encPassword, err := securityService.CryptMessage(password)
	if err != nil {
		return fmt.Errorf("error to encrypt secret: %w", err)
	}

	req.UserID = user.ID
	req.Password = encPassword
	l.client.CreateLoginSecret(ctx, &req)

	return nil
}
