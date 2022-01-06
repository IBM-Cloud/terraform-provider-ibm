---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Roles"
description: |-
    Retrieves a list of AppID Roles.
---

# ibm_appid_roles
Retrieve information about an IBM Cloud AppID Management Services roles. For more information, see [creating roles with API](https://cloud.ibm.com/docs/appid?topic=appid-access-control&interface=api#create-roles-api)

## Example usage

```terraform
data "ibm_appid_roles" "roles" {
    tenant_id = var.tenant_id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `roles` - (List of Object) A list of AppID roles
    
  Nested scheme for `roles`:

  - `name` - (String) Role name
  - `description` - (String) Role description
  - `access` - (Set of Object) A set of access policies that bind specific application scopes to the role

    Nested scheme for `access`:
      - `application_id` - (String) AppID application identifier
      - `scopes` - (List of String) A list of AppID application scopes
