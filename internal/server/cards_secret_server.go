package server

import (
	"context"
	"fmt"
	"gokeeper/internal/core/services"
	"gokeeper/internal/proto"
)

type CardsSecretServer struct {
	proto.UnimplementedCardSecretServiceServer
	secretService *services.CardReqSecretService
}

func NewCardsSecretServer(secretService *services.CardReqSecretService) *CardsSecretServer {
	return &CardsSecretServer{
		secretService: secretService,
	}
}

func (l *CardsSecretServer) CreateCardSecret(ctx context.Context, req *proto.CreateCardSecretRequest) (*proto.SecretCardResponse, error) {
	id, err := l.secretService.CreateCardSecret(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error to create login secret: %w", err)
	}

	return &proto.SecretCardResponse{
		ID: id,
	}, nil
}

func (l *CardsSecretServer) DeleteCardSecret(ctx context.Context, req *proto.DeleteCardSecretRequest) (*proto.SecretCardResponse, error) {
	err := l.secretService.DeleteCardSecret(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error to delete login secret: %w", err)
	}

	return &proto.SecretCardResponse{}, nil
}

func (l *CardsSecretServer) UpdateCardSecret(ctx context.Context, req *proto.UpdateCardSecretRequest) (*proto.SecretCardResponse, error) {
	err := l.secretService.UpdateCardSecret(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error to update login secret: %w", err)
	}

	return &proto.SecretCardResponse{}, nil
}
