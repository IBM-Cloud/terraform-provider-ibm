---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Google IDP"
description: |-
    Provides AppID Google IDP resource.
---

# ibm_appid_idp_google

Update or reset an IBM Cloud AppID Management Services Google IDP configuration. For more information, see [App ID social identity providers](https://cloud.ibm.com/docs/appid?topic=appid-social)

## Example usage

```terraform
resource "ibm_appid_idp_google" "gg" {
  tenant_id = var.tenant_id
  is_active = true
  config {
    application_id      = "test_id"
    application_secret 	= "test_secret"
  }
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `is_active` (Required, Boolean) Google IDP activation
- `config` (Optional, List of Object, Max: 1) Google IDP configuration

  Nested scheme for `config`:
    - `application_id` - (Required, String) Google application ID
    - `application_secret` - (Required, String) Google application secret

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created

- `redirect_url` - (String) Paste the URI into the Authorized redirect URIs field in the Google Developer Console

## Import

The `ibm_appid_idp_google` resource can be imported by using the AppID tenant ID.

**Syntax**

```bash
$ terraform import ibm_appid_idp_google.gg <tenant_id>
```
**Example**

```bash
$ terraform import ibm_appid_idp_google.gg 5fa344a8-d361-4bc2-9051-58ca253f4b2b
```
