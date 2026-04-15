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
	UserUseCase input_port.IUserUseCase
}

func NewDeps(userUseCase input_port.IUserUseCase) *Deps {
	return &Deps{
		UserUseCase: userUseCase,
	}
}

func SetupRouter(router *gin.Engine, deps *Deps) {
	api := humagin.New(router, huma.DefaultConfig("My API", "1.0.0"))

	// ハンドラーの初期化
	userHandler := NewUserHandler(deps.UserUseCase)

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
		Path:        "/login",
		Summary:     "Login",
		Description: "Login with email and password.",
		Tags:        []string{"Users"},
	}, userHandler.Login)

	huma.Register(api, huma.Operation{
		OperationID: "create-user",
		Method:      http.MethodPost,
		Path:        "/users",
		Summary:     "Create a new user",
		Description: "Create a new user account with email and password.",
		Tags:        []string{"Users"},
	}, userHandler.Create)

}
