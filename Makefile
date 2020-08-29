# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD)
GOINSTAL=$(GOCMD) install
BINARY_NAME=notebox

.PHONY: all
all: test build

build:
				$(GOBUILD) -o bin/$(BINARY_NAME)  cmd/notebox/notebox.go

.PHONY: clean
clean:
				$(GOCLEAN)
				rm -f bin/$(BINARY_NAME)

.PHONY: test
test:
				$(GOTEST) -v ./...

run:
				$(GORUN) cmd/notebox/notebox.go

get:
				$(GOGET)  -d -v ./...

install:
				$(GOINSTALL)  -d -v ./...
