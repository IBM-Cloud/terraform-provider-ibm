---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID MFA"
description: |-
    Retrieves AppID MFA Channel configuration.
---

# ibm_appid_mfa_channel
Retrieve an IBM Cloud AppID Management Services MFA channel configuration. For more information, see [multifactor authentication](https://cloud.ibm.com/docs/appid?topic=appid-cd-mfa)

## Example usage

```terraform
data "ibm_appid_mfa_channel" "mf" {
    tenant_id = var.tenant_id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `active` - (String) Shows which channel is currently active, possible values: `email`, `sms`
- `sms_config` - (List of Object, Max: 1) SMS channel configuration

  Nested scheme for `sms_config`:
  - `key` - (String) API key
  - `secret` - (String) API secret
  - `from` - (String) Sender's phone number


