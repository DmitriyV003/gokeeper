package client

import (
	"context"
	"github.com/rs/zerolog/log"
	"gokeeper/internal/proto"
	"gokeeper/pkg/crypt"
	"google.golang.org/grpc"
)

func NewUserClient(ctx context.Context, address string, sslCertPath string, sslKeyPath string) proto.UserClient {
	tlsCredential, err := crypt.LoadClientCertificate(sslCertPath, sslKeyPath)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("cert-path", sslCertPath).
			Str("key-path", sslKeyPath).
			Msg("Loading client TLS cert")
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(tlsCredential))
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

func NewLoginSecretClient(ctx context.Context, address string, sslCertPath string, sslKeyPath string) proto.LoginSecretServiceClient {
	tlsCredential, err := crypt.LoadClientCertificate(sslCertPath, sslKeyPath)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("cert-path", sslCertPath).
			Str("key-path", sslKeyPath).
			Msg("Loading client TLS cert")
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(tlsCredential))
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

func NewCardSecretClient(ctx context.Context, address string, sslCertPath string, sslKeyPath string) proto.CardSecretServiceClient {
	tlsCredential, err := crypt.LoadClientCertificate(sslCertPath, sslKeyPath)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("cert-path", sslCertPath).
			Str("key-path", sslKeyPath).
			Msg("Loading client TLS cert")
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(tlsCredential))
	if err != nil {
		log.Fatal().Err(err).Msg("Connecting to gRPC server")
	}

	go func() {
		<-ctx.Done()
		if err = conn.Close(); err != nil {
			log.Fatal().Err(err).Msg("Closing gRPC connection")
		}
	}()

	return proto.NewCardSecretServiceClient(conn)
}
