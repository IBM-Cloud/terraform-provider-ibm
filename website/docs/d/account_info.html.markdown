---
layout: "ibm"
page_title: "IBM : ibm_account_info"
description: |-
  Get information about an IBM Cloud account
subcategory: "Account Management"
---

# ibm_account_info

Provides a read-only data source to retrieve information about an IBM Cloud account. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

For more information about IBM Cloud accounts, see [Managing accounts](https://cloud.ibm.com/docs/account).

## Example Usage

```hcl
data "ibm_account_info" "account" {
  account_id = "your-account-id-here"
}

output "account_name" {
  value = data.ibm_account_info.account.name
}

output "account_owner" {
  value = data.ibm_account_info.account.owner
}

output "account_type" {
  value = data.ibm_account_info.account.type
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required, String) The unique identifier of the account you want to retrieve.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created:

* `id` - (String) The unique identifier of the account.
* `name` - (String) The name of the account.
* `owner` - (String) The owner of the account.
* `owner_userid` - (String) The user ID of the account owner.
* `owner_iamid` - (String) The IAM ID of the account owner.
* `type` - (String) The type of the account.
* `status` - (String) The status of the account.
* `linked_softlayer_account` - (String) The linked SoftLayer account ID, if applicable.
* `team_directory_enabled` - (Boolean) Indicates whether the team directory is enabled for the account.
* `traits` - (List) Account traits and characteristics.
  
  Nested schema for `traits`:
  * `eu_supported` - (Boolean) Indicates if the account supports EU data residency.
  * `poc` - (Boolean) Indicates if this is a proof of concept account.
  * `hippa` - (Boolean) Indicates if the account is HIPAA compliant.

## Import

The `ibm_account_info` data source does not support import as it is a read-only data source.