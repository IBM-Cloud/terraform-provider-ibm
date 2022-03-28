# Example for IbmToolchainV2ApiV2

This example illustrates how to use the IbmToolchainV2ApiV2

These types of resources are supported:

* toolchain
* toolchain_integration

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IbmToolchainV2ApiV2 resources

toolchain resource:

```hcl
resource "toolchain" "toolchain_instance" {
  post_toolchain_request = var.toolchain_post_toolchain_request
}
```
toolchain_integration resource:

```hcl
resource "toolchain_integration" "toolchain_integration_instance" {
  toolchain_id = var.toolchain_integration_toolchain_id
  service_id = var.toolchain_integration_service_id
  name = var.toolchain_integration_name
  parameters = var.toolchain_integration_parameters
}
```

## IbmToolchainV2ApiV2 Data sources


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
| post_toolchain_request | Body structure for creating a toolchain. | `` | true |
| toolchain_id | ID of the toolchain to bind integration to. | `string` | true |
| service_id | The unique short name of the service that should be provisioned. | `string` | true |
| name | Name of tool integration. | `string` | false |
| parameters | Arbitrary JSON data. | `map()` | false |

## Outputs

| Name | Description |
|------|-------------|
| toolchain | toolchain object |
| toolchain_integration | toolchain_integration object |
