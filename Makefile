GOCMD?=CGO_ENABLED=0 go
GO_LINT?=$(shell which golangci-lint)

.PHONY: test
test:
	$(GOCMD) test ./... -mod=vendor -count=1


.PHONY: lint
lint: vendor
	$(GO_LINT) -c build/golangci.yaml run

.PHONY: vendor
vendor:
	$(GOCMD) mod vendor

build:
	$(GOCMD) build -o deploy-tool github.com/voiapp/deploy-tool/cmd/deploytool