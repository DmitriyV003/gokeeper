package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gokeeper/internal/applicationerrors"
	"gokeeper/internal/core"
	"gokeeper/internal/requests"
	"net/http"
)

type LoginHandler struct {
	authService AuthService
}

type AuthService interface {
	LoginByUser(user *core.User) (*core.Token, error)
	Login(ctx context.Context, login string, password string) (*core.Token, error)
	ValidateToken(token string) (bool, error)
	ParseTokenWithClaims(token *core.Token) error
}

func NewLoginHandler(authService AuthService) *LoginHandler {
	return &LoginHandler{
		authService: authService,
	}
}

func (h LoginHandler) Handle() http.HandlerFunc {
	return func(res http.ResponseWriter, request *http.Request) {
		var loginRequest requests.LoginRequest

		validate := validator.New()

		if err := json.NewDecoder(request.Body).Decode(&loginRequest); err != nil {
			applicationerrors.SwitchError(&res, err, nil)
			return
		}

		if err := validate.Struct(&loginRequest); err != nil {
			applicationerrors.WriteHTTPError(&res, http.StatusBadRequest, err)
			return
		}

		token, err := h.authService.Login(request.Context(), loginRequest.Login, loginRequest.Password)
		if err != nil {
			applicationerrors.SwitchError(&res, err, nil)
			return
		}

		res.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token.Value))
		res.WriteHeader(http.StatusOK)
	}
}
