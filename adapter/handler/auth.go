package handler

import (
	"context"
	"errors"

	"go_practice/adapter/schema"
	"go_practice/usecase/input_port"
	"go_practice/usecase/interactor"

	"github.com/danielgtaylor/huma/v2"
)

type AuthHandler struct {
	AuthUC input_port.IAuthUseCase
}

func NewAuthHandler(authUC input_port.IAuthUseCase) *AuthHandler {
	return &AuthHandler{
		AuthUC: authUC,
	}
}

func (h *AuthHandler) Login(ctx context.Context, input *schema.LoginReq) (*schema.LoginRes, error) {
	user, token, err := h.AuthUC.Login(input.Body.Email, input.Body.Password)
	if err != nil {
		switch {
		case errors.Is(err, interactor.ErrKind.BadRequest):
			return nil, huma.Error400BadRequest("invalid email or password", err)

		case errors.Is(err, interactor.ErrKind.NotFound):
			return nil, huma.Error404NotFound("user not found", err)

		case errors.Is(err, interactor.ErrKind.InternalServerError):
			return nil, huma.Error500InternalServerError("internal server error", err)

		default:
			return nil, huma.Error500InternalServerError("internal server error", err)
		}
	}

	return schema.ToLoginResponse(user, token), nil
}

func (h *AuthHandler) Signup(ctx context.Context, input *schema.SignupReq) (*schema.SignupRes, error) {
	uc := input_port.SignupInput{
		Email:    input.Body.Email,
		Password: input.Body.Password,
		Name:     input.Body.Name,
	}

	res, err := h.AuthUC.Signup(uc)
	if err != nil {
		switch {
		case errors.Is(err, interactor.ErrKind.Validation):
			return nil, huma.Error400BadRequest("validation error", err)

		case errors.Is(err, interactor.ErrKind.Conflict):
			return nil, huma.Error409Conflict("user already exists", err)

		case errors.Is(err, interactor.ErrKind.DB):
			return nil, huma.Error500InternalServerError("database error", err)

		case errors.Is(err, interactor.ErrKind.InternalServerError):
			return nil, huma.Error500InternalServerError("internal server error", err)

		default:
			return nil, huma.Error500InternalServerError("internal server error", err)
		}
	}

	return schema.ToSignupResponse(res), nil
}
