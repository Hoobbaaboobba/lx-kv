run: build
	@./bin/lx-kv --listenAddr :6379
build:
	@go build -o bin/lx-kv .
test:
	@go test -v ./...
