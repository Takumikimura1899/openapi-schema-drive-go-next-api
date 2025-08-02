.PHONY: generate run clean build

# OpenAPI仕様からGoコードを生成
generate:
	@echo "Generating Go code from OpenAPI spec..."
	oapi-codegen -generate types,gin,spec -package generated -o generated/api.go schema/spec/openapi.yaml
	@echo "Code generation complete!"

# アプリケーションを実行
run:
	go run main.go

# 生成されたファイルをクリーンアップ
clean:
	rm -f generated/api.go

# アプリケーションをビルド
build:
	go build -o tmp/app .

# 依存関係を整理
tidy:
	go mod tidy

# 開発用：生成→ビルド→実行
dev: generate build run
