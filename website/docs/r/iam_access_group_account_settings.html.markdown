---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_access_group_account_settings"
description: |-
  Manages IAM Access Groups account level settings.
---

# ibm_iam_access_group_account_settings

Create, modify, or delete an `iam_access_group_account_settings` resources. Access groups can be used to define a set of permissions that you want to grant to a group of users. For more information, about IAM account settings, refer to [setting up your IBM Cloud](https://cloud.ibm.com/docs/account?topic=account-account-getting-started).

## Example usage

```terraform
resource "ibm_iam_access_group_account_settings" "iam_access_group_account_settings" {
  public_access_enabled = true
}
```


## Argument reference
Review the argument references that you can specify for your resource. 

- `public_access_enabled` - (Optional, Bool) Defines if the public groups are included in the response for access group listing.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `account_id` - (String) Unique ID of an account.
- `public_access_enabled` - (Bool) if the public groups are enabled for access group listing.

