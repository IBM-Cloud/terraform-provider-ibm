#!/bin/bash

set -e 

WORKING_DIR=$(pwd)

cp ${WORKING_DIR}/variables.tf ${WORKING_DIR}/ibm_odf_addon/variables.tf
cp ${WORKING_DIR}/input.tfvars ${WORKING_DIR}/ibm_odf_addon/input.tfvars 
cd ${WORKING_DIR}/ibm_odf_addon
terraform apply --auto-approve -var-file ${WORKING_DIR}/ibm_odf_addon/input.tfvars
