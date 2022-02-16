---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID User Roles"
description: |-
    Retrieves AppID User Role information.
---

# ibm_appid_user_roles
Retrieve information about an IBM Cloud AppID Management Services user roles. For more information, see [assigning roles to users with the API](https://cloud.ibm.com/docs/appid?topic=appid-access-control&interface=api#assign-roles-api)

## Example usage

```terraform
data "ibm_appid_user_roles" "roles" {
  tenant_id = var.tenant_id
  subject = ibm_appid_cloud_directory_user.test_user.subject
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID
- `subject` - (Required, String) The user's identifier ('subject' in identity token)

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `roles` - (Set of Object) A set of AppID user roles

  Nested scheme for `access`:
    - `id` - (String) Role ID
    - `name` - (String) Role name
