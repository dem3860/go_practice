package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	app, err := NewApp()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}
	defer app.Close()

	// serverの設定
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", app.Config.Port),
		Handler:        app.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("server starting on port %d", app.Config.Port)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to start server: %v", err)
	}
}
