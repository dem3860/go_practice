package httpadapter

import (
	"net/http"

	"go_practice/adapter/http/handler"
	"go_practice/adapter/http/middleware"
	"go_practice/usecase/port/input"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
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

func SetupRouter(router *gin.Engine, deps *Deps) {
	api := humagin.New(router, huma.DefaultConfig("My API", "1.0.0"))
	// ハンドラーの初期化
	authHandler := handler.NewAuthHandler(deps.AuthUseCase)
	authMiddleware := middleware.NewAuthMiddleware(api, deps.AuthUseCase, deps.UserUseCase)

	// OpenAPI ドキュメントにセキュリティスキームを追加
	api.OpenAPI().Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"Bearer": {
			Type:   "http",
			Scheme: "bearer",
		},
	}

	huma.Register(api, huma.Operation{
		OperationID: "login",
		Method:      http.MethodPost,
		Path:        "/auth/login",
		Summary:     "Login",
		Description: "Login with email and password.",
		Tags:        []string{"Auth"},
	}, authHandler.Login)

	huma.Register(api, huma.Operation{
		OperationID: "signup",
		Method:      http.MethodPost,
		Path:        "/auth/signup",
		Summary:     "Sign up",
		Description: "Create a new user account with email and password.",
		Tags:        []string{"Auth"},
	}, authHandler.Signup)

	// Apply these middlewares to protected operations when user APIs are added.
	_ = huma.Middlewares{authMiddleware.Authenticate, authMiddleware.RequireAdmin}
}
