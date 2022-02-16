---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID APM"
description: |-
    Retrieves AppID information for Advanced Password Management.
---

# ibm_appid_apm
Retrieve information about an IBM Cloud AppID Management Services APM. For more information, see [defining password policies](https://cloud.ibm.com/docs/appid?topic=appid-cd-strength).

~> **WARNING:** This feature is only available for AppID graduated tier plans.

## Example usage

```terraform
data "ibm_appid_apm" "app" {
    tenant_id = var.tenant_id   
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `enabled` (Boolean) `true` if APM is enabled
- `lockout_policy` (List of Object)
    Nested scheme for `lockout_policy`:
    - `enabled` - (Bool) Enable lockout policy
    - `lockout_time_sec` - (Int) Lockout timeout in seconds
    - `num_of_attempts` - (Int) Number of invalid attempts before lockout

- `min_password_change_interval` (List of Object)
    Nested scheme for `min_password_change_interval`:
    - `enabled` - (Bool) Enable minimum password change interval policy
    - `min_hours_to_change_password` (Int) Min amount of hours between password changes

- `password_expiration` (List of Object)
    Nested scheme for `password_expiration`:
    - `enabled` - (Bool) Enable password expiration policy
    - `days_to_expire` - (Int) Days until password expires

- `password_reuse` (List of Object)
    Nested scheme for `password_reuse`:
    - `enabled` - (Bool) Enable password reuse policy
    - `max_password_reuse` (Int) Maximum password reuse

- `prevent_password_with_username` (Boolean) `true` to prevent username in passwords
