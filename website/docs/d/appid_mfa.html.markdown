---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID MFA"
description: |-
    Retrieves AppID MFA activation status.
---

# ibm_appid_mfa
Retrieve an IBM Cloud AppID Management Services MFA activation status. For more information, see [multifactor authentication](https://cloud.ibm.com/docs/appid?topic=appid-cd-mfa)

## Example usage

```terraform
data "ibm_appid_mfa" "mf" {
    tenant_id = var.tenant_id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `is_active` - (Boolean) `true` if MFA is enabled
