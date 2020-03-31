# Parameters
BINARY_NAME=kubelint

all: test build

build:
	go build -o bin/$(BINARY_NAME) -v

test:
	go fmt ./...
	go test -v ./...

clean:
	go clean
	rm -f bin/$(BINARY_NAME)

docker-build:
	docker build . -t kubelint