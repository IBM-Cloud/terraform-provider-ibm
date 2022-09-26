# IBM Cloud VPC Gen 2 Worker Replace

This example shows how to replace & update the Kubernetes VPC Gen-2 worker to the latest patch in the specified cluster. For more information, about VPC worker updates, see [Updating VPC worker nodes](https://cloud.ibm.com/docs/containers?topic=containers-update&interface=ui#vpc_worker_node)
 
## Usage

To run this example you need to execute:

```sh
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform untaint ibm_container_vpc_worker.<resource_name>[index] to untaint the failed worker after fixing it manually to proceed with next set of workers
Run `terraform destroy` when you need to provide new set of worker list

## Example usage

Perform worker replace:

```terraform
resource "ibm_container_vpc_worker" "worker" {
  count                         = length(var.worker_list)
  name                          = var.cluster_name
  replace_worker                = element(var.worker_list, count.index)
  resource_group_id             = data.ibm_resource_group.group.id
  kube_config_path              = data.ibm_container_cluster_config.cluster_config.config_file_path
  check_ptx_status              = var.check_ptx_status
  ptx_timeout                   = (var.ptx_timeout != null ? var.ptx_timeout : null)

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

## Examples

* [Portworx VPC Gen 2 worker replace](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/portworx/vpc-worker-replace)

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
| name | Name of the cluster. | `string` | yes |
| replace_worker | ID of the worker to be replaced. | `string` | yes |
| resource_group_id | ID of the resousrce group | `string` | no |
| check_ptx_status | Whether to check ptx status on replaced workers | `bool` | no |
| kube_config_path | The Cluster config with absolute path | `string` | no |
| ptx_timeout | Timeout used while checking the portworx status | `string` | no
{: caption="inputs"}

## Note

This resource is different from all other resource of IBM Cloud. Worker replace has 2 operations, i.e. Delete old worker & Create a new worker. On `Create` of terraform, Replace operation is being handled where both the deletion & creation happens whereas on the `Delete` of terraform, only the state is cleared but not the actual resource.
