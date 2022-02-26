.PHONY: build run docker-run
.SILENT:

build:
	go build -o ./build/bin/file-service ./cmd/main.go

run: build
	./build/bin/file-service

docker-run:
	docker-compose up
