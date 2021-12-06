.PHONY: build run migrate-up migrate-down
.SILENT:

build:
	go build -o ./build/bin/file-service ./cmd/main.go

run: build
	./build/bin/file-service

migrate-up:
	migrate -path ./migrations -database \
	"postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_USE_SSL}" \
	up

migrate-down:
	migrate -path ./migrations -database \
	"postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_USE_SSL}" \
	down
