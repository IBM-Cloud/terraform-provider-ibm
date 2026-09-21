---
layout: "ibm"
page_title: "IBM : ibm_iam_idp"
description: |-
  Get information about an IAM Identity Provider (IdP).
subcategory: "IAM Identity Services"
---

# ibm_iam_idp

Provides a read-only data source to retrieve information about an IAM Identity Provider (IdP). You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

For more information, see the [IAM Identity Services API documentation](https://cloud.ibm.com/apidocs/iam-identity-token-api).

## Example Usage

```hcl
data "ibm_iam_idp" "my_idp" {
  idp_id = "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}
```

### Reference fields from the data source

```hcl
output "idp_name" {
  value = data.ibm_iam_idp.my_idp.name
}

output "idp_active" {
  value = data.ibm_iam_idp.my_idp.active
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `idp_id` - (Required, String) Unique identifier of the Identity Provider to retrieve.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the IDP (same value as `idp_id`).
* `account_id` - (String) Account where the IdP resides.
* `name` - (String) Speaking name of the Identity Provider.
* `type` - (String) Type of the IDP (e.g. `saml`, `appid`, `ldap`).
* `active` - (Boolean) Whether the IDP is active (enabled) for all consuming accounts.
* `entity_tag` - (String) Version of the IDP. Required when updating the IDP to avoid stale writes.
* `share_scope` - (List) List of targets (accounts or enterprises) that can consume the IdP.
Nested schema for **share_scope**:
  * `id` - (String) ID of the account or enterprise.
  * `type` - (String) Type of the share target. Possible values are `account`, `enterprise`.
* `created_at` - (String) Timestamp when the IDP was created, in ISO 8601 format.
* `modified_at` - (String) Timestamp when the IDP was last modified, in ISO 8601 format.
