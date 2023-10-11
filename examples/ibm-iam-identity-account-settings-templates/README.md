# Example for IamIdentityV1

This example illustrates how to use the IamIdentityV1

The following types of resources are supported:

* account_settings_template

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IamIdentityV1 resources

account_settings_template resource:

```hcl
resource "account_settings_template" "account_settings_template_instance" {
  account_id = var.account_settings_template_account_id
  name = var.account_settings_template_name
  description = var.account_settings_template_description
  account_settings = var.account_settings_template_account_settings
}
```

## IamIdentityV1 data sources

account_settings_template data source:

```hcl
data "account_settings_template" "account_settings_template_instance" {
  template_id = var.account_settings_template_template_id
  version = var.account_settings_template_version
  include_history = var.account_settings_template_include_history
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
| account_id | ID of the account where the template resides. | `string` | false |
| name | The name of the trusted profile template. This is visible only in the enterprise account. | `string` | false |
| description | The description of the trusted profile template. Describe the template for enterprise account users. | `string` | false |
| account_settings |  | `` | false |
| template_id | ID of the account settings template. | `string` | true |
| version | Version of the account settings template. | `string` | true |
| include_history | Defines if the entity history is included in the response. | `bool` | false |

## Outputs

| Name | Description |
|------|-------------|
| account_settings_template | account_settings_template object |
| account_settings_template | account_settings_template object |
