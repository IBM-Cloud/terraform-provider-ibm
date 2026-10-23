---
subcategory: "Account Management"
layout: "ibm"
page_title: "IBM : ibm_account_traits"
description: |-
  Reads account traits.
---

# ibm_account_traits (Data Source)

Reads account traits through the Account Management API (`GET /v1/accounts/{account_id}/traits`).

## Example usage

```terraform
data "ibm_account_traits" "current" {}
```

## Argument reference

- `account_id` - (Optional, String) Account ID. Defaults to the account of the provider's credentials.
- `endpoint` - (Optional, String) Account Management API host.

## Attribute reference

- `fs_ready` - Whether Financial Services Validated products are enabled.
- `team_directory_enabled` - Whether users can see the user list.
