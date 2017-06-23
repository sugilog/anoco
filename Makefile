.DEFAULT_GOAL := help

## Run tests
# test:
# 	go test

## Install dependencies
deps:
	go get -u github.com/vmihailenco/msgpack
	go get -u github.com/jakm/msgpack-cli
	go get -u github.com/google/uuid

## Build Releases
# release:
# 	  go build

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: test deps release
