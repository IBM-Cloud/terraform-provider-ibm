---
layout: "ibm"
page_title: "IBM : ibm_iam_idp"
description: |-
  Manages an IAM Identity Provider (IdP).
subcategory: "IAM Identity Services"
---

# ibm_iam_idp

Create, update, and delete an IAM Identity Provider (IdP) with this resource. For more information, see the [IAM Identity Services API documentation](https://cloud.ibm.com/apidocs/iam-identity-token-api).

## Example Usage

### Create a SAML Identity Provider

```hcl
resource "ibm_iam_idp" "saml_idp" {
  account_id = var.account_id
  name       = "my-saml-idp"
  type       = "saml"
  active     = true

  properties {
    idp {
      entity_id           = "https://idp.example.com/saml/metadata"
      redirect_binding_url = "https://idp.example.com/saml/sso"
      want_request_signed  = true
      logout_url           = "https://idp.example.com/saml/logout"
    }
    sp {
      want_assertion_signed          = true
      want_response_signed           = true
      encrypt_response               = false
      idp_initiated_login_enabled    = false
      logout_url_enabled_when_available = true
    }
  }

  # secrets block is required for SAML; leave empty to auto-generate SP certificates
  secrets {}
}
```

### Create an IdP and share it with another account

```hcl
resource "ibm_iam_idp" "shared_idp" {
  account_id = var.account_id
  name       = "shared-saml-idp"
  type       = "saml"
  active     = true

  properties {
    idp {
      entity_id           = "https://idp.example.com/saml/metadata"
      redirect_binding_url = "https://idp.example.com/saml/sso"
    }
  }

  secrets {}

  share_scope {
    id   = var.consumer_account_id
    type = "account"
  }
}
```

### Create an IdP and share with an entire enterprise

```hcl
resource "ibm_iam_idp" "enterprise_idp" {
  account_id = var.account_id
  name       = "enterprise-saml-idp"
  type       = "saml"

  properties {
    idp {
      entity_id           = "https://idp.example.com/saml/metadata"
      redirect_binding_url = "https://idp.example.com/saml/sso"
    }
  }

  secrets {}

  share_scope {
    id   = var.enterprise_id
    type = "enterprise"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `account_id` - (Required, Forces new resource, String) Account where the IdP resides. Changing this value creates a new resource.
* `name` - (Required, String) Speaking name of the Identity Provider.
* `type` - (Required, Forces new resource, String) Type of the IDP. Changing this value creates a new resource.
  * Constraints: Allowed values are `saml`, `appid`, `ldap`.
* `active` - (Optional, Boolean) Defines if the IDP is active (enabled) for all accounts, including those that have consumed the IdP. Default during creation is `true`.
* `share_scope` - (Optional, List) List of targets that can consume the IdP. Each entry specifies an account or enterprise that is allowed to bind this IdP.
Nested schema for **share_scope**:
  * `id` - (Optional, String) ID of the account or enterprise to share with.
  * `type` - (Optional, String) Type of the share target. Allowed values are `account`, `enterprise`.
* `secrets` - (Optional, Sensitive, List) Secrets of the IDP stored encrypted. Required for SAML type — use an empty `secrets {}` block to auto-generate SP certificates. Maximum of one block.
Nested schema for **secrets**:
  * `idp` - (Optional, List) Identity Provider secrets. Maximum of one block.
  Nested schema for **idp**:
    * `xml_import` - (Optional, Boolean) Flag indicating if secrets should be imported from a `metadata.xml` file.
  * `sp` - (Optional, List) Service Provider secrets. Leave empty to have IBM Cloud auto-generate SP signing certificates. Maximum of one block.
* `properties` - (Optional, List) Properties of the IDP stored in plain text. Required for SAML type. Maximum of one block.
Nested schema for **properties**:
  * `idp` - (Optional, List) Identity Provider (SAML IDP) configuration. Maximum of one block.
  Nested schema for **idp**:
    * `xml_import` - (Optional, Boolean) Flag indicating if IdP should be imported from a `metadata.xml` file. When `true`, `entity_id` and `redirect_binding_url` are derived from the imported XML.
    * `entity_id` - (Optional, String) SAML IDP entity ID. Required when `xml_import` is `false`.
    * `redirect_binding_url` - (Optional, String) SAML redirect binding URL (SSO endpoint). Required when `xml_import` is `false`.
    * `want_request_signed` - (Optional, Boolean) Indicates if the IDP requires authentication requests to be signed.
    * `logout_url` - (Optional, String) SAML IDP single logout URL.
  * `sp` - (Optional, List) Service Provider (IBM Cloud SP) configuration. Maximum of one block.
  Nested schema for **sp**:
    * `want_assertion_signed` - (Optional, Boolean) Indicates if the SP requires SAML assertions to be signed.
    * `want_response_signed` - (Optional, Boolean) Indicates if the SP requires SAML responses to be signed.
    * `encrypt_response` - (Optional, Boolean) Indicates if the SP requires SAML assertions to be encrypted.
    * `idp_initiated_login_enabled` - (Optional, Boolean) Enables IdP-initiated login (unsolicited SSO).
    * `logout_url_enabled_when_available` - (Optional, Boolean) Enables the SP to use the IdP logout URL when it is available.
    * `idp_initiated_urls` - (Optional, List of String) Target URLs for IdP-initiated login. Only applicable when `idp_initiated_login_enabled` is `true`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the IDP (same as `idp_id`).
* `idp_id` - (String) Unique identifier assigned to the IDP by the IAM Identity Service.
* `entity_tag` - (String) Version of the IDP. This value is required when updating the IDP to prevent stale writes.
* `created_at` - (String) Timestamp when the IDP was created, in ISO 8601 format.
* `modified_at` - (String) Timestamp when the IDP was last modified, in ISO 8601 format.

## Import

You can import the `ibm_iam_idp` resource by using `idp_id`.

# Syntax
<pre>
$ terraform import ibm_iam_idp.my_idp &lt;idp_id&gt;
</pre>

# Example
<pre>
$ terraform import ibm_iam_idp.my_idp a1b2c3d4-e5f6-7890-abcd-ef1234567890
</pre>
