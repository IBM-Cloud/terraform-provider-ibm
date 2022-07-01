---
layout: "ibm"
page_title: "IBM : ibm_hpcs_vault"
description: |-
  Get information about vault
subcategory: "Hyper Protect Crypto Services"
---

# ibm_hpcs_vault

Provides a read-only data source for vault. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_hpcs_vault" "vault" {
  instance_id = "76195d24-8a31-4c6d-9050-c35f09375cfb"
  region = "us-east"
  vault_id = "5295ad47-2ce9-43c3-b9e7-e5a9482c362b"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `instance_id` - (Required, String) ID of UKO Instance
  * Constraints: Must match the ID of the UKO instance you are trying to work with.
* `region` - (Required, String) Region of the UKO Instance
  * Constraints: Allowable values are: `au-syd`, `in-che`, `jp-osa`, `jp-tok`, `kr-seo`, `eu-de`, `eu-gb`, `ca-tor`, `us-south`, `us-south-test`, `us-east`, `br-sao`.
* `vault_id` - (Required, Forces new resource, String) UUID of the vault.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `vault_id` - The unique identifier of the vault.
* `created_at` - (String) Date and time when the vault was created.

* `created_by` - (String) ID of the user that created the vault.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z][A-Za-z0-9#@!$%'_-]*$/`.

* `description` - (Required, String) Description of the vault.
  * Constraints: The maximum length is `200` characters. The minimum length is `0` characters. The value must match regular expression `/.*/`.

* `href` - (String) A URL that uniquely identifies your cloud resource.
  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._~:\/?&=-]+$/`.

* `name` - (String) Name of the vault.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z][A-Za-z0-9#@!$% '_-]*$/`.

* `updated_at` - (String) Date and time when the vault was last updated.

* `updated_by` - (String) ID of the user that last updated the vault.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z][A-Za-z0-9#@!$% '_-]*$/`.

