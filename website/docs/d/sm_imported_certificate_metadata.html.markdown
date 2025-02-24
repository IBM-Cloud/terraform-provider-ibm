---
layout: "ibm"
page_title: "IBM : ibm_sm_imported_certificate_metadata"
description: |-
  Get information about ImportedCertificateMetadata
subcategory: "Secrets Manager"
---

# ibm_sm_imported_certificate_metadata

Provides a read-only data source for the metadata of an imported certificate. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_imported_certificate_metadata" "imported_certificate_metadata" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
  secret_id = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Optional, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
  * Constraints: Allowable values are: `private`, `public`.
* `secret_id` - (Required, String) The ID of the secret.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the data source.

* `common_name` - (String) The Common Name (AKA CN) represents the server name protected by the SSL certificate.
  * Constraints: The maximum length is `64` characters. The minimum length is `4` characters. The value must match regular expression `/^(\\*\\.)?(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])\\.?$/`.

* `created_at` - (String) The date when a resource was created. The date format follows RFC 3339.

* `created_by` - (String) The unique identifier that is associated with the entity that created the secret.
  * Constraints: The maximum length is `128` characters. The minimum length is `4` characters.

* `crn` - (String) A CRN that uniquely identifies an IBM Cloud resource.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.

* `csr` - (String) The certificate signing request generated based on the parameters in the `managed_csr` data. The value may differ from the `csr` attribute within `managed_csr` if the `managed_csr` attributes have been modified.

* `custom_metadata` - (Map) The secret metadata that a user can customize.

* `description` - (String) An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/(.*?)/`.

* `downloaded` - (Boolean) Indicates whether the secret data that is associated with a secret version was retrieved in a call to the service API.

* `expiration_date` - (String) The date a secret is expired. The date format follows RFC 3339.

* `intermediate_included` - (Boolean) Indicates whether the certificate was imported with an associated intermediate certificate.

* `issuer` - (String) The distinguished name that identifies the entity that signed and issued the certificate.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.

* `key_algorithm` - (String) The identifier for the cryptographic algorithm used to generate the public key that is associated with the certificate.
  * Constraints: The maximum length is `10` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.

* `labels` - (List) Labels that you can use to search for secrets in your instance.Up to 30 labels can be created.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `30` items. The minimum length is `0` items.

* `locks_total` - (Integer) The number of locks of the secret.
  * Constraints: The maximum value is `1000`. The minimum value is `0`.

* `managed_csr` - (List) The data specified to create the CSR and the private key.
  Nested scheme for **managed_csr**:
  * `alt_names` - (String) With the Subject Alternative Name field, you can specify additional hostnames to be protected by a single SSL certificate.
  * `client_flag` - (Boolean) This field indicates whether certificate is flagged for client use.
  * `code_signing_flag` - ( Boolean) This field indicates whether certificate is flagged for code signing use.
  * `common_name` - (String) The Common Name (CN) represents the server name protected by the SSL certificate.
  * `csr` - (String) The certificate signing request generated based on the parameters in the `managed_csr` data.
  * `country` - (List) The Country (C) values to define in the subject field of the resulting certificate.
  * `email_protection_flag` - (String) This field indicates whether certificate is flagged for email protection use.
  * `exclude_cn_from_sans` - (String) This parameter controls whether the common name is excluded from Subject Alternative Names (SANs).
  * `ext_key_usage` - (String) The allowed extended key usage constraint on certificate, in a comma-delimited list.
  * `ext_key_usage_oids` - (String) A comma-delimited list of extended key usage Object Identifiers (OIDs).
  * `ip_sans` - (String) The IP Subject Alternative Names to define for the certificate, in a comma-delimited list.
  * `key_bits` - (Integer) The number of bits to use to generate the private key.
  * `key_type` - (String) The type of private key to generate.
  * `key_usage` - (String) The allowed key usage constraint to define for certificate, in a comma-delimited list.
  * `locality` - (List) The Locality (L) values to define in the subject field of the resulting certificate.
  * `organization` - (List) The Organization (O) values to define in the subject field of the resulting certificate.
  * `other_sans` - (String) The custom Object Identifier (OID) or UTF8-string Subject Alternative Names to define for the certificate, in a comma-delimited list.
  * `ou` - (List) The Organizational Unit (OU) values to define in the subject field of the resulting certificate.
  * `policy_identifiers` - (String) A comma-delimited list of policy Object Identifiers (OIDs).
  * `postal_code` - (List) The postal code values to define in the subject field of the resulting certificate.
  * `province` - (List) The Province (ST) values to define in the subject field of the resulting certificate.
  * `require_cn` - (Boolean) If set to false, makes the common_name field optional while generating a certificate.
  * `rotate_keys` - (Boolean) This field indicates whether the private key will be rotated.
  * `server_flag` - (Boolean) This field indicates whether certificate is flagged for server use.
  * `street_address` - (List) The street address values to define in the subject field of the resulting certificate.
  * `uri_sans` - (String) The URI Subject Alternative Names to define for the certificate, in a comma-delimited list.
  * `user_ids` - (String) Specifies the list of requested User ID (OID 0.9.2342.19200300.100.1.1) Subject values to be placed on the signed certificate.

* `name` - (String) The human-readable name of your secret.
  * Constraints: The maximum length is `256` characters. The minimum length is `2` characters.

* `private_key_included` - (Boolean) Indicates whether the certificate was imported with an associated private key.

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

