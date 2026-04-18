package handler

import (
	"context"
	"errors"
	"go_practice/adapter/http/authctx"
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

func (h *UserHandler) UpdateByMe(ctx context.Context, req *schema.UpdateByMeReq) (*schema.UserRes, error) {
	// ctxからuserを取得
	user, err := authctx.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, huma.Error401Unauthorized("authentication required")
	}

	uu := inputport.UpdateByMeInput{
		ID:   user.ID,
		Name: req.Body.Name,
	}

	res, err := h.UserUC.UpdateByMe(uu)
	if err != nil {
		switch {
		case errors.Is(err, interactor.ErrKind.NotFound):
			return nil, huma.Error404NotFound("user not found", err)
		case errors.Is(err, interactor.ErrKind.Validation):
			return nil, huma.Error400BadRequest("validation error", err)
		case errors.Is(err, interactor.ErrKind.DB):
			return nil, huma.Error500InternalServerError("database error", err)
		case errors.Is(err, interactor.ErrKind.InternalServerError):
			return nil, huma.Error500InternalServerError("internal server error", err)
		default:
			return nil, huma.Error500InternalServerError("internal server error", err)
		}
	}

	return schema.ToUserResponse(res), nil
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

func (h *UserHandler) Delete(ctx context.Context, req *schema.DeleteUserReq) (*struct{}, error) {
	if err := h.UserUC.Delete(req.UserID); err != nil {
		switch {
		case errors.Is(err, interactor.ErrKind.NotFound):
			return nil, huma.Error404NotFound("user not found", err)
		case errors.Is(err, interactor.ErrKind.DB):
			return nil, huma.Error500InternalServerError("database error", err)
		case errors.Is(err, interactor.ErrKind.InternalServerError):
			return nil, huma.Error500InternalServerError("internal server error", err)
		default:
			return nil, huma.Error500InternalServerError("internal server error", err)
		}
	}

	return nil, nil
}
