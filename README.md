# OpenAPI Schema Driven Go API

このプロジェクトは、OpenAPI仕様書からGoのAPIサーバーコードを自動生成するサンプルプロジェクトです。

## 必要な依存関係

- Go 1.24+
- oapi-codegen v2

## セットアップ

### 1. oapi-codegenのインストール

```bash
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
```

### 2. 依存関係のインストール

```bash
go mod tidy
```

## 使用方法

### コード生成

OpenAPI仕様書（`schema/spec/openapi.yaml`）からGoコードを生成：

```bash
make generate
```

または直接コマンドで：

```bash
oapi-codegen -generate types,gin,spec -package generated -o generated/api.go schema/spec/openapi.yaml
```

### 開発とテスト

```bash
# アプリケーションの実行
make run

# ビルド
make build

# 生成されたファイルのクリーンアップ
make clean

# 開発用（生成→ビルド→実行）
make dev
```

### APIの確認

サーバーを起動後、以下のエンドポイントにアクセスできます：

- `GET /ping` - ヘルスチェック
- `GET /users` - ユーザー一覧の取得

```bash
# ヘルスチェック
curl http://localhost:8080/ping

# ユーザー一覧取得
curl http://localhost:8080/users
```

## プロジェクト構造

```tree
.
├── schema/
│   └── spec/
│       └── openapi.yaml          # OpenAPI仕様書
├── generated/
│   └── api.go                   # 自動生成されたGoコード
├── handlers/
│   └── users.go                 # APIハンドラーの実装
├── main.go                      # メインファイル
├── Makefile                     # ビルドタスク
├── docker-compose.yaml          # Docker設定
└── Dockerfile                   # Dockerイメージ設定
```

## OpenAPI仕様の更新

1. `schema/spec/openapi.yaml` を編集
2. `make generate` でコードを再生成
3. 必要に応じて `handlers/` でビジネスロジックを実装

## Docker使用

```bash
# Docker Composeでアプリケーションを起動
docker compose up --build

# バックグラウンドで起動
docker compose up -d --build
```

## 注意事項

- 現在のoapi-codegenはOpenAPI 3.1を完全にサポートしていないため、3.0.3を使用しています
- `generated/api.go` は自動生成されるファイルなので、直接編集しないでください
- ビジネスロジックは `handlers/` ディレクトリで実装してください
