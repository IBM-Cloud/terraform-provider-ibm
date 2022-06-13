---
layout: "ibm"
page_title: "IBM : ibm_hpcs_vault"
description: |-
  Get information about vault
subcategory: "Hyper Protect Crypto Service (HPCS)"
---

# ibm_hpcs_vault

Provides a read-only data source for vault. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_hpcs_vault" "vault" {
	instance_id = "instance_id"
  region = "region"
  id = "id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `instance_id` - (Required, String) ID of UKO Instance
  * Constraints: Must match the ID of the UKO instance you are trying to work with.
* `region` - (Required, String) Region of the UKO Instance
* `id` - (Required, Forces new resource, String) UUID of the vault.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the vault.
* `created_at` - (Optional, String) Date and time when the vault was created.

* `created_by` - (Optional, String) ID of the user that created the vault.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z][A-Za-z0-9#@!$%'_-]*$/`.

* `description` - (Required, String) Description of the vault.
  * Constraints: The maximum length is `200` characters. The minimum length is `0` characters. The value must match regular expression `/.*/`.

* `href` - (Optional, String) A URL that uniquely identifies your cloud resource.
  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._~:\/?&=-]+$/`.

* `name` - (Required, String) Name of the vault.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z][A-Za-z0-9#@!$% '_-]*$/`.

* `updated_at` - (Optional, String) Date and time when the vault was last updated.

* `updated_by` - (Optional, String) ID of the user that last updated the vault.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z][A-Za-z0-9#@!$% '_-]*$/`.

