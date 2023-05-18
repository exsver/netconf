GO_BIN ?= go
ENV_BIN ?= env

export PATH := $(PATH):/usr/local/go/bin

download:
	$(GO_BIN) get ./...
	$(GO_BIN) mod tidy

update:
	$(GO_BIN) get -u ./...
	$(GO_BIN) mod tidy

lint:
	# linter binary will be $(go env GOPATH)/bin/golangci-lint
	# curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
    # $(GO_BIN) get -u github.com/mgechev/revive
	golangci-lint run ./...
	revive -config revive.toml ./netconf/*.go
	revive -config revive.toml ./comware/*.go
	revive -config revive.toml ./junos/*.go