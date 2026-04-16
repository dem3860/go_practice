package handler

import (
	"context"
	"errors"
	"go_practice/adapter/http/schema"
	"go_practice/usecase/interactor"
	inputport "go_practice/usecase/port/input"

	"github.com/danielgtaylor/huma/v2"
)

type UserHandler struct {
	UserUC inputport.IUserUseCase
}

func NewUserHandler(userUC inputport.IUserUseCase) *UserHandler {
	return &UserHandler{
		UserUC: userUC,
	}
}

func (u *UserHandler) List(ctx context.Context, req *schema.ListUsersReq) (*schema.ListUsersRes, error) {
	list, total, nextPage, err := u.UserUC.List(req.ToQuery())
	if err != nil {
		switch {
		case errors.Is(err, interactor.ErrKind.DB):
			return nil, huma.Error500InternalServerError("database error", err)
		case errors.Is(err, interactor.ErrKind.InternalServerError):
			return nil, huma.Error500InternalServerError("internal server error", err)
		default:
			return nil, huma.Error500InternalServerError("internal server error", err)
		}
	}

	return schema.ToListUsersResponse(list, total, nextPage), nil
}
