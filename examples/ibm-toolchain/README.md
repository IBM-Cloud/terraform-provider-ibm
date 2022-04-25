# Example for ToolchainV2

This example illustrates how to use the ToolchainV2

These types of resources are supported:

* toolchain_tool_security_compliance

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## ToolchainV2 resources

toolchain_tool_security_compliance resource:

```hcl
resource "toolchain_tool_security_compliance" "toolchain_tool_security_compliance_instance" {
  toolchain_id = var.toolchain_tool_security_compliance_toolchain_id
  name = var.toolchain_tool_security_compliance_name
  parameters = var.toolchain_tool_security_compliance_parameters
  parameters_references = var.toolchain_tool_security_compliance_parameters_references
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
| parameters | Arbitrary JSON data. | `` | false |
| parameters_references | Decoded values used on provision in the broker that reference fields in the parameters. | `map()` | false |

## Outputs

| Name | Description |
|------|-------------|
| toolchain_tool_security_compliance | toolchain_tool_security_compliance object |
