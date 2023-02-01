package core

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"gokeeper/internal/requests"
	"gokeeper/pkg/crypt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

var loggedInUser *User

func SetLoggedInUser(user *User) {
	loggedInUser = user
}

func GetLoggedInUser() *User {
	return loggedInUser
}

type Token struct {
	Value  string
	Claims map[string]interface{}
}

type UserRepository interface {
	GetByLogin(ctx context.Context, login string) (*User, error)
	Create(ctx context.Context, user *User) error
}

type AuthService struct {
	users  UserRepository
	secret string
}

func NewAuthService(secret string, users UserRepository) *AuthService {
	return &AuthService{
		secret: secret,
		users:  users,
	}
}

func (s *AuthService) CreateUser(ctx context.Context, request *requests.RegistrationRequest) (*Token, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		return nil, fmt.Errorf("error to generate hash from password: %w", err)
	}

	aesSecret := make([]byte, 64)
	_, err = rand.Read(aesSecret)
	if err != nil {
		return nil, fmt.Errorf("error to generate random key: %w", err)
	}

	bobPrivateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, fmt.Errorf("error to generate RSA keys: %w", err)
	}

	encAESSecret, err := crypt.EncryptRSA(&bobPrivateKey.PublicKey, aesSecret)
	if err != nil {
		return nil, fmt.Errorf("error to encrypt AES key: %w", err)
	}

	user := User{
		Login:     request.Login,
		Password:  string(bytes),
		CreatedAt: time.Now(),
		AesSecret: string(encAESSecret),
		RsaSecret: string(crypt.ExportPrivateKeyAsPemBytes(bobPrivateKey)),
	}
	err = s.users.Create(ctx, &user)
	if err != nil {
		return nil, fmt.Errorf("unable to create user in db: %w", err)
	}

	dbUser, err := s.users.GetByLogin(ctx, request.Login)
	if err != nil {
		return nil, fmt.Errorf("error to get user by login: %w", err)
	}

	token, err := s.LoginByUser(dbUser)
	if err != nil {
		return nil, fmt.Errorf("login user programmly: %w", err)
	}

	return token, nil
}

func (s *AuthService) LoginByUser(user *User) (*Token, error) {
	token, err := s.generateJwt(user)
	if err != nil {
		return nil, fmt.Errorf("error to generate jwt to login user: %w", err)
	}

	return token, nil
}

func (s *AuthService) Login(ctx context.Context, login string, password string) (*Token, error) {
	user, err := s.users.GetByLogin(ctx, login)
	if err != nil {
		return nil, fmt.Errorf("error to get user by login: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("error to compare passwords: %w", err)
	}

	token, err := s.generateJwt(user)
	if err != nil {
		return nil, fmt.Errorf("error to generate jwt token: %w", err)
	}

	return token, nil
}

func (s *AuthService) ValidateToken(token string) (bool, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("error to validate token")
		}

		return []byte(s.secret), nil
	})

	return parsedToken.Valid, err
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

func (s *AuthService) generateJwt(user *User) (*Token, error) {
	token := jwt.New(jwt.SigningMethodHS512)

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(24 * 60 * 5 * time.Minute).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = strconv.FormatInt(user.ID, 10)
	claims["user_id"] = strconv.FormatInt(user.ID, 10)
	token.Claims = claims

	genToken, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return nil, fmt.Errorf("error sign token: %w", err)
	}

	return &Token{Value: genToken}, nil
}
