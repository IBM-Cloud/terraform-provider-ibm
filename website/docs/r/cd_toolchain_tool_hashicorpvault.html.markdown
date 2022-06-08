---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_hashicorpvault"
description: |-
  Manages cd_toolchain_tool_hashicorpvault.
subcategory: "CD Toolchain"
---

# ibm_cd_toolchain_tool_hashicorpvault

Provides a resource for cd_toolchain_tool_hashicorpvault. This allows cd_toolchain_tool_hashicorpvault to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cd_toolchain_tool_hashicorpvault" "cd_toolchain_tool_hashicorpvault" {
  parameters {
		name = "name"
		server_url = "server_url"
		authentication_method = "token"
		token = "token"
		role_id = "role_id"
		secret_id = "secret_id"
		dashboard_url = "dashboard_url"
		path = "path"
		secret_filter = "secret_filter"
		default_secret = "default_secret"
		username = "username"
		password = "password"
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
	* `authentication_method` - (Required, String) Choose the authentication method for your HashiCorp Vault instance.
	  * Constraints: Allowable values are: `token`, `approle`, `userpass`, `github`.
	* `dashboard_url` - (Required, String) Type the URL that you want to navigate to when you click the HashiCorp Vault integration tile.
	* `default_secret` - (Optional, String) Type a default secret name that will be selected or used if no list of secret names are returned from your HashiCorp Vault instance.
	* `name` - (Required, String) Enter a name for this tool integration. This name is displayed on your toolchain.
	* `password` - (Optional, String) Type or select the authentication password for your HashiCorp Vault instance.
	* `path` - (Required, String) Type the mount path where your secrets are stored in your HashiCorp Vault instance.
	* `role_id` - (Optional, String) Type or select the authentication role ID for your HashiCorp Vault instance.
	* `secret_filter` - (Optional, String) Type a regular expression to filter the list of secret names returned from your HashiCorp Vault instance.
	* `secret_id` - (Optional, String) Type or select the authentication secret ID for your HashiCorp Vault instance.
	* `server_url` - (Required, String) Type the server URL for your HashiCorp Vault instance.
	* `token` - (Optional, String) Type or select the authentication token for your HashiCorp Vault instance.
	* `username` - (Optional, String) Type or select the authentication username for your HashiCorp Vault instance.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind tool to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cd_toolchain_tool_hashicorpvault.
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

You can import the `ibm_cd_toolchain_tool_hashicorpvault` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `tool_id` in the following format:

```
<toolchain_id>/<tool_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind tool to.
* `tool_id`: A string. ID of the tool bound to the toolchain.

# Syntax
```
$ terraform import ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault <toolchain_id>/<tool_id>
```
