build:
		@go build -o bin/DSA

run: build
		@./bin/DSA

test:
		@go test -v ./...
