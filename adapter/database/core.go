package database

import (
	"fmt"

	"go_practice/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// PostgreSQL に接続する
func NewPostgreSQLDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := buildDSN(cfg)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect postgres: %w", err)
	}

	return db, nil
}

// PostgreSQL の接続をクローズする
func ClosePostgreSQLDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("failed to close db: %w", err)
	}

	return nil
}

func buildDSN(cfg *config.Config) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tokyo",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)
}
