package common

import (
	"go_practice/adapter/database"
	"go_practice/config"

	"gorm.io/gorm"
)

// DB 接続のエントリーポイント
func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
	return database.NewPostgreSQLDB(cfg)
}

// DB クローズ
func CloseDatabase(db *gorm.DB) error {
	return database.ClosePostgreSQLDB(db)
}
