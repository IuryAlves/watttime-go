GOCMD?=CGO_ENABLED=0 go
GO_LINT?=$(shell which golangci-lint)

.PHONY: test
test:
	$(GOCMD) test ./pkg -mod=vendor -count=1 -v

.PHONY: fmt
fmt:
	$(GOCMD) fmt ./...

.PHONY: lint
lint: vendor
	$(GO_LINT) -c build/golangci.yaml run

.PHONY: vendor
vendor:
	$(GOCMD) mod vendor

.PHONY: docs
docs:
	gomarkdoc -e -o README.md ./pkg
