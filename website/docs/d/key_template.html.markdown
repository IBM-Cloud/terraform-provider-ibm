---
layout: "ibm"
page_title: "IBM : ibm_hpcs_key_template"
description: |-
  Get information about key_template
subcategory: "Hyper Protect Crypto Service (HPCS)"
---

# ibm_hpcs_key_template

Provides a read-only data source for key_template. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_hpcs_key_template" "key_template" {
	instance_id = "instance_id"
  	region = "region"
	id = "id"
	uko_vault = ibm_hpcs_key_template.key_template.uko_vault
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `instance_id` - (Required, String) ID of UKO Instance
  * Constraints: Must match the ID of the UKO instance you are trying to work with.
* `region` - (Required, String) Region of the UKO Instance
* `id` - (Required, Forces new resource, String) UUID of the template.
* `uko_vault` - (Required, String) The UUID of the Vault in which the update is to take place.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the key_template.
* `created_at` - (Optional, String) Date and time when the key template was created.

* `created_by` - (Optional, String) ID of the user that created the key template.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-]+$/`.

* `description` - (Required, String) Description of the key template.
  * Constraints: The maximum length is `200` characters. The minimum length is `0` characters. The value must match regular expression `/.*/`.

* `href` - (Optional, String) A URL that uniquely identifies your cloud resource.
  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._~:\/?&=-]+$/`.

* `key` - (Required, List) Properties describing the properties of the managed key.
Nested scheme for **key**:
	* `activation_date` - (Required, String) Key activation date can be provided as a period definition (e.g. PY1 means 1 year).
	  * Constraints: The maximum length is `100` characters. The minimum length is `3` characters. The value must match regular expression `/P^[0-9YMWD]+$/`.
	* `algorithm` - (Required, String) The algorithm of the key.
	  * Constraints: Allowable values are: `aes`, `rsa`.
	* `expiration_date` - (Required, String) Key expiration date can be provided as a period definition (e.g. PY1 means 1 year).
	  * Constraints: The maximum length is `100` characters. The minimum length is `3` characters. The value must match regular expression `/P^[0-9YMWD]+$/`.
	* `size` - (Required, String) The size of the underlying cryptographic key or key pair. E.g. "256" for AES keys, or "2048" for RSA.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9]+$/`.
	* `state` - (Required, String) The state that the key will be in after generation.
	  * Constraints: The default value is `active`. Allowable values are: `pre_activation`, `active`.

* `keystores` - (Required, List) 
  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
Nested scheme for **keystores**:
	* `group` - (Required, String) Which keystore group to distribute the key to.
	  * Constraints: The maximum length is `200` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9][A-Za-z0-9-_ ]+$/`.
	* `type` - (Required, String) Type of keystore.
	  * Constraints: Allowable values are: `aws_kms`, `azure_key_vault`, `ibm_cloud_kms`.

* `name` - (Optional, String) Name of the key template.
  * Constraints: The maximum length is `30` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Z0-9-]+$/`.

* `updated_at` - (Optional, String) Date and time when the key template was updated.

* `updated_by` - (Optional, String) ID of the user that updated the key.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-]+$/`.

* `vault` - (Required, List) Reference to a vault.
Nested scheme for **vault**:
	* `href` - (Optional, String) A URL that uniquely identifies your cloud resource.
	  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._~:\/?&=-]+$/`.
	* `id` - (Optional, String) The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
	* `name` - (Optional, String) Name of the referenced vault.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z][A-Za-z0-9#@!$% '_-]*$/`.

* `version` - (Optional, Integer) Version of the key template. Every time the key template is updated, the version will be updated automatically.
  * Constraints: The maximum value is `2147483647`. The minimum value is `1`.

