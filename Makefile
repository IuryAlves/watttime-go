GOCMD?=CGO_ENABLED=0 go

.PHONY: test
test:
	$(GOCMD) test ./... -mod=vendor -count=1