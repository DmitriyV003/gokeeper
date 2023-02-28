package server

import (
	"context"
	"fmt"
	"gokeeper/internal/core/services"
	"gokeeper/internal/proto"
)

type LoginSecretServer struct {
	proto.UnimplementedLoginSecretServiceServer
	secretService *services.SecretService
}

func NewLoginSecretServer(secretService *services.SecretService) *LoginSecretServer {
	return &LoginSecretServer{
		secretService: secretService,
	}
}

func (l *LoginSecretServer) CreateLoginSecret(ctx context.Context, req *proto.CreateLoginSecretRequest) (*proto.SecretSecretResponse, error) {
	id, err := l.secretService.CreateLoginSecret(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error to create login secret: %w", err)
	}

	return &proto.SecretSecretResponse{
		ID: id,
	}, nil
}

func (l *LoginSecretServer) DeleteLoginSecret(ctx context.Context, req *proto.DeleteLoginSecretRequest) (*proto.SecretSecretResponse, error) {
	err := l.secretService.DeleteLoginSecret(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error to delete login secret: %w", err)
	}

	return &proto.SecretSecretResponse{}, nil
}

func (l *LoginSecretServer) UpdateLoginSecret(ctx context.Context, req *proto.UpdateLoginSecretRequest) (*proto.SecretSecretResponse, error) {
	err := l.secretService.UpdateLoginSecret(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error to update login secret: %w", err)
	}

	return &proto.SecretSecretResponse{}, nil
}
