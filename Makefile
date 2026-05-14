.PHONY: build run

# Build the Go-Go binary.
build:
	go build ./...
# Run the server with port 8080 and 4 workers.
run:
	go run ./cmd/server/main.go --port 8080 --workers 4
