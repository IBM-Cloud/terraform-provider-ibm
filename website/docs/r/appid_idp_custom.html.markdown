---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Custom IDP"
description: |-
    Provides AppID Custom IDP resource.
---

# ibm_appid_idp_custom

Create, update, or delete an IBM Cloud AppID Management Services Custom IDP resource. For more information, see [AppID custom identity](https://cloud.ibm.com/docs/appid?topic=appid-custom-identity)

## Example usage

```terraform
resource "ibm_appid_idp_custom" "idp" {
  tenant_id = var.tenant_id
  is_active = true
  public_key = file("path/to/public_key")
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `is_active` - (Boolean) `true` if custom IDP integration should be enabled
- `public_key` - (String) The public key used to validate signed JWT

## Import

The `ibm_appid_idp_custom` resource can be imported by using the AppID tenant ID.

**Syntax**

```bash
$ terraform import ibm_appid_idp_custom.idp <tenant_id>
```
**Example**

```bash
$ terraform import ibm_appid_idp_custom.idp 5fa344a8-d361-4bc2-9051-58ca253f4b2b
```
