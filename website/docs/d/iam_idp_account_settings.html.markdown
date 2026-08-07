---
layout: "ibm"
page_title: "IBM : ibm_iam_idp_account_settings"
description: |-
  Get information about IAM Identity Provider (IdP) account settings.
subcategory: "IAM Identity Services"
---

# ibm_iam_idp_account_settings

Provides a read-only data source to list IAM Identity Provider (IdP) account settings for a given account. Depending on the `type` argument, this data source returns either the IdPs that the account can consume (`consumable`) or the IdPs that the account is already consuming (`consumed`).

For more information, see the [IAM Identity Services API documentation](https://cloud.ibm.com/apidocs/iam-identity-token-api).

## Example Usage

### List IdPs available for an account to consume

```hcl
data "ibm_iam_idp_account_settings" "consumable" {
  account_id = var.account_id
  type       = "consumable"
}
```

### List IdPs already bound to an account

```hcl
data "ibm_iam_idp_account_settings" "consumed" {
  account_id = var.account_id
  type       = "consumed"
}
```

### Output the names of all bound IdPs

```hcl
output "bound_idp_names" {
  value = [for idp in data.ibm_iam_idp_account_settings.consumed.idps : idp.idp_name]
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `account_id` - (Required, String) Account ID to retrieve IDP settings for.
* `type` - (Required, String) The type of IDP settings to list.
  * Constraints: Allowed values are:
    * `consumable` - Returns IdPs that are shared with this account and available to be bound.
    * `consumed` - Returns IdPs that are already bound (active) in this account.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of this data source, formed as `<account_id>/<type>`.
* `idps` - (List) List of IDP account settings.
Nested schema for **idps**:
  * `idp_id` - (String) Identity provider ID.
  * `owner_account` - (String) ID of the account that owns (created) the IdP.
  * `owner_account_name` - (String) Display name of the account that owns the IdP.
  * `idp_name` - (String) Display name of the Identity Provider.
  * `idp_type` - (String) Type of the Identity Provider (e.g. `saml`, `appid`, `ldap`).
  * `cloud_user_strategy` - (String) Strategy for how Cloud User representatives are managed for users of this IdP. Possible values are `DYNAMIC`, `STATIC`, `NEVER`.
  * `active` - (Boolean) Whether the IdP is enabled for usage in this account.
  * `ui_default` - (Boolean) Whether the IdP is the default login option shown in the IBM Cloud UI for this account.
