package server

import (
	"context"
	"gokeeper/internal/proto"
)

type LoginSecretServer struct {
	proto.UnimplementedLoginServiceServer
}

func NewLoginSecretServer() *LoginSecretServer {
	return &LoginSecretServer{}
}

func (l *LoginSecretServer) CreateLoginSecret(ctx context.Context, req *proto.CreateLoginSecretRequest) (*proto.SecretResponse, error) {
	panic("Unimpemented")
}

func (l *LoginSecretServer) UpdateLoginSecret(context.Context, *proto.UpdateLoginSecretRequest) (*proto.SecretResponse, error) {
	panic("Unimpemented")
}

func (l *LoginSecretServer) DeleteLoginSecret(context.Context, *proto.DeleteLoginSecretRequest) (*proto.SecretResponse, error) {
	panic("Unimpemented")
}
