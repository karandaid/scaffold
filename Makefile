.PHONY: build run test test-race lint migrate clean

DB_DRIVER ?= sqlite

build:
	go build -o bin/cms cmd/server/main.go

run:
	DB_DRIVER=$(DB_DRIVER) go run cmd/server/main.go

test:
	go test ./... -count=1

test-race:
	go test ./... -race

lint:
	golangci-lint run

migrate:
	go run cmd/server/main.go migrate

clean:
	rm -rf bin/
