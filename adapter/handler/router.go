package handler

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Deps struct {
	DB *gorm.DB

	UserUseCase UserUseCase
	TaskUseCase TaskUseCase
}

func NewDeps(db *gorm.DB) *Deps {
	return &Deps{
		DB: db,
	}
}

type UserUseCase interface{}
type TaskUseCase interface{}

func SetupRouter(router *gin.Engine, deps *Deps) {
	router.Use(func(c *gin.Context) {
		c.Set("deps", deps)
		c.Next()
	})

	api := humagin.New(router, huma.DefaultConfig("My API", "1.0.0"))

	// OpenAPI ドキュメントにセキュリティスキームを追加
	api.OpenAPI().Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"Bearer": {
			Type:   "http",
			Scheme: "bearer",
		},
	}

	huma.Get(api, "/health", func(ctx context.Context, input *struct{}) (*struct {
		Body struct{} `json:"body"`
	}, error) {
		return &struct {
			Body struct{} `json:"body"`
		}{}, nil
	})

	huma.Get(api,"/", func(ctx context.Context, input *struct{}) (*struct {
		Body string `json:"body"`
	}, error) {
		return &struct {
			Body string `json:"body"`
		}{Body: "Hello Gin + Huma!"}, nil
	})

	huma.Get(api, "/doc", func(ctx context.Context, input *struct{}) (*struct {
		Body interface{} `json:"body"`
	}, error) {
		return &struct {
			Body interface{} `json:"body"`
		}{Body: api.OpenAPI()}, nil
	})

	registerUserRoutes(api, deps)
	registerTaskRoutes(api, deps)
}

func registerUserRoutes(api huma.API, deps *Deps) {}

func registerTaskRoutes(api huma.API, deps *Deps) {}
