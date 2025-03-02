build:
	@go build -o bin/lx-kv

run: build
	@./bin/lx-kv
