package services

import (
	"context"
	"github.com/golang/mock/gomock"
	"gokeeper/internal/core"
	"reflect"
	"testing"
	"time"
)

func TestUserService_Create(t *testing.T) {
	userRepo := NewMockUserRepo(gomock.NewController(t))
	keyService := NewMockKeyService(gomock.NewController(t))
	type args struct {
		ctx      context.Context
		login    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    *core.User
		wantErr bool
		prepare func()
	}{
		{
			name: "",
			args: args{
				ctx:      context.Background(),
				login:    "test_user",
				password: "password",
			},
			wantErr: false,
			prepare: func() {
				userRepo.EXPECT().GetByLogin(context.Background(), "test_user").Return(nil)
				keyService.EXPECT().GenerateKeys().Return("secret", "secret")
			},
		},
		{
			name: "",
			args: args{
				ctx:      context.Background(),
				login:    "test_user",
				password: "password",
			},
			wantErr: true,
			prepare: func() {
				userRepo.EXPECT().GetByLogin(context.Background(), "test_user").Return(&core.User{ID: 5})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				keysService: keyService,
				repo:        userRepo,
			}
			tt.prepare()
			got, err := s.Create(tt.args.ctx, tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_FindByLogin(t *testing.T) {
	userRepo := NewMockUserRepo(gomock.NewController(t))
	keyService := NewMockKeyService(gomock.NewController(t))
	type args struct {
		ctx   context.Context
		login string
	}
	tests := []struct {
		name    string
		args    args
		want    *core.User
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:   context.Background(),
				login: "test_user",
			},
			wantErr: false,
			want: &core.User{
				ID:        1,
				Login:     "test_user",
				Password:  "password",
				AesSecret: "dfsdg",
				RsaSecret: "dgdfgds",
				CreatedAt: time.Now(),
				UpdatedAt: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				keysService: keyService,
				repo:        userRepo,
			}
			userRepo.EXPECT().GetByLogin(context.Background(), tt.args.login).Return(&core.User{
				ID:        1,
				Login:     tt.args.login,
				Password:  "password",
				AesSecret: "dfsdg",
				RsaSecret: "dgdfgds",
				CreatedAt: time.Now(),
				UpdatedAt: nil,
			})
			got, err := s.FindByLogin(tt.args.ctx, tt.args.login)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByLogin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_hashPassword(t *testing.T) {
	userRepo := NewMockUserRepo(gomock.NewController(t))
	keyService := NewMockKeyService(gomock.NewController(t))
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "",
			args:    args{password: "secret"},
			want:    "$2a$08$aMSa62GGHKmnT0QPndKzWOo0TV59E/DVjNS1as3l4EISCnGUxFjfq",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				keysService: keyService,
				repo:        userRepo,
			}
			got, err := s.hashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("hashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hashPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}
