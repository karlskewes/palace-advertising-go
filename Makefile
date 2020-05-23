GORELEASER_BIN = ./bin/goreleaser
GOLANGCI_LINT_VERSION = v1.27.0
JQ_URL = https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64
GO_SWAGGER_URL = https://github.com/go-swagger/go-swagger/releases/download/v0.23.0/swagger_linux_amd64

SWAGGER_BIN = $${GOPATH}/bin/swagger
SWAGGER_URL_ADVERTISING = https://api.swaggerhub.com/apis/Palace/Advertising_Integration/1.0.0/

# Run docker image, mount in git repo, reset image's entrypoint so we can specify
DOCKER_RUN := docker run --rm -v $(PWD):/opt --workdir /opt --entrypoint=""

.PHONY: all
all: help

.PHONY: dep
dep: ## Install dependencies
	curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh
	# binary will be $(go env GOPATH)/bin/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
		sh -s -- -b $${GOPATH}/bin $(GOLANGCI_LINT_VERSION)
	curl -o $(SWAGGER_BIN) -L $(GO_SWAGGER_URL) && chmod +x $(SWAGGER_BIN)
	ls -l $${GOPATH}/bin
	go get -v -d ./...

.PHONY: test
test: ## Test code
	export failed=0; \
	test -z $$(go fmt $(shell go list ./... | grep -v /vendor/)) || export failed=1; \
	go vet $(shell go list ./... | grep -v /vendor/); \
	golint -set_exit_status $(shell go list ./... | grep -v /vendor/) || export failed=1; \
	golangci-lint run -v --enable-all || export failed=1; \
	go test -race $(shell go list ./... | grep -v /vendor/) || export failed=1; \
	if [ "$$failed" -eq 1 ]; then exit 1; fi

.PHONY: clean
clean: ## Clean artifacts directory
	$(RM) -r dist/

.PHONY: snapshot
snapshot: clean dep ## Generate a snapshot release.
	$(GORELEASER_BIN) --snapshot --skip-validate --skip-publish

.PHONY:
release: clean dep ## Generate a release, but don't publish to GitLab.
	$(GORELEASER_BIN) --skip-validate --skip-publish

.PHONY: publish
publish: clean dep ## Generate a release, and publish to GitLab.
	$(GORELEASER_BIN)


# Per: https://app.swaggerhub.com/help/apis/downloading-swagger-definition
.PHONY: fetch-swagger-spec
fetch-swagger-spec: ## Fetch Swagger Specs
	curl -o swagger.json $(SWAGGER_URL_ADVERTISING)
	$(SWAGGER_BIN) validate swagger.json


.PHONY: generate-swagger-client
generate-swagger-client: ## Generate Swagger client
	$(SWAGGER_BIN) generate client -f swagger.json

.PHONY: generate-swagger-server
generate-swagger-server: ## Generate Swagger server
	$(SWAGGER_BIN) generate server -f swagger.json



.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
