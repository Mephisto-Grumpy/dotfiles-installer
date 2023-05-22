.PHONY: build test lint

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOLINT=golangci-lint run

BINARY_NAME=dotfiles-installer

BIN=$(CURDIR)/bin
MAIN=$(CURDIR)/cmd/app

all: test build

build: build-linux build-darwin

build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BIN)/$(BINARY_NAME)-linux-amd64 $(MAIN)

build-darwin:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BIN)/$(BINARY_NAME)-darwin-amd64 $(MAIN)

test: 
	$(GOTEST) -v ./...

lint:
	$(GOLINT)
