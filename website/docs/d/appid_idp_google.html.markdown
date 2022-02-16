---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Google IDP"
description: |-
  Retrieves AppID Google IDP information.
---

# ibm_appid_idp_google
Retrieve information about an IBM Cloud AppID Google IDP. For more information, see [App ID social identity providers](https://cloud.ibm.com/docs/appid?topic=appid-social)

## Example usage

```terraform
data "ibm_appid_idp_google" "gg" {
    tenant_id = var.tenant_id   
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `is_active` (String) `true` if Google IDP is active
- `redirect_url` - (String) Paste the URI into the Authorized redirect URIs field in the Google Developer Console
- `config` (List of Object, Max: 1) current Google IDP configuration if active

  Nested scheme for `config`:
    - `application_id` - (String) Google application ID
    - `application_secret` - (String) Google application secret
