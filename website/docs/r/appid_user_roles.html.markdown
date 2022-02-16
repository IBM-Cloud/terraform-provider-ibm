---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID User Roles"
description: |-
    Provides AppID User Roles resource.
---

# ibm_appid_user_roles

Create, update, or delete an IBM Cloud AppID Management Services user roles resource. For more information, see [assigning roles to users with the API](https://cloud.ibm.com/docs/appid?topic=appid-access-control&interface=api#assign-roles-api)

## Example usage

```terraform
resource "ibm_appid_user_roles" "roles" {
  tenant_id = var.tenant_id
  subject = ibm_appid_cloud_directory_user.test_user.subject
  role_ids = [ibm_appid_role.test_role.role_id]
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `subject` - (Required, String) The user's identifier ('subject' in identity token)
- `role_ids` - (Required, List of String) The list of AppID role ids that you would like to assign to the user

## Import

The `ibm_appid_user_roles` resource can be imported by using the AppID tenant ID and user subject string.

**Syntax**

```bash
$ terraform import ibm_appid_user_roles.roles <tenant_id>/<subject>
```
**Example**

```bash
$ terraform import ibm_appid_user_roles.roles 5fa344a8-d361-4bc2-9051-58ca253f4b2b/03cd638a-b35a-43f2-a58a-c2d3fe26aaea
```
