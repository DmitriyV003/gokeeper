package services

import (
	"context"
	"github.com/golang/mock/gomock"
	"gokeeper/internal/core"
	"testing"
	"time"
)

func TestAuthService_GetSecret(t *testing.T) {
	type fields struct {
		secret string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "",
			fields: fields{
				secret: "sdgregergerger",
			},
			want: "sdgregergerger",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AuthService{
				userRepo: NewMockUserRepo(gomock.NewController(t)),
				secret:   tt.fields.secret,
			}
			if got := s.GetSecret(); got != tt.want {
				t.Errorf("GetSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_Login(t *testing.T) {
	type fields struct {
		secret string
	}
	type args struct {
		ctx      context.Context
		login    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "",
			fields: fields{secret: "sfsdfgdfgdffghfgreger"},
			args: args{
				ctx:      context.Background(),
				login:    "test_user",
				password: "password",
			},
			wantErr: false,
		},
		{
			name:   "",
			fields: fields{secret: "sfsdfgdfgdffghfgreger"},
			args: args{
				ctx:      context.Background(),
				login:    "test_user",
				password: "password",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := NewMockUserRepo(gomock.NewController(t))
			s := &AuthService{
				userRepo: userRepo,
				secret:   tt.fields.secret,
			}
			password := tt.args.password
			if tt.wantErr {
				password = "dfhfghfghdf"
			}

			userRepo.EXPECT().GetByLogin(context.Background(), "test_user").Return(&core.User{
				ID:        1,
				Login:     tt.args.login,
				Password:  password,
				AesSecret: "dfsdg",
				RsaSecret: "dgdfgds",
				CreatedAt: time.Now(),
				UpdatedAt: nil,
			})

			_, err := s.Login(tt.args.ctx, tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAuthService_checkPassword(t *testing.T) {
	type fields struct {
		secret string
	}
	type args struct {
		hashedPassword   string
		providedPassword string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				secret: "asefrgrthtyejetyj",
			},
			args: args{
				hashedPassword:   "$2a$08$aMSa62GGHKmnT0QPndKzWOo0TV59E/DVjNS1as3l4EISCnGUxFjfq",
				providedPassword: "secret",
			},
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				secret: "asefrgrthtyejetyj",
			},
			args: args{
				hashedPassword:   "$2a$08$aMSa62GGHKmnT0QPndKzWOo0TV59E/DVjNS1as3l4EISCnGUxFjfq",
				providedPassword: "secret1",
			},
			wantErr: true,
		},
		{
			name: "",
			fields: fields{
				secret: "asefrgrthtyejetyj",
			},
			args: args{
				hashedPassword:   "$2m$08$aMSa62GGHKmnT0QPndKzWOo0TV59E/DVjNS1as3l4EISCnGUxFjfq",
				providedPassword: "secret",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := NewMockUserRepo(gomock.NewController(t))
			s := &AuthService{
				userRepo: userRepo,
				secret:   tt.fields.secret,
			}
			if err := s.checkPassword(tt.args.hashedPassword, tt.args.providedPassword); (err != nil) != tt.wantErr {
				t.Errorf("checkPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
