# This Module is used to create satellite ROKS Cluster

This module creates `satellite cluster and worker pool` for the specified location.

## Prerequisite

* Set up the IBM Cloud command line interface (CLI), the Satellite plug-in, and other related CLIs.
* Install cli and plugin package
```console
    ibmcloud plugin install container-service
```
## Usage

```
terraform init
```
```
terraform plan
```
```
terraform apply
```
```
terraform destroy
```
## Example Usage

``` hcl
module "satellite-cluster" {
  source               = "./modules/cluster"

  cluster              = var.cluster
  location             = module.satellite-host.location
  kube_version         = var.kube_version
  default_wp_labels    = var.default_wp_labels
  zones                = var.cluster_zones
  resource_group       = var.resource_group
  worker_pool_name     = var.worker_pool_name
  worker_count         = var.worker_count
  workerpool_labels    = var.workerpool_labels
  cluster_tags         = var.cluster_tags
  host_labels          = var.host_labels
  zone_name            = var.zone_name
}
```
<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name                          | Description                                                       | Type     | Default | Required |
|-------------------------------|-------------------------------------------------------------------|----------|---------|----------|
| ibmcloud_api_key              | IBM Cloud API Key.                                                | string   | n/a     | yes      |
| resource_group                | Resource Group Name that has to be targeted.                      | string   | n/a     | no       |
| ibm_region                    | The location or the region in which VM instance exists.           | string   | us-east | yes      |
| cluster                       | Name of the ROKS Cluster that has to be created                   | string   | n/a     | yes      |
| location                      | Name of the Location that has to be created                       | string   | n/a     | yes      |
| zones                         | Allocate your hosts across these three zones                      | set      | n/a     | yes      |
| kube_version                  | Kuber version                                                     | string   | 4.6_openshift   | yes      |
| cluster                       | The name for the new IBM Cloud Satellite cluster                  | string   | n/a     | no       |
| kube_version                  | The OpenShift Container Platform version                          | string   | n/a     | no       |
| default_wp_labels             | Labels on the default worker pool                                 | map      | n/a     | no       |
| worker_pool_name              | Public SSH key used to provision Host/VSI                         | string   | n/a     | no       |
| host_labels                   | List of host labels to assign host to cluter                      | list     | n/a     | no       |
| workerpool_labels             | Labels on the worker pool                                         | map      | n/a     | no       |
| cluster_tags                  | List of tags for the cluster resource                             | list     | n/a     | no       |
| zone_name                     | creates new zone on workerpool                                    | string   | n/a     | no       |


<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->