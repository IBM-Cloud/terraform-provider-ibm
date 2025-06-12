---
layout: "ibm"
page_title: "IBM : ibm_sm_iam_credentials_configuration"
description: |-
  Get information about IAMCredentialsConfiguration
subcategory: "Secrets Manager"

---

# ibm_sm_iam_credentials_configuration

Provides a read-only data source for IAM Credentials configuration properties. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_iam_credentials_configuration" "sm_iam_credentials_configuration" {
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
* `name` - (Required, Forces new resource, String) The name of the configuration.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9][A-Za-z0-9]*(?:_?-?\\.?[A-Za-z0-9]+)*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `api_key` - (String) An IBM Cloud API key that can create and manage service IDs. The API key must be assigned the Editor platform role on the Access Groups Service and the Operator platform role on the IAM Identity Service. For more information, see the [docs](https://cloud.ibm.com/docs/secrets-manager?topic=secrets-manager-configure-iam-engine).
  * Constraints: The maximum length is `60` characters. The minimum length is `5` characters. The value must match regular expression `/^(?:[A-Za-z0-9_\\-]{4})*(?:[A-Za-z0-9_\\-]{2}==|[A-Za-z0-9_\\-]{3}=)?$/`.

* `config_type` - (String) The configuration type.
  * Constraints: Allowable values are: `public_cert_configuration_ca_lets_encrypt`, `public_cert_configuration_dns_classic_infrastructure`, `public_cert_configuration_dns_cloud_internet_services`, `iam_credentials_configuration`, `private_cert_configuration_root_ca`, `private_cert_configuration_intermediate_ca`, `private_cert_configuration_template`.

* `created_at` - (String) The date when the resource was created. The date format follows `RFC 3339`.

* `created_by` - (String) The unique identifier that is associated with the entity that created the secret.
  * Constraints: The maximum length is `128` characters. The minimum length is `4` characters.

* `disabled` - (Boolean) Indicates whether the API key configuration is disabled. If it is set to `true`, the IAM credentials engine doesn't use the configured API key for credentials management.
 
* `secret_type` - (String) The secret type. Supported types are arbitrary, certificates (imported, public, and private), IAM credentials, key-value, and user credentials.
  * Constraints: Allowable values are: `arbitrary`, `imported_cert`, `public_cert`, `iam_credentials`, `kv`, `username_password`, `private_cert`.

* `updated_at` - (String) The date when a resource was modified. The date format follows `RFC 3339`.

