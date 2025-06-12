---
layout: "ibm"
page_title: "IBM : ibm_sm_public_certificate_configuration_ca_lets_encrypt"
description: |-
  Manages PublicCertificateConfigurationCALetsEncrypt.
subcategory: "Secrets Manager"
---

# ibm_sm_public_certificate_configuration_ca_lets_encrypt

Provides a resource for PublicCertificateConfigurationCALetsEncrypt. This allows PublicCertificateConfigurationCALetsEncrypt to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_sm_public_certificate_configuration_ca_lets_encrypt" "sm_public_certificate_configuration_ca_lets_encrypt" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
  name          = "lets-encrypt-config"
  lets_encrypt_environment = "production"
  lets_encrypt_private_key = "-----BEGIN PRIVATE KEY-----\nMIIEowIBAAKCAQEAqcRbzV1wp0nVrPtEpMtnWMO6Js1q3rhREZluKZfu0Q8SY4H3\n-----END PRIVATE KEY-----",
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Optional, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
  * Constraints: Allowable values are: `private`, `public`.
* `lets_encrypt_environment` - (Required, String) The configuration of the Let's Encrypt CA environment.
  * Constraints: Allowable values are: `production`, `staging`.
* `lets_encrypt_preferred_chain` - (Optional, String) This field supports only the chains that Let's Encrypt provides. Keep empty to use the default or supply a valid Let's Encrypt-provided value. For a list of supported chains, see: https://letsencrypt.org/certificates/.
  * Constraints: The value must match regular expression `/(.*?)/`.
* `lets_encrypt_private_key` - (Required, String) The PEM encoded private key of your Lets Encrypt account.
  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/(^-----BEGIN PRIVATE KEY-----.*?)/`.
* `name` - (Required, String) A human-readable unique name to assign to your configuration.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the PublicCertificateConfigurationCALetsEncrypt.
* `config_type` - (String) The configuration type.
  * Constraints: Allowable values are: `public_cert_configuration_ca_lets_encrypt`, `public_cert_configuration_dns_classic_infrastructure`, `public_cert_configuration_dns_cloud_internet_services`, `iam_credentials_configuration`, `private_cert_configuration_root_ca`, `private_cert_configuration_intermediate_ca`, `private_cert_configuration_template`.
* `created_at` - (String) The date when a resource was created. The date format follows RFC 3339.
* `created_by` - (String) The unique identifier that is associated with the entity that created the secret.
  * Constraints: The maximum length is `128` characters. The minimum length is `4` characters.
* `secret_type` - (String) The secret type. Supported types are arbitrary, certificates (imported, public, and private), IAM credentials, key-value, and user credentials.
  * Constraints: Allowable values are: `arbitrary`, `imported_cert`, `public_cert`, `iam_credentials`, `kv`, `username_password`, `private_cert`.
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

You can import the `ibm_sm_public_certificate_configuration_ca_lets_encrypt` resource by using `region`, `instance_id`, and `name`.
For more information, see [the documentation](https://cloud.ibm.com/docs/secrets-manager)

# Syntax
```bash
$ terraform import ibm_sm_public_certificate_configuration_ca_lets_encrypt.sm_public_certificate_configuration_ca_lets_encrypt <region>/<instance_id>/<name>
```

# Example
```bash
$ terraform import ibm_sm_public_certificate_configuration_ca_lets_encrypt.sm_public_certificate_configuration_ca_lets_encrypt us-east/6ebc4224-e983-496a-8a54-f40a0bfa9175/lets-encrypt-config
```
