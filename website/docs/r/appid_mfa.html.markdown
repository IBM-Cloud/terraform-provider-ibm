---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID MFA"
description: |-
    Provides AppID MFA resource.
---

# ibm_appid_mfa

Create, update, or delete an IBM Cloud AppID Management Services MFA resource. For more information, see [multifactor authentication](https://cloud.ibm.com/docs/appid?topic=appid-cd-mfa)

## Example usage

```terraform
resource "ibm_appid_mfa" "mf" {
  tenant_id = var.tenant_id
  is_active = true
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `is_active` - (Boolean) `true` if MFA should be enabled

## Import

The `ibm_appid_mfa` resource can be imported by using the AppID tenant ID.

**Syntax**

```bash
$ terraform import ibm_appid_mfa.mf <tenant_id>
```
**Example**

```bash
$ terraform import ibm_appid_mfa.mf 5fa344a8-d361-4bc2-9051-58ca253f4b2b
```
