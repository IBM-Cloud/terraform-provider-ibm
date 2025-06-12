# Example for EnterpriseManagementV1

This example illustrates how to use the EnterpriseManagementV1

These types of resources are supported:

* enterprise
* enterprise_account_group
* enterprise_account

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## EnterpriseManagementV1 resources

enterprise resource:

```hcl
resource "enterprise" "enterprise_instance" {
  source_account_id = var.enterprise_source_account_id
  name = var.enterprise_name
  primary_contact_iam_id = var.enterprise_primary_contact_iam_id
  domain = var.enterprise_domain
}
```
enterprise_account_group resource:

```hcl
resource "enterprise_account_group" "enterprise_account_group_instance" {
  parent = var.enterprise_account_group_parent
  name = var.enterprise_account_group_name
  primary_contact_iam_id = var.enterprise_account_group_primary_contact_iam_id
}
```
enterprise_account resource:

```hcl
resource "enterprise_account" "enterprise_account_instance" {
  parent = var.enterprise_account_parent
  name = var.enterprise_account_name
  owner_iam_id = var.enterprise_account_owner_iam_id
  traits = var.enterprise_account_traits
  options = var.enterprise_account_options
}
```

## EnterpriseManagementV1 Data sources

enterprises data source:

```hcl
data "enterprises" "enterprises_instance" {
  name = var.enterprises_name
}
```
account_groups data source:

```hcl
data "account_groups" "account_groups_instance" {
  name = var.account_groups_name
}
```
accounts data source:

```hcl
data "accounts" "accounts_instance" {
  name = var.accounts_name
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
| source_account_id | The ID of the account that is used to create the enterprise. | `string` | true |
| name | The name of the enterprise. This field must have 3 - 60 characters. | `string` | true |
| primary_contact_iam_id | The IAM ID of the enterprise primary contact, such as `IBMid-0123ABC`. The IAM ID must already exist. | `string` | true |
| domain | A domain or subdomain for the enterprise, such as `example.com` or `my.example.com`. | `string` | false |
| parent | The CRN of the parent under which the account group will be created. The parent can be an existing account group or the enterprise itself. | `string` | true |
| name | The name of the account group. This field must have 3 - 60 characters. | `string` | true |
| primary_contact_iam_id | The IAM ID of the primary contact for this account group, such as `IBMid-0123ABC`. The IAM ID must already exist. | `string` | true |
| parent | The CRN of the parent under which the account will be created. The parent can be an existing account group or the enterprise itself. | `string` | true |
| name | The name of the account. This field must have 3 - 60 characters. | `string` | true |
| owner_iam_id | The IAM ID of the account owner, such as `IBMid-0123ABC`. The IAM ID must already exist. | `string` | true |
| name | The name of the enterprise. | `string` | false |
| name | The name of the account group. | `string` | false |
| name | The name of the account. | `string` | false |
| traits | The traits object can be used to opt-out of Multi-Factor Authenticatin '`mfa` or for setting enterprise IAM settings `enterprise_iam_managed` setting when creating a child account in the enterprise. | `set` | false |
| options | The options object can be used to set properties on child accounts of an enterprise. You can pass a field to to create IAM service id with IAM api key when creating a child account in the enterprise. | `set` | false |

## Outputs

| Name | Description |
|------|-------------|
| enterprise | enterprise object |
| enterprise_account_group | enterprise_account_group object |
| enterprise_account | enterprise_account object |
| enterprises | enterprises object |
| account_groups | account_groups object |
| accounts | accounts object |
