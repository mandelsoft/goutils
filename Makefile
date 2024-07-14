
REPO_ROOT:= $(dir $(realpath $(lastword $(MAKEFILE_LIST))))../..

PLATFORMS = linux/amd64 linux/arm64 darwin/arm64 darwin/amd64 windows/amd64

SRCS=$(shell find $(REPO_ROOT) -type f) Makefile
GEN = $(REPO_ROOT)/gen/$(shell basename $(realpath .))

.PHONY: build
build: $(GEN)/build

$(GEN)/build: $(GEN)/.exists $(SRCS) $(SRCS)
	@for i in $(PLATFORMS); do \
    echo GOARCH=$$(basename $$i) GOOS=$$(dirname $$i) CGO_ENABLED=0 go build  ./...; \
    GOARCH=$$(basename $$i) GOOS=$$(dirname $$i) CGO_ENABLED=0 go build ./...; \
    done
	@touch $(GEN)/build

.PHONY: test
test:
	go test ./...

$(GEN)/.exists:
	@mkdir -p $(GEN)
	@touch $@
