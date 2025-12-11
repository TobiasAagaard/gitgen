.PHONY: build install test clean version run coverage

VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "v0.0.0-dev")
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS := -X github.com/TobiasAagaard/gitgen/pkg/version.Version=$(VERSION) \
           -X github.com/TobiasAagaard/gitgen/pkg/version.Commit=$(COMMIT) \
           -X github.com/TobiasAagaard/gitgen/pkg/version.Date=$(DATE)

build:
	@mkdir -p bin
	go build -ldflags "$(LDFLAGS)" -o bin/gitgen .

install:
	go install -ldflags "$(LDFLAGS)" .

run:
	go run -ldflags "$(LDFLAGS)" . $(filter-out $@,$(MAKECMDGOALS))

test:
	@mkdir -p .coverage
	go test -v -race -coverprofile=.coverage/coverage.out ./...

coverage: test
	go tool cover -html=.coverage/coverage.out -o .coverage/coverage.html
	@echo "Coverage report: .coverage/coverage.html"

clean:
	rm -rf bin/ .coverage/

version:
	@echo "Version: $(VERSION)"
	@echo "Commit:  $(COMMIT)"
	@echo "Date:    $(DATE)"

# Prevent make from treating args as targets
%:
