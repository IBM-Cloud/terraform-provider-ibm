#!/usr/bin/env bash

set -e

WORKING_DIR=$(pwd)

cp ${WORKING_DIR}/variables.tf ${WORKING_DIR}/ocscluster/variables.tf
cp ${WORKING_DIR}/schematics.tfvars ${WORKING_DIR}/ocscluster/schematics.tfvars
cd ${WORKING_DIR}/ocscluster
terraform init
if [ -e ${WORKING_DIR}/ocscluster/terraform.tfstate ]
then
    echo "ok"
else
    terraform import -var-file=${WORKING_DIR}/ocscluster/schematics.tfvars kubernetes_manifest.ocscluster_ocscluster_auto "apiVersion=ocs.ibm.io/v1,kind=OcsCluster,namespace=openshift-storage,name=ocscluster-auto"
fi

terraform apply --auto-approve -var-file ${WORKING_DIR}/ocscluster/schematics.tfvars

sed -i'' -e  "s|ocsUpgrade = \"true\"|ocsUpgrade = \"false\"|g" ${WORKING_DIR}/schematics.tfvars
sed -i'' -e  "s|ocsUpgrade = \"true\"|ocsUpgrade = \"false\"|g" ${WORKING_DIR}/ocscluster/schematics.tfvars
rm -f ${WORKING_DIR}/schematics.tfvars-e
rm -f ${WORKING_DIR}/ocscluster/schematics.tfvars-e

terraform apply --auto-approve -var-file ${WORKING_DIR}/ocscluster/schematics.tfvars