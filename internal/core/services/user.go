package services

import (
	"context"
	"fmt"
	"gokeeper/internal/proto"
	"strings"
)

var ErrLoginIsTaken = fmt.Errorf("this login is already taken")
var ErrCredentialsDontMatch = fmt.Errorf("the credentials don't match any of our records")

type UserService struct {
	client       proto.UserClient
	settingsRepo Settings
}

func NewUserService(client proto.UserClient, settingsRepo Settings) *UserService {
	return &UserService{
		client:       client,
		settingsRepo: settingsRepo,
	}
}

func (s *UserService) Register(ctx context.Context, login, password string) error {
	request := &proto.RegisterRequest{
		Login:    login,
		Password: password,
	}

	response, err := s.client.Register(ctx, request)
	if err != nil {
		if strings.Contains(err.Error(), "taken") {
			return ErrLoginIsTaken
		}

		return err
	}

	if err = s.rememberMe(ctx, response.Token, response.AesSecret, response.PrivateKey); err != nil {
		return fmt.Errorf("register remember me: %w", err)
	}

	return nil
}

func (s *UserService) Login(ctx context.Context, login, password string) error {
	request := &proto.LoginRequest{
		Login:    login,
		Password: password,
	}

	response, err := s.client.Login(ctx, request)
	if err != nil {
		if strings.Contains(err.Error(), "credentials") {
			return ErrCredentialsDontMatch
		}

		return err
	}

	if err = s.rememberMe(ctx, response.Token, response.AesSecret, response.PrivateKey); err != nil {
		return fmt.Errorf("login remember me: %w", err)
	}

	return nil
}

func (s *UserService) Delete(ctx context.Context) error {
	if err := s.settingsRepo.Truncate(ctx); err != nil {
		return err
	}

	return nil
}

func (s *UserService) rememberMe(ctx context.Context, jwt, aesSecret, privateKey string) error {
	if _, err := s.settingsRepo.Set(ctx, "jwt", jwt); err != nil {
		return err
	}

	if _, err := s.settingsRepo.Set(ctx, "aes_secret", aesSecret); err != nil {
		return err
	}

	if _, err := s.settingsRepo.Set(ctx, "private_key", privateKey); err != nil {
		return err
	}

	return nil
}
