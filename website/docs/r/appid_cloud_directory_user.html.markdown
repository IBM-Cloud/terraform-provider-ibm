---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Cloud Directory User"
description: |-
    Provides AppID Cloud Directory User resource.
---

# ibm_appid_cloud_directory_user

Create, update, or delete an IBM Cloud AppID Management Services Cloud Directory user resource. For more information, see [managing users](https://cloud.ibm.com/docs/appid?topic=appid-cd-users)

Note: depending on your AppID Cloud Directory settings, new user creation may trigger user verification email.

## Example usage

```terraform
resource "ibm_appid_cloud_directory_user" "user" {
  tenant_id = var.tenant_id

  email {
    value = "test_user@mail.com"
    primary = true
  }

  active = false
  locked_until = 1631034316584

  password = "P@ssw0rd"

  display_name = "Test TF User"
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `active` - (Optional, Boolean) Determines if the user account is active or not (Default: true)
- `create_profile` - (Optional, Boolean) A boolean indication if a profile should be created for the Cloud Directory user
- `locked_until` - (Optional, Integer) Epoch time in milliseconds, determines till when the user account will be locked
- `display_name` - (Optional, String) Optional user's display name, defaults to user's email
- `user_name` - (Optional, String) Username
- `password` - (Required, String) Password
- `status` - (Optional, String) `PENDING` or `CONFIRMED` (Default: `PENDING`)
- `email` - (Required, Set of Object) A set of user emails

  Nested scheme for `email`:
  - `value` - (Required, String) An email string
  - `primary` - (Boolean) `true` if this is primary email

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created

- `user_id` - (String) User identifier
- `subject` - (String) The user's identifier ('subject' in identity token)
- `meta` - (List of Object) User metadata

  Nested scheme for `meta`:
  - `created` - (String) User creation date
  - `last_modified` - (String) Last modification date
## Import

The `ibm_appid_cloud_directory_user` resource can be imported by using the AppID tenant ID and user ID.

**Syntax**

```bash
$ terraform import ibm_appid_cloud_directory_user.user <tenant_id>/<user_id>
```
**Example**

```bash
$ terraform import ibm_appid_cloud_directory_user.user 5fa344a8-d361-4bc2-9051-58ca253f4b2b/03dde38a-b35a-43f2-a58a-c2d3fe26aaea
```
