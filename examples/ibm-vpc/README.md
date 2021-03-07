# Example for VpcV1

This example illustrates how to use the VpcV1

These types of resources are supported:

* DedicatedHostGroup

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## VpcV1 resources

is_dedicated_host_group resource:

```hcl
resource "is_dedicated_host_group" "is_dedicated_host_group_instance" {
  class = var.is_dedicated_host_group_class
  family = var.is_dedicated_host_group_family
  name = var.is_dedicated_host_group_name
  resource_group = var.is_dedicated_host_group_resource_group
  zone = var.is_dedicated_host_group_zone
}
```

## VpcV1 Data sources


## Assumptions

1. TODO

## Notes

1. TODO

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| class | The dedicated host profile class for hosts in this group. | `string` | false |
| family | The dedicated host profile family for hosts in this group. | `string` | false |
| name | The unique user-defined name for this dedicated host group. If unspecified, the name will be a hyphenated list of randomly-selected words. | `string` | false |
| resource_group | The resource group to use. If unspecified, the account's [default resourcegroup](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used. | `` | false |
| zone | The zone this dedicated host group will reside in. | `` | false |

## Outputs

| Name | Description |
|------|-------------|
| is_dedicated_host_group | is_dedicated_host_group object |
