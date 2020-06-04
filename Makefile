COVEROUT := ./coverage.out
APP_NAME := srv-users
OS := linux
#OS := darwin

VERSION_PATH := gitlab.com/rtemb/srv-users/pkg/version.ServiceVersion

.PHONY: deps
deps:
	$(info Installing dependencies)
	GO111MODULE=on go mod download
	GO111MODULE=on go mod vendor

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=$(OS) go build -a -ldflags "-X '$(VERSION_PATH)=$(VERSION)'" -installsuffix cgo -o main ./cmd/...

.PHONY: test
test:
	go test --cover -covermode=atomic -coverprofile=$(COVEROUT) --race -count=1 ./...

.PHONY: lint
lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint -v --deadline=180s run