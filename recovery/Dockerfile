FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /recovery-service /app/recovery/recovery.go

CMD [ "/recovery-service" ]
