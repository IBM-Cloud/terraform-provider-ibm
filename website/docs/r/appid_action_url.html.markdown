---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Action URL"
description: |-
    Provides AppID Action URL resource.
---

# ibm_appid_action_url

Create, update, or delete an IBM Cloud AppID Management Services Action URL.

## Example usage

```terraform
resource "ibm_appid_action_url" "url" {
  tenant_id = var.tenant_id
  action = "on_user_verified"
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `action` - (Required, String) The type of the action: `on_user_verified` - the URL of your custom user verified page, `on_reset_password` - the URL of your custom reset password page
- `url` - (String) The action URL

## Import

The `ibm_appid_action_url` resource can be imported by using the AppID tenant ID and action type string.

**Syntax**

```bash
$ terraform import ibm_appid_action_url.url <tenant_id>/<action_type>
```
**Example**

```bash
$ terraform import ibm_appid_action_url.url 5fa344a8-d361-4bc2-9051-58ca253f4b2b/on_reset_password
```
