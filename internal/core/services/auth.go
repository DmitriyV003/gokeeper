package services

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"gokeeper/internal/data"
)

type AuthService struct {
	secret       string
	jwtSecret    string
	settingsRepo Settings
}

type Settings interface {
	Get(ctx context.Context, key string) (string, bool, error)
	Set(ctx context.Context, key, value string) (bool, error)
	Delete(ctx context.Context, key string) (bool, error)
	Truncate(ctx context.Context) error
}

func NewAuthService(secret string, settingsRepo Settings) *AuthService {
	return &AuthService{
		secret:       secret,
		settingsRepo: settingsRepo,
	}
}

func (s *AuthService) CheckAuthorized(ctx context.Context) (bool, error) {
	tokenString, existed, err := s.settingsRepo.Get(ctx, "jwt")
	if err != nil {
		return false, fmt.Errorf("getting JWT from setting: %w", err)
	}
	if !existed {
		return false, nil
	}

	token, err := s.ParseJWT(tokenString)
	if err != nil {
		return false, fmt.Errorf("parse jwt: %w", err)
	}

	if !token.Valid {
		return false, fmt.Errorf("invalid jwt stored in settings")
	}

	return true, data.ErrLoggedInAlready
}

func (s *AuthService) ParseJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secret), nil
	})
}

type Token struct {
	Value  string
	Claims map[string]interface{}
}

func (s *AuthService) ParseTokenWithClaims(token *Token) error {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil {
		return fmt.Errorf("error to parse jwt token: %w", err)
	}

	token.Claims = claims
	return nil
}
