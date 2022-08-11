---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Token Configuration"
description: |-
        Retrieves AppID Token Configuration.
---

# ibm_appid_token_config
Retrieve information about an IBM Cloud AppID Management Services token configuration. For more information, refer to [Customizing AppID tokens](https://cloud.ibm.com/docs/appid?topic=appid-customizing-tokens).

## Example usage

```terraform
data "ibm_appid_token_config" "tc" {
    tenant_id = var.tenant_id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `access_token_claim` - (Set of Object) A set of objects that are created when claims that are related to access tokens are mapped

    Nested scheme for `access_token_claim`:
    - `destination_claim` - (String) Defines the custom attribute that can override the current claim in token
    - `source` - (String) Defines the source of the claim. Options include: `saml`, `cloud_directory`, `facebook`, `google`, `appid_custom`, and `attributes`
    - `source_claim` - (String) Defines the claim as provided by the source. It can refer to the identity provider's user information or the user's App ID custom attributes

- `access_token_expires_in` - (Number) The length of time for which access tokens are valid in seconds
- `anonymous_access_enabled` - (Bool) Enable anonymous access
- `anonymous_token_expires_in` - (Number) The length of time for which an anonymous token is valid in seconds
- `id_token_claim` - (Set of Object) A set of objects that are created when claims that are related to identity tokens are mapped

    Nested scheme for `id_token_claim`:
    - `destination_claim` - (String) Defines the custom attribute that can override the current claim in token
    - `source` - (String) Defines the source of the claim. Options include: `saml`, `cloud_directory`, `facebook`, `google`, `appid_custom`, and `attributes`
    - `source_claim` - (String) Defines the claim as provided by the source. It can refer to the identity provider's user information or the user's App ID custom attributes
    
- `refresh_token_enabled` - (Bool) Enable refresh token
- `refresh_token_expires_in` - (Number) The length of time for which refresh tokens are valid in seconds
