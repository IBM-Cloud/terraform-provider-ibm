# OpenShift-Data-Foundation VPC ROKS Worker Replace

This example shows how to replace & update the Kubernetes VPC Gen-2 worker installed with Openshift-Data-Foundation to the latest patch in the specified cluster.

For more information, about VPC worker updates, see [Updating VPC worker nodes](https://cloud.ibm.com/docs/containers?topic=containers-update&interface=ui#vpc_worker_node)

## Usage

To run this example you need to execute:

```sh
$ terraform init
$ terraform plan -var-file input.tfvars
$ terraform apply -var-file input.tfvars
```

* Run `terraform untaint ibm_container_vpc_worker.<resource_name>[index]` to untaint the failed worker after fixing it manually to proceed with next set of workers
* Run `terraform destroy` when you need to provide new set of worker list

A Sample `input.tfvars` is as follows -

```
worker_list = [ "worker_id_1 ","worker_id_2","worker_id_3" ]
cluster_name = ""
ibmcloud_api_key = ""
resource_group = ""
```
In-case you want to perform worker update or replace on non ODF nodes, remove the sds and sds_timeout parameters from the main.tf file.

## Example usage

Perform worker replace:

```terraform
resource "ibm_container_vpc_worker" "worker" {
  count                         = length(var.worker_list)
  cluster_name                  = var.cluster_name
  replace_worker                = element(var.worker_list, count.index)
  resource_group_id             = data.ibm_resource_group.group.id
  kube_config_path              = data.ibm_container_cluster_config.cluster_config.config_file_path
  sds                           = "ODF"
  sds_timeout                   = (var.sds_timeout != null ? var.sds_timeout : null)

  timeouts {
    create = (var.create_timeout != null ? var.create_timeout : null)
    delete = (var.delete_timeout != null ? var.delete_timeout : null)
  }
}
```

```terraform
data ibm_resource_group group {
  name = var.resource_group
}

data ibm_container_cluster_config cluster_config {
    cluster_name_id   = var.cluster_name
    resource_group_id = data.ibm_resource_group.group.id
}
```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

## Requirements

| Name | Version |
|------|---------|
| terraform | >=1.0.0, <2.0 |

## Providers

| Name | Version |
|------|---------|
| ibm  | latest |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| cluster_name | Name of the cluster. | `string` | yes |
| replace_worker | ID of the worker to be replaced. | `string` | yes |
| resource_group_id | ID of the resousrce group | `string` | no |
| check_ptx_status | Whether to check ptx status on replaced workers | `bool` | no |
| kube_config_path | The Cluster config with absolute path | `string` | no |
| ptx_timeout | Timeout used while checking the portworx status | `string` | no
| sds | Software Defined Storage - Only `ODF` is currently supported | `string` | no
| sds_timeout | Timeout used while checking the sds status/deployment | `string` | no

## Note

* This resource is different from all other resource of IBM Cloud. Worker replace has 2 operations, i.e. Delete old worker & Create a new worker. On `terraform apply`, Replace operation is being handled where both the deletion & creation happens whereas on the `terraform destroy`, only the state is cleared but not the actual resource.
* When the worker list is being provided as inputs, the list must be user generated and should not be passed from the `ibm_container_cluster` data source.
* If `terraform apply` fails during worker replace or while checking the sds status, perform any one of the following actions before retrying.
   * Resolve the issue manually and perform `terraform untaint` to proceed with the subsequent workers in the list.
   * If worker replace is still needed, update the input list by replacing the existing worker id with the new worker id.
   


* Please note, currently only `ODF` is a supported value for the `sds` input
