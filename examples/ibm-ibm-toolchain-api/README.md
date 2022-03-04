# Example for IbmToolchainApiV2

This example illustrates how to use the IbmToolchainApiV2

These types of resources are supported:

* toolchain_tool_sonarqube

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IbmToolchainApiV2 resources

toolchain_tool_sonarqube resource:

```hcl
resource "toolchain_tool_sonarqube" "toolchain_tool_sonarqube_instance" {
  toolchain_id = var.toolchain_tool_sonarqube_toolchain_id
  parameters = var.toolchain_tool_sonarqube_parameters
  parameters_references = var.toolchain_tool_sonarqube_parameters_references
  container = var.toolchain_tool_sonarqube_container
}
```

## IbmToolchainApiV2 Data sources


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
| toolchain_id |  | `string` | true |
| parameters |  | `` | false |
| parameters_references | Decoded values used on provision in the broker that reference fields in the parameters. | `map()` | false |
| container |  | `` | false |

## Outputs

| Name | Description |
|------|-------------|
| toolchain_tool_sonarqube | toolchain_tool_sonarqube object |
