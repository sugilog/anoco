.DEFAULT_GOAL := help

## Run tests
# test:
# 	go test

## Install dependencies
deps:
	go get github.com/vmihailenco/msgpack

## Build Releases
# release:
# 	  go build
demo.agent:
	go run cmd/anoco-agent.go

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: test deps release
