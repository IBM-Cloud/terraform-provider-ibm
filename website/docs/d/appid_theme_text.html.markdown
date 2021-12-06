---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Theme Text"
description: |-
    Retrieves AppID Theme Text configuration.
---

# ibm_appid_theme_text

Retrieve an IBM Cloud AppID Management Services theme text configuration. For more information, see [customizing the login widget](https://cloud.ibm.com/docs/appid?topic=appid-login-widget&interface=api#widget-customize)

## Example usage

```terraform
data "ibm_appid_theme_text" "text" {
    tenant_id = var.tenant_id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `tab_title` - (String) The tab name that will be displayed in the browser
- `footnote` - (String) Footnote
