package authctx

import (
	"context"
	"fmt"

	"go_practice/domain/entity"

	"github.com/danielgtaylor/huma/v2"
)

type userContextKey string

const authenticatedUserKey userContextKey = "auth_user"

func SetAuthenticatedUser(ctx huma.Context, user entity.User) huma.Context {
	return huma.WithValue(ctx, authenticatedUserKey, user)
}

func GetAuthenticatedUser(ctx context.Context) (entity.User, error) {
	value := ctx.Value(authenticatedUserKey)
	user, ok := value.(entity.User)
	if !ok {
		return entity.User{}, fmt.Errorf("authenticated user not found in context")
	}

	return user, nil
}
