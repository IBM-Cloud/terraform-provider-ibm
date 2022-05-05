---
layout: "ibm"
page_title: "IBM : ibm_toolchain"
description: |-
  Manages toolchain.
subcategory: "Toolchain"
---

# ibm_toolchain

Provides a resource for toolchain. This allows toolchain to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_toolchain" "toolchain" {
  description = "A sample toolchain to test the API"
  name = "TestToolchainV2"
  resource_group_id = "6a9a01f2cff54a7f966f803d92877123"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `description` - (Optional, String) Describes the toolchain.
  * Constraints: The maximum length is `500` characters.
* `name` - (Required, String) Toolchain name.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9-._ ])+$/`.
* `resource_group_id` - (Required, Forces new resource, String) 
  * Constraints: The value must match regular expression `/^[0-9a-f]{32}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain.
* `account_id` - (Required, String) 
* `created_at` - (Required, String) 
* `created_by` - (Required, String) 
* `crn` - (Required, String) 
* `href` - (Required, String) 
* `location` - (Required, String) 
* `tags` - (Required, List) 
* `updated_at` - (Required, String) 

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

You can import the `ibm_toolchain` resource by using `id`. The unique identifier of the toolchain.

# Syntax
```
$ terraform import ibm_toolchain.toolchain <id>
```

# Example
```
$ terraform import ibm_toolchain.toolchain ec58a911-c217-4e56-a40b-93482cd18706
```
