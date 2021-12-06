---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Custom IDP"
description: |-
    Retrieves AppID Custom IDP information.
---

# ibm_appid_idp_custom
Retrieve information about an IBM Cloud AppID Management Services Custom IDP. For more information, see [AppID custom identity](https://cloud.ibm.com/docs/appid?topic=appid-custom-identity)

## Example usage

```terraform
data "ibm_appid_idp_custom" "idp" {
    tenant_id = var.tenant_id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `is_active` - (Boolean) `true` if custom IDP integration is enabled
- `public_key` - (String) The public key used to validate signed JWT
