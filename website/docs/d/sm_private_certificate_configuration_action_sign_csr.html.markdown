---
layout: "ibm"
page_title: "IBM : ibm_sm_private_certificate_configuration_action_sign_csr" (Beta)
description: |-
  Get information about PrivateCertificateConfigurationActionSignCSR
subcategory: "IBM Cloud Secrets Manager API"
---

# ibm_sm_private_certificate_configuration_action_sign_csr

Provides a read-only data source for PrivateCertificateConfigurationActionSignCSR. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_private_certificate_configuration_action_sign_csr" {
  instance_id   = "6ebc4224-e983-496a-8a54-f40a0bfa9175"
  region        = "us-south"
  config_action_prototype = {"action_type":"private_cert_configuration_action_sign_csr"}
  name = "configuration-name"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `config_action_prototype` - (Required, List) The request body to specify the properties of the action to create a configuration.
Nested scheme for **config_action_prototype**:
	* `action_type` - (Optional, String) The type of configuration action.
	  * Constraints: Allowable values are: `private_cert_configuration_action_rotate_crl`, `private_cert_configuration_action_sign_intermediate`, `private_cert_configuration_action_sign_csr`, `private_cert_configuration_action_set_signed`, `private_cert_configuration_action_revoke_ca_certificate`.
	* `alt_names` - (Computed, Forces new resource, List) With the Subject Alternative Name field, you can specify additional host names to be protected by a single SSL certificate.
	  * Constraints: The list items must match regular expression `/^(.*?)$/`. The maximum length is `99` items. The minimum length is `0` items.
	* `certificate` - (Computed, Forces new resource, String) The PEM-encoded contents of your certificate.
	  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.
	* `common_name` - (Computed, Forces new resource, String) The Common Name (AKA CN) represents the server name that is protected by the SSL certificate.
	  * Constraints: The maximum length is `128` characters. The minimum length is `4` characters. The value must match regular expression `/(.*?)/`.
	* `country` - (Computed, Forces new resource, List) The Country (C) values to define in the subject field of the resulting certificate.
	  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
	* `csr` - (Computed, Forces new resource, String) The certificate signing request.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `2` characters. The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.
	* `exclude_cn_from_sans` - (Computed, Forces new resource, Boolean) Controls whether the common name is excluded from Subject Alternative Names (SANs).If the common name set to `true`, it is not included in DNS or Email SANs if they apply. This field can be useful if the common name is a human-readable identifier, instead of a hostname or an email address.
	* `format` - (Computed, Forces new resource, String) The format of the returned data.
	  * Constraints: Allowable values are: `pem`, `pem_bundle`.
	* `intermediate_certificate_authority` - (Computed, String) The unique name of your configuration.
	  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.
	* `ip_sans` - (Computed, Forces new resource, String) The IP Subject Alternative Names to define for the CA certificate, in a comma-delimited list.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.
	* `locality` - (Computed, Forces new resource, List) The Locality (L) values to define in the subject field of the resulting certificate.
	  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
	* `max_path_length` - (Computed, Forces new resource, Integer) The maximum path length to encode in the generated certificate. `-1` means no limit.If the signing certificate has a maximum path length set, the path length is set to one less than that of the signing certificate. A limit of `0` means a literal path length of zero.
	* `organization` - (Computed, Forces new resource, List) The Organization (O) values to define in the subject field of the resulting certificate.
	  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
	* `other_sans` - (Computed, Forces new resource, List) The custom Object Identifier (OID) or UTF8-string Subject Alternative Names to define for the CA certificate.The alternative names must match the values that are specified in the `allowed_other_sans` field in the associated certificate template. The format is the same as OpenSSL: `<oid>:<type>:<value>` where the current valid type is `UTF8`.
	  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.
	* `ou` - (Computed, Forces new resource, List) The Organizational Unit (OU) values to define in the subject field of the resulting certificate.
	  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
	* `permitted_dns_domains` - (Computed, Forces new resource, List) The allowed DNS domains or subdomains for the certificates that are to be signed and issued by this CA certificate.
	  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.
	* `postal_code` - (Computed, Forces new resource, List) The postal code values to define in the subject field of the resulting certificate.
	  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
	* `province` - (Computed, Forces new resource, List) The Province (ST) values to define in the subject field of the resulting certificate.
	  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
	* `serial_number` - (Computed, Forces new resource, String) The serial number to assign to the generated certificate. To assign a random serial number, you can omit this field.
	  * Constraints: The maximum length is `64` characters. The minimum length is `32` characters. The value must match regular expression `/[^a-fA-F0-9]/`.
	* `street_address` - (Computed, Forces new resource, List) The street address values to define in the subject field of the resulting certificate.
	  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
	* `ttl` - (Optional, Forces new resource, String) The time-to-live (TTL) to assign to a private certificate.The value can be supplied as a string representation of a duration in hours, for example '12h'. The value can't exceed the `max_ttl` that is defined in the associated certificate template.
	  * Constraints: The maximum length is `10` characters. The minimum length is `2` characters. The value must match regular expression `/^[0-9]+[s,m,h,d]{0,1}$/`.
	* `uri_sans` - (Computed, Forces new resource, String) The URI Subject Alternative Names to define for the CA certificate, in a comma-delimited list.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.
	* `use_csr_values` - (Optional, Boolean) Determines whether to use values from a certificate signing request (CSR) to complete a `private_cert_configuration_action_sign_csr` action. If it is set to `true`, then:1) Subject information, including names and alternate names, are preserved from the CSR rather than by using the values that are provided in the other parameters to this operation.2) Any key usage, for example, non-repudiation, that are requested in the CSR are added to the basic set of key usages used for CA certificates that are signed by the intermediate authority.3) Extensions that are requested in the CSR are copied into the issued private certificate.
* `name` - (Required, Forces new resource, String) The name of the configuration.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9][A-Za-z0-9]*(?:_?-?\\.?[A-Za-z0-9]+)*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the PrivateCertificateConfigurationActionSignCSR.
* `action_type` - (String) The type of configuration action.
  * Constraints: Allowable values are: `private_cert_configuration_action_rotate_crl`, `private_cert_configuration_action_sign_intermediate`, `private_cert_configuration_action_sign_csr`, `private_cert_configuration_action_set_signed`, `private_cert_configuration_action_revoke_ca_certificate`.

* `alt_names` - (Forces new resource, List) With the Subject Alternative Name field, you can specify additional host names to be protected by a single SSL certificate.
  * Constraints: The list items must match regular expression `/^(.*?)$/`. The maximum length is `99` items. The minimum length is `0` items.

* `common_name` - (Forces new resource, String) The Common Name (AKA CN) represents the server name that is protected by the SSL certificate.
  * Constraints: The maximum length is `128` characters. The minimum length is `4` characters. The value must match regular expression `/(.*?)/`.

* `country` - (Forces new resource, List) The Country (C) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `csr` - (Forces new resource, String) The certificate signing request.
  * Constraints: The maximum length is `4096` characters. The minimum length is `2` characters. The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.

* `data` - (List) The data that is associated with the root certificate authority.
Nested scheme for **data**:
	* `ca_chain` - (List) The chain of certificate authorities that are associated with the certificate.
	  * Constraints: The list items must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`. The maximum length is `16` items. The minimum length is `1` item.
	* `certificate` - (Forces new resource, String) The PEM-encoded contents of your certificate.
	  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.
	* `expiration` - (Integer) The certificate expiration time.
	* `issuing_ca` - (String) The PEM-encoded certificate of the certificate authority that signed and issued this certificate.
	  * Constraints: The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.

* `exclude_cn_from_sans` - (Forces new resource, Boolean) Controls whether the common name is excluded from Subject Alternative Names (SANs).If the common name set to `true`, it is not included in DNS or Email SANs if they apply. This field can be useful if the common name is a human-readable identifier, instead of a hostname or an email address.

* `format` - (Forces new resource, String) The format of the returned data.
  * Constraints: Allowable values are: `pem`, `pem_bundle`.

* `ip_sans` - (Forces new resource, String) The IP Subject Alternative Names to define for the CA certificate, in a comma-delimited list.
  * Constraints: The maximum length is `2048` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.

* `locality` - (Forces new resource, List) The Locality (L) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `max_path_length` - (Forces new resource, Integer) The maximum path length to encode in the generated certificate. `-1` means no limit.If the signing certificate has a maximum path length set, the path length is set to one less than that of the signing certificate. A limit of `0` means a literal path length of zero.

* `organization` - (Forces new resource, List) The Organization (O) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `other_sans` - (Forces new resource, List) The custom Object Identifier (OID) or UTF8-string Subject Alternative Names to define for the CA certificate.The alternative names must match the values that are specified in the `allowed_other_sans` field in the associated certificate template. The format is the same as OpenSSL: `<oid>:<type>:<value>` where the current valid type is `UTF8`.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.

* `ou` - (Forces new resource, List) The Organizational Unit (OU) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `permitted_dns_domains` - (Forces new resource, List) The allowed DNS domains or subdomains for the certificates that are to be signed and issued by this CA certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.

* `postal_code` - (Forces new resource, List) The postal code values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `province` - (Forces new resource, List) The Province (ST) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `serial_number` - (Forces new resource, String) The serial number to assign to the generated certificate. To assign a random serial number, you can omit this field.
  * Constraints: The maximum length is `64` characters. The minimum length is `32` characters. The value must match regular expression `/[^a-fA-F0-9]/`.

* `street_address` - (Forces new resource, List) The street address values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.

* `ttl` - (Forces new resource, String) The time-to-live (TTL) to assign to a private certificate.The value can be supplied as a string representation of a duration in hours, for example '12h'. The value can't exceed the `max_ttl` that is defined in the associated certificate template.
  * Constraints: The maximum length is `10` characters. The minimum length is `2` characters. The value must match regular expression `/^[0-9]+[s,m,h,d]{0,1}$/`.

* `uri_sans` - (Forces new resource, String) The URI Subject Alternative Names to define for the CA certificate, in a comma-delimited list.
  * Constraints: The maximum length is `2048` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.

* `use_csr_values` - (Boolean) Determines whether to use values from a certificate signing request (CSR) to complete a `private_cert_configuration_action_sign_csr` action. If it is set to `true`, then:1) Subject information, including names and alternate names, are preserved from the CSR rather than by using the values that are provided in the other parameters to this operation.2) Any key usage, for example, non-repudiation, that are requested in the CSR are added to the basic set of key usages used for CA certificates that are signed by the intermediate authority.3) Extensions that are requested in the CSR are copied into the issued private certificate.

