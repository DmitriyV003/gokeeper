package services

import (
	"context"
	"fmt"
	"gokeeper/internal/core"
	"gokeeper/internal/data/sqlite"
	"gokeeper/internal/proto"
)

type LoginSecretService struct {
	authService     *AuthService
	client          proto.LoginSecretServiceClient
	settingsRepo    *sqlite.SettingsRepository
	masterPassword  string
	keysService     *KeysService
	loginSecretRepo *sqlite.LoginSecretRepository
}

func NewLoginSecretService(
	authService *AuthService,
	client proto.LoginSecretServiceClient,
	settingsRepo *sqlite.SettingsRepository,
	masterPassword string,
	keysService *KeysService,
	loginSecretRepo *sqlite.LoginSecretRepository,
) *LoginSecretService {
	return &LoginSecretService{
		authService:     authService,
		client:          client,
		settingsRepo:    settingsRepo,
		masterPassword:  masterPassword,
		keysService:     keysService,
		loginSecretRepo: loginSecretRepo,
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

	id := token.Claims["sub"]
	pId := id.(int64)
	aesSecret, _, err := l.settingsRepo.Get(ctx, "aes_secret")
	if err != nil {
		return fmt.Errorf("error to get user: %w", err)
	}

	rsaSecret, _, err := l.settingsRepo.Get(ctx, "private_key")
	if err != nil {
		return fmt.Errorf("error to get user: %w", err)
	}

	securityService := NewSecurityService(aesSecret, rsaSecret, l.masterPassword, l.keysService)
	encPassword, err := securityService.CryptMessage(password)
	if err != nil {
		return fmt.Errorf("error to encrypt secret: %w", err)
	}

	req.UserID = pId
	req.Password = encPassword
	res, _ := l.client.CreateLoginSecret(ctx, &req)

	l.loginSecretRepo.Create(ctx, core.LoginSecret{
		Name:           name,
		Username:       username,
		Website:        website,
		Password:       encPassword,
		AdditionalData: additionalData,
		ID:             res.ID,
	})

	return nil
}
