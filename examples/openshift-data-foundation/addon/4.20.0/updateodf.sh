#!/bin/bash

set -e 

WORKING_DIR=$(pwd)

cp ${WORKING_DIR}/variables.tf ${WORKING_DIR}/ibm_odf_addon/variables.tf
cp ${WORKING_DIR}/schematics.tfvars ${WORKING_DIR}/ibm_odf_addon/schematics.tfvars
cd ${WORKING_DIR}/ibm_odf_addon
terraform init
terraform apply --auto-approve -var-file ${WORKING_DIR}/ibm_odf_addon/schematics.tfvars
