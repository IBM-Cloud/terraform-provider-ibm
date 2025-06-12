---
layout: "ibm"
page_title: "IBM : ibm_sm_public_certificate_configuration_dns_classic_infrastructure"
description: |-
  Manages PublicCertificateConfigurationDNSClassicInfrastructure.
subcategory: "Secrets Manager"
---

# ibm_sm_public_certificate_configuration_dns_classic_infrastructure

Provides a resource for PublicCertificateConfigurationDNSClassicInfrastructure. This allows PublicCertificateConfigurationDNSClassicInfrastructure to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_sm_public_certificate_configuration_dns_classic_infrastructure" "sm_public_certificate_configuration_dns_classic_infrastructure_instance" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
  name          = "my_DNS_config"
  classic_infrastructure_password = "username"
  classic_infrastructure_username = "password"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Optional, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
  * Constraints: Allowable values are: `private`, `public`.
* `name` - A human-readable unique name to assign to your DNS provider configuration name.
* `classic_infrastructure_password` - (Required, String) Your classic infrastructure API key.For information about viewing and accessing your classic infrastructure API key, see the [docs](https://cloud.ibm.com/docs/account?topic=account-classic_keys).
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.
* `classic_infrastructure_username` - (Required, String) The username that is associated with your classic infrastructure account.In most cases, your classic infrastructure username is your `<account_id>_<email_address>`. For more information, see the [docs](https://cloud.ibm.com/docs/account?topic=account-classic_keys).
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the PublicCertificateConfigurationDNSClassicInfrastructure.
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

You can import the `ibm_sm_public_certificate_configuration_dns_classic_infrastructure` resource by using `region`, `instance_id`, and `name`.
For more information, see [the documentation](https://cloud.ibm.com/docs/secrets-manager)

# Syntax
```bash
$ terraform import ibm_sm_public_certificate_configuration_dns_classic_infrastructure.sm_public_certificate_configuration_dns_classic_infrastructure <region>/<instance_id>/<name>
```

# Example
```bash
$ terraform import ibm_sm_public_certificate_configuration_dns_classic_infrastructure.sm_public_certificate_configuration_dns_classic_infrastructure us-east/6ebc4224-e983-496a-8a54-f40a0bfa9175/my_DNS_config
```
