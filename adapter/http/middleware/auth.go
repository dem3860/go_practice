package middleware

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"go_practice/domain/entity"
	"go_practice/usecase/interactor"
	"go_practice/usecase/port/input"

	"github.com/danielgtaylor/huma/v2"
)

type authUserContextKey string

const authUserKey authUserContextKey = "auth_user"

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
	token, err := bearerToken(ctx.Header("Authorization"))
	if err != nil {
		_ = huma.WriteErr(m.api, ctx, 401, "missing or invalid bearer token", err)
		return
	}

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

	next(SetAuthenticatedUser(ctx, user))
}

func (m *AuthMiddleware) RequireAdmin(ctx huma.Context, next func(huma.Context)) {
	user, err := GetAuthenticatedUser(ctx.Context())
	if err != nil {
		_ = huma.WriteErr(m.api, ctx, 401, "authentication required", err)
		return
	}
	if user.Role != entity.RoleAdmin {
		_ = huma.WriteErr(m.api, ctx, 403, "admin role required")
		return
	}

	next(ctx)
}

func SetAuthenticatedUser(ctx huma.Context, user entity.User) huma.Context {
	return huma.WithValue(ctx, authUserKey, user)
}

func GetAuthenticatedUser(ctx context.Context) (entity.User, error) {
	value := ctx.Value(authUserKey)
	user, ok := value.(entity.User)
	if !ok {
		return entity.User{}, fmt.Errorf("authenticated user not found in context")
	}

	return user, nil
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
