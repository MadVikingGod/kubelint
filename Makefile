# Parameters
BINARY_NAME=kubelint

all: test lint fmt build

build:
	go build -o bin/$(BINARY_NAME) -v ./cmd

lint:
	go vet ./...
	golint ./...
	golangci-lint run ./...

fmt:
	go fmt ./...

test:
	go test -v ./...

clean:
	go clean
	rm -f bin/$(BINARY_NAME)

docker-build:
	docker build . -t kubelint