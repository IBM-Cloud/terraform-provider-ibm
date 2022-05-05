---
layout: "ibm"
page_title: "IBM : ibm_toolchain_tool_appconfig"
description: |-
  Manages toolchain_tool_appconfig.
subcategory: "Toolchain"
---

# ibm_toolchain_tool_appconfig

Provides a resource for toolchain_tool_appconfig. This allows toolchain_tool_appconfig to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_toolchain_tool_appconfig" "toolchain_tool_appconfig" {
  parameters {
		name = "name"
		region = "region"
		resource-group = "resource-group"
		instance-name = "instance-name"
		environment-name = "environment-name"
		collection-name = "collection-name"
		integration-status = "integration-status"
  }
  toolchain_id = "toolchain_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `name` - (Optional, String) Name of tool integration.
* `parameters` - (Optional, List) Parameters to be used to create the integration.
Nested scheme for **parameters**:
	* `collection_name` - (Required, String) App Configuration collection.
	  * Constraints: The value must match regular expression `/\\S/`.
	* `environment_name` - (Required, String) App Configuration environment.
	  * Constraints: The value must match regular expression `/\\S/`.
	* `instance_name` - (Required, String) The name of your App Configuration instance. You should choose an entry from the list provided based on the selected region and resource group. e.g: App Configuration-01.
	  * Constraints: The value must match regular expression `/\\S/`.
	* `integration_status` - (Optional, String)
	* `name` - (Required, String) Type a name for this tool integration, for example: my-appconfig. This name displays on your toolchain.
	* `region` - (Required, String) Region.
	* `resource_group` - (Required, String) Resource group.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind integration to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_tool_appconfig.
* `crn` - (Required, String) 
* `instance_id` - (Required, String) 
* `href` - (Required, String) 
* `referent` - (Required, List) 
Nested scheme for **referent**:
	* `api_href` - (Optional, String)
	* `ui_href` - (Optional, String)
* `resource_group_id` - (Required, String) 
* `state` - (Required, String) 
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.
* `toolchain_crn` - (Required, String) 
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

You can import the `ibm_toolchain_tool_appconfig` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `integration_id` in the following format:

```
<toolchain_id>/<integration_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind integration to.
* `integration_id`: A string. ID of the tool integration to be deleted.

# Syntax
```
$ terraform import ibm_toolchain_tool_appconfig.toolchain_tool_appconfig <toolchain_id>/<integration_id>
```
