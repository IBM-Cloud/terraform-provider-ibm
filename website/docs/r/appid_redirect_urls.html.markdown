---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Cloud Directory Redirect URLs"
description: |-
    Provides AppID Cloud Directory Redirect URLs resource.
---

# ibm_appid_redirect_urls
Create, update, or delete an IBM Cloud AppID Management Services Cloud Directory redirect URLs. For more information, see [adding redirect URIs](https://cloud.ibm.com/docs/appid?topic=appid-managing-idp#add-redirect-uri)

## Example usage

```terraform
resource "ibm_appid_redirect_urls" "urls" {
    tenant_id = var.tenant_id 
    urls = [
      "https://test-application-1.com/login",
      "https://test-application-2.com/login",
      "https://test-application-3.com/login"
    ]
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `urls` - (Required, List of String) A list of redirect URLs

## Import

The `ibm_appid_redirect_urls` resource can be imported by using the AppID tenant ID.

**Syntax**

```bash
$ terraform import ibm_appid_redirect_urls.urls <tenant_id>
```
**Example**

```bash
$ terraform import ibm_appid_redirect_urls.urls 5fa344a8-d361-4bc2-9051-58ca253f4b2b
```
