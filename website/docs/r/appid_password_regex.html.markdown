---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Password Regex"
description: |-
    Provides AppID Password Regex resource.
---

# ibm_appid_password_regex

Update or reset an IBM Cloud AppID Management Services Password Regex configuration. For more information, see [defining password policies](https://cloud.ibm.com/docs/appid?topic=appid-cd-strength)

## Example usage

```terraform
resource "ibm_appid_password_regex" "rgx" {
  tenant_id = var.tenant_id
  regex = "^(?:(?=.*\\d)(?=.*[a-z])(?=.*[A-Z]).*)$"
  error_message = "test error"
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `error_message` (Optional, String) Custom error message
- `regex` (Required, String) The escaped regex expression rule for acceptable password

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created

- `base64_encoded_regex` (String) The regex expression rule for acceptable password encoded in base64

## Import

The `ibm_appid_password_regex` resource can be imported by using the AppID tenant ID.

**Syntax**

```bash
$ terraform import ibm_appid_password_regex.rgx <tenant_id>
```
**Example**

```bash
$ terraform import ibm_appid_password_regex.rgx 5fa344a8-d361-4bc2-9051-58ca253f4b2b
```
