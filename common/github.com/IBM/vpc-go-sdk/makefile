# Makefile

all: build lint tidy

travis-ci: build unittest lint tidy

build:
	go build ./vpcclassicv1
	go build ./vpcv1

unittest:
	cd vpcclassicv1 && go test -run TestVpcClassicV1 && cd ..
	cd vpcv1 && go test -run TestVpcV1 && cd ..

alltestgen1:
	go test `go list ./... | grep vpcclassicv1` -v -tags=integration -skipForMockTesting -testCount

alltestgen2:
	go test `go list ./... | grep vpcv1` -v -tags=integration -skipForMockTesting -testCount

lint:
	golangci-lint run

tidy:
	go mod tidy