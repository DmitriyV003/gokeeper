package server

import (
	"context"
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
	id, _ := l.secretService.CreateLoginSecret(ctx, req)

	return &proto.SecretSecretResponse{
		ID: id,
	}, nil
}

func (l *LoginSecretServer) UpdateLoginSecret(context.Context, *proto.UpdateLoginSecretRequest) (*proto.SecretSecretResponse, error) {
	panic("Unimpemented")
}

func (l *LoginSecretServer) DeleteLoginSecret(context.Context, *proto.DeleteLoginSecretRequest) (*proto.SecretSecretResponse, error) {
	panic("Unimpemented")
}
