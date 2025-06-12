---
layout: "ibm"
page_title: "IBM : ibm_sm_private_certificate_configuration_action_sign_csr"
description: |-
  Manages PrivateCertificateConfigurationActionSignCsr.
subcategory: "Secrets Manager"
---

# ibm_sm_private_certificate_configuration_action_sign_csr

Provides a resource for PrivateCertificateConfigurationActionSignCsr. This allows PrivateCertificateConfigurationActionSignCsr to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_sm_private_certificate_configuration_action_sign_csr" "sign_csr_action" {
  instance_id           = "6ebc4224-e983-496a-8a54-f40a0bfa9175"
  region                = "us-south"
  name    = "my_configuration"
  csr                   = "-----BEGIN CERTIFICATE REQUEST-----\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\nAAAAAAAAAAAAAAAAAAAAAA==\n-----END CERTIFICATE REQUEST-----"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Required, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
    * Constraints: Allowable values are: `private`, `public`.
* `name` - (Required, Forces new resource, String) The name that uniquely identifies the configuration that will be used to sign the CSR.
* `csr` - (Required, Forces new resource, String) The certificate signing request.
* `common_name` - (Optional, Forces new resource, String) The Common Name (AKA CN) represents the server name that is protected by the SSL certificate.
    * Constraints: The maximum length is `128` characters. The minimum length is `4` characters. The value must match regular expression `/(.*?)/`.
* `alt_names` - (Optional, Forces new resource, List) With the Subject Alternative Name field, you can specify additional host names to be protected by a single SSL certificate.
    * Constraints: The list items must match regular expression `/^(.*?)$/`. The maximum length is `99` items. The minimum length is `0` items.
* `ip_sans` - (Optional, Forces new resource, String) The IP Subject Alternative Names to define for the CA certificate, in a comma-delimited list.
    * Constraints: The maximum length is `2048` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.
* `uri_sans` - (Optional, Forces new resource, String) The URI Subject Alternative Names to define for the CA certificate, in a comma-delimited list.
    * Constraints: The maximum length is `2048` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.
* `other_sans` - (Optional, Forces new resource, List) The custom Object Identifier (OID) or UTF8-string Subject Alternative Names to define for the CA certificate.The alternative names must match the values that are specified in the `allowed_other_sans` field in the associated certificate template. The format is the same as OpenSSL: `<oid>:<type>:<value>` where the current valid type is `UTF8`.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.
* `ttl` - (Optional, Forces new resource, String) The requested time-to-live (TTL) for certificates that are created by this CA. This field's value cannot be longer than the `max_ttl` limit.The value can be supplied as a string representation of a duration in hours, for example '8760h'. In the API response, this value is returned in seconds (integer).
    * Constraints: The maximum length is `10` characters. The minimum length is `2` characters. The value must match regular expression `/^[0-9]+[s,m,h,d]{0,1}$/`.
* `format` - (Optional, Forces new resource, String) The format of the returned data.
    * Constraints: Allowable values are: `pem`, `pem_bundle`.
* `max_path_length` - (Optional, Forces new resource, Integer) The maximum path length to encode in the generated certificate. `-1` means no limit.If the signing certificate has a maximum path length set, the path length is set to one less than that of the signing certificate. A limit of `0` means a literal path length of zero.
* `exclude_cn_from_sans` - (Optional, Forces new resource, Boolean) Controls whether the common name is excluded from Subject Alternative Names (SANs).If the common name set to `true`, it is not included in DNS or Email SANs if they apply. This field can be useful if the common name is a human-readable identifier, instead of a hostname or an email address.
* `permitted_dns_domains` - (Optional, Forces new resource, List) The allowed DNS domains or subdomains for the certificates that are to be signed and issued by this CA certificate.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.
* `use_csr_values` - (Optional, Forces new resource, Boolean) Determines whether to use values from a certificate signing request (CSR) to complete a `private_cert_configuration_action_sign_csr` action.
* `ou` - (Optional, Forces new resource, List) The Organizational Unit (OU) values to define in the subject field of the resulting certificate.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `organization` - (Optional, Forces new resource, List) The Organization (O) values to define in the subject field of the resulting certificate.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `country` - (Optional, Forces new resource, List) The Country (C) values to define in the subject field of the resulting certificate.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `locality` - (Optional, Forces new resource, List) The Locality (L) values to define in the subject field of the resulting certificate.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `province` - (Optional, Forces new resource, List) The Province (ST) values to define in the subject field of the resulting certificate.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `street_address` - (Optional, Forces new resource, List) The street address values to define in the subject field of the resulting certificate.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `postal_code` - (Optional, Forces new resource, List) The postal code values to define in the subject field of the resulting certificate.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `serial_number` - (Optional, Forces new resource, String) The serial number to assign to the generated certificate. To assign a random serial number, you can omit this field.
    * Constraints: The maximum length is `64` characters. The minimum length is `32` characters. The value must match regular expression `/[^a-fA-F0-9]/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `data` - (List) The configuration data of your Private Certificate.
  Nested scheme for **data**:
    * `certificate` - (String) The PEM-encoded contents of your certificate.
        * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.
    * `issuing_ca` - (String) The PEM-encoded certificate of the certificate authority that signed and issued this certificate.
        * Constraints: The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.
    * `ca_chain` - (List) The chain of certificate authorities that are associated with the certificate.
        * Constraints: The list items must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`. The maximum length is `16` items. The minimum length is `1` item.
    * `expiration` - (Integer) The certificate expiration time.
