.PHONY: init
init:
ifeq ($(shell uname -s),Darwin)
	@grep -r -l golang-cli-template * | xargs sed -i "" "s/golang-cli-template/$$(basename `git rev-parse --show-toplevel`)/"
else
	@grep -r -l golang-cli-template * | xargs sed -i "s/golang-cli-template/$$(basename `git rev-parse --show-toplevel`)/"
endif

# GOBIN ?= $(shell go env GOPATH)/bin

# .PHONY: lint test install

# lint: $(GOBIN)/staticcheck
# 	staticcheck ./...

# test: lint
# 	go test ./...

# install:
# 	go install github.com/nukokusa/golang-cli-template/cmd/app

# $(GOBIN)/staticcheck:
# 	@go install honnef.co/go/tools/cmd/staticcheck@latest
