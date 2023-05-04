#!/bin/bash

set -e

WORKING_DIR=$(pwd)

cp ${WORKING_DIR}/variables.tf ${WORKING_DIR}/ocscluster/variables.tf
cp ${WORKING_DIR}/input.tfvars ${WORKING_DIR}/ocscluster/input.tfvars 
cd ${WORKING_DIR}/ocscluster
terraform init
terraform apply --auto-approve -var-file ${WORKING_DIR}/ocscluster/input.tfvars