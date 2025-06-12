---
layout: "ibm"
page_title: "IBM : ibm_sm_private_certificate_configuration_root_ca"
description: |-
  Manages PrivateCertificateConfigurationRootCA.
subcategory: "Secrets Manager"
---

# ibm_sm_private_certificate_configuration_root_ca

Provides a resource for an internally signed root certificate authority. This allows a root CA to be created, updated and deleted. Note that a root CA cannot be deleted if there are intermediate CAs signed by it. Therefore, arguments that are marked as `Forces new resource` should not be modified if there are dependent intermediate CAs.

## Example Usage

```hcl
resource "ibm_sm_private_certificate_configuration_root_ca" "private_certificate_root_CA" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
  name          = "my_root_ca"
  common_name   = "ibm.com"
  alt_names     = ["alt-name-1", "alt-name-2"]
  permitted_dns_domains = ["exampleString"]
  ou            = ["example_ou"]
  organization  = ["example_organization"]
  country       = ["US"]
  locality      = ["example_locality"]
  province      = ["example_province"]
  street_address  = ["example street address"]
  postal_code   = ["example_postal_code"]
  ip_sans       = "127.0.0.1"
  uri_sans      = "https://www.example.com/test"
  other_sans    = ["1.2.3.5.4.3.201.10.4.3;utf8:test@example.com"]
  exclude_cn_from_sans = false
  ttl           = "2190h"
  max_ttl       = "8760h"
  max_path_length = -1
  issuing_certificates_urls_encoded = true
  key_type      = "rsa"
  key_bits      = 4096
  format        = "pem"
  private_key_format = "der"
  crl_expiry    = "72h"
  crl_disable   = false
  crl_distribution_points_encoded   = true
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Optional, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
    * Constraints: Allowable values are: `private`, `public`.
* `alt_names` - (Optional, Forces new resource, List) With the Subject Alternative Name field, you can specify additional host names to be protected by a single SSL certificate.
    * Constraints: The list items must match regular expression `/^(.*?)$/`. The maximum length is `99` items. The minimum length is `0` items.
* `common_name` - (Required, Forces new resource, String) The Common Name (AKA CN) represents the server name that is protected by the SSL certificate.
    * Constraints: The maximum length is `128` characters. The minimum length is `4` characters. The value must match regular expression `/(.*?)/`.
* `country` - (Optional, Forces new resource, List) The Country (C) values to define in the subject field of the resulting certificate.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `crl_disable` - (Optional, Boolean) Disables or enables certificate revocation list (CRL) building.If CRL building is disabled, a signed but zero-length CRL is returned when downloading the CRL. If CRL building is enabled, it will rebuild the CRL.
* `crl_distribution_points_encoded` - (Optional, Boolean) Determines whether to encode the certificate revocation list (CRL) distribution points in the certificates that are issued by this certificate authority.
* `crl_expiry` - (Optional, Boolean) The time until the certificate revocation list (CRL) expires.The value can be supplied as a string representation of a duration in hours, such as `48h`. The default is 72 hours.
* `exclude_cn_from_sans` - (Optional, Forces new resource, Boolean) Controls whether the common name is excluded from Subject Alternative Names (SANs).If the common name set to `true`, it is not included in DNS or Email SANs if they apply. This field can be useful if the common name is a human-readable identifier, instead of a hostname or an email address.
* `expiration_date` - (Optional, Forces new resource, String) The date a secret is expired. The date format follows RFC 3339.
* `format` - (Optional, Forces new resource, String) The format of the returned data.
    * Constraints: Allowable values are: `pem`, `pem_bundle`.
* `ip_sans` - (Optional, Forces new resource, String) The IP Subject Alternative Names to define for the CA certificate, in a comma-delimited list.
    * Constraints: The maximum length is `2048` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.
* `issuing_certificates_urls_encoded` - (Optional, Boolean) Determines whether to encode the URL of the issuing certificate in the certificates that are issued by this certificate authority.
* `key_bits` - (Optional, Forces new resource, Integer) The number of bits to use to generate the private key.Allowable values for RSA keys are: `2048` and `4096`. Allowable values for EC keys are: `224`, `256`, `384`, and `521`. The default for RSA keys is `2048`. The default for EC keys is `256`.
* `key_type` - (Optional, Forces new resource, String) The type of private key to generate.
    * Constraints: Allowable values are: `rsa`, `ec`.
* `locality` - (Optional, Forces new resource, List) The Locality (L) values to define in the subject field of the resulting certificate.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `max_path_length` - (Optional, Forces new resource, Integer) The maximum path length to encode in the generated certificate. `-1` means no limit.If the signing certificate has a maximum path length set, the path length is set to one less than that of the signing certificate. A limit of `0` means a literal path length of zero.
* `max_ttl` - (Required, String) The maximum time-to-live (TTL) for certificates that are created by this CA.
* `name` - (Required, String) A human-readable unique name to assign to the root CA configuration.
* `organization` - (Optional, Forces new resource, List) The Organization (O) values to define in the subject field of the resulting certificate.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `other_sans` - (Optional, Forces new resource, List) The custom Object Identifier (OID) or UTF8-string Subject Alternative Names to define for the CA certificate.The alternative names must match the values that are specified in the `allowed_other_sans` field in the associated certificate template. The format is the same as OpenSSL: `<oid>:<type>:<value>` where the current valid type is `UTF8`.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.
* `ou` - (Optional, Forces new resource, List) The Organizational Unit (OU) values to define in the subject field of the resulting certificate.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `permitted_dns_domains` - (Optional, Forces new resource, List) The allowed DNS domains or subdomains for the certificates that are to be signed and issued by this CA certificate.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `100` items. The minimum length is `0` items.
* `postal_code` - (Optional, Forces new resource, List) The postal code values to define in the subject field of the resulting certificate.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `private_key_format` - (Optional, Forces new resource, String) The format of the generated private key.
    * Constraints: The default value is `der`. Allowable values are: `der`, `pkcs8`.
* `province` - (Optional, Forces new resource, List) The Province (ST) values to define in the subject field of the resulting certificate.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `street_address` - (Optional, Forces new resource, List) The street address values to define in the subject field of the resulting certificate.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `10` items. The minimum length is `0` items.
* `ttl` - (Optional, String) The requested time-to-live (TTL) for certificates that are created by this CA. This field's value cannot be longer than the `max_ttl` limit.The value can be supplied as a string representation of a duration in hours, for example '8760h'. In the API response, this value is returned in seconds (integer).
  * Constraints: The maximum length is `10` characters. The minimum length is `2` characters. The value must match regular expression `/^[0-9]+[s,m,h,d]{0,1}$/`.
* `uri_sans` - (Optional, Forces new resource, String) The URI Subject Alternative Names to define for the CA certificate, in a comma-delimited list.
    * Constraints: The maximum length is `2048` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.
* `crypto_key` - (Optional, Forces new resource, List) The data that is associated with a cryptographic key.
  Nested scheme for **crypto_key**:
    * `provider` - (Required, Forces new resource, List) The data that is associated with a cryptographic provider.
      Nested scheme for **provider**:
        * `type` - (Required, Forces new resource, String) The type of cryptographic provider.
            * Constraints: Allowable values are: `hyper_protect_crypto_services`.
        * `instance_crn` - (Required, Forces new resource, String) The HPCS instance CRN.
            * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@/]|%[0-9A-Z]{2})*){8}$`.
        * `pin_iam_credentials_secret_id` - (Required, Forces new resource, String) The secret Id of iam credentials with api key to access HPCS instance.
            * Constraints: Value length should be 36. The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
        * `private_keystore_id` - (Required, Forces new resource, String) The HPCS private key store space id.
            * Constraints: Value length should be 36. The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
    * `id` - (Optional, Forces new resource, String) The ID of a PKCS#11 key to use. If the key does not exist and generation is enabled, this ID is given to the generated key. If the key exists, and generation is disabled, then this ID is used to look up the key. This value or the crypto key label must be specified.
        * Constraints: Value length should be 36. The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
    * `label` - (Optional, Forces new resource, String) The label of the key to use. If the key does not exist and generation is enabled, this field is the label that is given to the generated key. If the key exists, and generation is disabled, then this label is used to look up the key. This value or the crypto key ID must be specified.
        * Constraints: The maximum length is `255` characters. The minimum length is `1` characters. The value must match regular expression `/^[A-Za-z0-9._ /-]+$/`.
    * `allow_generate_key` - (Optional, Forces new resource, Boolean) The indication of whether a new key is generated by the crypto provider if the given key name cannot be found. Default is `false`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the PrivateCertificateConfigurationRootCA.
* `config_type` - (String) The configuration type.
    * Constraints: Allowable values are: `public_cert_configuration_ca_lets_encrypt`, `public_cert_configuration_dns_classic_infrastructure`, `public_cert_configuration_dns_cloud_internet_services`, `iam_credentials_configuration`, `private_cert_configuration_root_ca`, `private_cert_configuration_intermediate_ca`, `private_cert_configuration_template`.
* `created_at` - (String) The date when a resource was created. The date format follows RFC 3339.
* `created_by` - (String) The unique identifier that is associated with the entity that created the secret.
  * Constraints: The maximum length is `128` characters. The minimum length is `4` characters.
* `crl_expiry_seconds` - (Integer) The time until the certificate revocation list (CRL) expires, in seconds.
* `data` - (List) The configuration data of your Private Certificate.
Nested scheme for **data**:
	* `ca_chain` - (List) The chain of certificate authorities that are associated with the certificate.
	  * Constraints: The list items must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`. The maximum length is `16` items. The minimum length is `1` item.
	* `certificate` - (Forces new resource, String) The PEM-encoded contents of your certificate.
	  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.
	* `csr` - (Forces new resource, String) The certificate signing request.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `2` characters. The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.
	* `expiration` - (Integer) The certificate expiration time.
	* `issuing_ca` - (String) The PEM-encoded certificate of the certificate authority that signed and issued this certificate.
	  * Constraints: The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.
	* `private_key` - (Forces new resource, String) (Optional) The PEM-encoded private key to associate with the certificate.
	  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/^(-{5}BEGIN.+?-{5}[\\s\\S]+-{5}END.+?-{5})$/`.
	* `private_key_type` - (Forces new resource, String) The type of private key to generate.
	  * Constraints: Allowable values are: `rsa`, `ec`.
* `max_ttl_seconds` - (Integer) The maximum time-to-live (TTL) for certificates that are created by this CA in seconds.
* `secret_type` - (String) The secret type. Supported types are arbitrary, certificates (imported, public, and private), IAM credentials, key-value, and user credentials.
  * Constraints: Allowable values are: `arbitrary`, `imported_cert`, `public_cert`, `iam_credentials`, `kv`, `username_password`, `private_cert`.
* `serial_number` - (String) The unique serial number that was assigned to a certificate by the issuing certificate authority.
    * Constraints: The maximum length is `64` characters. The minimum length is `32` characters. The value must match regular expression `/[^a-fA-F0-9]/`.
* `status` - (String) The status of the certificate authority. The status of a root certificate authority is either `configured` or `expired`. For intermediate certificate authorities, possible statuses include `signing_required`,`signed_certificate_required`, `certificate_template_required`, `configured`, `expired` or `revoked`.
  * Constraints: Allowable values are: `signing_required`, `signed_certificate_required`, `certificate_template_required`, `configured`, `expired`, `revoked`.
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

You can import the `ibm_sm_private_certificate_configuration_root_ca` resource by using `region`, `instance_id`, and `name`.
For more information, see [the documentation](https://cloud.ibm.com/docs/secrets-manager)

# Syntax
```bash
$ terraform import ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca <region>/<instance_id>/<name>
```

# Example
```bash
$ terraform import ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca us-east/6ebc4224-e983-496a-8a54-f40a0bfa9175/my_root_ca
```
