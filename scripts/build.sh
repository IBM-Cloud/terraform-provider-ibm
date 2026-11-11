#!/usr/bin/env bash
#
# This script builds the application from source for multiple platforms.

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR"

# Package which has the version information, required to set the Version, GitCommit info
VERSION_PACKAGE="github.com/terraform-providers/terraform-provider-ibm/version"

# Get the git commit
GIT_COMMIT=$(git rev-parse HEAD)
GIT_DIRTY=$(test -n "`git status --porcelain`" && echo "+CHANGES" || true)

# Determine the arch/os combos we're building for
XC_ARCH=${XC_ARCH:-"amd64" "arm64" "arm"}
XC_OS=${XC_OS:-linux darwin windows}
XC_EXCLUDE_OSARCH="!darwin/386 !windows/arm64 !windows/arm !darwin/arm"

# Delete the old dir
echo "==> Removing old directory..."
rm -f bin/*
rm -rf pkg/*
mkdir -p bin/

# If its dev mode, only build for ourself
if [ "${TF_DEV}x" != "x" ]; then
    XC_OS=$(go env GOOS)
    XC_ARCH=$(go env GOARCH)
fi

if ! which gox > /dev/null; then
    echo "==> Installing gox..."
    go get -u github.com/mitchellh/gox
fi

# instruct gox to build statically linked binaries
export CGO_ENABLED=0

# Allow LD_FLAGS to be appended during development compilations
LD_FLAGS="-X ${VERSION_PACKAGE}.GitCommit=${GIT_COMMIT}${GIT_DIRTY} $LD_FLAGS"

# In release mode we don't want debug information in the binary
if [[ -n "${TF_RELEASE}" ]]; then
    LD_FLAGS="-X ${VERSION_PACKAGE}.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X ${VERSION_PACKAGE}.VersionPrerelease= -s -w"
fi

# Build!
VERSION="1.64.0"
echo "==> Building..."
gox \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -osarch="${XC_EXCLUDE_OSARCH}" \
    -ldflags "${LD_FLAGS}" \
    -output "pkg/{{.OS}}_{{.Arch}}/terraform-provider-ibm" \
    .

# Resolve the path to the bundled native KMS crypto library.
# The upstream keyprotect-go-client module ships pre-built shared libraries
# under dedicated/internal/lib/{os}-{arch}/.  We copy the correct one next
# to the provider binary so the runtime search in getLibraryPath() finds it.
KP_MODULE_DIR="$(go env GOPATH)/pkg/mod/github.com/!i!b!m/keyprotect-go-client@v0.17.2/dedicated/internal/lib"

copy_kms_lib() {
    local os_arch="$1"   # e.g. linux_amd64
    local dest_dir="$2"
    # Convert go os_arch (underscore) to module dir style (hyphen)
    local lib_dir="${KP_MODULE_DIR}/${os_arch/_/-}"
    if [[ -d "$lib_dir" ]]; then
        cp "$lib_dir"/ibmkmscrypto* "$dest_dir"/ 2>/dev/null || true
    fi
}

# Move all the compiled things to the $GOPATH/bin
GOPATH=${GOPATH:-$(go env GOPATH)}
case $(uname) in
    CYGWIN*)
        GOPATH="$(cygpath $GOPATH)"
        ;;
esac
OLDIFS=$IFS
IFS=: MAIN_GOPATH=($GOPATH)
IFS=$OLDIFS

# Create GOPATH/bin if it's doesn't exists
if [ ! -d $MAIN_GOPATH/bin ]; then
    echo "==> Creating GOPATH/bin directory..."
    mkdir -p $MAIN_GOPATH/bin
fi

# Copy our OS/Arch binary and native lib to the bin/ directory
DEV_PLATFORM="./pkg/$(go env GOOS)_$(go env GOARCH)"
if [[ -d "${DEV_PLATFORM}" ]]; then
    for F in $(find ${DEV_PLATFORM} -mindepth 1 -maxdepth 1 -type f); do
        cp ${F} bin/
        cp ${F} ${MAIN_GOPATH}/bin/
    done
    copy_kms_lib "$(go env GOOS)_$(go env GOARCH)" bin/
    copy_kms_lib "$(go env GOOS)_$(go env GOARCH)" ${MAIN_GOPATH}/bin/
fi

if [ "${TF_DEV}x" = "x" ]; then
    # Zip and copy to the dist dir (include native lib in each platform archive)
    echo "==> Packaging..."
    for PLATFORM in $(find ./pkg -mindepth 1 -maxdepth 1 -type d); do
        OSARCH=$(basename ${PLATFORM})
        echo "--> ${OSARCH}"

        copy_kms_lib "${OSARCH}" "${PLATFORM}"

        pushd $PLATFORM >/dev/null 2>&1
        zip ../${OSARCH}.zip ./*
        popd >/dev/null 2>&1
    done
fi

# Done!
echo
echo "==> Results:"
ls -hl bin/