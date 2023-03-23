---
layout: "ibm"
page_title: "IBM : ibm_hpcs_key_template"
description: |-
  Get information about key_template
subcategory: "Hyper Protect Crypto Services"
---

# ibm_hpcs_key_template

Provides a read-only data source for key_template. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_hpcs_key_template" "key_template" {
  instance_id = "76195d24-8a31-4c6d-9050-c35f09375cfb"
  region = "us-east"
  template_id = "d8cc1ef7-d13b-4731-95be-1f7c98c9f524"
  uko_vault = ibm_hpcs_vault.vault.vault_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `instance_id` - (Required, String) ID of UKO Instance
  * Constraints: Must match the ID of the UKO instance you are trying to work with.
* `region` - (Required, String) Region of the UKO Instance
  * Constraints: Allowable values are: `au-syd`, `in-che`, `jp-osa`, `jp-tok`, `kr-seo`, `eu-de`, `eu-gb`, `ca-tor`, `us-south`, `us-south-test`, `us-east`, `br-sao`.
* `template_id` - (Required, Forces new resource, String) UUID of the template.
* `uko_vault` - (Required, String) The UUID of the Vault in which the update is to take place.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `template_id` - The unique identifier of the key_template.
* `created_at` - (String) Date and time when the key template was created.

* `created_by` - (String) ID of the user that created the key template.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-]+$/`.

* `description` - (String) Description of the key template.
  * Constraints: The maximum length is `200` characters. The minimum length is `0` characters. The value must match regular expression `/.*/`.

* `href` - (String) A URL that uniquely identifies your cloud resource.
  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._~:\/?&=-]+$/`.

* `key` - (List) Properties describing the properties of the managed key.
Nested scheme for **key**:
	* `activation_date` - (String) Key activation date can be provided as a period definition (e.g. PY1 means 1 year).
	  * Constraints: The maximum length is `100` characters. The minimum length is `3` characters. The value must match regular expression `/P^[0-9YMWD]+$/`.
	* `algorithm` - (String) The algorithm of the key.
	  * Constraints: Allowable values are: `aes`, `rsa`, `hmac`, `ec`.
	* `expiration_date` - (String) Key expiration date can be provided as a period definition (e.g. PY1 means 1 year).
	  * Constraints: The maximum length is `100` characters. The minimum length is `3` characters. The value must match regular expression `/P^[0-9YMWD]+$/`.
	* `size` - (String) The size of the underlying cryptographic key or key pair. E.g. "256" for AES keys, or "2048" for RSA.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9]+$/`.
	* `state` - (String) The state that the key will be in after generation.
	  * Constraints: The default value is `active`. Allowable values are: `pre_activation`, `active`.

* `keystores` - (List) 
  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
Nested scheme for **keystores**:
	* `google_key_protection_level` - (String)
	  * Constraints: Allowable values are: `software`, `hsm`.
	* `google_key_purpose` - (String)
	  * Constraints: Allowable values are: `encrypt_decrypt`, `asymmetric_decrypt`, `asymmetric_sign`, `mac`.
	* `google_kms_algorithm` - (String)
	  * Constraints: Allowable values are: `google_symmetric_encryption`, `ec_sign_p256_sha256`, `ec_sign_p384_sha384`, `ec_sign_secp256k1_sha256`, `rsa_sign_pss_2048_sha256`, `rsa_sign_pss_3072_sha256`, `rsa_sign_pss_4096_sha256`, `rsa_sign_pss_4096_sha512`, `rsa_sign_pkcs1_2048_sha256`, `rsa_sign_pkcs1_3072_sha256`, `rsa_sign_pkcs1_4096_sha256`, `rsa_sign_pkcs1_4096_sha512`, `rsa_sign_raw_pkcs1_2048`, `rsa_sign_raw_pkcs1_3072`, `rsa_sign_raw_pkcs1_4096`, `rsa_decrypt_oaep_2048_sha1`, `rsa_decrypt_oaep_2048_sha256`, `rsa_decrypt_oaep_3072_sha1`, `rsa_decrypt_oaep_3072_sha256`, `rsa_decrypt_oaep_4096_sha1`, `rsa_decrypt_oaep_4096_sha256`, `rsa_decrypt_oaep_4096_sha512`, `hmac_sha256`.
	* `group` - (String) Which keystore group to distribute the key to.
	  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9][A-Za-z0-9-_ ]+$/`.
	* `type` - (String) Type of keystore.
	  * Constraints: Allowable values are: `aws_kms`, `azure_key_vault`, `ibm_cloud_kms`, `google_kms`.

* `name` - (ForceNew, String) Name of the key template. Updating the name is not supported in this version, so changing this value will require recreating the resource, which may be impossible while other resources rely on it.
  * Constraints: The maximum length is `30` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Z0-9-]+$/`.

* `updated_at` - (String) Date and time when the key template was updated.

* `updated_by` - (String) ID of the user that updated the key.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-]+$/`.

* `vault` - (List) Reference to a vault.
Nested scheme for **vault**:
	* `href` - (String) A URL that uniquely identifies your cloud resource.
	  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._~:\/?&=-]+$/`.
	* `id` - (String) The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
	* `name` - (String) Name of the referenced vault.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9#@!$%'_-][A-Za-z0-9#@!$% '_-]*$/`.

* `version` - (Integer) Version of the key template. Every time the key template is updated, the version will be updated automatically.
  * Constraints: The maximum value is `2147483647`. The minimum value is `1`.

