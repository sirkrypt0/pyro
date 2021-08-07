PROJECT_NAME := "pyro"
ROOT := "github.com/sirkrypt0/$(PROJECT_NAME)/cmd"
UNIT_TESTS = $(shell go list ./... | grep -v /e2e)
OUT = "out/"

default: help

.PHONY: all
all: build

.PHONY: bootstrap
bootstrap: deps lint-deps ## Install all dependencies

install-buf:
	@go install github.com/bufbuild/buf/cmd/buf@v0.48.2

.PHONY: deps
deps: install-buf ## Get the dependencies
	@go get -v -d ./...
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

.PHONY: proto
proto: ## Generate protobuf bindings
	@buf --config tools/buf/buf.yaml --template tools/buf/buf.gen.yaml generate

.PHONY: build
build: deps proto ## Build the binaries
	@go build -o $(OUT)/pyro-agent -v $(ROOT)/pyro-agent
	@go build -o $(OUT)/pyro -v $(ROOT)/pyro

.PHONY: clean
clean: ## Clean artifacts
	@rm -rf $(OUT)

.PHONY: lint-deps
lint-deps: ## Install linter dependencies
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: golangci-lint
golangci-lint: ## Lint the source code using golangci-lint
	@golangci-lint run ./... --timeout=3m

.PHONY: buf-lint
buf-lint: install-buf ## Lint the protobuf files
	@buf lint --config tools/buf/buf.yaml .
	@buf breaking --config tools/buf/buf.yaml --against-config tools/buf/buf.yaml --against '.git#branch=main'

.PHONY: lint
lint: golangci-lint buf-lint ## Lint the source code using all linters

.PHONY: test
test: deps ## Run unit tests
	@go test -count=1 -short $(UNIT_TESTS)

.PHONY: race
race: deps ## Run data race detector
	@go test -race -count=1 -short $(UNIT_TESTS)

.PHONY: coverage
coverage: deps ## Generate code coverage report
	@go test $(UNIT_TESTS) -v -coverprofile coverage.cov
	# exclude mock files from coverage
	@cat coverage.cov | grep -v _mock.go > coverage_cleaned.cov || true
	@go tool cover -func=coverage_cleaned.cov

.PHONY: coverhtml
coverhtml: coverage ## Generate HTML coverage report
	@go tool cover -html=coverage_cleaned.cov -o coverage_unit.html

.PHONY: help
HELP_FORMAT="    \033[36m%-25s\033[0m %s\n"
help: ## Display this help screen
	@echo "Valid targets:"
	@grep -E '^[^ ]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; \
			{printf $(HELP_FORMAT), $$1, $$2}'
