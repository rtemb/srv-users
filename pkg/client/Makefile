all: deps mocks

.PHONY: mocks
mocks:
	@go generate ./...

.PHONY: deps
deps:
	@go mod download
