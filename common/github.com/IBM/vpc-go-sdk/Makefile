# Makefile

all: build lint tidy

travis-ci: build lint tidy unittest

build:
	go build ./...

unittest:
	go test ./...

alltestgen1:
	go test `go list ./... | grep vpcclassicv1` -v -tags=integration -skipForMockTesting -testCount

alltestgen2:
	go test `go list ./... | grep vpcv1` -v -tags=integration -skipForMockTesting -testCount

lint:
	golangci-lint run

tidy:
	go mod tidy
