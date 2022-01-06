---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Languages"
description: |-
    Provides AppID Languages resource.
---

# ibm_appid_languages

Create, update, or delete an IBM Cloud AppID Management Services languages. For more information, see [supported languages](https://cloud.ibm.com/docs/appid?topic=appid-cd-types#cd-languages)

## Example usage

```terraform
resource "ibm_appid_languages" "lang" {
  tenant_id = var.tenant_id
  languages = ["en", "es-ES", "fr-FR"]
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `languages` - (Required, String) The list of languages that can be used to customize email templates for Cloud Directory

## Import

The `ibm_appid_languages` resource can be imported by using the AppID tenant ID.

**Syntax**

```bash
$ terraform import ibm_appid_languages.lang <tenant_id>
```
**Example**

```bash
$ terraform import ibm_appid_languages.lang 5fa344a8-d361-4bc2-9051-58ca253f4b2b
```
