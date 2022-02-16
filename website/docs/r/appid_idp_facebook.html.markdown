---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Facebook IDP"
description: |-
    Provides AppID Facebook IDP resource.
---

# ibm_appid_idp_facebook

Update or reset an IBM Cloud AppID Management Services Facebook IDP configuration. For more information, see [App ID social identity providers](https://cloud.ibm.com/docs/appid?topic=appid-social)

## Example usage

```terraform
resource "ibm_appid_idp_facebook" "fb" {
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
- `is_active` (Required, Boolean) Facebook IDP activation
- `config` (Optional, List of Object, Max: 1) Facebook IDP configuration

  Nested scheme for `config`:
    - `application_id` - (Required, String) Facebook application ID
    - `application_secret` - (Required, String) Facebook application secret

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created

- `redirect_url` - (String) Paste the URI into the Valid OAuth redirect URIs field in the Facebook Login section of the Facebook Developers Portal

## Import

The `ibm_appid_idp_facebook` resource can be imported by using the AppID tenant ID.

**Syntax**

```bash
$ terraform import ibm_appid_idp_facebook.fb <tenant_id>
```
**Example**

```bash
$ terraform import ibm_appid_idp_facebook.fb 5fa344a8-d361-4bc2-9051-58ca253f4b2b
```
