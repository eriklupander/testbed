GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

default: format build test

format:
	go fmt

build:
	mkdir -p dist
	export GO111MODULE=on
	go build -o dist/testbed

test:
	@go test -v $(GOPACKAGES)

integrationtest:
	@go test -tags=testtools integration

e2etest:
	@go test -tags=testtools e2e

run: build
	./dist/testbed
