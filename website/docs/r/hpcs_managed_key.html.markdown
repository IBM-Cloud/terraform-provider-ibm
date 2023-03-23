---
layout: "ibm"
page_title: "IBM : ibm_hpcs_managed_key"
description: |-
  Manages managed_key.
subcategory: "Hyper Protect Crypto Services"
---

# ibm_hpcs_managed_key

Provides a resource for managed_key. This allows managed_key to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_hpcs_managed_key" "managed_key_instance" {
  instance_id = ibm_hpcs_vault.vault_instance.instance_id
  region      = ibm_hpcs_vault.vault_instance.region
  uko_vault   = ibm_hpcs_vault.vault_instance.vault_id
  vault {
    id = ibm_hpcs_vault.vault_instance.vault_id
  }
  label         = "terraformKey"
  description   = "example key"
  template_name = ibm_hpcs_key_template.key_template_instance.name
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, String) ID of UKO Instance
  * Constraints: Must match the ID of the UKO instance you are trying to work with.
* `region` - (Required, String) Region of the UKO Instance
  * Constraints: Must match the region of the UKO instance you are trying to work with. Allowable values are: `au-syd`, `in-che`, `jp-osa`, `jp-tok`, `kr-seo`, `eu-de`, `eu-gb`, `ca-tor`, `us-south`, `us-south-test`, `us-east`, `br-sao`.
* `description` - (Optional, String) Description of the managed key.
  * Constraints: The maximum length is `200` characters. The minimum length is `0` characters. The value must match regular expression `/(.|\\n)*/`.
* `label` - (Required, String) The label of the key.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._ \/-]+$/`.
* `tags` - (Optional, List) Key-value pairs associated with the key.
  * Constraints: The maximum length is `128` items. The minimum length is `0` items.
Nested scheme for **tags**:
	* `name` - (Required, String) Name of a tag.
	  * Constraints: The maximum length is `254` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9 -_]+$/`.
	* `value` - (Required, String) Value of a tag.
	  * Constraints: The maximum length is `8192` characters. The minimum length is `0` characters. The value must match regular expression `/^(\\w|\\s)*$/`.
* `template_name` - (Required, String) Name of the key template to use when creating a key.
  * Constraints: The maximum length is `30` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z][A-Za-z0-9-]+$/`.
* `uko_vault` - (Required, String) The UUID of the Vault in which the update is to take place.
* `vault` - (Required, List) ID of the Vault where the entity is to be created in.
Nested scheme for **vault**:
	* `id` - (Required, String) The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `key_id` - The unique identifier of the managed_key.
* `activation_date` - (String) First day when the key is active.
* `algorithm` - (String) The algorithm of the key.
  * Constraints: Allowable values are: `aes`, `rsa`, `hmac`, `ec`.
* `created_at` - (String) Date and time when the key was created.
* `created_by` - (String) ID of the user that created the key.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-]+$/`.
* `expiration_date` - (String) Last day when the key is active.
* `href` - (String) A URL that uniquely identifies your cloud resource.
  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._~:\/?&=-]+$/`.
* `instances` - (List) key instances.
  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
Nested scheme for **instances**:
	* `google_key_protection_level` - (String)
	  * Constraints: Allowable values are: `software`, `hsm`.
	* `google_key_purpose` - (String)
	  * Constraints: Allowable values are: `encrypt_decrypt`, `asymmetric_decrypt`, `asymmetric_sign`, `mac`.
	* `google_kms_algorithm` - (String)
	  * Constraints: Allowable values are: `google_symmetric_encryption`, `ec_sign_p256_sha256`, `ec_sign_p384_sha384`, `ec_sign_secp256k1_sha256`, `rsa_sign_pss_2048_sha256`, `rsa_sign_pss_3072_sha256`, `rsa_sign_pss_4096_sha256`, `rsa_sign_pss_4096_sha512`, `rsa_sign_pkcs1_2048_sha256`, `rsa_sign_pkcs1_3072_sha256`, `rsa_sign_pkcs1_4096_sha256`, `rsa_sign_pkcs1_4096_sha512`, `rsa_sign_raw_pkcs1_2048`, `rsa_sign_raw_pkcs1_3072`, `rsa_sign_raw_pkcs1_4096`, `rsa_decrypt_oaep_2048_sha1`, `rsa_decrypt_oaep_2048_sha256`, `rsa_decrypt_oaep_3072_sha1`, `rsa_decrypt_oaep_3072_sha256`, `rsa_decrypt_oaep_4096_sha1`, `rsa_decrypt_oaep_4096_sha256`, `rsa_decrypt_oaep_4096_sha512`, `hmac_sha256`.
	* `id` - (String) The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
	* `keystore` - (List) Description of properties of a key within the context of keystores.
	Nested scheme for **keystore**:
		* `group` - (String)
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9][A-Za-z0-9_ -]+$/`.
		* `type` - (String) Type of keystore.
		  * Constraints: Allowable values are: `aws_kms`, `azure_key_vault`, `ibm_cloud_kms`, `google_kms`.
	* `label_in_keystore` - (String) The label of the key.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._ \/-]+$/`.
	* `type` - (String) Type of the key instance.
	  * Constraints: Allowable values are: `public_key`, `private_key`, `key_pair`, `secret_key`.
* `referenced_keystores` - (List) referenced keystores.
  * Constraints: The maximum length is `128` items. The minimum length is `0` items.
Nested scheme for **referenced_keystores**:
	* `href` - (String) A URL that uniquely identifies your cloud resource.
	  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._~:\/?&=-]+$/`.
	* `id` - (String) The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
	* `name` - (String) Name of the target keystore.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9][A-Za-z0-9 ._-]*$/`.
	* `type` - (String) Type of keystore.
	  * Constraints: Allowable values are: `aws_kms`, `azure_key_vault`, `ibm_cloud_kms`, `google_kms`.
* `size` - (String) The size of the underlying cryptographic key or key pair. E.g. "256" for AES keys, or "2048" for RSA.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9]+$/`.
* `state` - (String) The state of the key.
  * Constraints: The default value is `active`. Allowable values are: `pre_activation`, `active`, `deactivated`, `destroyed`.
* `template` - (List) Reference to a key template.
Nested scheme for **template**:
	* `href` - (String) A URL that uniquely identifies your cloud resource.
	  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._~:\/?&=-]+$/`.
	* `id` - (String) The v4 UUID used to uniquely identify the resource, as specified by RFC 4122.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
	* `name` - (String) Name of the key template.
	  * Constraints: The maximum length is `30` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z][A-Za-z0-9-]*$/`.
* `updated_at` - (String) Date and time when the key was last updated.
* `updated_by` - (String) ID of the user that last updated the key.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-]+$/`.
* `verification_patterns` - (List) A list of verification patterns of the key (e.g. public key hash for RSA keys).
  * Constraints: The maximum length is `16` items. The minimum length is `1` item.
Nested scheme for **verification_patterns**:
	* `method` - (String) The method used for calculating the verification pattern.
	  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9-]+$/`.
	* `value` - (String) The calculated value.
	  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9+\/=]+$/`.

* `etag` - ETag identifier for managed_key.

## Import

You can import the `ibm_hpcs_managed_key` resource by using `region`, `instance_id`, `vault_id`, and `key_id`.

# Syntax
```bash
$ terraform import ibm_hpcs_managed_key.key <region>/<instance_id>/<vault_id>/<key_id>
```

# Example
```
$ terraform import ibm_hpcs_managed_key.key us-east/76195d24-8a31-4c6d-9050-c35f09375cfb/5295ad47-2ce9-43c3-b9e7-e5a9482c362b/d8cc1ef7-d13b-4731-95be-1f7c98c9f524
```
