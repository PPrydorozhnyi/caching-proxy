.PHONY: build run test lint

build:
	go build -o bin/api ./cmd/api

run:
	go run ./cmd/api

test:
	go test -race -v ./...

lint:
	gofmt -w .
	staticcheck ./...
	govulncheck ./...