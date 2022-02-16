---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Theme Color"
description: |-
    Provides AppID Theme Color resource.
---

# ibm_appid_theme_color

Create, update, or delete an IBM Cloud AppID Management Services theme color resource. For more information, see [customizing the login widget](https://cloud.ibm.com/docs/appid?topic=appid-login-widget&interface=api#widget-customize)

## Example usage

```terraform
resource "ibm_appid_theme_color" "theme" {
  tenant_id = var.tenant_id
  header_color = "#000000"
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `header_color` - (Required, String) Header color for AppID login screen

## Import

The `ibm_appid_theme_color` resource can be imported by using the AppID tenant ID.

**Syntax**

```bash
$ terraform import ibm_appid_theme_color.theme <tenant_id>
```
**Example**

```bash
$ terraform import ibm_appid_theme_color.theme 5fa344a8-d361-4bc2-9051-58ca253f4b2b
```
