---
layout: "ibm"
page_title: "IBM : ibm_sm_public_certificate_configuration_ca_lets_encrypt"
description: |-
  Get information about PublicCertificateConfigurationCALetsEncrypt
subcategory: "Secrets Manager"
---

# ibm_sm_public_certificate_configuration_ca_lets_encrypt

Provides a read-only data source for a Let's Encrypt CA configuration. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_public_certificate_configuration_ca_lets_encrypt" "ca_lets_encrypt" {
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
* `config_type` - (String) Th configuration type.
  * Constraints: Allowable values are: `public_cert_configuration_ca_lets_encrypt`, `public_cert_configuration_dns_classic_infrastructure`, `public_cert_configuration_dns_cloud_internet_services`, `iam_credentials_configuration`, `private_cert_configuration_root_ca`, `private_cert_configuration_intermediate_ca`, `private_cert_configuration_template`.

* `created_at` - (String) The date when a resource was created. The date format follows RFC 3339.

* `created_by` - (String) The unique identifier that is associated with the entity that created the secret.
  * Constraints: The maximum length is `128` characters. The minimum length is `4` characters.

* `lets_encrypt_environment` - (String) The configuration of the Let's Encrypt CA environment.
  * Constraints: Allowable values are: `production`, `staging`.

* `lets_encrypt_preferred_chain` - (String) Prefer the chain with an issuer matching this Subject Common Name.
  * Constraints: The maximum length is `30` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.

* `lets_encrypt_private_key` - (String) The PEM encoded private key of your Lets Encrypt account.
  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/(^-----BEGIN PRIVATE KEY-----.*?)/`.

* `secret_type` - (String) The secret type. Supported types are arbitrary, certificates (imported, public, and private), IAM credentials, key-value, and user credentials.
  * Constraints: Allowable values are: `arbitrary`, `imported_cert`, `public_cert`, `iam_credentials`, `kv`, `username_password`, `private_cert`.

* `updated_at` - (String) The date when a resource was recently modified. The date format follows RFC 3339.

