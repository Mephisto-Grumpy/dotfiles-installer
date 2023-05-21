.PHONY: build test lint

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOLINT=golangci-lint run

BINARY_NAME=dotfiles-installer

all: test build

build: 
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/app

test: 
	$(GOTEST) -v ./...

lint:
	$(GOLINT)
