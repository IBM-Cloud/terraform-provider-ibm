---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Application Scopes"
description: |-
    Provides AppID Application Scopes resource.
---

# ibm_appid_application_scopes

Create, update, or delete an IBM Cloud AppID Management Services application scopes resource. For more information, see [controlling access](https://cloud.ibm.com/docs/appid?topic=appid-access-control&interface=api)

## Example usage

```terraform
resource "ibm_appid_application_scopes" "scopes" {
  tenant_id = var.tenant_id
  client_id = var.client_id // AppID application client_id
  scopes = ["scope_1", "scope_2", "scope_3"]
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `client_id` - (Required, String) The AppID application identifier
- `scopes` - (Required, List of String) A `scope` is a runtime action in your application that you register with IBM Cloud App ID to create access permission

## Import

The `ibm_appid_application_scopes` resource can be imported by using the AppID tenant ID and application client ID.

**Syntax**

```bash
$ terraform import ibm_appid_application_scopes.scopes <tenant_id>/<client_id>
```
**Example**

```bash
$ terraform import ibm_appid_application_scopes.scopes 4be72312-63b7-45fa-9b58-3ae6cd2c90e7/ace469ef-5e1a-4991-8a65-2201b1c5c362
```
