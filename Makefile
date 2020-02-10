GO_BIN ?= go
ENV_BIN ?= env

export PATH := $(PATH):/usr/local/go/bin

update:
	$(GO_BIN) get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	$(GO_BIN) get -u github.com/mgechev/revive

lint:
	golangci-lint run ./netconf/*.go
	golangci-lint run ./comware/*.go
	golangci-lint run ./junos/*.go
	revive -config revive.toml ./netconf/*.go
	revive -config revive.toml ./comware/*.go
	revive -config revive.toml ./junos/*.goq