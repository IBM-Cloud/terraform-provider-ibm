# Example for VpcV1

This example illustrates how to use the VpcV1

These types of resources are supported:


## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## VpcV1 resources


## VpcV1 Data sources

is_dedicated_host_profile data source:

```hcl
data "is_dedicated_host_profile" "is_dedicated_host_profile_instance" {
  name = var.is_dedicated_host_profile_name
}
```
is_dedicated_host_profiles data source:

```hcl
data "is_dedicated_host_profiles" "is_dedicated_host_profiles_instance" {
  name = var.is_dedicated_host_profiles_name
}
```

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
| name | The globally unique name for this virtual server instance profile. | `string` | false |
| name | The globally unique name for this dedicated host profile. | `string` | false |

## Outputs

| Name | Description |
|------|-------------|
| is_dedicated_host_profile | is_dedicated_host_profile object |
| is_dedicated_host_profiles | is_dedicated_host_profiles object |
