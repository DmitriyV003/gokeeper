package services

import (
	"fmt"
	"gokeeper/pkg/crypt"
)

type SecurityService struct {
	aesEncPrivateKey string
	rsaEncPrivateKey string
	masterPassword   string
	keysService      *KeysService
}

func NewSecurityService(aesEncPrivateKey string, rsaEncPrivateKey string, masterPassword string, keysService *KeysService) *SecurityService {
	return &SecurityService{
		aesEncPrivateKey: aesEncPrivateKey,
		rsaEncPrivateKey: rsaEncPrivateKey,
		masterPassword:   masterPassword,
		keysService:      keysService,
	}
}

func (s *SecurityService) DecryptMessage(mes string) (string, error) {
	aes, _, err := s.keysService.DecodeKeys(s.aesEncPrivateKey, s.rsaEncPrivateKey)
	if err != nil {
		return "", fmt.Errorf("error to decode keys: %w", err)
	}

	res, err := crypt.DecryptAES(aes, mes)
	if err != nil {
		return "", fmt.Errorf("error to decrypt AES message: %w", err)
	}

	return res, nil
}

func (s *SecurityService) CryptMessage(mes string) (string, error) {
	aes, _, err := s.keysService.DecodeKeys(s.aesEncPrivateKey, s.rsaEncPrivateKey)
	if err != nil {
		return "", fmt.Errorf("error to decode keys: %w", err)
	}

	res, err := crypt.EncryptAES(aes, mes)
	if err != nil {
		return "", fmt.Errorf("error to crypt AES message: %w", err)
	}

	return res, nil
}
