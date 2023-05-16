---
layout: "ibm"
page_title: "IBM : ibm_sm_private_certificate_configuration_root_ca"
description: |-
  Get information about PrivateCertificateConfigurationRootCA
subcategory: "Secrets Manager"
---

# ibm_sm_private_certificate_configuration_root_ca

Provides a read-only data source for the configuraion of a root CA. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_private_certificate_configuration_root_ca" "root_ca" {
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
* `alt_names` - (List) With the Subject Alternative Name field, you can specify additional host names to be protected by a single SSL certificate.
  * Constraints: The list items must match regular expression `/^(.*?)$/`. The maximum length is `99` items. The minimum length is `0` items.

* `common_name` - (String) The Common Name (AKA CN) represents the server name that is protected by the SSL certificate.
  * Constraints: The maximum length is `128` characters. The minimum length is `4` characters. The value must match regular expression `/(.*?)/`.

* `config_type` - (String) Th configuration type.
  * Constraints: Allowable values are: `public_cert_configuration_ca_lets_encrypt`, `public_cert_configuration_dns_classic_infrastructure`, `public_cert_configuration_dns_cloud_internet_services`, `iam_credentials_configuration`, `private_cert_configuration_root_ca`, `private_cert_configuration_intermediate_ca`, `private_cert_configuration_template`.

* `country` - (List) The Country (C) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `created_at` - (String) The date when a resource was created. The date format follows RFC 3339.

* `created_by` - (String) The unique identifier that is associated with the entity that created the secret.
  * Constraints: The maximum length is `128` characters. The minimum length is `4` characters.

* `crl_disable` - (Boolean) Disables or enables certificate revocation list (CRL) building.If CRL building is disabled, a signed but zero-length CRL is returned when downloading the CRL. If CRL building is enabled, it will rebuild the CRL.

* `crl_distribution_points_encoded` - (Boolean) Determines whether to encode the certificate revocation list (CRL) distribution points in the certificates that are issued by this certificate authority.

* `crl_expiry_seconds` - (Integer) The time until the certificate revocation list (CRL) expires, in seconds.

* `data` - (List) The configuration data of your Private Certificate.
Nested scheme for **data**:
	* `ca_chain` - (List) The chain of certificate authorities that are associated with the certificate.
	  * Constraints: The list items must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`. The maximum length is `16` items. The minimum length is `1` item.
	* `certificate` - (String) The PEM-encoded contents of your certificate.
	  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.
	* `csr` - (String) The certificate signing request.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `2` characters. The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.
	* `expiration` - (Integer) The certificate expiration time.
	* `issuing_ca` - (String) The PEM-encoded certificate of the certificate authority that signed and issued this certificate.
	  * Constraints: The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.
	* `private_key` - (String) (Optional) The PEM-encoded private key to associate with the certificate.
	  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.
	* `private_key_type` - (String) The type of private key to generate.
	  * Constraints: Allowable values are: `rsa`, `ec`.

* `exclude_cn_from_sans` - (Boolean) Controls whether the common name is excluded from Subject Alternative Names (SANs).If the common name set to `true`, it is not included in DNS or Email SANs if they apply. This field can be useful if the common name is a human-readable identifier, instead of a hostname or an email address.

* `expiration_date` - (String) The date a secret is expired. The date format follows RFC 3339.

* `format` - (String) The format of the returned data.
  * Constraints: Allowable values are: `pem`, `pem_bundle`.

* `ip_sans` - (String) The IP Subject Alternative Names to define for the CA certificate, in a comma-delimited list.
  * Constraints: The maximum length is `2048` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.

* `issuing_certificates_urls_encoded` - (Boolean) Determines whether to encode the URL of the issuing certificate in the certificates that are issued by this certificate authority.

* `key_bits` - (Integer) The number of bits to use to generate the private key.Allowable values for RSA keys are: `2048` and `4096`. Allowable values for EC keys are: `224`, `256`, `384`, and `521`. The default for RSA keys is `2048`. The default for EC keys is `256`.

* `key_type` - (String) The type of private key to generate.
  * Constraints: Allowable values are: `rsa`, `ec`.

* `locality` - (List) The Locality (L) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `max_path_length` - (Integer) The maximum path length to encode in the generated certificate. `-1` means no limit.If the signing certificate has a maximum path length set, the path length is set to one less than that of the signing certificate. A limit of `0` means a literal path length of zero.

* `max_ttl_seconds` - (Integer) The maximum time-to-live (TTL) for certificates that are created by this CA in seconds.

* `organization` - (List) The Organization (O) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `other_sans` - (List) The custom Object Identifier (OID) or UTF8-string Subject Alternative Names to define for the CA certificate.The alternative names must match the values that are specified in the `allowed_other_sans` field in the associated certificate template. The format is the same as OpenSSL: `<oid>:<type>:<value>` where the current valid type is `UTF8`.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.

* `ou` - (List) The Organizational Unit (OU) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `permitted_dns_domains` - (List) The allowed DNS domains or subdomains for the certificates that are to be signed and issued by this CA certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.

* `postal_code` - (List) The postal code values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `private_key_format` - (String) The format of the generated private key.
  * Constraints: The default value is `der`. Allowable values are: `der`, `pkcs8`.

* `province` - (List) The Province (ST) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `secret_type` - (String) The secret type. Supported types are arbitrary, certificates (imported, public, and private), IAM credentials, key-value, and user credentials.
  * Constraints: Allowable values are: `arbitrary`, `imported_cert`, `public_cert`, `iam_credentials`, `kv`, `username_password`, `private_cert`.

* `serial_number` - (String) The serial number to assign to the generated certificate. To assign a random serial number, you can omit this field.
  * Constraints: The maximum length is `64` characters. The minimum length is `32` characters. The value must match regular expression `/[^a-fA-F0-9]/`.

* `status` - (String) The status of the certificate authority. The status of a root certificate authority is either `configured` or `expired`. For intermediate certificate authorities, possible statuses include `signing_required`,`signed_certificate_required`, `certificate_template_required`, `configured`, `expired` or `revoked`.
  * Constraints: Allowable values are: `signing_required`, `signed_certificate_required`, `certificate_template_required`, `configured`, `expired`, `revoked`.

* `street_address` - (List) The street address values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `ttl` - (String) The requested time-to-live (TTL) for certificates that are created by this CA. This field's value cannot be longer than the `max_ttl` limit.The value can be supplied as a string representation of a duration in hours, for example '8760h'. In the API response, this value is returned in seconds (integer).
  * Constraints: The maximum length is `10` characters. The minimum length is `2` characters. The value must match regular expression `/^[0-9]+[s,m,h,d]{0,1}$/`.

* `updated_at` - (String) The date when a resource was recently modified. The date format follows RFC 3339.

* `uri_sans` - (String) The URI Subject Alternative Names to define for the CA certificate, in a comma-delimited list.
  * Constraints: The maximum length is `2048` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.

