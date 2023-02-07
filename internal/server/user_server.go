package server

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"gokeeper/internal/proto"
	services2 "gokeeper/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errInternalServerError = status.Error(codes.Unknown, "Internal server error")
var ErrLoginTaken = errors.New("login already taken in")
var ErrCredentials = errors.New("invalid credentials")

type UserServer struct {
	authService *services2.AuthService
	userService *services2.UserService
	proto.UnimplementedUserServer
}

func NewUserServer(authService *services2.AuthService, userService *services2.UserService) *UserServer {
	return &UserServer{
		authService: authService,
		userService: userService,
	}
}

func (u *UserServer) Register(ctx context.Context, in *proto.RegisterRequest) (*proto.AuthResponse, error) {
	user, err := u.userService.Create(ctx, in.Login, in.Password)
	if err != nil {
		if errors.Is(err, ErrLoginTaken) {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		log.Error().Err(err).Msg("Creating user on register")
		return nil, errInternalServerError
	}

	token, err := u.authService.LoginByUser(user)
	if err != nil {
		log.Error().Err(err).Msg("Authorizing user")
		return nil, errInternalServerError
	}

	log.Debug().Str("login", in.Login).Msg("User created successfully")

	return &proto.AuthResponse{
		Token:      token,
		AesSecret:  user.AesSecret,
		PrivateKey: user.RsaSecret,
	}, nil
}

func (u *UserServer) Login(ctx context.Context, in *proto.LoginRequest) (*proto.AuthResponse, error) {
	token, err := u.authService.Login(ctx, in.Login, in.Password)
	if err != nil {
		if errors.Is(err, ErrCredentials) {
			return nil, status.Error(codes.Unauthenticated, "The credentials don't match any of our records")
		}

		log.Error().Err(err).Msg("Authorizing by credentials")
		return nil, errInternalServerError
	}

	user, err := u.userService.FindByLogin(ctx, in.Login)
	if err != nil {
		return nil, errInternalServerError
	}

	return &proto.AuthResponse{
		Token:      token,
		AesSecret:  user.AesSecret,
		PrivateKey: user.RsaSecret,
	}, nil
}
