---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Languages"
description: |-
    Retrieves a list of supported AppID languages.
---

# ibm_appid_languages

Retrieve information about an IBM Cloud AppID Management Services languages. For more information, see [supported languages](https://cloud.ibm.com/docs/appid?topic=appid-cd-types#cd-languages)

## Example usage

```terraform
data "ibm_appid_languages" "lang" {
  tenant_id = var.tenant_id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `languages` - (List of String) The list of languages that can be used to customize email templates for Cloud Directory
