# Makefile

# Variables
GO_CMD = go run
MAIN = cmd/main.go
FILE = file.txt
GO_TEST= go test -v ./...

# Default target
run:
	@$(GO_CMD) $(MAIN) $(FILE)


test:
	@$(GO_TEST)

# Help target (optional)
help:
	@echo "Makefile for running Go Lem in application"
	@echo "Usage:"
	@echo "  make run      - Run the Go application"
	@echo "make test       - Run the Go test files within the project"
	@echo "  make help     - Display this help message"
