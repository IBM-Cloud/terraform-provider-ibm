#!/bin/bash
export HOME=/home/jenkins
mkdir $HOME/go
export GOPATH=$HOME/go
export GOX_VERSION="v1.0.1"

function check_response {
    if [[ ${1} != 0 ]]; then
        exit ${1}
    fi
}

echo "Step 1: Navigate to terraform directory..."
cd ${WORKING_DIRECTORY}

echo "Step 2: Set git config url..."
git config --global url."git@github.ibm.com:".insteadOf "https://github.ibm.com/"
go env -w GOPRIVATE=github.ibm.com

echo "Step 3: running go clean and mod tidy"
go clean -modcache
go mod tidy

echo "Step 4: Running make tools..."
make tools
check_response ${?}

echo "Step 5: Install gox lib..."
go install github.com/mitchellh/gox@${GOX_VERSION}
export PATH=$PATH:$GOPATH/bin

# add dependencies to current module and install them
echo "Step 6: Running go get..."
go get
check_response ${?}

echo "Step 7: Running make bin..."
make bin
check_response ${?}
