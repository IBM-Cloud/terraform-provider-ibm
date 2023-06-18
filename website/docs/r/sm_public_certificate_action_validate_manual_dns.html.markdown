---
layout: "ibm"
page_title: "IBM : ibm_sm_public_certificate_action_validate_manual_dns"
description: |-
Manages PublicCertificateActionValidateManualDns.
subcategory: "Secrets Manager"
---

# ibm_sm_public_certificate_action_validate_manual_dns

Provides a resource for PublicCertificateActionValidateManualDns. This allows PublicCertificateActionValidateManualDns to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_sm_public_certificate_action_validate_manual_dns" "validate_manual_dns_action" {
  instance_id           = "6ebc4224-e983-496a-8a54-f40a0bfa9175"
  region                = "us-south"
  secret_id             = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Required, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
  * Constraints: Allowable values are: `private`, `public`.
* `secret_id` - (Required, String) The ID of the secret.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
  

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