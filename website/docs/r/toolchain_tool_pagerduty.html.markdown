---
layout: "ibm"
page_title: "IBM : ibm_toolchain_tool_pagerduty"
description: |-
  Manages toolchain_tool_pagerduty.
subcategory: "Toolchain"
---

# ibm_toolchain_tool_pagerduty

Provides a resource for toolchain_tool_pagerduty. This allows toolchain_tool_pagerduty to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_toolchain_tool_pagerduty" "toolchain_tool_pagerduty" {
  parameters {
		key_type = "api"
		api_key = "api_key"
		service_name = "service_name"
		user_email = "user_email"
		user_phone = "user_phone"
		service_url = "service_url"
		service_key = "service_key"
		service_id = "service_id"
  }
  toolchain_id = "toolchain_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `name` - (Optional, String) Name of tool integration.
* `parameters` - (Optional, List) Parameters to be used to create the integration.
Nested scheme for **parameters**:
	* `api_key` - (Optional, String) Type your API access key. You can find or create this key on the Configuration/API Access section of the PagerDuty website. [PagerDuty Support article on how to get API Key](https://support.pagerduty.com/hc/en-us/articles/202829310-Generating-an-API-Key).
	* `key_type` - (Required, String) Select whether to integrate at the account level with an API key or at the service level with an integration key.
	  * Constraints: Allowable values are: `api`, `service`.
	* `service_id` - (Optional, String) service_id.
	* `service_key` - (Optional, String) Type your integration key. You can find or create this key in the Integrations section of the PagerDuty service page.
	* `service_name` - (Optional, String) Type the name of the PagerDuty service to post alerts to. If you want alerts to be posted to a new service, type a new name. PagerDuty will create the service.
	* `service_url` - (Optional, String) Type the URL of the PagerDuty service to post alerts to.
	* `user_email` - (Optional, String) Type the email address of the user to contact when an alert is posted. If you want alerts to be sent to a new email address, type the address and PagerDuty will create a user.
	* `user_phone` - (Optional, String) Type the phone number of the user to contact when an alert is posted. Include the national code followed by a space and a 10-digit number; for example: +1 1234567890. If you omit the national code, it is set to +1 by default.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind integration to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_tool_pagerduty.
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

You can import the `ibm_toolchain_tool_pagerduty` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `integration_id` in the following format:

```
<toolchain_id>/<integration_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind integration to.
* `integration_id`: A string. ID of the tool integration to be deleted.

# Syntax
```
$ terraform import ibm_toolchain_tool_pagerduty.toolchain_tool_pagerduty <toolchain_id>/<integration_id>
```
