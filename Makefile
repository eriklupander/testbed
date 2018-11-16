GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)
TEST_RESULTS=/tmp/test-results

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
	mkdir -p ${TEST_RESULTS}
	@go test -coverprofile=${TEST_RESULTS}/unittest.out -v $(GOPACKAGES)
	@go tool cover -html=${TEST_RESULTS}/unittest.out -o ${TEST_RESULTS}/unittest-coverage.html
	rm -f ${TEST_RESULTS}/unittest.out

integrationtest:
	mkdir -p ${TEST_RESULTS}
	@go test -coverprofile=${TEST_RESULTS}/integrationtest.out -tags="testtools integration"
	@go tool cover -html=${TEST_RESULTS}/integrationtest.out -o ${TEST_RESULTS}/integrationtest-coverage.html
	rm -f ${TEST_RESULTS}/integrationtest.out

e2etest:
	mkdir -p ${TEST_RESULTS}
	@go test -coverprofile=${TEST_RESULTS}/e2etest.out -tags="testtools e2e"
	@go tool cover -html=${TEST_RESULTS}/e2etest.out -o ${TEST_RESULTS}/e2etest-coverage.html
	rm -f ${TEST_RESULTS}/e2etest.out

run: build
	./dist/testbed
