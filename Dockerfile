# 1. Goのビルド環境を設定
FROM golang AS builder

WORKDIR /app

  COPY . .
RUN go mod download && go mod tidy

CMD ["go", "run", "main.go"]
