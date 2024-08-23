#!/bin/bash

function check_response {
    if [[ ${1} != 0 ]]; then
        exit ${1}
    fi
}

architectures=( linux_amd64 linux_arm64 darwin_amd64 darwin_arm64 )

for arc in "${architectures[@]}";
do
    echo "Uploading $arc provider zip to artifactory..."
    PROVIDER_FILE_NAME=${TF_BIN_FILE_NAME}_${PROVIDER_VERSION}_$arc.zip
    echo "Provider file name is $PROVIDER_FILE_NAME"

    curl -u${ARTIFACTORY_USERNAME}:${ARTIFACTORY_TOKEN} \
    -T "${WORKING_DIRECTORY}/pkg/$arc.zip" \
    -XPUT "${TF_REPO_URL}/${PROVIDER_VERSION}/${PROVIDER_FILE_NAME}"

    check_response ${?}
done
