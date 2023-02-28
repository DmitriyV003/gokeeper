package services

import (
	"context"
	"fmt"
	"gokeeper/internal/core"
	"gokeeper/internal/data/postgres"
	"gokeeper/internal/proto"
)

type CardReqSecretService struct {
	secretsRepo *postgres.CardSecretRepository
}

func NewCardReqSecretService(secretsRepo *postgres.CardSecretRepository) *CardReqSecretService {
	return &CardReqSecretService{
		secretsRepo: secretsRepo,
	}
}

func (s *CardReqSecretService) CreateCardSecret(ctx context.Context, req *proto.CreateCardSecretRequest) (int64, error) {
	cardSecret := core.CardSecret{
		CardholderName: req.CardholderName,
		Type:           req.Type,
		ExpireDate:     req.ExpireDate,
		ValidFrom:      req.ValidFrom,
		Number:         req.Number,
		AdditionalData: req.AdditionalData,
		SecretCode:     req.SecretCode,
		UserID:         req.UserID,
	}

	id, err := s.secretsRepo.Create(ctx, cardSecret)
	if err != nil {
		return 0, fmt.Errorf("error to save login: %w", err)
	}

	return id, nil
}

func (s *CardReqSecretService) UpdateCardSecret(ctx context.Context, req *proto.UpdateCardSecretRequest) error {
	cardSecret := core.CardSecret{
		CardholderName: req.CardholderName,
		Type:           req.Type,
		ExpireDate:     req.ExpireDate,
		ValidFrom:      req.ValidFrom,
		Number:         req.Number,
		AdditionalData: req.AdditionalData,
		SecretCode:     req.SecretCode,
		UserID:         req.UserID,
	}

	err := s.secretsRepo.UpdateByID(ctx, cardSecret)
	if err != nil {
		return fmt.Errorf("error to update login: %w", err)
	}

	return nil
}

func (s *CardReqSecretService) DeleteCardSecret(ctx context.Context, req *proto.DeleteCardSecretRequest) error {
	err := s.secretsRepo.DeleteById(ctx, req.Id)
	if err != nil {
		return fmt.Errorf("error to delete login: %w", err)
	}

	return nil
}
