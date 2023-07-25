---
layout: "ibm"
page_title: "IBM : ibm_sm_configurations"
description: |-
  Get information about sm_configurations
subcategory: "Secrets Manager"
---

# ibm_sm_configurations

Provides a read-only data source for the list of configuration metadata. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_configurations" "configurations" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Optional, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
	* Constraints: Allowable values are: `private`, `public`.
	
## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the data source.
* `configurations` - (List) A collection of configuration metadata.
	* Constraints: The maximum length is `1000` items. The minimum length is `0` items.
	  Nested scheme for **configurations**:
		* `config_type` - (String) Th configuration type.
			* Constraints: Allowable values are: `public_cert_configuration_ca_lets_encrypt`, `public_cert_configuration_dns_classic_infrastructure`, `public_cert_configuration_dns_cloud_internet_services`, `iam_credentials_configuration`, `private_cert_configuration_root_ca`, `private_cert_configuration_intermediate_ca`, `private_cert_configuration_template`.
		* `created_at` - (String) The date when a resource was created. The date format follows RFC 3339.
		* `created_by` - (String) The unique identifier that is associated with the entity that created the secret.
			* Constraints: The maximum length is `128` characters. The minimum length is `4` characters.
		* `name` - (String) The unique name of your configuration.
			* Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.
		* `secret_type` - (String) The secret type. Supported types are arbitrary, certificates (imported, public, and private), IAM credentials, key-value, and user credentials.
			* Constraints: Allowable values are: `arbitrary`, `imported_cert`, `public_cert`, `iam_credentials`, `kv`, `username_password`, `private_cert`.
		* `updated_at` - (String) The date when a resource was recently modified. The date format follows RFC 3339.
		* `lets_encrypt_environment` - (String) The configuration of the Let's Encrypt CA environment.
			* Constraints: Allowable values are: `production`, `staging`.
		* `lets_encrypt_preferred_chain` - (String) Prefer the chain with an issuer matching this Subject Common Name.
			* Constraints: The maximum length is `30` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.
		* `common_name` - (String) The Common Name (AKA CN) represents the server name that is protected by the SSL certificate.
			* Constraints: The maximum length is `128` characters. The minimum length is `4` characters. The value must match regular expression `/(.*?)/`.
		* `crl_distribution_points_encoded` - (Boolean) Determines whether to encode the certificate revocation list (CRL) distribution points in the certificates that are issued by this certificate authority.
		* `expiration_date` - (String) The date a secret is expired. The date format follows RFC 3339.
		* `key_type` - (String) The type of private key to generate.
			* Constraints: Allowable values are: `rsa`, `ec`.
		* `key_bits` - (Integer) The number of bits to use to generate the private key.Allowable values for RSA keys are: `2048` and `4096`. Allowable values for EC keys are: `224`, `256`, `384`, and `521`. The default for RSA keys is `2048`. The default for EC keys is `256`.
		* `status` - (String) The status of the certificate authority. The status of a root certificate authority is either `configured` or `expired`. For intermediate certificate authorities, possible statuses include `signing_required`,`signed_certificate_required`, `certificate_template_required`, `configured`, `expired` or `revoked`.
			* Constraints: Allowable values are: `signing_required`, `signed_certificate_required`, `certificate_template_required`, `configured`, `expired`, `revoked`.
		* `issuer` - (String) The distinguished name that identifies the entity that signed and issued the certificate.
			* Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.
		* `signing_method` - (String) The signing method to use with this certificate authority to generate private certificates.You can choose between internal or externally signed options. For more information, see the [docs](https://cloud.ibm.com/docs/secrets-manager?topic=secrets-manager-intermediate-certificate-authorities).
			* Constraints: Allowable values are: `internal`, `external`.
		* `certificate_authority` - (String) The name of the intermediate certificate authority.
			* Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9][A-Za-z0-9]*(?:_?-?\\.?[A-Za-z0-9]+)*$/`.
