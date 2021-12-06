---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID IDP SAML Metadata"
description: |-
    Retrieves AppID SAML Metadata.
---

# ibm_appid_idp_saml_metadata
Retrieve an IBM Cloud AppID Management Services SAML metadata. For more information, see [SAML](https://cloud.ibm.com/docs/appid?topic=appid-enterprise)

## Example usage

```terraform
data "ibm_appid_idp_saml_metadata" "saml" {
    tenant_id = var.tenant_id   
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `metadata` - (String) SAML Metadata
