.PHONY: build run test clean

up:
	docker-compose up -d

build:
	@go build -o bin/main cmd/main.go

run: build
	./bin/main

test:
	go test -v

clean:
	rm -rf bin