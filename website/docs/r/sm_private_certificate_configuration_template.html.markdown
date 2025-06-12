---
layout: "ibm"
page_title: "IBM : ibm_sm_private_certificate_configuration_template"
description: |-
  Manages PrivateCertificateConfigurationTemplate.
subcategory: "Secrets Manager"
---

# ibm_sm_private_certificate_configuration_template

Provides a resource for a certificate template for private certificate secrets. This allows a certificate template to be created, updated and deleted. Note that a certificate template cannot be deleted if one or more private certificates exist that were generated with this template. Therefore, arguments that are marked as `Forces new resource` should not be modified if secrets generated with this template exist.

## Example Usage

```hcl
resource "ibm_sm_private_certificate_configuration_template" "certificate_template" {
  depends_on     = [ ibm_sm_private_certificate_configuration_intermediate_ca.intermediate_CA ]
  instance_id    = ibm_resource_instance.sm_instance.guid
  region                = "us-south"
  name                  = "my_template"
  certificate_authority = ibm_sm_private_certificate_configuration_intermediate_ca.intermediate_CA.name
  ou            = ["example_ou"]
  organization  = ["example_organization"]
  country       = ["US"]
  locality      = ["example_locality"]
  province      = ["example_province"]
  street_address  = ["example street address"]
  postal_code   = ["example_postal_code"]
  ttl           = "2190h"
  max_ttl       = "8760h"
  key_type      = "rsa"
  key_bits      = 4096
  allowed_domains    = ["example.com"]
  allow_any_name    = true
  allow_bare_domains = false
  allow_glob_domains = false
  allow_ip_sans      = true
  allow_localhost    = true
  allow_subdomains   = false
  allowed_domains_template = false
  allowed_other_sans = []
  allowed_uri_sans   = ["https://www.example.com/test"]
  enforce_hostnames  = false
  server_flag           = false
  client_flag           = false
  code_signing_flag     = false
  email_protection_flag = false
  key_usage           = ["DigitalSignature","KeyAgreement","KeyEncipherment"]
  use_csr_common_name = true
  use_csr_sans        = true
  require_cn          = true
  basic_constraints_valid_for_non_ca = false
  not_before_duration = "30s"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Optional, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
  * Constraints: Allowable values are: `private`, `public`.
* `allow_any_name` - (Optional, Boolean) Determines whether to allow clients to request a private certificate that matches any common name.
* `allow_bare_domains` - (Optional, Boolean) Determines whether to allow clients to request private certificates that match the value of the actual domains on the final certificate.For example, if you specify `example.com` in the `allowed_domains` field, you grant clients the ability to request a certificate that contains the name `example.com` as one of the DNS values on the final certificate.**Important:** In some scenarios, allowing bare domains can be considered a security risk.
* `allow_glob_domains` - (Optional, Boolean) Determines whether to allow glob patterns, for example, `ftp*.example.com`, in the names that are specified in the `allowed_domains` field.If set to `true`, clients are allowed to request private certificates with names that match the glob patterns.
* `allow_ip_sans` - (Optional, Boolean) Determines whether to allow clients to request a private certificate with IP Subject Alternative Names.
* `allow_localhost` - (Optional, Boolean) Determines whether to allow `localhost` to be included as one of the requested common names.
* `allow_subdomains` - (Optional, Boolean) Determines whether to allow clients to request private certificates with common names (CN) that are subdomains of the CNs that are allowed by the other certificate template options. This includes wildcard subdomains.For example, if `allowed_domains` has a value of `example.com` and `allow_subdomains`is set to `true`, then the following subdomains are allowed: `foo.example.com`, `bar.example.com`, `*.example.com`.**Note:** This field is redundant if you use the `allow_any_name` option.
* `allowed_domains` - (Optional, List) The domains to define for the certificate template. This property is used along with the `allow_bare_domains` and `allow_subdomains` options.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.
* `allowed_domains_template` - (Optional, Boolean) Determines whether to allow the domains that are supplied in the `allowed_domains` field to contain access control list (ACL) templates.
* `allowed_other_sans` - (Optional, List) The custom Object Identifier (OID) or UTF8-string Subject Alternative Names (SANs) to allow for private certificates.The format for each element in the list is the same as OpenSSL: `<oid>:<type>:<value>` where the current valid type is `UTF8`. To allow any value for an OID, use `*` as its value. Alternatively, specify a single `*` to allow any `other_sans` input.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.
* `allowed_secret_groups` - (Optional, String) Scopes the creation of private certificates to only the secret groups that you specify.This field can be supplied as a comma-delimited list of secret group IDs.
  * Constraints: The maximum length is `1024` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.
* `allowed_uri_sans` - (Optional, List) The URI Subject Alternative Names to allow for private certificates.Values can contain glob patterns, for example `spiffe://hostname/_*`.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.
* `basic_constraints_valid_for_non_ca` - (Optional, Boolean) Determines whether to mark the Basic Constraints extension of an issued private certificate as valid for non-CA certificates.
* `certificate_authority` - (Required, String) The name of the intermediate certificate authority.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9][A-Za-z0-9]*(?:_?-?\\.?[A-Za-z0-9]+)*$/`.
* `client_flag` - (Optional, Boolean) Determines whether private certificates are flagged for client use.
* `code_signing_flag` - (Optional, Boolean) Determines whether private certificates are flagged for code signing use.
* `config_type` - (Optional, String) Th configuration type.
  * Constraints: Allowable values are: `public_cert_configuration_ca_lets_encrypt`, `public_cert_configuration_dns_classic_infrastructure`, `public_cert_configuration_dns_cloud_internet_services`, `iam_credentials_configuration`, `private_cert_configuration_root_ca`, `private_cert_configuration_intermediate_ca`, `private_cert_configuration_template`.
* `country` - (Otional, Forces new resource, List) The Country (C) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `email_protection_flag` - (Optional, Boolean) Determines whether private certificates are flagged for email protection use.
* `enforce_hostnames` - (Optional, Boolean) Determines whether to enforce only valid host names for common names, DNS Subject Alternative Names, and the host section of email addresses.
* `ext_key_usage` - (Optional, List) The allowed extended key usage constraint on private certificates.You can find valid values in the [Go x509 package documentation](https://golang.org/pkg/crypto/x509/#ExtKeyUsage). Omit the `ExtKeyUsage` part of the value. Values are not case-sensitive. To specify no key usage constraints, set this field to an empty list.
  * Constraints: The list items must match regular expression `/^[a-zA-Z]+$/`. The maximum length is `100` items. The minimum length is `0` items.
* `ext_key_usage_oids` - (Optional, List) A list of extended key usage Object Identifiers (OIDs).
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.
* `key_bits` - (Optional, Forces new resource, Integer) The number of bits to use to generate the private key.Allowable values for RSA keys are: `2048` and `4096`. Allowable values for EC keys are: `224`, `256`, `384`, and `521`. The default for RSA keys is `2048`. The default for EC keys is `256`.
* `key_type` - (Optional, Forces new resource, String) The type of private key to generate.
  * Constraints: Allowable values are: `rsa`, `ec`.
* `key_usage` - (Optional, List) The allowed key usage constraint to define for private certificates.You can find valid values in the [Go x509 package documentation](https://pkg.go.dev/crypto/x509#KeyUsage). Omit the `KeyUsage` part of the value. Values are not case-sensitive. To specify no key usage constraints, set this field to an empty list.
  * Constraints: The list items must match regular expression `/^[a-zA-Z]+$/`. The maximum length is `100` items. The minimum length is `0` items.
* `locality` - (Optional, Forces new resource, List) The Locality (L) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `max_ttl` - (Optional, String) The maximum time-to-live (TTL) for certificates that are created by this template.
* `name` - (Required, String) A human-readable unique name to assign to your configuration.
* `organization` - (Optional, Forces new resource, List) The Organization (O) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `ou` - (Optional, Forces new resource, List) The Organizational Unit (OU) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `policy_identifiers` - (Optional, List) A list of policy Object Identifiers (OIDs).
* `postal_code` - (Optional, Forces new resource, List) The postal code values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `province` - (Optional, Forces new resource, List) The Province (ST) values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.
* `require_cn` - (Optional, Boolean) Determines whether to require a common name to create a private certificate.By default, a common name is required to generate a certificate. To make the `common_name` field optional, set the `require_cn` option to `false`.
* `server_flag` - (Optional, Boolean) Determines whether private certificates are flagged for server use.
* `serial_number` - (Optional, Forces new resource, String) Deprecated. Unused field. 
  * Constraints: The maximum length is `64` characters. The minimum length is `32` characters. The value must match regular expression `/[^a-fA-F0-9]/`.
* `street_address` - (Optional, Forces new resource, List) The street address values to define in the subject field of the resulting certificate.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `ttl` - The requested time-to-live (TTL) for certificates that are created by this template. This field's value can't be longer than the max_ttl limit.
* `use_csr_common_name` - (Optional, Boolean) When used with the `private_cert_configuration_action_sign_csr` action, this field determines whether to use the common name (CN) from a certificate signing request (CSR) instead of the CN that's included in the data of the certificate.Does not include any requested Subject Alternative Names (SANs) in the CSR. To use the alternative names, include the `use_csr_sans` property.
* `use_csr_sans` - (Optional, Boolean) When used with the `private_cert_configuration_action_sign_csr` action, this field determines whether to use the Subject Alternative Names(SANs) from a certificate signing request (CSR) instead of the SANs that are included in the data of the certificate.Does not include the common name in the CSR. To use the common name, include the `use_csr_common_name` property.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the PrivateCertificateConfigurationTemplate.
* `created_at` - (String) The date when a resource was created. The date format follows RFC 3339.
* `created_by` - (String) The unique identifier that is associated with the entity that created the secret.
  * Constraints: The maximum length is `128` characters. The minimum length is `4` characters.
* `max_ttl_seconds` - (Integer) The maximum time-to-live (TTL) for certificates that are created by this CA in seconds.
* `not_before_duration_seconds` - (Integer) The duration in seconds by which to backdate the `not_before` property of an issued private certificate.
* `secret_type` - (String) The secret type. Supported types are arbitrary, certificates (imported, public, and private), IAM credentials, key-value, and user credentials.
  * Constraints: Allowable values are: `arbitrary`, `imported_cert`, `public_cert`, `iam_credentials`, `kv`, `username_password`, `private_cert`.
* `ttl_seconds` - (Integer) The requested Time To Live, after which the certificate will be expired.
* `updated_at` - (String) The date when a resource was recently modified. The date format follows RFC 3339.

## Provider Configuration

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

To find which credentials are required for this resource, see the service table [here](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters).

### Static credentials

You can provide your static credentials by adding the `ibmcloud_api_key`, `iaas_classic_username`, and `iaas_classic_api_key` arguments in the IBM Cloud provider block.

Usage:
```
provider "ibm" {
    ibmcloud_api_key = ""
    iaas_classic_username = ""
    iaas_classic_api_key = ""
}
```

### Environment variables

You can provide your credentials by exporting the `IC_API_KEY`, `IAAS_CLASSIC_USERNAME`, and `IAAS_CLASSIC_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud Classic Infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```
provider "ibm" {}
```

Usage:
```
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
terraform plan
```

Note:

1. Create or find your `ibmcloud_api_key` and `iaas_classic_api_key` [here](https://cloud.ibm.com/iam/apikeys).
  - Select `My IBM Cloud API Keys` option from view dropdown for `ibmcloud_api_key`
  - Select `Classic Infrastructure API Keys` option from view dropdown for `iaas_classic_api_key`
2. For iaas_classic_username
  - Go to [Users](https://cloud.ibm.com/iam/users)
  - Click on user.
  - Find user name in the `VPN password` section under `User Details` tab

For more informaton, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).

## Import

You can import the `ibm_sm_private_certificate_configuration_template` resource by using `region`, `instance_id`, and `name`.
For more information, see [the documentation](https://cloud.ibm.com/docs/secrets-manager)

# Syntax
```bash
$ terraform import ibm_sm_private_certificate_configuration_template.sm_private_certificate_configuration_template <region>/<instance_id>/<name>
```

# Example
```bash
$ terraform import ibm_sm_private_certificate_configuration_template.sm_private_certificate_configuration_template us-east/6ebc4224-e983-496a-8a54-f40a0bfa9175/my_template
```
