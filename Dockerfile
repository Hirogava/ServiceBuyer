FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata

RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN swag init -g cmd/main.go --parseDependency --parseInternal

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o servicebuyer cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

RUN adduser -D -s /bin/sh appuser

WORKDIR /app

COPY --from=builder /app/servicebuyer .

COPY --from=builder /app/docs ./docs

COPY --from=builder /app/internal/repository/migrations ./internal/repository/migrations

RUN chown -R appuser:appuser /app

USER appuser

EXPOSE 8080

CMD ["./servicebuyer"]
