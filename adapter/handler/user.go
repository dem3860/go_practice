package handler

import (
	"context"
	"errors"

	"go_practice/adapter/schema"
	"go_practice/usecase/input_port"
	"go_practice/usecase/interactor"

	"github.com/danielgtaylor/huma/v2"
)

type UserHandler struct {
	UserUC input_port.IUserUseCase
}

func NewUserHandler(userUC input_port.IUserUseCase) *UserHandler {
	return &UserHandler{
		UserUC: userUC,
	}
}

func (h *UserHandler) Login(ctx context.Context, input *schema.LoginReq) (*schema.LoginRes, error) {
	user, token, err := h.UserUC.Login(input.Body.Email, input.Body.Password)
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

func (h *UserHandler) Create(ctx context.Context, input *schema.CreateUserReq) (*schema.CreateUserRes, error) {
	uc := input_port.UserCreate{
		Email:    input.Body.Email,
		Password: input.Body.Password,
		Name:     input.Body.Name,
		Role:     input.Body.Role,
	}

	res, err := h.UserUC.Create(uc)
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

	return schema.ToCreateUserResponse(res), nil
}
