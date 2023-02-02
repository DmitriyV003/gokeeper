package core

import (
	"context"
	"fmt"
	"gokeeper/internal/data"
	"gokeeper/internal/proto"
)

type SecretService struct {
	secretsRepo *data.LoginSecretRepository
}

func NewSecretService(secretsRepo *data.LoginSecretRepository) *SecretService {
	return &SecretService{
		secretsRepo: secretsRepo,
	}
}

func (s *SecretService) CreateLoginSecret(ctx context.Context, req *proto.CreateLoginSecretRequest) (string, error) {
	loginSecret := LoginSecret{
		Name:           req.Name,
		Username:       req.Username,
		Website:        req.Website,
		Password:       req.Password,
		AdditionalData: req.AdditionalData,
		ID:             0,
	}

	err := s.secretsRepo.Create(ctx, loginSecret)
	if err != nil {
		return "", fmt.Errorf("error to save login: %w", err)
	}

	return "", nil
}

func (s *SecretService) UpdateLoginSecret(ctx context.Context, req *proto.UpdateLoginSecretRequest) error {
	loginSecret := LoginSecret{
		Name:           req.Name,
		Username:       req.Username,
		Website:        req.Website,
		Password:       req.Password,
		AdditionalData: req.AdditionalData,
		ID:             req.ID,
	}

	err := s.secretsRepo.UpdateByID(ctx, loginSecret)
	if err != nil {
		return fmt.Errorf("error to update login: %w", err)
	}

	return nil
}

func (s *SecretService) DeleteLoginSecret(ctx context.Context, req *proto.DeleteLoginSecretRequest) error {
	err := s.secretsRepo.DeleteById(ctx, req.Id)
	if err != nil {
		return fmt.Errorf("error to delete login: %w", err)
	}

	return nil
}
