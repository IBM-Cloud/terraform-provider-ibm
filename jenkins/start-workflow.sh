#!/bin/bash

set -e

build_version=$(date +"%Y%m%d")"-${BUILD_NUMBER}"

echo
echo "Arguments:"
echo " - app_version = '$build_version'"
echo " - branch = '$branch'"
echo

/home/jenkins/gosteps -template "./terraform-provider-ibm/jenkins/workflow.xml" \
                      -argument "build-version=$build_version" \
                      -argument "provider-version=$provider_version" \
                      -argument "build-number=${BUILD_NUMBER}"