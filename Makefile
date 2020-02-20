.DEFAULT_GOAL=build

test:
	@go run scrupts/make.go --tests

build:
	@go run scripts/make.go

.PHONY: test build
