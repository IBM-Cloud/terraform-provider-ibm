# Example for IamIdentityV1

This example illustrates how to use the IamIdentityV1

The following types of resources are supported:

* trusted_profile_template

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IamIdentityV1 resources

trusted_profile_template resource:

```hcl
resource "trusted_profile_template" "trusted_profile_template_instance" {
  account_id = var.trusted_profile_template_account_id
  name = var.trusted_profile_template_name
  description = var.trusted_profile_template_description
  profile = var.trusted_profile_template_profile
  policy_template_references = var.trusted_profile_template_policy_template_references
}
```

## IamIdentityV1 data sources

trusted_profile_template data source:

```hcl
data "trusted_profile_template" "trusted_profile_template_instance" {
  template_id = var.trusted_profile_template_template_id
  version = var.trusted_profile_template_version
  include_history = var.trusted_profile_template_include_history
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
| profile | Input body parameters for the TemplateProfileComponent. | `` | false |
| policy_template_references | Existing policy templates that you can reference to assign access in the trusted profile component. | `list()` | false |
| template_id | ID of the trusted profile template. | `string` | true |
| version | Version of the Profile Template. | `string` | true |
| include_history | Defines if the entity history is included in the response. | `bool` | false |

## Outputs

| Name | Description |
|------|-------------|
| trusted_profile_template | trusted_profile_template object |
| trusted_profile_template | trusted_profile_template object |
