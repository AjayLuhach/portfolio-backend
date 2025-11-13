GO_VERSION ?= 1.23.0
SERVICES := gateway auth blog interaction notification analytics bookmark search worker

.PHONY: help install-go tidy fmt test run-% dev-up

help:
	@printf "Available targets:\n"
	@printf "  install-go   Install Go toolchain version $(GO_VERSION)\n"
	@printf "  tidy         go mod tidy\n"
	@printf "  fmt          go fmt ./...\n"
	@printf "  test         go test ./...\n"
	@printf "  run-<svc>    Run a service (e.g. make run-auth)\n"

install-go:
	@bash tools/install-go.sh $(GO_VERSION)

ENV ?= .env

fmt:
	@GOENV=$(ENV) gofmt -w $$(find . -name '*.go' -not -path './vendor/*')

tidy:
	@GOENV=$(ENV) go mod tidy

 test:
	@GOENV=$(ENV) go test ./...

$(SERVICES:%=run-%):
	@SERVICE=$(patsubst run-%,%,$@) && \
	echo "Starting $$SERVICE" && \
	GOENV=$(ENV) go run ./cmd/$$SERVICE
