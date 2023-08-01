# Example for IamIdentityV1

This example illustrates how to use the IamIdentityV1

The following types of resources are supported:

* account_settings_template_assignment

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IamIdentityV1 resources

account_settings_template_assignment resource:

```hcl
resource "account_settings_template_assignment" "account_settings_template_assignment_instance" {
  template_id = var.account_settings_template_assignment_template_id
  template_version = var.account_settings_template_assignment_template_version
  target_type = var.account_settings_template_assignment_target_type
  target = var.account_settings_template_assignment_target
}
```

## IamIdentityV1 data sources

account_settings_template_assignment data source:

```hcl
data "account_settings_template_assignment" "account_settings_template_assignment_instance" {
  assignment_id = var.account_settings_template_assignment_assignment_id
  include_history = var.account_settings_template_assignment_include_history
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
| template_id | Template Id. | `string` | true |
| template_version | Template version. | `number` | true |
| target_type | Assignment target type. | `string` | true |
| target | Assignment target. | `string` | true |
| assignment_id | ID of the Assignment Record. | `string` | true |
| include_history | Defines if the entity history is included in the response. | `bool` | false |

## Outputs

| Name | Description |
|------|-------------|
| account_settings_template_assignment_instance | account_settings_template_assignment_instance object |
| account_settings_template_assignment_instance | account_settings_template_assignment_instance object |
