# Dockerfile
FROM golang:1.21

# 作業ディレクトリの設定
WORKDIR /app

# SQLCのインストール
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# 必要なファイルをコピー
COPY go.mod go.sum ./
RUN go mod download

# パッケージのインストール
RUN go mod download

# 必要なファイルをコピー
COPY . .
