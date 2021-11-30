---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Theme Text"
description: |-
    Provides AppID Theme Text resource.
---

# ibm_appid_theme_text

Create, update, or delete an IBM Cloud AppID Management Services theme text resource. For more information, see [customizing the login widget](https://cloud.ibm.com/docs/appid?topic=appid-login-widget&interface=api#widget-customize)

## Example usage

```terraform
resource "ibm_appid_theme_text" "text" {
  tenant_id = var.tenant_id
  tab_title = "App Login"
  footnote = "Powered by IBM AppID"
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `tab_title` - (Optional, String) The tab name that will be displayed in the browser
- `footnote` - (Optional, String) Footnote

## Import

The `ibm_appid_theme_text` resource can be imported by using the AppID tenant ID.

**Syntax**

```bash
$ terraform import ibm_appid_theme_text.text <tenant_id>
```
**Example**

```bash
$ terraform import ibm_appid_theme_text.text 5fa344a8-d361-4bc2-9051-58ca253f4b2b
```
