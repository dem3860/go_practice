package handler

import (
	"net/http"

	"go_practice/usecase/input_port"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
)

// 依存関係をまとめる構造体
type Deps struct {
	AuthUseCase input_port.IAuthUseCase
}

func NewDeps(authUseCase input_port.IAuthUseCase) *Deps {
	return &Deps{
		AuthUseCase: authUseCase,
	}
}

func SetupRouter(router *gin.Engine, deps *Deps) {
	api := humagin.New(router, huma.DefaultConfig("My API", "1.0.0"))

	// ハンドラーの初期化
	authHandler := NewAuthHandler(deps.AuthUseCase)

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

}
