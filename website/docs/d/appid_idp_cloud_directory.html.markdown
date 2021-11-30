---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Cloud Directory IDP"
description: |-
    Retrieves AppID Cloud Directory IDP information.
---

# ibm_appid_idp_cloud_directory
Retrieve information about an IBM Cloud AppID Cloud Directory IDP. For more information, see [configuring Cloud Directory](https://cloud.ibm.com/docs/appid?topic=appid-cloud-directory)

## Example usage

```terraform
data "ibm_appid_idp_cloud_directory" "cd" {
    tenant_id = var.tenant_id   
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `identity_confirm_access_mode` (String) Example values: `FULL`, `RESTRICTIVE`, `OFF`
- `identity_confirm_methods` (List of String) Example: `email`
- `identity_field` (String) Example values: `email`, `userName`
- `is_active` (Boolean) Cloud Directory IDP activation
- `reset_password_enabled` (Boolean) Enable password resets
- `reset_password_notification_enabled` (Boolean) Enable password reset notification emails
- `self_service_enabled` (Boolean) Let users change their password, edit user details
- `signup_enabled` (Boolean) Allow users to sign-up
- `welcome_enabled` (Boolean) Send welcome email to new users
