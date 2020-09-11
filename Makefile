build:
	docker run --rm -v "${PWD}":/app golang:latest sh -c 'cd /app; make build-go'

build-go:
	go build -o bin/server cmd/dmserver/*.go