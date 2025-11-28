# Variables
APP_NAME := echo-demo
BUILD_DIR := tmp
BIN := $(BUILD_DIR)/main.exe

# Default target
all: build

# Build the project
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BIN) .

# Run the project
run: build
	@echo "Running $(APP_NAME)..."
	$(BIN)

# Run with Air (live reload)
live:
	@echo "Starting live reload..."
	air

# Test the project
test:
	@echo "Running tests..."
	go test ./...

# Clean build files
clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

.PHONY: all build run live test clean fmt
