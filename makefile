# Makefile

# Variables
GO_CMD = go run
MAIN = cmd/main.go
FILE = file.txt

# Default target
run:
	@$(GO_CMD) $(MAIN) $(FILE)

# Clean target (optional)
clean:
	@echo "Cleaning up..."
	@rm -f $(MAIN)

# Help target (optional)
help:
	@echo "Makefile for running Go Lem in application"
	@echo "Usage:"
	@echo "  make run      - Run the Go application"
	@echo "  make clean    - Clean up files (if needed)"
	@echo "  make help     - Display this help message"
