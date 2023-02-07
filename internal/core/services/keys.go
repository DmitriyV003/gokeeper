package services

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"fmt"
	"gokeeper/pkg/crypt"
)

type KeysService struct {
	masterPassword string
}

func NewKeysService(masterPassword string) *KeysService {
	return &KeysService{
		masterPassword: masterPassword,
	}
}

// GenerateKeys создаёт AES ключ и приватный ключ, который в свою очередь шифрует AES.
//
// Возвращает ключи в зашифрованном виде строкой, где первый ключ - AES, а второй - приватный.
func (s *KeysService) GenerateKeys() (string, string, error) {
	aesSecret := make([]byte, 32)
	_, err := rand.Read(aesSecret)
	if err != nil {
		return "", "", fmt.Errorf("generating random key as AES secret: %w", err)
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", "", fmt.Errorf("generating RSA keys: %w", err)
	}

	encAesSecretBytes, err := crypt.EncryptRSA(&privateKey.PublicKey, aesSecret)
	if err != nil {
		return "", "", fmt.Errorf("encrypting AES via private key: %w", err)
	}
	encAesSecret := hex.EncodeToString(encAesSecretBytes)

	privateKeyInHexString := hex.EncodeToString(crypt.PrivateKeyToPemBytes(privateKey))

	masterPassword, err := hex.DecodeString(s.masterPassword)
	if err != nil {
		return "", "", fmt.Errorf("decoding hexed master-password: %w", err)
	}

	encPrivateKey, err := crypt.EncryptAES(masterPassword, privateKeyInHexString)
	if err != nil {
		return "", "", fmt.Errorf("encrypting private key via AES using master-password: %w", err)
	}

	return encAesSecret, encPrivateKey, nil
}

func (s *KeysService) DecodeKeys(encAesSecret, encPrivateKey string) ([]byte, *rsa.PrivateKey, error) {
	privateKeyInHexString, err := crypt.DecryptAES([]byte(s.masterPassword), encPrivateKey)
	if err != nil {
		return []byte{}, nil, fmt.Errorf("decrypting private key via aes: %w", err)
	}

	privateKeyPemBytes, err := hex.DecodeString(privateKeyInHexString)
	if err != nil {
		return []byte{}, nil, fmt.Errorf("decoding hex to private key pem bytes: %w", err)
	}

	privateKey, err := crypt.PrivateKeyFromPemBytes(privateKeyPemBytes)
	if err != nil {
		return []byte{}, nil, fmt.Errorf("getting private key from pem bytes: %w", err)
	}

	encAesSecretBytes, err := hex.DecodeString(encAesSecret)
	if err != nil {
		return []byte{}, nil, fmt.Errorf("decoding hex to : %w", err)
	}

	aesSecret, err := crypt.DecryptRSA(privateKey, encAesSecretBytes)
	if err != nil {
		return []byte{}, nil, fmt.Errorf("decrypting AES secret via private key: %w", err)
	}

	return aesSecret, privateKey, nil
}
