package services

import (
	"context"
	"fmt"
	"gokeeper/internal/core"
	"gokeeper/internal/data/postgres"
	"gokeeper/internal/proto"
)

type SecretService struct {
	secretsRepo *postgres.LoginSecretRepository
}

func NewSecretService(secretsRepo *postgres.LoginSecretRepository) *SecretService {
	return &SecretService{
		secretsRepo: secretsRepo,
	}
}

func (s *SecretService) CreateLoginSecret(ctx context.Context, req *proto.CreateLoginSecretRequest) error {
	loginSecret := core.LoginSecret{
		Name:           req.Name,
		Username:       req.Username,
		Website:        req.Website,
		Password:       req.Password,
		AdditionalData: req.AdditionalData,
		ID:             0,
	}

	err := s.secretsRepo.Create(ctx, loginSecret)
	if err != nil {
		return fmt.Errorf("error to save login: %w", err)
	}

	return nil
}

func (s *SecretService) UpdateLoginSecret(ctx context.Context, req *proto.UpdateLoginSecretRequest) error {
	loginSecret := core.LoginSecret{
		Name:           req.Name,
		Username:       req.Username,
		Website:        req.Website,
		Password:       req.Password,
		AdditionalData: req.AdditionalData,
		ID:             req.ID,
		UserID:         req.UserID,
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
