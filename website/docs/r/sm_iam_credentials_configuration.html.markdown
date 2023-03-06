---
layout: "ibm"
page_title: "IBM : ibm_sm_iam_credentials_configuration"
description: |-
  Manages IAMCredentialsConfiguration.
subcategory: "Secrets Manager"
---

# ibm_sm_iam_credentials_configuration

Provides a resource for IAMCredentialsConfiguration. This allows IAMCredentialsConfiguration to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_sm_iam_credentials_configuration" "sm_iam_credentials_configuration_instance" {
	instance_id   = "6ebc4224-e983-496a-8a54-f40a0bfa9175"
	region        = "us-south"
	name = "my-example-engine-config"
	api_key = "my-api-key"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `name` - (Required, String) A human-readable unique name to assign to your IAM Credentials configuration.
* `api_key` - (Required, String) The API key that is generated for this secret.After the secret reaches the end of its lease (see the `ttl` field), the API key is deleted automatically. If you want to continue to use the same API key for future read operations, see the `reuse_api_key` field.
	* Constraints: The maximum length is `60` characters. The minimum length is `5` characters. The value must match regular expression `/^(?:[A-Za-z0-9_\\-]{4})*(?:[A-Za-z0-9_\\-]{2}==|[A-Za-z0-9_\\-]{3}=)?$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the IAMCredentialsConfiguration.
* `config_type` - (String) The configuration type.
  * Constraints: Allowable values are: `public_cert_configuration_ca_lets_encrypt`, `public_cert_configuration_dns_classic_infrastructure`, `public_cert_configuration_dns_cloud_internet_services`, `iam_credentials_configuration`, `private_cert_configuration_root_ca`, `private_cert_configuration_intermediate_ca`, `private_cert_configuration_template`.
* `created_at` - (String) The date when the resource was created. The date format follows `RFC 3339`.
* `created_by` - (String) The unique identifier that is associated with the entity that created the secret.
  * Constraints: The maximum length is `128` characters. The minimum length is `4` characters.
* `secret_type` - (String) The secret type. Supported types are arbitrary, certificates (imported, public, and private), IAM credentials, key-value, and user credentials.
  * Constraints: Allowable values are: `arbitrary`, `imported_cert`, `public_cert`, `iam_credentials`, `kv`, `username_password`, `private_cert`.
* `updated_at` - (String) The date when a resource was modified. The date format follows `RFC 3339`.

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

You can import the `ibm_sm_iam_credentials_configuration` resource by using `region`, `instance_id`, and `name`.
For more information, see [the documentation](https://cloud.ibm.com/docs/secrets-manager)

# Syntax
```
$ terraform import ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration <region>/<instance_id>/<name>
```

# Example
```
$ terraform import ibm_sm_iam_credentials_configuration.sm_iam_credentials_configuration us-east/6ebc4224-e983-496a-8a54-f40a0bfa9175/my-secret-engine-config
```
