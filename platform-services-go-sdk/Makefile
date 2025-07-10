# Makefile to build the project
GO=go
LINT=golangci-lint
COVERAGE = -coverprofile=coverage.txt -covermode=atomic

all: tidy test lint
travis-ci: tidy test-cov lint

test:
	${GO} test ./...

test-cov:
	${GO} test ./... ${COVERAGE}

test-int:
	${GO} test ./... -tags=integration

test-int-cov:
	${GO} test ./... -tags=integration ${COVERAGE}

lint:
	${LINT} run --build-tags=integration,examples

tidy:
	${GO} mod tidy
