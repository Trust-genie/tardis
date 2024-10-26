build:
	@go build -o bin/main ./cmd/main

run: build
	./bin/main

test:
	@go test -v ./...

bench:
	@go test -bench =. ./...

	