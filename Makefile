GOLANGCI_LINT_VERSION = v1.49.0
GOLANGCI_LINT_BIN = /tmp/golangci-lint
GO_SWAGGER_URL = https://github.com/go-swagger/go-swagger/releases/download/v0.23.0/swagger_linux_amd64
GO_SWAGGER_BIN = /tmp/swagger
SWAGGER_URL_ADVERTISING = https://api.swaggerhub.com/apis/Palace/Advertising_Integration/2.0.0/

# Run docker image, mount in git repo, reset image's entrypoint so we can specify
DOCKER_RUN := docker run --rm -v $(PWD):/opt --workdir /opt --entrypoint=""

.PHONY: all
all: help

.PHONY: dep
dep: ## Install dependencies
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
		sh -s -- -b /tmp $(GOLANGCI_LINT_VERSION)
	curl -o $(GO_SWAGGER_BIN) -L $(GO_SWAGGER_URL) && chmod +x $(GO_SWAGGER_BIN)
	go get -v -d ./...

.PHONY: test
test: ## Test code
	export failed=0; \
	test -z $$(go fmt $(shell go list ./... | grep -v /vendor/)) || export failed=1; \
	go vet $(shell go list ./... | grep -v /vendor/); \
	$(GOLANGCI_LINT_BIN) run -v --enable-all || export failed=1; \
	go test -race $(shell go list ./... | grep -v /vendor/) || export failed=1; \
	if [ "$$failed" -eq 1 ]; then exit 1; fi

.PHONY: clean
clean: ## Clean artifacts directory
	$(RM) -r dist/*

.PHONY: build
build: ## Build example(s)
	mkdir -p dist
	go build -o dist ./examples/...


# Per: https://app.swaggerhub.com/help/apis/downloading-swagger-definition
.PHONY: fetch-swagger-spec
fetch-swagger-spec: ## Fetch Swagger Specs
	curl -o swagger.json $(SWAGGER_URL_ADVERTISING)
	$(GO_SWAGGER_BIN) validate swagger.json


.PHONY: generate-swagger-client
generate-swagger-client: ## Generate Swagger client
	$(GO_SWAGGER_BIN) generate client -f swagger.json

.PHONY: generate-swagger-server
generate-swagger-server: ## Generate Swagger server
	$(GO_SWAGGER_BIN) generate server -f swagger.json

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
