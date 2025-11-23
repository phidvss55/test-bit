build:
	@go build -o bin/myapp main.go

run: build
	./bin/myapp

test:
	go test -v