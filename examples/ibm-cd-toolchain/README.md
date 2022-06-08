# Example for CdToolchainV2

This example illustrates how to use the CdToolchainV2

These types of resources are supported:

* cd_toolchain_tool_sonarqube

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## CdToolchainV2 resources

cd_toolchain_tool_sonarqube resource:

```hcl
resource "cd_toolchain_tool_sonarqube" "cd_toolchain_tool_sonarqube_instance" {
  toolchain_id = var.cd_toolchain_tool_sonarqube_toolchain_id
  name = var.cd_toolchain_tool_sonarqube_name
  parameters = var.cd_toolchain_tool_sonarqube_parameters
}
```

## CdToolchainV2 Data sources

cd_toolchain_tool_sonarqube data source:

```hcl
data "cd_toolchain_tool_sonarqube" "cd_toolchain_tool_sonarqube_instance" {
  toolchain_id = var.cd_toolchain_tool_sonarqube_toolchain_id
  tool_id = var.cd_toolchain_tool_sonarqube_tool_id
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
| toolchain_id | ID of the toolchain to bind tool to. | `string` | true |
| name | Name of tool. | `string` | false |
| parameters | Parameters to be used to create the tool. | `` | false |
| toolchain_id | ID of the toolchain. | `string` | true |
| tool_id | ID of the tool bound to the toolchain. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| cd_toolchain_tool_sonarqube | cd_toolchain_tool_sonarqube object |
| cd_toolchain_tool_sonarqube | cd_toolchain_tool_sonarqube object |
