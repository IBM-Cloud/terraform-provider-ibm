---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Cloud Directory Redirect URLs"
description: |-
        Retrieves AppID Cloud Directory Redirect URLs.
---

# ibm_appid_redirect_urls
Retrieve IBM Cloud AppID Management Services Cloud Directory redirect URLs. For more information, see [adding redirect URIs](https://cloud.ibm.com/docs/appid?topic=appid-managing-idp#add-redirect-uri)

## Example usage

```terraform
data "ibm_appid_redirect_urls" "urls" {
    tenant_id = var.tenant_id   
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `urls` - (List of String) A list of redirect URLs
