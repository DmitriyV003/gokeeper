package crypt

import (
	"crypto/tls"
	"google.golang.org/grpc/credentials"
)

func LoadClientCertificate(sslCertPath, sslKeyPath string) (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair(sslCertPath, sslKeyPath)
	if err != nil {
		return nil, err
	}

	return credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}), nil
}

func LoadServerCertificate(sslCertPath, sslKeyPath string) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(sslCertPath, sslKeyPath)
	if err != nil {
		return nil, err
	}

	return &tls.Config{Certificates: []tls.Certificate{cert}}, nil
}
