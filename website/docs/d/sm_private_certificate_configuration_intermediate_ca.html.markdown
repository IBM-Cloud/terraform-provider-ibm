---
layout: "ibm"
page_title: "IBM : ibm_sm_private_certificate_configuration_intermediate_ca"
description: |-
  Get information about PrivateCertificateConfigurationIntermediateCA
subcategory: "Secrets Manager"
---

# ibm_sm_private_certificate_configuration_intermediate_ca

Provides a read-only data source for the configuraion of an intermediate CA. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_private_certificate_configuration_intermediate_ca" "intermediate_ca" {
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

* `issuer` - (String) The distinguished name that identifies the entity that signed and issued the certificate.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.

* `issuing_certificates_urls_encoded` - (Boolean) Determines whether to encode the URL of the issuing certificate in the certificates that are issued by this certificate authority.

* `key_bits` - (Integer) The number of bits to use to generate the private key.Allowable values for RSA keys are: `2048` and `4096`. Allowable values for EC keys are: `224`, `256`, `384`, and `521`. The default for RSA keys is `2048`. The default for EC keys is `256`.

* `key_type` - (String) The type of private key to generate.
  * Constraints: Allowable values are: `rsa`, `ec`.

* `locality` - (List) The Locality (L) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `max_ttl_seconds` - (Integer) The maximum time-to-live (TTL) for certificates that are created by this CA in seconds.

* `organization` - (List) The Organization (O) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `other_sans` - (List) The custom Object Identifier (OID) or UTF8-string Subject Alternative Names to define for the CA certificate.The alternative names must match the values that are specified in the `allowed_other_sans` field in the associated certificate template. The format is the same as OpenSSL: `<oid>:<type>:<value>` where the current valid type is `UTF8`.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.

* `ou` - (List) The Organizational Unit (OU) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

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

* `signing_method` - (String) The signing method to use with this certificate authority to generate private certificates.You can choose between internal or externally signed options. For more information, see the [docs](https://cloud.ibm.com/docs/secrets-manager?topic=secrets-manager-intermediate-certificate-authorities).
  * Constraints: Allowable values are: `internal`, `external`.

* `status` - (String) The status of the certificate authority. The status of a root certificate authority is either `configured` or `expired`. For intermediate certificate authorities, possible statuses include `signing_required`,`signed_certificate_required`, `certificate_template_required`, `configured`, `expired` or `revoked`.
  * Constraints: Allowable values are: `signing_required`, `signed_certificate_required`, `certificate_template_required`, `configured`, `expired`, `revoked`.

* `street_address` - (List) The street address values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `updated_at` - (String) The date when a resource was recently modified. The date format follows RFC 3339.

* `uri_sans` - (String) The URI Subject Alternative Names to define for the CA certificate, in a comma-delimited list.
  * Constraints: The maximum length is `2048` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.

* `crypto_key` - (List) The data that is associated with a cryptographic key.
  Nested scheme for **crypto_key**:
     * `provider` - (List) The data that is associated with a cryptographic provider.
        Nested scheme for **provider**:
          * `type` - (String) The type of cryptographic provider.
            * Constraints: Allowable values are: `hyper_protect_crypto_services`.
          * `instance_crn` - (String) The HPCS instance CRN.
             * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@/]|%[0-9A-Z]{2})*){8}$`.
          * `pin_iam_credentials_secret_id` - (String) The secret Id of iam credentials with api key to access HPCS instance.
             * Constraints: Value length should be 36. The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
          * `private_keystore_id` - (String) The HPCS private key store space id.
             * Constraints: Value length should be 36. The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
     * `id` - (String) The ID of a PKCS#11 key to use. If the key does not exist and generation is enabled, this ID is given to the generated key. If the key exists, and generation is disabled, then this ID is used to look up the key. This value or the crypto key label must be specified.
       * Constraints: Value length should be 36. The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
     * `label` - (String) The label of the key to use. If the key does not exist and generation is enabled, this field is the label that is given to the generated key. If the key exists, and generation is disabled, then this label is used to look up the key. This value or the crypto key ID must be specified.
       * Constraints: The maximum length is `255` characters. The minimum length is `1` characters. The value must match regular expression `/^[A-Za-z0-9._ /-]+$/`.
     * `allow_generate_key` - (Boolean) The indication of whether a new key is generated by the crypto provider if the given key name cannot be found. Default is `false`.

