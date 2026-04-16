# ==============================
# 基本設定
# ==============================
APP_NAME := app
PORT := 8080

# Docker Compose
DC := docker compose

# ==============================
# 開発用コマンド
# ==============================

## 起動（バックグラウンド）
up:
	$(DC) up -d --build

## 停止
down:
	$(DC) down

## ログ確認
logs:
	$(DC) logs -f

## APIコンテナのログだけ
logs-api:
	$(DC) logs -f api

## DBコンテナのログだけ
logs-db:
	$(DC) logs -f pg

## 再起動
restart:
	$(DC) down && $(DC) up -d --build

## コンテナに入る
sh:
	$(DC) exec api sh

## DBに入る
psql:
	$(DC) exec pg psql -U postgres -d task_app

# ==============================
# Go関連（ローカル or コンテナ）
# ==============================

## フォーマット
fmt:
	go fmt ./...

## tidy
tidy:
	go mod tidy

## ビルド（ローカル）
build:
	go build -o bin/main .

## 実行（ローカル）
run:
	go run main.go

## OpenAPI JSON を生成
openapi:
	@echo "Fetching OpenAPI JSON from http://localhost:$(PORT)/openapi.json..."
	@curl -s http://localhost:$(PORT)/openapi.json | jq . > openapi.json
	@echo "✓ OpenAPI JSON generated: openapi.json"

# ==============================
# DB関連
# ==============================

## DBリセット（危険）
reset-db:
	$(DC) down -v
	$(DC) up -d

## DBデータを完全削除（bind mount含む）
wipe-db:
	$(DC) down
	rm -rf ./mnt/pgdata
	mkdir -p ./mnt/pgdata
	$(DC) up -d

# ==============================
# 開発補助
# ==============================

## コンテナ再ビルド（キャッシュなし）
rebuild:
	$(DC) build --no-cache

## 全削除（イメージ含む）
clean:
	$(DC) down -v --rmi all --remove-orphans

# ==============================
# ヘルプ
# ==============================

help:
	@echo "make up          - コンテナ起動"
	@echo "make down        - コンテナ停止"
	@echo "make logs        - ログ確認"
	@echo "make sh          - apiコンテナに入る"
	@echo "make psql        - DBに入る"
	@echo "make openapi     - OpenAPI JSON を生成"
	@echo "make reset-db    - DBリセット（注意）"
	@echo "make wipe-db     - DBデータを完全削除（注意）"
	@echo "make rebuild     - キャッシュなしビルド"
	@echo "make clean       - 完全削除"
