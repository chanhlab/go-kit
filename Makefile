SHELL := /bin/bash -o pipefail

UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)

.PHONY: env
env:
	@echo "GOPATH: $(GOPATH)"

.PHONY: lint
lint:
	@echo "## Run GolangCI Lint"
	golangci-lint run

.PHONY: test
test:
	@echo "## Run Unit Tests"
	go test -covermode=atomic -coverprofile=coverage.out ./... -v
	go tool cover -html=coverage.out -o coverage.html
