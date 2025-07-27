# ベースイメージ
FROM golang:1.24-alpine

# Air (ライブリローダー)をインストール
RUN go install github.com/air-verse/air@latest

# 作業ディレクトリを作成・設定
WORKDIR /app

# 依存関係のレイヤーをキャッシュするため、先にgo.modをコピー
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# アプリケーションのコードをコピー
COPY . .

# Airを使ってサーバーを起動
CMD ["air"]
