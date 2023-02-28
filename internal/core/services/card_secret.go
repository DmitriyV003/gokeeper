package services

import (
	"context"
	"fmt"
	"gokeeper/internal/core"
	"gokeeper/internal/data/sqlite"
	"gokeeper/internal/proto"
	"gokeeper/pkg/crypt"
)

type CardSecretService struct {
	authService    *AuthService
	client         proto.CardSecretServiceClient
	settingsRepo   *sqlite.SettingsRepository
	masterPassword string
	keysService    *KeysService
	secretRepo     *sqlite.CardSecretRepository
}

func NewCardSecretService(
	authService *AuthService,
	client proto.CardSecretServiceClient,
	settingsRepo *sqlite.SettingsRepository,
	masterPassword string,
	keysService *KeysService,
	loginSecretRepo *sqlite.CardSecretRepository,
) *CardSecretService {
	return &CardSecretService{
		authService:    authService,
		client:         client,
		settingsRepo:   settingsRepo,
		masterPassword: masterPassword,
		keysService:    keysService,
		secretRepo:     loginSecretRepo,
	}
}

func (l *CardSecretService) Create(ctx context.Context, cardholderName, typ, expireDate, validFrom, additionalData, number, secretCode string) error {
	jwt, _, _ := l.settingsRepo.Get(ctx, "jwt")
	req := proto.CreateCardSecretRequest{
		CardholderName: cardholderName,
		Type:           typ,
		ExpireDate:     expireDate,
		ValidFrom:      validFrom,
		AdditionalData: additionalData,
		Number:         number,
		SecretCode:     secretCode,
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
	encPassword, err := securityService.CryptMessage(secretCode)
	if err != nil {
		return fmt.Errorf("error to encrypt secret: %w", err)
	}

	req.UserID = pId
	req.SecretCode = encPassword
	res, _ := l.client.CreateCardSecret(ctx, &req)

	l.secretRepo.Create(ctx, core.CardSecret{
		CardholderName: cardholderName,
		Type:           typ,
		ExpireDate:     expireDate,
		ValidFrom:      validFrom,
		Number:         number,
		AdditionalData: additionalData,
		SecretCode:     encPassword,
		ID:             res.ID,
	})

	return nil
}

func (l *CardSecretService) Delete(ctx context.Context, id int64) error {
	if err := l.secretRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete logins secret: %w", err)
	}

	jwt, _, _ := l.settingsRepo.Get(ctx, "jwt")
	token := Token{
		Value:  jwt,
		Claims: map[string]interface{}{},
	}
	l.authService.ParseTokenWithClaims(&token)

	userId := token.Claims["sub"]
	pUserId := userId.(int64)

	req := proto.DeleteCardSecretRequest{
		Id:     id,
		UserID: pUserId,
	}
	_, err := l.client.DeleteCardSecret(ctx, &req)
	if err != nil {
		return fmt.Errorf("error to delete login secret: %w", err)
	}

	return nil
}

func (l *CardSecretService) GetAll(ctx context.Context, userID int) ([]*core.CardSecret, error) {
	secrets, err := l.secretRepo.GetAll(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get logins secret by id: %w", err)
	}

	return secrets, nil
}

func (l *CardSecretService) Get(ctx context.Context, id int64) (*core.CardSecret, error) {
	secret, err := l.secretRepo.GetById(ctx, id)
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

	password, err := crypt.DecryptAES([]byte(aesSecret), secret.SecretCode)
	if err != nil {
		return nil, fmt.Errorf("decrypt password on getting login secret: %w", err)
	}

	secret.SecretCode = password

	return secret, nil
}
