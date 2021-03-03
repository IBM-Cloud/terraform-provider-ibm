# Example for VpcV1

This example illustrates how to use the VpcV1

These types of resources are supported:

* DedicatedHost

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## VpcV1 resources

is_dedicated_host resource:

```hcl
resource "is_dedicated_host" "is_dedicated_host_instance" {
  dedicated_host_prototype = var.is_dedicated_host_dedicated_host_prototype
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
| dedicated_host_prototype | The dedicated host prototype object. | `` | true |

## Outputs

| Name | Description |
|------|-------------|
| is_dedicated_host | is_dedicated_host object |
