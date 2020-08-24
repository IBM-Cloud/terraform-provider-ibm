# Example for VPC Flow Logs resources

This example shows how to create Flow Logs for VPC resources.

Following types of resources are supported:

* [Flow Logs](https://cloud.ibm.com/docs/terraform)


## Terraform versions

Terraform 0.12. Pin module version to `~> v1.5.1`. Branch - `master`.

Terraform 0.11. Pin module version to `~> v0.27.0`. Branch - `terraform_v0.11.x`.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Example Usage

Create a Flow log:

```hcl
data "ibm_resource_group" "cos_group" {
  name = var.resource_group
}

data "ibm_is_instance" "ds_instance" {
  name        = "vpc1-instance"
}

resource "ibm_resource_instance" "instance1" {
  name              = "cos-instance"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "cloud-object-storage"
  plan              = "standard"
  location          = "global"
}

resource "ibm_cos_bucket" "bucket1" {
   bucket_name          = "us-south-bucket-vpc1"
   resource_instance_id = ibm_resource_instance.instance1.id
   region_location = var.region
   storage_class = "standard"
}

resource ibm_is_flow_log test_flowlog {
  depends_on = [ibm_cos_bucket.bucket1]
  name = "test-instance-flow-log"
  target = data.ibm_is_instance.ds_instance.id
  active = true
  storage_bucket = ibm_cos_bucket.bucket1.bucket_name
}
```

## Examples

* [ Flow Log ](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-is-flow-log)

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | n/a |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| name | The unique user-defined name for this flow log collector. | `string` | yes |
| target | The id of the target this collector is to collect flow logs for. If the target is an instance, subnet, or VPC, flow logs will not be collected for any network interfaces within the target that are themselves the target of a more specific flow log collector. | `string` | yes |
| storage\_bucket | The name of the Cloud Object Storage bucket where the collected flows will be logged. The bucket must exist and an IAM service authorization must grant IBM Cloud Flow Logs resources of VPC Infrastructure Services writer access to the bucket. | `string` | yes |
| active | Indicates whether this collector is active. If false, this collector is created in inactive mode. Default is true. | `boolean` | no |
| resource\_group | The resource group ID where the flow log is to be created. | `string` | no |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
