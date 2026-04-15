package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	// サーバー設定
	Port int `env:"PORT" envDefault:"8080"`

	// JWT設定
	JWTSecret string `env:"JWT_SECRET" envDefault:"mysecretkey"`
	JWTExpire int    `env:"JWT_EXPIRE" envDefault:"3600"`

	// データベース設定
	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBPort     int    `env:"DB_PORT" envDefault:"5432"`
	DBUser     string `env:"DB_USER" envDefault:"postgres"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"postgres"`
	DBName     string `env:"DB_NAME" envDefault:"task_app"`
	DBSSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
