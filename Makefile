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

## Demo task for run agent
demo.agent:
	go run cmd/anoco-agent.go

## Demo task for request with valid message
demo.command.valid:
	msgpack-cli encode sample/command.valid.json > /tmp/anoco.bin
	curl -D /dev/stderr localhost:15525/job -XPOST -d@/tmp/anoco.bin

## Demo task for request with invalid message
demo.command.invalid:
	msgpack-cli encode sample/command.invalid.json > /tmp/anoco.bin
	curl -D /dev/stderr localhost:15525/job -XPOST -d@/tmp/anoco.bin

## Demo task for request status
demo.command.status:
	curl -D /dev/stderr localhost:15525/job/status

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: test deps release
