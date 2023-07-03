---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain"
description: |-
  Manages cd_toolchain.
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain

Provides a resource for cd_toolchain. This allows cd_toolchain to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cd_toolchain" "cd_toolchain_instance" {
  description = "A sample toolchain to test the API"
  name = "TestToolchainV2"
  resource_group_id = "6a9a01f2cff54a7f966f803d92877123"
  tags = ["tag1", "tag2"]
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `description` - (Optional, String) Describes the toolchain.
  * Constraints: The maximum length is `500` characters. The minimum length is `0` characters. The value must match regular expression `/^(.*?)$/`.
* `name` - (Required, String) Toolchain name.
  * Constraints: The maximum length is `128` characters. The minimum length is `0` characters. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9-._ ])+$/`.
* `resource_group_id` - (Required, Forces new resource, String) Resource group where toolchain will be created.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[0-9a-f]{32}$/`.
* `tags` - (Optional, Array of Strings) Tags associated with the toolchain.


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cd_toolchain.
* `account_id` - (String) Account ID where toolchain can be found.
* `created_at` - (String) Toolchain creation timestamp.
* `created_by` - (String) Identity that created the toolchain.
* `crn` - (String) Toolchain CRN.
* `href` - (String) URI that can be used to retrieve toolchain.
* `location` - (String) Toolchain region.
* `ui_href` - (String) URL of a user-facing user interface for this toolchain.
* `updated_at` - (String) Latest toolchain update timestamp.

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

You can import the `ibm_cd_toolchain` resource by using `id`. Toolchain ID.

# Syntax
```
$ terraform import ibm_cd_toolchain.cd_toolchain <id>
```
