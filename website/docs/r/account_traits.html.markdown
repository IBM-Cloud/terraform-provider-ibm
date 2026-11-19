---
subcategory: "Account Management"
layout: "ibm"
page_title: "IBM : ibm_account_traits"
description: |-
  Manages account traits such as Financial Services Validated products and user list visibility.
---

# ibm_account_traits

Manages account traits through the Account Management API (`PATCH /v1/accounts/{account_id}/traits`):

- `fs_ready`: use of Financial Services Validated products in the account.
- `team_directory_enabled`: whether users can see the user list. Set `false` to restrict user list visibility.

Only the traits you set are managed. Traits are account settings, so destroying this resource leaves them as they are.


## Example usage

```terraform
resource "ibm_account_traits" "this" {
  fs_ready               = true
  team_directory_enabled = false # restrict user list visibility
}
```

## Argument reference

At least one of `fs_ready` or `team_directory_enabled` must be set.

- `account_id` - (Optional, String) Account ID. Defaults to the account of the provider's credentials.
- `endpoint` - (Optional, String) Account Management API host.
- `fs_ready` - (Optional, Boolean) Enable Financial Services Validated products. Not managed when omitted.
- `team_directory_enabled` - (Optional, Boolean) Let users see the user list. Not managed when omitted.

## Attribute reference

- `id` - Same as `account_id`.

## Import

```
terraform import ibm_account_traits.this <account_id>
```
