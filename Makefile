.DEFAULT_GOAL := help

## Run tests
# test:
# 	go test

## Install dependencies
deps:
	go get -u github.com/vmihailenco/msgpack
	go get -u github.com/jakm/msgpack-cli

## Build Releases
# release:
# 	  go build
demo.agent:
	go run cmd/anoco-agent.go

demo.command.valid:
	msgpack-cli encode sample/command.valid.json > /tmp/anoco.bin
	curl -D /dev/stderr localhost:15525/job -XPOST -d@/tmp/anoco.bin

demo.command.invalid:
	msgpack-cli encode sample/command.invalid.json > /tmp/anoco.bin
	curl -D /dev/stderr localhost:15525/job -XPOST -d@/tmp/anoco.bin


## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: test deps release
