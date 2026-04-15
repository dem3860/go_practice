package main

import (
	"go_practice/adapter/auth"
	"go_practice/adapter/database/repository"
	"go_practice/adapter/handler"
	"go_practice/common"
	"go_practice/config"
	"go_practice/usecase/interactor"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Config *config.Config
	DB     *gorm.DB
	Router *gin.Engine
}

func NewApp() (*App, error) {
	// 環境変数の読み込み
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}

	// データベースの接続
	db, err := common.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	// レポジトリとユースケースのDI
	userRepo := repository.NewUserRepository(db)
	tokenProvider := auth.NewJWTProvider(cfg.JWTSecret, time.Duration(cfg.JWTExpire)*time.Second)
	userUC := interactor.NewUserUseCase(userRepo, tokenProvider)

	deps := handler.NewDeps(userUC)

	router := gin.Default()
	handler.SetupRouter(router, deps)

	return &App{
		Config: cfg,
		DB:     db,
		Router: router,
	}, nil
}

func (a *App) Close() {
	if a.DB == nil {
		return
	}
	if err := common.CloseDatabase(a.DB); err != nil {
		log.Printf("failed to close DB: %v", err)
	}
}
