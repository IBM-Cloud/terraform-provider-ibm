# Example for IbmToolchainApiV2

This example illustrates how to use the IbmToolchainApiV2

These types of resources are supported:

* toolchain_tool_git

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IbmToolchainApiV2 resources

toolchain_tool_git resource:

```hcl
resource "toolchain_tool_git" "toolchain_tool_git_instance" {
  git_provider = var.toolchain_tool_git_git_provider
  toolchain_id = var.toolchain_tool_git_toolchain_id
  initialization = var.toolchain_tool_git_initialization
  parameters = var.toolchain_tool_git_parameters
  container = var.toolchain_tool_git_container
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
| git_provider |  | `string` | true |
| toolchain_id |  | `string` | true |
| initialization |  | `` | true |
| parameters |  | `` | false |
| container |  | `` | false |

## Outputs

| Name | Description |
|------|-------------|
| toolchain_tool_git | toolchain_tool_git object |
