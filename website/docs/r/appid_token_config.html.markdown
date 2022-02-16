---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID Token Configuration"
description: |-
    Provides AppID Token Configuration resource.
---

# ibm_appid_token_config

Create, update, or delete an IBM Cloud AppID Management Services token configuration resource. This resource is associated with an IBM Cloud AppID Management Services instance. For more information, about AppID token configuration, see [Customizing AppID tokens](https://cloud.ibm.com/docs/appid?topic=appid-customizing-tokens).

## Example usage

```terraform
resource "ibm_appid_token_config" "tc" {
  tenant_id = var.tenant_id  
  access_token_expires_in = 7200    
  anonymous_access_enabled = true
  anonymous_token_expires_in = 3200    
  refresh_token_enabled = false 
  
  access_token_claim {
    source = "roles"
    destination_claim = "groupIds"
  }

  access_token_claim {
    source = "appid_custom"
    source_claim = "employeeId"
    destination_claim = "employeeId"
  }

  access_token_claim {
    source = "saml"
    source_claim = "attributes.uid"
    destination_claim = "employeeId"
  }
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `tenant_id` - (Required, Forces new resource, String) The AppID instance GUID
- `access_token_claim` - (Optional, Set of Object) A set of objects that are created when claims that are related to access tokens are mapped

  Nested scheme for `access_token_claim`:
    - `destination_claim` - (Optional, String) Defines the custom attribute that can override the current claim in token
    - `source` - (Required, String) Defines the source of the claim. Options include: `saml`, `cloud_directory`, `facebook`, `google`, `appid_custom`,`ibmid`, `roles` and `attributes`
    - `source_claim` - (Optional, String) Defines the claim as provided by the source. It can refer to the identity provider's user information or the user's App ID custom attributes

- `access_token_expires_in` - (Optional, Number) The length of time for which access tokens are valid in seconds
- `anonymous_access_enabled` - (Optional, Bool) Enable anonymous access
- `anonymous_token_expires_in` - (Optional, Number) The length of time for which an anonymous token is valid in seconds
- `id_token_claim` - (Optional, Set of Object) A set of objects that are created when claims that are related to identity tokens are mapped

  Nested scheme for `id_token_claim`:
    - `destination_claim` - (Optional, String) Defines the custom attribute that can override the current claim in token
    - `source` - (Required, String) Defines the source of the claim. Options include: `saml`, `cloud_directory`, `facebook`, `google`, `appid_custom`,`ibmid`, `roles` and `attributes`
    - `source_claim` - (Optional, String) Defines the claim as provided by the source. It can refer to the identity provider's user information or the user's App ID custom attributes

- `refresh_token_enabled` - (Optional, Bool) Enable refresh token
- `refresh_token_expires_in` - (Optional, Number) The length of time for which refresh tokens are valid in seconds

## Import

The `ibm_appid_token_config` resource can be imported by using the AppID tenant ID.

**Syntax**

```bash
$ terraform import ibm_appid_token_config.tc <tenant_id>
```
**Example**

```bash
$ terraform import ibm_appid_token_config.tc 5fa344a8-d361-4bc2-9051-58ca253f4b2b
```
