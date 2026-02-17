# Makefile para comandos comunes

run:
	go run ./cmd/main.go

test:
	go test ./...

lint:
	golangci-lint run

migrate:
	migrate -path ./migrations -database "$(DB_URL)" up

build:
	go build -o bin/app ./cmd/main.go

swagger:
	swagger generate spec -o ./docs/swagger.json --scan-models
