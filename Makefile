run: build
	@./bin/go-blockchain

build: 
	@go build -o bin/go-blockchain

test:
	@go test ./...
