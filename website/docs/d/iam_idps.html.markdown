---
layout: "ibm"
page_title: "IBM : ibm_iam_idps"
description: |-
  Get information about all IAM Identity Providers (IdPs) in an account.
subcategory: "IAM Identity Services"
---

# ibm_iam_idps

Provides a read-only data source to list all IAM Identity Providers (IdPs) in a given account. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

For more information, see the [IAM Identity Services API documentation](https://cloud.ibm.com/apidocs/iam-identity-token-api).

## Example Usage

```hcl
data "ibm_iam_idps" "all_idps" {
  account_id = var.account_id
}
```

### Iterate over returned IdPs

```hcl
output "idp_names" {
  value = [for idp in data.ibm_iam_idps.all_idps.idps : idp.name]
}

output "active_idps" {
  value = [for idp in data.ibm_iam_idps.all_idps.idps : idp if idp.active == true]
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `account_id` - (Required, String) Account ID to list Identity Providers for.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of this data source (set to the `account_id`).
* `idps` - (List) List of Identity Providers in the account.
Nested schema for **idps**:
  * `idp_id` - (String) Unique identifier of the IDP.
  * `account_id` - (String) Account where the IdP resides.
  * `name` - (String) Speaking name of the Identity Provider.
  * `type` - (String) Type of the IDP (e.g. `saml`, `appid`, `ldap`).
  * `active` - (Boolean) Whether the IDP is active for all consuming accounts.
  * `entity_tag` - (String) Version of the IDP.
  * `share_scope` - (List) List of targets that can consume the IdP.
  Nested schema for **share_scope**:
    * `id` - (String) ID of the account or enterprise.
    * `type` - (String) Type of the share target. Possible values are `account`, `enterprise`.
  * `created_at` - (String) Timestamp when the IDP was created, in ISO 8601 format.
  * `modified_at` - (String) Timestamp when the IDP was last modified, in ISO 8601 format.
