# IBM Cloud resource instance example

The following example creates an instance of IBM Cloud resource. Instance could be of any resource, for example cloud object staorage, activity tracker, metrics monitor etc. By specifying the right value to argument `service`, we can provision respective resource instance.
Document reference http://servicedata.mybluemix.net/

Following types of resources are supported:

* [ Resource Instance](https://cloud.ibm.com/docs/terraform?topic=terraform-resource-mgmt-resources#resource-instance)


## Terraform versions

Terraform 0.12. Pin module version to `~> v1.7.1`. Branch - `master`.

Terraform 0.11. Pin module version to `~> v0.29.1`. Branch - `terraform_v0.11.x`.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Example Usage

Create an IBM Cloud Object Storage bucket resource instance. 

```hcl

data "ibm_resource_group" "cos_group" {
  name = var.resource_group_name
}

resource "ibm_resource_instance" "cos_instance" {
  name              = "cos-instance"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = var.service_type
  plan              = var.plan_type
  location          = var.location
}

data "ibm_resource_instance" "test" {
  name    = ibm_resource_instance.cos_instance.name
  service = var.service_type
}

```

## Examples

* [ Resource Istance  ](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-resource-instance)

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
| name | Name of the resource group. | `string` | yes |
| service_name | The name of the service offering. | `string` | yes |
| service_type | The type of the service offering. | `string` | yes |
| plan| The name of the plan type supported by service.| `string` | yes |
| location | Target location or environment to create the resource instance. | `string` | yes |


