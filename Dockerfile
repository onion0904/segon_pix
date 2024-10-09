# 1. Goのビルド環境を設定
FROM golang AS builder

WORKDIR /app

# # 2. Goモジュールの依存関係を追加
# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

# 3. アプリケーションのソースコードを追加
COPY . .
RUN go mod download && go mod tidy

CMD ["go", "run", "main.go"]
