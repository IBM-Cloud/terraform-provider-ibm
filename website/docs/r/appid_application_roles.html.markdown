---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Application Roles"
description: |-
    Provides AppID Application Roles resource.
---

# ibm_appid_application_roles

Create, update, or delete an IBM Cloud AppID Management Services application roles resource. For more information, see [controlling access](https://cloud.ibm.com/docs/appid?topic=appid-access-control&interface=api)

## Example usage

```terraform
resource "ibm_appid_application_roles" "roles" {
  tenant_id = var.tenant_id
  client_id = var.client_id // AppID application client_id
  roles = [
    "cf9bb562-8639-46f0-aa8c-0068e4162519", 
    "a330db5f-fa42-4c42-9134-821535728f57"
  ]
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `client_id` - (Required, String) The AppID application identifier
- `roles` - (Required, List of String) A list of AppID role identifiers

## Import

The `ibm_appid_application_roles` resource can be imported by using the AppID tenant ID and application client ID.

**Syntax**

```bash
$ terraform import ibm_appid_application_roles.roles <tenant_id>/<client_id>
```
**Example**

```bash
$ terraform import ibm_appid_application_roles.roles 4be72312-63b7-45fa-9b58-3ae6cd2c90e7/ace469ef-5e1a-4991-8a65-2201b1c5c362
```
