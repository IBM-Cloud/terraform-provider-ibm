# Terraform for OpenShift Data Foundation

OpenShift Data Foundation is a highly available storage solution that you can use to manage persistent storage for your containerized workloads in Red Hat® OpenShift® on IBM Cloud® clusters. This folder contains the different ways you can deploy and manage OpenShift Data Foundation (ODF) on IBM Cloud through terraform.

## Deploying & Managing OpenShift Data Foundation on ROKS VPC

If you'd like to Deploy and Manage the different configurations for ODF on a Red Hat OpenShift Cluster (VPC) head over to the [addon](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/openshift-data-foundation/addon) folder.

## Updating or replacing worker nodes that use OpenShift Data Foundation on VPC clusters

If you'd like to update or replace the different worker nodes with ODF enabled, head over to the [vpc-worker-replace](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/openshift-data-foundation/vpc-worker-replace) folder. This inherently covers the worker replace steps of sequential cordon, drain, and replace.

## Deploying & Managing OpenShift Data Foundation on ROKS Satellite

If you'd like to Deploy and Manage ODF on a Red Hat OpenShift on a Satellite environment head over to the [satellite](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/openshift-data-foundation/satellite) folder.