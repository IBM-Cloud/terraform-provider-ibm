# Example for IBMCloudShellV1

This example illustrates how to use the IBMCloudShellV1

These types of resources are supported:

* cloud_shell_account_settings

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

`terraform destroy` is not required, Cloud Shell account settings cannot be deleted.

## IBMCloudShellV1 Data sources

cloud_shell_account_settings data source:

```hcl
data "cloud_shell_account_settings" "cloud_shell_account_settings" {
  account_id = var.cloud_shell_account_settings_account_id
}
```

## IBMCloudShellV1 resources

cloud_shell_account_settings resource:

```hcl
resource "cloud_shell_account_settings" "cloud_shell_account_settings" {
  account_id = var.cloud_shell_account_settings_account_id
  rev = data.ibm_cloud_shell_account_settings.cloud_shell_account_settings
  default_enable_new_features = var.cloud_shell_account_settings_default_enable_new_features
  default_enable_new_regions = var.cloud_shell_account_settings_default_enable_new_regions
  enabled = var.cloud_shell_account_settings_enabled
  features = var.cloud_shell_account_settings_features
  regions = var.cloud_shell_account_settings_regions
  tags = var.cloud_shell_account_settings_tags
}
```

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
| account_id | The account ID in which the account settings belong to. | `string` | true |
| rev | Unique revision number for the settings object. | `string` | false |
| default_enable_new_features | You can choose which Cloud Shell features are available in the account and whether any new features are enabled as they become available. The feature settings apply only to the enabled Cloud Shell locations. | `bool` | false |
| default_enable_new_regions | Set whether Cloud Shell is enabled in a specific location for the account. The location determines where user and session data are stored. By default, users are routed to the nearest available location. | `bool` | false |
| enabled | When enabled, Cloud Shell is available to all users in the account. | `bool` | false |
| features | List of Cloud Shell features. | `list()` | false |
| regions | List of Cloud Shell region settings. | `list()` | false |

## Outputs

| Name | Description |
|------|-------------|
| cloud_shell_account_settings | cloud_shell_account_settings object |
