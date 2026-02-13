.PHONY: build run test clean

APP_NAME=check-network
VERSION=1.0.0

build:
	go build -o bin/$(APP_NAME) -ldflags="-X 'main.Version=$(VERSION)'" cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/
	go clean

deps:
	go mod tidy
	go mod download

docker-build:
	docker build -t $(APP_NAME):$(VERSION) .

docker-run:
	docker run -p 8080:8080 $(APP_NAME):$(VERSION)