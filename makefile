# Makefile shamelessly "inspired" by https://github.com/helm/helm
GO        ?= go
TAGS      :=
LDFLAGS	  := -w -s
TESTS     := .
TESTFLAGS :=
GOFLAGS   :=
RUNFLAGS  :=
BINDIR    := $(CURDIR)/bin

GOOS := linux
GOARCH := amd64

#
# CI/CD Commands
#
all: server

server:
	@rm -rf ./main \
	export GOPRIVATE="git.dochq.co.uk/packages" && \
	git config --global url."https://github.com:${CI_REGISTRY_PASSWORD}@github.com".insteadOf "https://git.dochq.co.uk/packages"
	CGO_ENABLED=0 GOOS=$(shell uname -s | tr '[:upper:]' '[:lower:]') GOARCH=amd64 $(GO) build -a -ldflags="$(LDFLAGS)" -v ./src/main/main.go;

.PHONY: clean
clean:
	@rm -rf ./server
