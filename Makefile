SHELL := /bin/bash
ROOT := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
TITLE := $(shell basename $(ROOT))
APP_PKG := $(shell go list -m)
export BIN_OUT := $(ROOT)/bin

all: print test build
.PHONY: print test build

print:
	@echo SHELL:$(SHELL)
	@echo ROOT:$(ROOT)
	@echo TITLE:$(TITLE)
	@echo APP_PKG:$(APP_PKG)
	@echo BIN_OUT:$(BIN_OUT)
	@echo -e "\n"

test:
	go test ./... -v -cover
	@echo -e "\n"

build:
	tool/build.sh ${BIN_OUT}/${TITLE} ${APP_PKG}/pkg/version ${APP_PKG}
	@echo -e "\n"
