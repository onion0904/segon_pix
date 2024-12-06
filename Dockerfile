FROM golang AS builder

WORKDIR /app

COPY . .

RUN go mod download && go mod tidy

CMD ["go", "run", "main.go"]
