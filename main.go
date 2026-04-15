package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"go_practice/adapter/handler"
	"go_practice/common"
	"go_practice/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// 環境変数を読み込む
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// データベースに接続する
	db, err := common.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// アプリケーション終了時にデータベース接続をクローズする
	defer func() {
		if err := common.CloseDatabase(db); err != nil {
			log.Printf("failed to close database: %v", err)
		}
	}()

	deps := handler.NewDeps(db)

	r := gin.Default()
	handler.SetupRouter(r, deps)

	// HTTP サーバーの設定
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("server starting on port %d", cfg.Port)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to start server: %v", err)
	}
}
