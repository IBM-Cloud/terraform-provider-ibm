---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_jira"
description: |-
  Manages cd_toolchain_tool_jira.
subcategory: "CD Toolchain"
---

# ibm_cd_toolchain_tool_jira

Provides a resource for cd_toolchain_tool_jira. This allows cd_toolchain_tool_jira to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cd_toolchain_tool_jira" "cd_toolchain_tool_jira" {
  parameters {
		type = "new"
		project_key = "project_key"
		project_name = "project_name"
		project_admin = "project_admin"
		api_url = "api_url"
		username = "username"
		password = "password"
		enable_traceability = true
  }
  toolchain_id = "toolchain_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `name` - (Optional, String) Name of tool.
  * Constraints: The maximum length is `128` characters. The minimum length is `0` characters. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9-._ ])+$/`.
* `parameters` - (Optional, List) Parameters to be used to create the tool.
Nested scheme for **parameters**:
	* `api_url` - (Required, String) Type the base API URL for your JIRA instance. To find that value, from the header of your JIRA instance, click the **Administration** icon, which looks like a gear, and then click **System**.
	* `enable_traceability` - (Optional, Boolean) Select this check box to track the deployment of code changes by creating tags, labels and comments on commits, pull requests and referenced issues.
	  * Constraints: The default value is `false`.
	* `password` - (Optional, String) Your api token is required only if you are connecting to a private JIRA instance or if you are connecting to a public instance and want to receive traceability information or if you are creating a new project. Otherwise, you do not need to enter your api token.
	* `project_admin` - (Optional, String)
	* `project_key` - (Required, String)
	* `project_name` - (Optional, String)
	* `type` - (Required, String)
	  * Constraints: Allowable values are: `new`, `existing`.
	* `username` - (Optional, String) Your user name is required only if you are connecting to a private JIRA instance or if you are connecting to a public instance and want to receive traceability information or if you are creating a new project. Otherwise, you do not need to enter your user name.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind tool to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cd_toolchain_tool_jira.
* `crn` - (Required, String) Tool CRN.
* `get_tool_by_id_response_id` - (Required, String) Tool ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.
* `href` - (Required, String) URI representing the tool.
* `referent` - (Required, List) Information on URIs to access this resource through the UI or API.
Nested scheme for **referent**:
	* `api_href` - (Optional, String) URI representing the this resource through an API.
	* `ui_href` - (Optional, String) URI representing the this resource through the UI.
* `resource_group_id` - (Required, String) Resource group where tool can be found.
* `state` - (Required, String) Current configuration state of the tool.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.
* `toolchain_crn` - (Required, String) CRN of toolchain which the tool is bound to.
* `updated_at` - (Required, String) Latest tool update timestamp.

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

You can import the `ibm_cd_toolchain_tool_jira` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `tool_id` in the following format:

```
<toolchain_id>/<tool_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind tool to.
* `tool_id`: A string. ID of the tool bound to the toolchain.

# Syntax
```
$ terraform import ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira <toolchain_id>/<tool_id>
```
