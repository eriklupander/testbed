GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

default: format build test vet

format:
	go fmt

vet:
	go vet ./...

build:
	mkdir -p dist
	export GO111MODULE=on
	go build -o dist/testbed

test:
	@go test -coverprofile=unittest.out -v $(GOPACKAGES)
	@go tool cover -html=unittest.out -o unittest-coverage.html

integrationtest:
	@go test -coverprofile=integrationtest.out -tags="testtools integration"
	@go tool cover -html=integrationtest.out -o integrationtest-coverage.html

e2etest:
	@go test -coverprofile=e2etest.out -tags="testtools e2e"
	@go tool cover -html=e2etest.out -o e2etest-coverage.html

run: build
	./dist/testbed
