GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

default: build test

build:
	mkdir -p dist
	export GO111MODULE=on
	go build -o dist/testbed

test:
	@go test -v $(GOPACKAGES)
