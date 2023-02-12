package services

import (
	"context"
	"fmt"
	"gokeeper/internal/core"
	"gokeeper/internal/data/sqlite"
	"gokeeper/internal/proto"
	"gokeeper/pkg/crypt"
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

func (l *LoginSecretService) Create(ctx context.Context, username, website, password, additionalData string) error {
	jwt, _, _ := l.settingsRepo.Get(ctx, "jwt")
	req := proto.CreateLoginSecretRequest{
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
		Username:       username,
		Website:        website,
		Password:       encPassword,
		AdditionalData: additionalData,
		ID:             res.ID,
	})

	return nil
}

func (l *LoginSecretService) Delete(ctx context.Context, id int64) error {
	if err := l.loginSecretRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete logins secret: %w", err)
	}

	return nil
}

func (l *LoginSecretService) GetAll(ctx context.Context, userID int) ([]*core.LoginSecret, error) {
	logins, err := l.loginSecretRepo.GetAll(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get logins secret by id: %w", err)
	}

	return logins, nil
}

func (l *LoginSecretService) Get(ctx context.Context, id int64) (*core.LoginSecret, error) {
	secret, err := l.loginSecretRepo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get login secret by id: %w", err)
	}
	if secret == nil {
		return nil, nil
	}

	aesSecret, _, err := l.settingsRepo.Get(ctx, "aes_secret")
	if err != nil {
		return nil, fmt.Errorf("get decrypted AES secret: %w", err)
	}

	password, err := crypt.DecryptAES([]byte(aesSecret), secret.Password)
	if err != nil {
		return nil, fmt.Errorf("decrypt password on getting login secret: %w", err)
	}

	secret.Password = password

	return secret, nil
}
