---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Application Roles"
description: |-
    Retrieves AppID Application Roles.
---

# ibm_appid_application_roles
Retrieve IBM Cloud AppID Management Services application roles. For more information, see [controlling access](https://cloud.ibm.com/docs/appid?topic=appid-access-control&interface=api)

## Example usage

```terraform
data "ibm_appid_application_roles" "roles" {
    tenant_id = var.tenant_id
    client_id = var.client_id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID
- `client_id` - (Required, String) The AppID application identifier

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `roles` - (Set of Object) A set of roles that are assigned to the application

  Nested scheme for `roles`:
    - `id` - (String) AppID role ID
    - `name` - (String) AppID role name
