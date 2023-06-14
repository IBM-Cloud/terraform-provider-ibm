#!/usr/bin/env bash

set -e

WORKING_DIR=$(pwd)

cp ${WORKING_DIR}/variables.tf ${WORKING_DIR}/ocscluster/variables.tf
cp ${WORKING_DIR}/input.tfvars ${WORKING_DIR}/ocscluster/input.tfvars 
cd ${WORKING_DIR}/ocscluster
terraform apply --auto-approve -var-file ${WORKING_DIR}/ocscluster/input.tfvars

sed -i'' -e  "s|ocsUpgrade = \"true\"|ocsUpgrade = \"false\"|g" ${WORKING_DIR}/input.tfvars
sed -i'' -e  "s|ocsUpgrade = \"true\"|ocsUpgrade = \"false\"|g" ${WORKING_DIR}/ocscluster/input.tfvars
rm -f ${WORKING_DIR}/input.tfvars-e
rm -f ${WORKING_DIR}/ocscluster/input.tfvars-e

terraform apply --auto-approve -var-file ${WORKING_DIR}/ocscluster/input.tfvars
