---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID APM"
description: |-
  Provides AppID Advanced Password Management resource.
---

# ibm_appid_apm

Update or reset an IBM Cloud AppID Management Services APM configuration. For more information, see [defining password policies](https://cloud.ibm.com/docs/appid?topic=appid-cd-strength).

~> **WARNING:** This feature is only available for AppID graduated tier plans.

## Example usage

```terraform
resource "ibm_appid_apm" "apm" {
  tenant_id = var.tenant_id
  enabled = true
  prevent_password_with_username = true

  password_reuse {
    enabled = true
    max_password_reuse = 4
  }

  password_expiration {
    enabled = true
    days_to_expire = 25
  }

  lockout_policy {
    enabled = true
    lockout_time_sec = 2600
    num_of_attempts = 4
  }

  min_password_change_interval {
    enabled = true
    min_hours_to_change_password = 1
  }
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `enabled` (Required, Boolean) `true` if APM is enabled
- `lockout_policy` (Required, List of Object, Max: 1)
  Nested scheme for `lockout_policy`:
    - `enabled` - (Required, Bool) Enable lockout policy
    - `lockout_time_sec` - (Optional, Int) Lockout timeout in seconds (Default: 1800)
    - `num_of_attempts` - (Optional, Int) Number of invalid attempts before lockout (Default: 3)

- `min_password_change_interval` (Required, List of Object, Max: 1)
  Nested scheme for `min_password_change_interval`:
    - `enabled` - (Required, Bool) Enable minimum password change interval policy
    - `min_hours_to_change_password` (Optional, Int) Min amount of hours between password changes

- `password_expiration` (Required, List of Object, Max: 1)
  Nested scheme for `password_expiration`:
    - `enabled` - (Required, Bool) Enable password expiration policy
    - `days_to_expire` - (Optional, Int) Days until password expires (Default: 30)

- `password_reuse` (Required, List of Object, Max: 1)
  Nested scheme for `password_reuse`:
    - `enabled` - (Required, Bool) Enable password reuse policy
    - `max_password_reuse` (Optional, Int) Maximum password reuse (Default: 8)

- `prevent_password_with_username` (Optional, Boolean) `true` to prevent username in passwords (Default: false)


## Import

The `ibm_appid_apm` resource can be imported by using the AppID tenant ID.

**Syntax**

```bash
$ terraform import ibm_appid_apm.apm <tenant_id>
```
**Example**

```bash
$ terraform import ibm_appid_apm.apm 5fa344a8-d361-4bc2-9051-58ca253f4b2b
```
