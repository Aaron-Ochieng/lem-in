# Makefile

# Variables
GO_CMD = go run
MAIN = cmd/main.go
FILE = file.txt

# Default target
run:
	@$(GO_CMD) $(MAIN) $(FILE)


# Help target (optional)
help:
	@echo "Makefile for running Go Lem in application"
	@echo "Usage:"
	@echo "  make run      - Run the Go application"
	@echo "  make help     - Display this help message"
