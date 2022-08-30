# Makefile

all: build lint tidy

travis-ci: build test-unit lint tidy

build:
	go build ./vpcv1

test-unit:
	go test `go list ./... | grep vpcv1` -v -tags=unit

test-integration:
	go test `go list ./... | grep vpcv1` -v -tags=integration -skipForMockTesting -testCount

test-examples:
	go test `go list ./... | grep vpcv1` -v -tags=examples

lint:
	golangci-lint run

tidy:
	go mod tidy
