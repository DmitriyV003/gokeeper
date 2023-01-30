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

type RegisterHandler struct {
	userService UserService
}

type UserService interface {
	CreateUser(ctx context.Context, request *requests.RegistrationRequest) (*core.Token, error)
}

func NewRegisterHandler(userService UserService) *RegisterHandler {
	return &RegisterHandler{
		userService: userService,
	}
}

func (h *RegisterHandler) Handle() http.HandlerFunc {
	return func(res http.ResponseWriter, request *http.Request) {
		var regRequest requests.RegistrationRequest

		validate := validator.New()

		if err := json.NewDecoder(request.Body).Decode(&regRequest); err != nil {
			applicationerrors.SwitchError(&res, err, nil, "error decode body")
			return
		}

		if err := validate.Struct(&regRequest); err != nil {
			applicationerrors.WriteHTTPError(&res, http.StatusBadRequest, err)
			return
		}

		token, err := h.userService.CreateUser(request.Context(), &regRequest)
		if err != nil {
			applicationerrors.SwitchError(&res, err, nil, "error to create user")
			return
		}

		res.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token.Value))
		res.WriteHeader(http.StatusOK)
	}
}
