SHELL = sh

default: test

.PHONY: test
test:
	@echo "--> Running tests ..."
	@go test -count=1 -v -race ./...

.PHONY: copywrite
copywrite:
	@echo "--> Checking Copywrite ..."
	copywrite \
		--config .github/workflows/scripts/copywrite.hcl headers \
		--spdx "BSD-3-Clause"

.PHONY: vet
vet:
	@echo "--> Vet Go sources ..."
	@go vet ./...

.PHONY: lint
lint: vet
	@echo "--> Lint ..."
	@golangci-lint run --config .github/workflows/scripts/golangci.yaml
