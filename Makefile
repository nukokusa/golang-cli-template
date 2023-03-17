.PHONY: init
init:
	@mv cmd/golang-cli-template cmd/$$(basename `git rev-parse --show-toplevel`)
	@grep -r -l golang-cli-template * .goreleaser.yml | xargs gsed -i "s/golang-cli-template/$$(basename `git rev-parse --show-toplevel`)/g"

# GOBIN ?= $(shell go env GOPATH)/bin

# .PHONY: lint test install

# lint: $(GOBIN)/staticcheck
# 	staticcheck ./...

# test: lint
# 	go test ./...

# install:
# 	go install github.com/nukokusa/golang-cli-template/cmd/golang-cli-template

# $(GOBIN)/staticcheck:
# 	@go install honnef.co/go/tools/cmd/staticcheck@latest
