---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Application Scopes"
description: |-
        Retrieves AppID Application Scopes.
---

# ibm_appid_application_scopes
Retrieve IBM Cloud AppID Management Services application scopes. For more information, see [controlling access](https://cloud.ibm.com/docs/appid?topic=appid-access-control&interface=api)

## Example usage

```terraform
data "ibm_appid_application_scopes" "scopes" {
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

- `scopes` - (List of String) A `scope` is a runtime action in your application that you register with IBM Cloud App ID to create access permission
