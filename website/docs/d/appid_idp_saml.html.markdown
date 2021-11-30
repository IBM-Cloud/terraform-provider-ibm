---
subcategory: "App ID Management"
layout: "ibm"
page_title: "IBM: AppID SAML IDP"
description: |-
    Retrieves AppID SAML IDP information.
---

# ibm_appid_idp_saml
Retrieve information about an IBM Cloud AppID SAML IDP. For more information, see [SAML](https://cloud.ibm.com/docs/appid?topic=appid-enterprise)

## Example usage

```terraform
data "ibm_appid_idp_saml" "saml" {
    tenant_id = var.tenant_id   
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `tenant_id` - (Required, String) The AppID instance GUID

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created

- `is_active` (String) `true` if SAML IDP is active
- `config` (List of Object, Max: 1) current SAML IDP configuration if active

    Nested scheme for `config`:
    - `entity_id` - (String) Unique name for an Identity Provider
    - `sign_in_url` - (String) SAML SSO url
    - `certificates` - (List of String) List of certificates, primary and optional secondary
    - `display_name` - (String) Optional provider name
    - `encrypt_response` - (Bool) `true` if SAML responses should be encrypted
    - `sign_request` - (Bool) `true` if SAML requests should be signed
    - `include_scoping` - (Bool) `true` if scopes are included
    - `authn_context` - (List of Object, Max: 1) SAML authNContext configuration

      Nested scheme for `authn_context`:
      `class` - (List of String) List of `authnContext` classes
      `comparison` - (String) Example values: `exact`, `maximum`, `minimum`, `better`
