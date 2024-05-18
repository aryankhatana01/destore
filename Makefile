build:
	@go build -o bin/destore

run: build
	@./bin/destore

test:
	@go test -v ./...