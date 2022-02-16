---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Cloud Directory User"
description: |-
    Retrieves AppID Cloud Directory user information.
---

# ibm_appid_cloud_directory_user
Retrieve information about an IBM Cloud AppID Management Services Cloud Directory User. For more information, see [managing users](https://cloud.ibm.com/docs/appid?topic=appid-cd-users)

## Example usage

```terraform
data "ibm_appid_cloud_directory_user" "user" {
    tenant_id = var.tenant_id
    user_id = var.user_id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID
- `user_id` - (Required, String) The AppID Cloud Directory user ID

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `active` - (Boolean) Determines if the user account is active or not
- `locked_until` - (Integer) Epoch time in milliseconds, determines till when the user account will be locked
- `display_name` - (String) Optional user's display name
- `subject` - (String) The user's identifier ('subject' in identity token)
- `user_name` - (String) Username
- `status` - (String) `PENDING` or `CONFIRMED`
- `email` - (Set of Object) A set of user emails

  Nested scheme for `email`:
    - `value` - (String) An email string
    - `primary` - (Boolean) `true` if this is primary email

- `meta` - (List of Object) User metadata

  Nested scheme for `meta`:
    - `created` - (String) User creation date
    - `last_modified` - (String) Last modification date
