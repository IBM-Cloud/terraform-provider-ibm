---
layout: "ibm"
page_title: "IBM : ibm_sm_public_certificate"
description: |-
  Get information about PublicCertificate
subcategory: "Secrets Manager"
---

# ibm_sm_public_certificate

Provides a read-only data source for a public certificate. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.
The data source can be defined by providing the secret ID or the secret and secret group names.

## Example Usage

By secret id
```hcl
data "ibm_sm_public_certificate" "public_certificate" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
  secret_id = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}
```

By secret name and group name
```hcl
data "ibm_sm_public_certificate" "public_certificate" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
  name          = "secret-name"
  secret_group_name = "group-name"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Optional, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
    * Constraints: Allowable values are: `private`, `public`.
* `secret_id` - (Optional, String) The ID of the secret.
    * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
* `name` - (Optional, String) The human-readable name of your secret. To be used in combination with `secret_group_name`.
    * Constraints: The maximum length is `256` characters. The minimum length is `2` characters. The value must match regular expression `^[A-Za-z0-9][A-Za-z0-9]*(?:_*-*\\.*[A-Za-z0-9]+)*$`.
* `secret_group_name` - (Optional, String) The name of your existing secret group. To be used in combination with `name`.
    * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the data source.
* `alt_names` - (List) With the Subject Alternative Name field, you can specify additional host names to be protected by a single SSL certificate.
  * Constraints: The list items must match regular expression `/^(.*?)$/`. The maximum length is `99` items. The minimum length is `0` items.

* `bundle_certs` - (Boolean) Indicates whether the issued certificate is bundled with intermediate certificates.

* `ca` - (String) The name of the certificate authority configuration..

* `certificate` - (String) The PEM-encoded contents of your certificate.
  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.

* `common_name` - (String) The Common Name (AKA CN) represents the server name protected by the SSL certificate.
  * Constraints: The maximum length is `64` characters. The minimum length is `4` characters. The value must match regular expression `/^(\\*\\.)?(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])\\.?$/`.

* `created_at` - (String) The date when a resource was created. The date format follows RFC 3339.

* `created_by` - (String) The unique identifier that is associated with the entity that created the secret.
  * Constraints: The maximum length is `128` characters. The minimum length is `4` characters.

* `crn` - (String) A CRN that uniquely identifies an IBM Cloud resource.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.

* `custom_metadata` - (Map) The secret metadata that a user can customize.

* `description` - (String) An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/(.*?)/`.

* `dns` - (String) The name of the DNS provider configuration.

* `downloaded` - (Boolean) Indicates whether the secret data that is associated with a secret version was retrieved in a call to the service API.

* `expiration_date` - (String) The date a secret is expired. The date format follows RFC 3339.

* `intermediate` - (String) (Optional) The PEM-encoded intermediate certificate to associate with the root certificate.
  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.

* `issuance_info` - (List) Issuance information that is associated with your certificate.
Nested scheme for **issuance_info**:
	* `auto_rotated` - (Boolean) Indicates whether the issued certificate is configured with an automatic rotation policy.
	* `challenges` - (List) The set of challenges. It is returned only when ordering public certificates by using manual DNS configuration.
	  * Constraints: The maximum length is `100` items. The minimum length is `1` item.
	Nested scheme for **challenges**:
		* `domain` - (String) The challenge domain.
		* `expiration` - (String) The challenge expiration date. The date format follows RFC 3339.
		* `status` - (String) The challenge status.
		* `txt_record_name` - (String) The TXT record name.
		* `txt_record_value` - (String) The TXT record value.
	* `dns_challenge_validation_time` - (String) The date that a user requests to validate DNS challenges for certificates that are ordered with a manual DNS provider. The date format follows RFC 3339.
	* `error_code` - (String) A code that identifies an issuance error.This field, along with `error_message`, is returned when Secrets Manager successfully processes your request, but the certificate authority is unable to issue a certificate.
	* `error_message` - (String) A human-readable message that provides details about the issuance error.
	* `ordered_on` - (String) The date when the certificate is ordered. The date format follows RFC 3339.
	* `state` - (Integer) The secret state that is based on NIST SP 800-57. States are integers and correspond to the `Pre-activation = 0`, `Active = 1`,  `Suspended = 2`, `Deactivated = 3`, and `Destroyed = 5` values.
	  * Constraints: Allowable values are: `0`, `1`, `2`, `3`, `5`.
	* `state_description` - (String) A text representation of the secret state.
	  * Constraints: Allowable values are: `pre_activation`, `active`, `suspended`, `deactivated`, `destroyed`.

* `issuer` - (String) The distinguished name that identifies the entity that signed and issued the certificate.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.

* `key_algorithm` - (String) The identifier for the cryptographic algorithm to be used to generate the public key that is associated with the certificate.The algorithm that you select determines the encryption algorithm (`RSA` or `ECDSA`) and key size to be used to generate keys and sign certificates. For longer living certificates, it is recommended to use longer keys to provide more encryption protection. Allowed values:  RSA2048, RSA4096, EC256, EC384.
  * Constraints: The default value is `RSA2048`. The maximum length is `7` characters. The minimum length is `5` characters. The value must match regular expression `/^(RSA2048|RSA4096|EC256|EC384)$/`.

* `labels` - (List) Labels that you can use to search for secrets in your instance.Up to 30 labels can be created.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `30` items. The minimum length is `0` items.

* `locks_total` - (Integer) The number of locks of the secret.
  * Constraints: The maximum value is `1000`. The minimum value is `0`.

* `name` - (String) The human-readable name of your secret.
  * Constraints: The maximum length is `256` characters. The minimum length is `2` characters.

* `private_key` - (String) (Optional) The PEM-encoded private key to associate with the certificate.
  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.

* `rotation` - (List) Determines whether Secrets Manager rotates your secrets automatically.
Nested scheme for **rotation**:
	* `auto_rotate` - (Boolean) Determines whether Secrets Manager rotates your secret automatically.Default is `false`.
	* `rotate_keys` - (Boolean) Determines whether Secrets Manager rotates the private key for your public certificate automatically.Default is `false`. If it is set to `true`, the service generates and stores a new private key for your rotated certificate.

* `secret_group_id` - (String) A UUID identifier, or `default` secret group.
  * Constraints: The maximum length is `36` characters. The minimum length is `7` characters. The value must match regular expression `/^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|default)$/`.

* `secret_type` - (String) The secret type. Supported types are arbitrary, certificates (imported, public, and private), IAM credentials, key-value, and user credentials.
  * Constraints: Allowable values are: `arbitrary`, `imported_cert`, `public_cert`, `iam_credentials`, `kv`, `username_password`, `private_cert`.

* `serial_number` - (String) The unique serial number that was assigned to a certificate by the issuing certificate authority.
  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[^a-fA-F0-9]/`.

* `signing_algorithm` - (String) The identifier for the cryptographic algorithm that was used by the issuing certificate authority to sign a certificate.
  * Constraints: The maximum length is `64` characters. The minimum length is `4` characters.

* `state` - (Integer) The secret state that is based on NIST SP 800-57. States are integers and correspond to the `Pre-activation = 0`, `Active = 1`,  `Suspended = 2`, `Deactivated = 3`, and `Destroyed = 5` values.
  * Constraints: Allowable values are: `0`, `1`, `2`, `3`, `5`.

* `state_description` - (String) A text representation of the secret state.
  * Constraints: Allowable values are: `pre_activation`, `active`, `suspended`, `deactivated`, `destroyed`.

* `updated_at` - (String) The date when a resource was recently modified. The date format follows RFC 3339.

* `validity` - (List) The date and time that the certificate validity period begins and ends.
Nested scheme for **validity**:
	* `not_after` - (String) The date-time format follows RFC 3339.
	* `not_before` - (String) The date-time format follows RFC 3339.

* `versions_total` - (Integer) The number of versions of the secret.
  * Constraints: The maximum value is `50`. The minimum value is `0`.

