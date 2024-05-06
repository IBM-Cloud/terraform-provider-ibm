---
subcategory: "Enterprise Management"
layout: "ibm"
page_title: "IBM : enterprise_accounts"
description: |-
  Get information about accounts
---

# ibm_enterprise_accounts

Retrieve an information from an `enterprise_accounts` data source. For more information, about enterprise account, refer to [setting up accounts to an enterprise](https://cloud.ibm.com/docs/account?topic=account-enterprise-add).


## Example usage

```terraform
data "ibm_enterprise_accounts" "accounts" {
}
```


## Argument reference
Review the argument reference that you can specify to your data source. 

- `name` - (Optional, String)  The name of an account..

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your data source is created. 

- `accounts` - (List) A list of  Nested `resources` blocks has the following structure.

  Nested scheme for `accounts`:
  - `crn` - (String) The Cloud Resource Name (CRN) of an account.
  - `created_at` - (Timestamp) The time stamp at which an account is created.
  - `created_by` - (String) The IAM ID of the user or service that created an account.
  - `enterprise_account_id` - (String) The enterprise account ID.
  - `enterprise_id` - (String) The enterprise ID that an account is a part of.
  - `enterprise_path` - (String) The path from an enterprise to the particular account.
  - `id` - (String) The account ID.
  - `is_enterprise_account` - (String) The flag to indicate whether the account is an enterprise account or not.
  - `name` - (String) The name of an enterprise.
  - `owner_iam_id` - (String) The IAM ID of the owner of an account.
  - `owner_email` - (String) The Email address of the owner of an account.
  - `parent` - (String) The CRN of the parent of an account.
  - `paid` - (String) The type of account, whether it is `free`, or `paid`.
  - `state` - (String) The state of an account.
  - `updated_at` - (Timestamp) The time stamp at which an account was last updated.
  - `updated_by` - (String) The IAM ID of the user or service that updated an account.
  - `url` - (String) The URL of an account.
  - `iam_apikey` - (String) The IAM API KEY of the account with owner IAM policies, will be used to create resources in enterprise child account.
  - `iam_apikey_id` - (String) The ID of IAM_API_KEY which has owner IAM policies.
  - `iam_service_id` - (String) The IAM Service ID of the account will be used to create IAM_API_KEY with owner IAM policies.
- `id` - (String) The unique identifier of an accounts.
