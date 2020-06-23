# Example for Direct Link resources

This example shows how to create Direct Link resources.

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
| name | The unique user-defined name for this gateway. | `string` | yes |
| location |  Transit Gateway location. | `string` | yes |
| global | Gateways with global routing (true) can connect to networks outside their associated region. | `boolean` | no |


<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
