---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Cloud Directory IDP"
description: |-
    Provides AppID Cloud Directory IDP resource.
---

# ibm_appid_idp_cloud_directory

Update or reset an IBM Cloud AppID Management Services Cloud Directory IDP configuration. For more information, see [configuring Cloud Directory](https://cloud.ibm.com/docs/appid?topic=appid-cloud-directory)

## Example usage

```terraform
resource "ibm_appid_idp_cloud_directory" "cd" {
  tenant_id = var.tenant_id
  is_active = true
  identity_confirm_methods = [
    "email"
  ]
  identity_field = "email"
  self_service_enabled = false
  signup_enabled = false
  welcome_enabled = true
  reset_password_enabled = false
  reset_password_notification_enabled = false
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `is_active` (Required, Boolean) Cloud Directory IDP activation
- `identity_confirm_access_mode` (Optional, String) Allowed values: `FULL`, `RESTRICTIVE`, `OFF`
- `identity_confirm_methods` (Optional, List of String) Allowed value: `email`
- `identity_field` (Optional, String) Allowed values: `email`, `userName`
- `reset_password_enabled` (Optional, Boolean) Enable password resets
- `reset_password_notification_enabled` (Optional, Boolean) Enable password reset notification emails
- `self_service_enabled` (Optional, Boolean) Let users change their password, edit user details
- `signup_enabled` (Optional, Boolean) Allow users to sign-up
- `welcome_enabled` (Optional, Boolean) Send welcome email to new users


## Import

The `ibm_appid_idp_cloud_directory` resource can be imported by using the AppID tenant ID.

**Syntax**

```bash
$ terraform import ibm_appid_idp_cloud_directory.cd <tenant_id>
```
**Example**

```bash
$ terraform import ibm_appid_idp_cloud_directory.cd 5fa344a8-d361-4bc2-9051-58ca253f4b2b
```
