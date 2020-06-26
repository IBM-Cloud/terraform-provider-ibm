# Example for Transit Gateway resources

This example shows how to create Transit Gateway resources.

Following types of resources are supported:

* [Transit Gateways](https://cloud.ibm.com/docs/terraform)


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

Create a transit gateway:

```hcl
resource "ibm_tg_gateway" "new_tg_gw"{
name="tg-gateway-1"
location="us-south"
global=true
} 
```

## Examples

* [ Transit Gateway](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-transit-gateway)

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
| name | Name of Transit Gateway Services. | `string` | yes |
| location |  Location of Transit Gateway Services. | `string` | yes |
| global | Allow global routing for a Transit Gateway. If unspecified, the default value is false. | `boolean` | no |


<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
