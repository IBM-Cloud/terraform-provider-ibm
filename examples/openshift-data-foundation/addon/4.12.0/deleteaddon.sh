#!/usr/bin/env bash

set -e

WORKING_DIR=$(pwd)

cp ${WORKING_DIR}/variables.tf ${WORKING_DIR}/ibm_odf_addon/variables.tf
cp ${WORKING_DIR}/schematics.tfvars ${WORKING_DIR}/ibm_odf_addon/schematics.tfvars
cd ${WORKING_DIR}/ibm_odf_addon
terraform init
if [ -e ${WORKING_DIR}/ibm_odf_addon/terraform.tfstate ]
then
    echo "ok"
else
   terraform apply --auto-approve -var-file=${WORKING_DIR}/ibm_odf_addon/schematics.tfvars
fi
terraform destroy --auto-approve -var-file=${WORKING_DIR}/ibm_odf_addon/schematics.tfvars