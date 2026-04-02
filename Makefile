.PHONY: build run test clean

up:
	docker-compose up -d

build:
	@go build -o bin/myapp main.go

run: build
	./bin/myapp

test:
	go test -v

clean:
	rm -rf bin