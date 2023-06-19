---
layout: "ibm"
page_title: "IBM : ibm_sm_private_certificate_configuration_template"
description: |-
  Get information about PrivateCertificateConfigurationTemplate
subcategory: "Secrets Manager"
---

# ibm_sm_private_certificate_configuration_template

Provides a read-only data source for the configuration of a private certificate template. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_private_certificate_configuration_template" "private_certificate_template" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
  name = "configuration-name"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Optional, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
  * Constraints: Allowable values are: `private`, `public`.
* `name` - (Required, String) The name of the configuration.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9][A-Za-z0-9]*(?:_?-?\\.?[A-Za-z0-9]+)*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the data source.
* `allow_any_name` - (Boolean) Determines whether to allow clients to request a private certificate that matches any common name.

* `allow_bare_domains` - (Boolean) Determines whether to allow clients to request private certificates that match the value of the actual domains on the final certificate.For example, if you specify `example.com` in the `allowed_domains` field, you grant clients the ability to request a certificate that contains the name `example.com` as one of the DNS values on the final certificate.**Important:** In some scenarios, allowing bare domains can be considered a security risk.

* `allow_glob_domains` - (Boolean) Determines whether to allow glob patterns, for example, `ftp*.example.com`, in the names that are specified in the `allowed_domains` field.If set to `true`, clients are allowed to request private certificates with names that match the glob patterns.

* `allow_ip_sans` - (Boolean) Determines whether to allow clients to request a private certificate with IP Subject Alternative Names.

* `allow_localhost` - (Boolean) Determines whether to allow `localhost` to be included as one of the requested common names.

* `allow_subdomains` - (Boolean) Determines whether to allow clients to request private certificates with common names (CN) that are subdomains of the CNs that are allowed by the other certificate template options. This includes wildcard subdomains.For example, if `allowed_domains` has a value of `example.com` and `allow_subdomains`is set to `true`, then the following subdomains are allowed: `foo.example.com`, `bar.example.com`, `*.example.com`.**Note:** This field is redundant if you use the `allow_any_name` option.

* `allowed_domains` - (List) The domains to define for the certificate template. This property is used along with the `allow_bare_domains` and `allow_subdomains` options.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.

* `allowed_domains_template` - (Boolean) Determines whether to allow the domains that are supplied in the `allowed_domains` field to contain access control list (ACL) templates.

* `allowed_other_sans` - (List) The custom Object Identifier (OID) or UTF8-string Subject Alternative Names (SANs) to allow for private certificates.The format for each element in the list is the same as OpenSSL: `<oid>:<type>:<value>` where the current valid type is `UTF8`. To allow any value for an OID, use `*` as its value. Alternatively, specify a single `*` to allow any `other_sans` input.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.

* `allowed_secret_groups` - (String) Scopes the creation of private certificates to only the secret groups that you specify.This field can be supplied as a comma-delimited list of secret group IDs.
  * Constraints: The maximum length is `1024` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.

* `allowed_uri_sans` - (List) The URI Subject Alternative Names to allow for private certificates.Values can contain glob patterns, for example `spiffe://hostname/_*`.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.

* `basic_constraints_valid_for_non_ca` - (Boolean) Determines whether to mark the Basic Constraints extension of an issued private certificate as valid for non-CA certificates.

* `certificate_authority` - (String) The name of the intermediate certificate authority.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9][A-Za-z0-9]*(?:_?-?\\.?[A-Za-z0-9]+)*$/`.

* `client_flag` - (Boolean) Determines whether private certificates are flagged for client use.

* `code_signing_flag` - (Boolean) Determines whether private certificates are flagged for code signing use.

* `config_type` - (String) Th configuration type.
  * Constraints: Allowable values are: `public_cert_configuration_ca_lets_encrypt`, `public_cert_configuration_dns_classic_infrastructure`, `public_cert_configuration_dns_cloud_internet_services`, `iam_credentials_configuration`, `private_cert_configuration_root_ca`, `private_cert_configuration_intermediate_ca`, `private_cert_configuration_template`.

* `country` - (List) The Country (C) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `created_at` - (String) The date when a resource was created. The date format follows RFC 3339.

* `created_by` - (String) The unique identifier that is associated with the entity that created the secret.
  * Constraints: The maximum length is `128` characters. The minimum length is `4` characters.

* `email_protection_flag` - (Boolean) Determines whether private certificates are flagged for email protection use.

* `enforce_hostnames` - (Boolean) Determines whether to enforce only valid host names for common names, DNS Subject Alternative Names, and the host section of email addresses.

* `ext_key_usage` - (List) The allowed extended key usage constraint on private certificates.You can find valid values in the [Go x509 package documentation](https://golang.org/pkg/crypto/x509/#ExtKeyUsage). Omit the `ExtKeyUsage` part of the value. Values are not case-sensitive. To specify no key usage constraints, set this field to an empty list.
  * Constraints: The list items must match regular expression `/^[a-zA-Z]+$/`. The maximum length is `100` items. The minimum length is `0` items.

* `ext_key_usage_oids` - (List) A list of extended key usage Object Identifiers (OIDs).
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.

* `key_bits` - (Integer) The number of bits to use to generate the private key.Allowable values for RSA keys are: `2048` and `4096`. Allowable values for EC keys are: `224`, `256`, `384`, and `521`. The default for RSA keys is `2048`. The default for EC keys is `256`.

* `key_type` - (String) The type of private key to generate.
  * Constraints: Allowable values are: `rsa`, `ec`.

* `key_usage` - (List) The allowed key usage constraint to define for private certificates.You can find valid values in the [Go x509 package documentation](https://pkg.go.dev/crypto/x509#KeyUsage). Omit the `KeyUsage` part of the value. Values are not case-sensitive. To specify no key usage constraints, set this field to an empty list.
  * Constraints: The list items must match regular expression `/^[a-zA-Z]+$/`. The maximum length is `100` items. The minimum length is `0` items.

* `locality` - (List) The Locality (L) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `max_ttl_seconds` - (Integer) The maximum time-to-live (TTL) for certificates that are created by this CA in seconds.

* `not_before_duration_seconds` - (Integer) The duration in seconds by which to backdate the `not_before` property of an issued private certificate.

* `organization` - (List) The Organization (O) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `ou` - (List) The Organizational Unit (OU) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `policy_identifiers` - (List) A list of policy Object Identifiers (OIDs).
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.

* `postal_code` - (List) The postal code values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `province` - (List) The Province (ST) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `require_cn` - (Boolean) Determines whether to require a common name to create a private certificate.By default, a common name is required to generate a certificate. To make the `common_name` field optional, set the `require_cn` option to `false`.

* `secret_type` - (String) The secret type. Supported types are arbitrary, certificates (imported, public, and private), IAM credentials, key-value, and user credentials.
  * Constraints: Allowable values are: `arbitrary`, `imported_cert`, `public_cert`, `iam_credentials`, `kv`, `username_password`, `private_cert`.

* `serial_number` - (String) The serial number to assign to the generated certificate. To assign a random serial number, you can omit this field.
  * Constraints: The maximum length is `64` characters. The minimum length is `32` characters. The value must match regular expression `/[^a-fA-F0-9]/`.

* `server_flag` - (Boolean) Determines whether private certificates are flagged for server use.

* `street_address` - (List) The street address values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `ttl_seconds` - (Integer) The requested Time To Live, after which the certificate will be expired.

* `updated_at` - (String) The date when a resource was recently modified. The date format follows RFC 3339.

* `use_csr_common_name` - (Boolean) When used with the `private_cert_configuration_action_sign_csr` action, this field determines whether to use the common name (CN) from a certificate signing request (CSR) instead of the CN that's included in the data of the certificate.Does not include any requested Subject Alternative Names (SANs) in the CSR. To use the alternative names, include the `use_csr_sans` property.

* `use_csr_sans` - (Boolean) When used with the `private_cert_configuration_action_sign_csr` action, this field determines whether to use the Subject Alternative Names(SANs) from a certificate signing request (CSR) instead of the SANs that are included in the data of the certificate.Does not include the common name in the CSR. To use the common name, include the `use_csr_common_name` property.

