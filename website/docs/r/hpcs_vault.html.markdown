---
layout: "ibm"
page_title: "IBM : ibm_hpcs_vault"
description: |-
  Manages vault.
subcategory: "Hyper Protect Crypto Services"
---

# ibm_hpcs_vault

Provides a resource for vault. This allows vault to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_hpcs_vault" "vault_instance" {
  instance_id = "76195d24-8a31-4c6d-9050-c35f09375cfb"
  region      = "us-east"
  name        = "terraformVault"
  description = "example vault"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, String) ID of UKO Instance
  * Constraints: Must match the ID of the UKO instance you are trying to work with.
* `region` - (Required, String) Region of the UKO Instance
  * Constraints: Must match the region of the UKO instance you are trying to work with. Allowable values are: `au-syd`, `in-che`, `jp-osa`, `jp-tok`, `kr-seo`, `eu-de`, `eu-gb`, `ca-tor`, `us-south`, `us-south-test`, `us-east`, `br-sao`.
* `description` - (Optional, String) Description of the vault.
  * Constraints: The maximum length is `200` characters. The minimum length is `0` characters. The value must match regular expression `/(.|\\n)*/`.
* `name` - (Required, String) A human-readable name to assign to your vault. To protect your privacy, do not use personal data, such as your name or location.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9#@!$%'_-][A-Za-z0-9#@!$% '_-]*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `vault_id` - The unique identifier of the vault.
* `created_at` - (String) Date and time when the vault was created.
* `created_by` - (String) ID of the user that created the vault.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z][A-Za-z0-9#@!$%'_-]*$/`.
* `href` - (String) A URL that uniquely identifies your cloud resource.
  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._~:\/?&=-]+$/`.
* `updated_at` - (String) Date and time when the vault was last updated.
* `updated_by` - (String) ID of the user that last updated the vault.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z][A-Za-z0-9#@!$% '_-]*$/`.

* `etag` - ETag identifier for hpcs_vault.

## Import

You can import the `ibm_hpcs_vault` resource by using `region`, `instance_id`, and `vault_id`.

# Syntax
```bash
$ terraform import ibm_hpcs_vault.vault <region>/<instance_id>/<vault_id>
```

# Example
```
$ terraform import ibm_hpcs_vault.vault us-east/76195d24-8a31-4c6d-9050-c35f09375cfb/5295ad47-2ce9-43c3-b9e7-e5a9482c362b
```
