---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Role"
description: |-
    Retrieves AppID Role information.
---

# ibm_appid_role
Retrieve information about an IBM Cloud AppID Management Services role. For more information, see [creating roles with API](https://cloud.ibm.com/docs/appid?topic=appid-access-control&interface=api#create-roles-api)

## Example usage

```terraform
data "ibm_appid_role" "role" {
    tenant_id = var.tenant_id
    role_id = var.role_id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID
- `role_id` - (Required, String) The AppID role identifier

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `name` - (String) Role name
- `description` - (String) Role description
- `access` - (Set of Object) A set of access policies that bind specific application scopes to the role

  Nested scheme for `access`:
    - `application_id` - (String) AppID application identifier
    - `scopes` - (List of String) A list of AppID application scopes
