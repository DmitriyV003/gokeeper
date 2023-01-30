package server

import (
	"context"
	"gokeeper/internal/core"
	"net/http"
	"strconv"
	"strings"
)

type AuthMiddleware struct {
	authService AuthService
	users       UserRepository
}

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (*core.User, error)
}

type AuthService interface {
	ParseTokenWithClaims(token *core.Token) error
	ValidateToken(token string) (bool, error)
}

func NewAuthMiddleware(authService AuthService, users UserRepository) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
		users:       users,
	}
}

func (m *AuthMiddleware) Pipe() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenHeader := r.Header.Get("Authorization")
			if tokenHeader == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			token := strings.Split(tokenHeader, " ")
			if len(token) < 2 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			isValid, err := m.authService.ValidateToken(token[1])
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if !isValid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			newToken := core.Token{
				Value:  token[1],
				Claims: map[string]interface{}{},
			}
			err = m.authService.ParseTokenWithClaims(&newToken)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			parsedUserID, err := strconv.ParseInt(newToken.Claims["user_id"].(string), 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			user, err := m.users.GetByID(context.Background(), parsedUserID)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			core.SetLoggedInUser(user)

			next.ServeHTTP(w, r.WithContext(context.Background()))
		})
	}
}
