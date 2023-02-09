package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
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
	data := []byte(mes)
	aesSecret, _, err := s.keysService.DecodeKeys(s.aesEncPrivateKey, s.rsaEncPrivateKey)
	if err != nil {
		return "", fmt.Errorf("error to decode keys: %w", err)
	}

	c, err := aes.NewCipher(aesSecret)
	if err != nil {
		return "", fmt.Errorf("error to generate new cipher: %w", err)
	}

	gcmDecrypt, err := cipher.NewGCM(c)
	if err != nil {
		return "", fmt.Errorf("error to generate NewGCM: %w", err)
	}

	nonceSize := gcmDecrypt.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("message len in less than nonsize: %w", err)
	}

	nonce, encryptedMessage := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcmDecrypt.Open(nil, nonce, encryptedMessage, nil)
	if err != nil {
		return "", fmt.Errorf("error to decrypt: %w", err)
	}
	fmt.Println(string(plaintext))

	return string(plaintext), nil
}

func (s *SecurityService) CryptMessage(mes string) (string, error) {
	aesSecret, _, err := s.keysService.DecodeKeys(s.aesEncPrivateKey, s.rsaEncPrivateKey)
	if err != nil {
		return "", fmt.Errorf("error to decode keys: %w", err)
	}

	cphr, err := aes.NewCipher(aesSecret)
	if err != nil {
		return "", fmt.Errorf("error to generate new cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(cphr)
	if err != nil {
		return "", fmt.Errorf("error to generate NewGCM: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("error to readFull: %w", err)
	}
	enc := gcm.Seal(nonce, nonce, []byte(mes), nil)

	return string(enc), nil
}
