package server

import (
	"context"
	"gokeeper/internal/core"
	"gokeeper/internal/proto"
)

type LoginSecretServer struct {
	proto.UnimplementedLoginServiceServer
	secretService *core.SecretService
}

func NewLoginSecretServer(secretService *core.SecretService) *LoginSecretServer {
	return &LoginSecretServer{
		secretService: secretService,
	}
}

func (l *LoginSecretServer) CreateLoginSecret(ctx context.Context, req *proto.CreateLoginSecretRequest) (*proto.SecretResponse, error) {
	l.secretService.CreateLoginSecret(ctx, req)

	return &proto.SecretResponse{}, nil
}

func (l *LoginSecretServer) UpdateLoginSecret(context.Context, *proto.UpdateLoginSecretRequest) (*proto.SecretResponse, error) {
	panic("Unimpemented")
}

func (l *LoginSecretServer) DeleteLoginSecret(context.Context, *proto.DeleteLoginSecretRequest) (*proto.SecretResponse, error) {
	panic("Unimpemented")
}
