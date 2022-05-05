# Example for ToolchainV2

This example illustrates how to use the ToolchainV2

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


## ToolchainV2 resources

toolchain_tool_sonarqube resource:

```hcl
resource "toolchain_tool_sonarqube" "toolchain_tool_sonarqube_instance" {
  toolchain_id = var.toolchain_tool_sonarqube_toolchain_id
  name = var.toolchain_tool_sonarqube_name
  parameters = var.toolchain_tool_sonarqube_parameters
}
```

## ToolchainV2 Data sources


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
| toolchain_id | ID of the toolchain to bind integration to. | `string` | true |
| name | Name of tool integration. | `string` | false |
| parameters | Parameters to be used to create the integration. | `` | false |

## Outputs

| Name | Description |
|------|-------------|
| toolchain_tool_sonarqube | toolchain_tool_sonarqube object |
