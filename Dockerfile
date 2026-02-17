# Multi-stage Dockerfile para desarrollo y producción
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o sugarerpgo ./cmd/main.go

FROM gcr.io/distroless/base-debian10
WORKDIR /app
COPY --from=builder /app/sugarerpgo .
CMD ["./sugarerpgo"]
