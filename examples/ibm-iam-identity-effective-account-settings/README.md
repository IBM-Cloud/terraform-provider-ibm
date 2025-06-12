# Examples for IAM Identity Services

These examples illustrate how to use the resources and data sources associated with IAM Identity Services.

The following data sources are supported:
* ibm_iam_effective_account_settings

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## IAM Identity Services data sources

### Data source: ibm_iam_effective_account_settings

```hcl
data "ibm_iam_effective_account_settings" "iam_effective_account_settings_instance" {
  account_id = var.iam_effective_account_settings_account_id
  include_history = var.iam_effective_account_settings_include_history
  resolve_user_mfa = var.iam_effective_account_settings_resolve_user_mfa
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| account_id | Unique ID of the account. | `string` | true |
| include_history | Defines if the entity history is included in the response. | `bool` | false |
| resolve_user_mfa | Enrich MFA exemptions with user information. | `bool` | false |

#### Outputs

| Name | Description |
|------|-------------|
| context | Context with key properties for problem determination. |
| effective |  |
| account |  |
| assigned_templates | assigned template section. |

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
