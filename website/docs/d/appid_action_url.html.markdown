---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Action URL"
description: |-
    Retrieves AppID Action URL.
---

# ibm_appid_action_url
Retrieve an IBM Cloud AppID Management Services action URL - the custom url to redirect to when Cloud Directory action is executed

## Example usage

```terraform
data "ibm_appid_action_url" "url" {
    tenant_id = var.tenant_id
    action = "on_user_verified" 
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID
- `action` - (Required, String) The type of the action: `on_user_verified` - the URL of your custom user verified page, `on_reset_password` - the URL of your custom reset password page

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `url` - (String) The action URL
