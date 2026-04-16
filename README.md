# Go 練習レポジトリ

Goの学習用プロジェクト

## 使用技術

- **言語**: Go 1.22+
- **Webフレームワーク**: [Huma](https://huma.rocks/) - 宣言的なHTTP APIフレームワーク
- **ルーター**: Gin
- **データベース**: PostgreSQL
- **ORM**: GORM
- **認証**: bcrypt（パスワード暗号化）
- **ID生成**: ULID

## プロジェクト構成

```text
go_practice/
├── adapter/
│   ├── auth/            # JWT など外部認証実装
│   ├── database/
│   │   ├── core.go      # DB接続
│   │   ├── model/       # DBモデル
│   │   └── repository/  # Repository実装
│   └── http/
│       ├── handler/     # HTTPハンドラー
│       ├── middleware/  # HTTPミドルウェア
│       ├── schema/      # HTTPリクエスト/レスポンス
│       └── router.go    # ルーティング設定
├── domain/
│   ├── entity/          # ドメインエンティティ
│   ├── factory/         # エンティティ生成ロジック
│   └── validation/      # ドメインバリデーション
├── usecase/
│   ├── interactor/      # ユースケース実装
│   └── port/
│       ├── input/       # 入力ポート
│       └── output/      # 出力ポート
├── config/              # 設定管理
├── common/              # 共通ユーティリティ
└── utils/               # ヘルパー関数
```

## セットアップ

### 必要な環境

- Go 1.22+
- Docker & Docker Compose
- Make

### 実行方法

1. **環境起動**

   ```bash
   make up
   ```

   Docker Compose で PostgreSQL コンテナを起動

2. **ローカル開発サーバー起動**

   ```bash
   make run
   ```

3. **ログ確認**

   ```bash
   make logs-api
   ```

4. **データベース接続**
   ```bash
   make psql
   ```

## アーキテクチャ

このプロジェクトは **クリーンアーキテクチャ** に基づいています：

- **Domain層**: ビジネスロジック・ドメインルール
- **UseCase層**: アプリケーションロジック・ワークフロー
- **Adapter層**: 外部インテグレーション（DB、HTTP等）

## Huma について

Huma は、OpenAPI 互換の宣言的なHTTP APIフレームワークです。本プロジェクトでは Huma を使用して以下を実現しています：

- 型安全なリクエスト/レスポンス処理
- 自動的なバリデーション
- OpenAPI スキーマ自動生成
- 構造化されたエラーハンドリング

ドキュメント
https://huma.rocks/
