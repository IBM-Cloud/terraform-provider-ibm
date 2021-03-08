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

is_dedicated_hosts data source:

```hcl
data "is_dedicated_hosts" "is_dedicated_hosts_instance" {
  name = var.is_dedicated_hosts_name
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
| name | The unique user-defined name for this dedicated host. If unspecified, the name will be a hyphenated list of randomly-selected words. | `string` | false |

## Outputs

| Name | Description |
|------|-------------|
| is_dedicated_hosts | is_dedicated_hosts object |
