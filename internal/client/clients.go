package client

import (
	"context"
	"github.com/rs/zerolog/log"
	"gokeeper/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(ctx context.Context, address string) proto.UserClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err).Msg("Connecting to gRPC server")
	}

	go func() {
		<-ctx.Done()
		if err = conn.Close(); err != nil {
			log.Fatal().Err(err).Msg("Closing gRPC connection")
		}
	}()

	return proto.NewUserClient(conn)
}

func NewLoginSecretClient(ctx context.Context, address string) proto.LoginSecretServiceClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err).Msg("Connecting to gRPC server")
	}

	go func() {
		<-ctx.Done()
		if err = conn.Close(); err != nil {
			log.Fatal().Err(err).Msg("Closing gRPC connection")
		}
	}()

	return proto.NewLoginSecretServiceClient(conn)
}
