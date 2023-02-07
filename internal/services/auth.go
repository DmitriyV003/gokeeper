package services

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"gokeeper/internal/core"
	"gokeeper/internal/data/postgres"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var ErrCredentials = errors.New("credentials don't match")

type AuthService struct {
	userRepo *postgres.UserRepository
	secret   string
}

func NewAuthService(userRepo *postgres.UserRepository, secret string) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		secret:   secret,
	}
}

func (s *AuthService) GetSecret() string {
	return s.secret
}

func (s *AuthService) LoginByUser(user *core.User) (string, error) {
	return s.generateJWT(user)
}

func (s *AuthService) Login(ctx context.Context, login, password string) (string, error) {
	user, err := s.userRepo.GetByLogin(ctx, login)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", ErrCredentials
	}

	if err = s.checkPassword(user.Password, password); err != nil {
		return "", ErrCredentials
	}

	return s.generateJWT(user)
}

func (s *AuthService) generateJWT(user *core.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(24 * 60 * time.Minute).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = user.Login

	tokenString, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) checkPassword(hashedPassword, providedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword)); err != nil {
		return err
	}

	return nil
}
