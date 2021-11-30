---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID MFA"
description: |-
    Provides AppID MFA Channel resource.
---

# ibm_appid_mfa_channel

Create, update, or delete an IBM Cloud AppID Management Services MFA Channel resource. For more information, see [multifactor authentication](https://cloud.ibm.com/docs/appid?topic=appid-cd-mfa)

## Example usage

```terraform
resource "ibm_appid_mfa_channel" "mf" {
  tenant_id = var.tenant_id
  active = "sms"

  sms_config {
    key = "<nexmo key>"
    secret = "<nexmo secret>"
    from = "+11112223333"
  }
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, String) The AppID instance GUID
- `active` - (Required, String) Determines which channel is currently active, allowed values: `email`, `sms`. **Note**: in addition, AppID MFA should be enabled, see `ibm_appid_mfa` resource
- `sms_config` - (Optional, List of Object, Max: 1) SMS channel configuration. After signing up for a [Vonage](https://dashboard.nexmo.com/sign-up) account, you can get your API key and secret on the dashboard.

  Nested scheme for `sms_config`:
    - `key` - (Required, String) API key
    - `secret` - (Required, String) API secret
    - `from` - (Required, String) Sender's phone number

## Import

The `ibm_appid_mfa_channel` resource can be imported by using the AppID tenant ID.

**Syntax**

```bash
$ terraform import ibm_appid_mfa_channel.mf <tenant_id>
```
**Example**

```bash
$ terraform import ibm_appid_mfa_channel.mf 5fa344a8-d361-4bc2-9051-58ca253f4b2b
```
