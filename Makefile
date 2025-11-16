.PHONY: help build test lint fmt deps install clean

# Binary name
BINARY_NAME=test-repo-standard-check

# Build variables
BUILD_DIR=bin
GO=go
GOFLAGS=-v

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

deps: ## Download dependencies
	$(GO) mod download
	$(GO) mod verify

build: ## Build the binary
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .

test: ## Run tests with coverage
	$(GO) test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

test-short: ## Run tests without coverage
	$(GO) test -v -short ./...

lint: ## Run linters
	@which golangci-lint > /dev/null || (echo "golangci-lint not found. Install from https://golangci-lint.run/usage/install/" && exit 1)
	golangci-lint run ./...

fmt: ## Format code
	$(GO) fmt ./...
	@which goimports > /dev/null && goimports -w . || true

install: ## Install binary to GOPATH/bin
	$(GO) install $(GOFLAGS) .

clean: ## Remove build artifacts
	rm -rf $(BUILD_DIR)
	rm -f coverage.txt
	$(GO) clean

run: build ## Build and run the binary
	./$(BUILD_DIR)/$(BINARY_NAME)

.DEFAULT_GOAL := help
