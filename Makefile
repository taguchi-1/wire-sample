# Option
#===============================================================
ENV           := local
OS            := $(shell uname | tr A-Z a-z )
SHELL         := /bin/bash
BUILD_OPTIONS := -tags netgo -installsuffix netgo
PREFIX        := /usr/local
INSTALL_BIN   :=
TAG_OPTION    := local-latest

# Const
#===============================================================
name                 := wire-sample
bin_dir              := bin


# Task
#===============================================================

setup:
ifeq ($(shell command -v golint 2> /dev/null),)
	go get -u golang.org/x/lint/golint
endif
ifeq ($(shell command -v goimports 2> /dev/null),)
	go get -u golang.org/x/tools/cmd/goimports
endif
ifeq ($(shell command -v wire 2> /dev/null),)
	go get github.com/google/wire/cmd/wire
endif

wire:
	wire infra/container/wire.go infra/container/providers.go

fmt:
	goimports -w cmd/$(name)/main.go
	for pkg in $$(go list -f {{.Dir}} ./... | grep -v "$(name)$$"); do \
		goimports -w $$pkg; \
	done

lint:
	for pkg in $$(go list ./...); do \
		golint -set_exit_status $$pkg || exit $$?; \
	done

test:
	go test $$(go list ./... | grep -v /test/ | grep -v /proto| tr '\n' ' ')

build:
	$(eval revision := $(shell if [[ $$REV = "" ]]; then git rev-parse --short HEAD; else echo $$REV;fi;))
	$(eval ldflags  := -X 'main.revision=$(revision)' -extldflags '-static')
	GOOS=$(OS) GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "$(ldflags)" -o $(bin_dir)/$(name)_$(OS)_amd64 $(BUILD_OPTIONS) cmd/$(name)/main.go

release: wire fmt lint test build

install:
ifeq ($(INSTALL_BIN),)
	$(eval bin := $(name)_$(OS)_amd64)
else
	$(eval bin := $(INSTALL_BIN))
endif
	chmod +x $(bin_dir)/$(name)_$(OS)_amd64
	if [ ! -d $(PREFIX)/bin ]; then mkdir -p $(PREFIX)/bin; fi
	cp -a $(bin_dir)/$(name)_$(OS)_amd64 $(PREFIX)/bin/$(bin)

.PHONY: setup fmt lint test build release install
.DEFAULT_GOAL := release