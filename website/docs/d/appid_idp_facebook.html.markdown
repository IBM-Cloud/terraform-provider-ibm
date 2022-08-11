---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Facebook IDP"
description: |-
    Retrieves AppID Facebook IDP information.
---

# ibm_appid_idp_facebook
Retrieve information about an IBM Cloud AppID Facebook IDP. For more information, see [App ID social identity providers](https://cloud.ibm.com/docs/appid?topic=appid-social)

## Example usage

```terraform
data "ibm_appid_idp_facebook" "fb" {
    tenant_id = var.tenant_id   
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `is_active` (String) `true` if Facebook IDP is active
- `redirect_url` - (String) Paste the URI into the Valid OAuth redirect URIs field in the Facebook Login section of the Facebook Developers Portal
- `config` (List of Object, Max: 1) current Facebook IDP configuration if active

  Nested scheme for `config`:
    - `application_id` - (String) Facebook application ID
    - `application_secret` - (String) Facebook application secret
