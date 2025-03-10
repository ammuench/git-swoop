# Build variables
SWOOPVERSION := $(shell git describe --tags 2>/dev/null || echo "dev")
LDFLAGS += -X "main.swoopVersion=$(SWOOPVERSION)"
LDFLAGS += -X "main.goVersion=$(shell go version | sed -r 's/go version go(.*)\ .*/\1/')"
BINARY_NAME := git-swoop

# Go commands
GO := go
GOBUILD := $(GO) build
GOINSTALL := $(GO) install
GOCLEAN := $(GO) clean
GOMOD := $(GO) mod
GOVET := $(GO) vet
GOFMT := gofmt

.PHONY: build install clean test fmt vet tidy help

build:
	@echo "Building git-swoop 🐦..."
	$(GOBUILD) -ldflags '$(LDFLAGS)' -o $(BINARY_NAME)

install:
	@echo "Installing git-swoop 🐦..."
	$(GOINSTALL) -ldflags '$(LDFLAGS)'

clean:
	@echo "Cleaning git-swoop 🐦..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Add when we have tests to run 
# test:
# 	@echo "Running tests..."
# 	$(GOTEST) -v ./...

fmt:
	@echo "Formatting code for git-swoop 🐦..."
	$(GOFMT) -s -w .

vet:
	@echo "Vetting git-swoop code 🐦..."
	$(GOVET) ./...

tidy:
	@echo "Tidying git-swoop dependencies 🐦..."
	$(GOMOD) tidy

help:
	@echo "Available commands:"
	@echo "  make build    - Build the binary"
	@echo "  make install  - Install the binary to GOPATH/bin"
	@echo "  make clean    - Remove build artifacts"
	@echo "  make fmt      - Format code"
	@echo "  make vet      - Run go vet"
	@echo "  make tidy     - Tidy go.mod"
	@echo "  make all      - Clean, test, and build"
	@echo "  make help     - Show this help"

