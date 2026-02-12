.PHONY: run build clean test install

# Run the scheduler
run:
	go run main.go

# Build the binary
build:
	go build -o scheduler main.go

# Install dependencies
install:
	go mod download

# Clean build artifacts
clean:
	rm -f scheduler
	go clean

# Run tests (when you add them)
test:
	go test ./...

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	golangci-lint run

# Build for multiple platforms
build-all:
	GOOS=linux GOARCH=amd64 go build -o scheduler-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o scheduler-windows-amd64.exe main.go
	GOOS=darwin GOARCH=amd64 go build -o scheduler-darwin-amd64 main.go
