---
layout: "ibm"
page_title: "IBM : ibm_iam_idp_account_setting"
description: |-
  Manages the binding of an IAM Identity Provider (IdP) to an IBM Cloud account.
subcategory: "IAM Identity Services"
---

# ibm_iam_idp_account_setting

Create, update, and delete the binding of a shared IAM Identity Provider (IdP) to a consuming IBM Cloud account. This resource uses the IDP sharing feature — the IdP must first exist in an owner account (created via [`ibm_iam_idp`](iam_idp.html.markdown)) and its `share_scope` must include the target account before it can be bound.

For more information, see the [IAM Identity Services API documentation](https://cloud.ibm.com/apidocs/iam-identity-token-api).

## Example Usage

### Bind a shared IdP to a consumer account

```hcl
# The IdP is created in the owner account and shared with the consumer
resource "ibm_iam_idp" "shared_idp" {
  account_id = var.owner_account_id
  name       = "shared-saml-idp"
  type       = "saml"
  active     = true

  share_scope {
    id   = var.consumer_account_id
    type = "account"
  }
}

# Bind the shared IdP to the consumer account
resource "ibm_iam_idp_account_setting" "consumer_binding" {
  account_id           = var.consumer_account_id
  idp_id               = ibm_iam_idp.shared_idp.idp_id
  cloud_user_strategy  = "DYNAMIC"
  active               = true
  ui_default           = true
}
```

### Bind with static user strategy

```hcl
resource "ibm_iam_idp_account_setting" "static_binding" {
  account_id           = var.account_id
  idp_id               = var.shared_idp_id
  cloud_user_strategy  = "STATIC"
  active               = true
  ui_default           = false
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `account_id` - (Required, Forces new resource, String) Account to bind the IdP to. Changing this value creates a new resource.
* `idp_id` - (Required, Forces new resource, String) Identity provider ID to bind to the account. The IdP must be shared with this account via its `share_scope`. Changing this value creates a new resource.
* `cloud_user_strategy` - (Required, String) Strategy for how Cloud User representatives are managed for the IdP users.
  * Constraints: Allowed values are:
    * `DYNAMIC` - Cloud User records are created automatically on first login.
    * `STATIC` - Cloud User records must be pre-provisioned.
    * `NEVER` - No Cloud User records are created; users cannot log in to IBM Cloud resources.
* `active` - (Required, Boolean) Specifies if the IdP is enabled for usage in the given account context.
* `ui_default` - (Required, Boolean) Specifies if the IdP is presented as the default login option in the IBM Cloud UI for this account.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the binding, formed as `<account_id>/<idp_id>`.
* `owner_account` - (String) ID of the account that owns (created) the IdP.
* `owner_account_name` - (String) Display name of the account that owns the IdP.
* `idp_name` - (String) Display name of the Identity Provider.
* `idp_type` - (String) Type of the Identity Provider (e.g. `saml`, `appid`, `ldap`).

## Import

You can import the `ibm_iam_idp_account_setting` resource by using `id`.
The `id` property is formed from `account_id` and `idp_id` in the following format:

<pre>
&lt;account_id&gt;/&lt;idp_id&gt;
</pre>

* `account_id`: A string. Account bound to the IDP.
* `idp_id`: A string. Identity provider ID.

# Syntax
<pre>
$ terraform import ibm_iam_idp_account_setting.my_binding &lt;account_id&gt;/&lt;idp_id&gt;
</pre>

# Example
<pre>
$ terraform import ibm_iam_idp_account_setting.my_binding abc123def456/a1b2c3d4-e5f6-7890-abcd-ef1234567890
</pre>
