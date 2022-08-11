---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Role"
description: |-
    Provides AppID Role resource.
---

# ibm_appid_role

Create, update, or delete an IBM Cloud AppID Management Services role resource. For more information, see [creating roles with API](https://cloud.ibm.com/docs/appid?topic=appid-access-control&interface=api#create-roles-api)

## Example usage

```terraform
resource "ibm_appid_role" "role" {
  tenant_id = var.tenant_id
  name = "example role"

  access {
    application_id = var.client_id
    scopes = ["scope1", "scope2"]
  }
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `name` - (Required, String) The AppID role name
- `access` - (Optional, Set of Object) A set of access policies that bind specific application scopes to the role

  Nested scheme for `access`:
    - `application_id` - (Required, String) AppID application identifier
    - `scopes` - (Required, List of String) A list of AppID application scopes

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created

- `role_id` - (String) The AppID role identifier

## Import

The `ibm_appid_role` resource can be imported by using the AppID tenant ID and role ID.

**Syntax**

```bash
$ terraform import ibm_appid_role.role <tenant_id>/<role_id>
```
**Example**

```bash
$ terraform import ibm_appid_role.role 5fa344a8-d361-4bc2-9051-58ca253f4b2b/03cd638a-b35a-43f2-a58a-c2d3fe26aaea
```
