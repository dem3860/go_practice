package httpadapter

import (
	"net/http"

	"go_practice/adapter/http/handler"
	"go_practice/adapter/http/middleware"
	"go_practice/usecase/port/input"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

// 依存関係をまとめる構造体
type Deps struct {
	AuthUseCase input.IAuthUseCase
	UserUseCase input.IUserUseCase
}

func NewDeps(authUseCase input.IAuthUseCase, userUseCase input.IUserUseCase) *Deps {
	return &Deps{
		AuthUseCase: authUseCase,
		UserUseCase: userUseCase,
	}
}

func SetupRouter(router *http.ServeMux, deps *Deps) {
	api := humago.New(router, huma.DefaultConfig("My API", "1.0.0"))
	// ハンドラーの初期化
	authHandler := handler.NewAuthHandler(deps.AuthUseCase)
	userHandler := handler.NewUserHandler(deps.UserUseCase)
	authMiddleware := middleware.NewAuthMiddleware(api, deps.AuthUseCase, deps.UserUseCase)

	// OpenAPI ドキュメントにセキュリティスキームを追加
	api.OpenAPI().Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"Bearer": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}

	// ログイン
	huma.Register(api, huma.Operation{
		OperationID: "login",
		Method:      http.MethodPost,
		Path:        "/auth/login",
		Summary:     "Login",
		Description: "Login with email and password.",
		Tags:        []string{"Auth"},
	}, authHandler.Login)

	// サインアップ
	huma.Register(api, huma.Operation{
		OperationID: "signup",
		Method:      http.MethodPost,
		Path:        "/auth/signup",
		Summary:     "Sign up",
		Description: "Create a new user account with email and password.",
		Tags:        []string{"Auth"},
	}, authHandler.Signup)

	// 自分のユーザ情報更新
	huma.Register(api, huma.Operation{
		OperationID: "update-my-user",
		Method:      http.MethodPatch,
		Path:        "/me",
		Summary:     "Update my profile",
		Description: "Update the authenticated user's profile.",
		Tags:        []string{"Users"},
		Security:    []map[string][]string{{"Bearer": {}}},
		Middlewares: huma.Middlewares{
			authMiddleware.Authenticate,
		},
	}, userHandler.UpdateByMe)

	huma.Register(api, huma.Operation{
		OperationID: "delete-user",
		Method:      http.MethodDelete,
		Path:        "/admin/users/{userID}",
		Summary:     "Delete user",
		Description: "Delete a user account for administrators.",
		Tags:        []string{"Admin"},
		Security:    []map[string][]string{{"Bearer": {}}},
		Middlewares: huma.Middlewares{
			authMiddleware.Authenticate,
			authMiddleware.RequireAdmin,
		},
	}, userHandler.Delete)

	// ユーザ一覧
	huma.Register(api, huma.Operation{
		OperationID: "list-users",
		Method:      http.MethodGet,
		Path:        "/admin/users",
		Summary:     "List users",
		Description: "List users with search and pagination for administrators.",
		Tags:        []string{"Admin"},
		Security:    []map[string][]string{{"Bearer": {}}},
		Middlewares: huma.Middlewares{
			authMiddleware.Authenticate,
			authMiddleware.RequireAdmin,
		},
	}, userHandler.List)
}
