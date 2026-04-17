package middleware

import (
	"errors"
	"strings"

	"go_practice/adapter/http/authctx"
	"go_practice/domain/entity"
	"go_practice/usecase/interactor"
	"go_practice/usecase/port/input"

	"fmt"
	"github.com/danielgtaylor/huma/v2"
)

type AuthMiddleware struct {
	api    huma.API
	authUC input.IAuthUseCase
	userUC input.IUserUseCase
}

func NewAuthMiddleware(api huma.API, authUC input.IAuthUseCase, userUC input.IUserUseCase) *AuthMiddleware {
	return &AuthMiddleware{
		api:    api,
		authUC: authUC,
		userUC: userUC,
	}
}

func (m *AuthMiddleware) Authenticate(ctx huma.Context, next func(huma.Context)) {
	// Authorizationヘッダーからtokenを抽出する
	token, err := bearerToken(ctx.Header("Authorization"))
	if err != nil {
		_ = huma.WriteErr(m.api, ctx, 401, "missing or invalid bearer token", err)
		return
	}

	// tokenが検証され、userIDを返ってくる
	userID, err := m.authUC.Authenticate(token)
	if err != nil {
		_ = huma.WriteErr(m.api, ctx, 401, "invalid or expired token", err)
		return
	}

	user, err := m.userUC.FindByID(userID)
	if err != nil {
		status := 500
		message := "failed to load authenticated user"
		if errors.Is(err, interactor.ErrKind.NotFound) {
			status = 401
			message = "authenticated user not found"
		}
		_ = huma.WriteErr(m.api, ctx, status, message, err)
		return
	}

	// contextにuserを入れ、新しいctxを渡す
	next(authctx.SetAuthenticatedUser(ctx, user))
}

func (m *AuthMiddleware) RequireAdmin(ctx huma.Context, next func(huma.Context)) {
	user, err := authctx.GetAuthenticatedUser(ctx.Context())
	if err != nil {
		_ = huma.WriteErr(m.api, ctx, 401, "authentication required", err)
		return
	}
	// userのロールがadminかどうかを検証する
	if user.Role != entity.RoleAdmin {
		_ = huma.WriteErr(m.api, ctx, 403, "admin role required")
		return
	}

	next(ctx)
}

func bearerToken(header string) (string, error) {
	if header == "" {
		return "", fmt.Errorf("authorization header is missing")
	}
	if !strings.HasPrefix(header, "Bearer ") {
		return "", fmt.Errorf("authorization header must use bearer scheme")
	}

	token := strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))
	if token == "" {
		return "", fmt.Errorf("bearer token is missing")
	}

	return token, nil
}
