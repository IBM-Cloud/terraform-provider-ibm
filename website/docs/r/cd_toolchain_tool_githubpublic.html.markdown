---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_githubpublic"
description: |-
  Manages cd_toolchain_tool_githubpublic.
subcategory: "CD Toolchain"
---

# ibm_cd_toolchain_tool_githubpublic

Provides a resource for cd_toolchain_tool_githubpublic. This allows cd_toolchain_tool_githubpublic to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cd_toolchain_tool_githubpublic" "cd_toolchain_tool_githubpublic" {
  initialization {
		legal = true
		repo_name = "repo_name"
		repo_url = "repo_url"
		source_repo_url = "source_repo_url"
		type = "new"
		private_repo = true
  }
  parameters {
		legal = true
		git_id = "git_id"
		title = "title"
		api_root_url = "api_root_url"
		default_branch = "default_branch"
		root_url = "root_url"
		access_token = "access_token"
		blind_connection = true
		owner_id = "owner_id"
		repo_name = "repo_name"
		repo_url = "repo_url"
		source_repo_url = "source_repo_url"
		token_url = "token_url"
		type = "new"
		private_repo = true
		has_issues = true
		auto_init = true
		enable_traceability = true
		authorized = "authorized"
		integration_owner = "integration_owner"
		auth_type = "oauth"
		api_token = "api_token"
  }
  toolchain_id = "toolchain_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `initialization` - (Optional, List) 
Nested scheme for **initialization**:
	* `legal` - (Optional, Forces new resource, Boolean)
	  * Constraints: The default value is `true`.
	* `private_repo` - (Optional, Forces new resource, Boolean) Select this check box to make this repository private.
	  * Constraints: The default value is `false`.
	* `repo_name` - (Optional, Forces new resource, String)
	* `repo_url` - (Optional, Forces new resource, String) Type the URL of the repository that you are linking to.
	* `source_repo_url` - (Optional, Forces new resource, String) Type the URL of the repository that you are forking or cloning.
	* `type` - (Required, Forces new resource, String)
	  * Constraints: Allowable values are: `new`, `fork`, `clone`, `link`.
* `name` - (Optional, String) Name of tool.
  * Constraints: The maximum length is `128` characters. The minimum length is `0` characters. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9-._ ])+$/`.
* `parameters` - (Optional, List) Parameters to be used to create the tool.
Nested scheme for **parameters**:
	* `access_token` - (Optional, String)
	* `api_root_url` - (Optional, String) e.g. https://api.github.example.com.
	* `api_token` - (Optional, String) Personal Access Token.
	* `auth_type` - (Optional, String)
	  * Constraints: Allowable values are: `oauth`, `pat`.
	* `authorized` - (Optional, String)
	* `auto_init` - (Optional, Boolean) Select this checkbox to initialize this repository with a README.
	  * Constraints: The default value is `false`.
	* `blind_connection` - (Optional, Boolean) Select this checkbox only if the server is not addressable on the public internet. IBM Cloud will not be able to validate the connection details you provide. Certain functionality that requires API access to the git server will be disabled. Delivery pipeline will only work using a private worker that has network access to the git server.
	  * Constraints: The default value is `false`.
	* `default_branch` - (Optional, String) e.g. main.
	* `enable_traceability` - (Optional, Boolean) Select this check box to track the deployment of code changes by creating tags, labels and comments on commits, pull requests and referenced issues.
	  * Constraints: The default value is `false`.
	* `git_id` - (Optional, String)
	* `has_issues` - (Optional, Boolean) Select this check box to enable GitHub Issues for lightweight issue tracking.
	  * Constraints: The default value is `true`.
	* `integration_owner` - (Optional, String) Select the user which git operations will be performed as.
	* `legal` - (Optional, Boolean)
	  * Constraints: The default value is `false`.
	* `owner_id` - (Optional, String)
	* `private_repo` - (Optional, Boolean) Select this check box to make this repository private.
	  * Constraints: The default value is `false`.
	* `repo_name` - (Optional, String)
	* `repo_url` - (Optional, String) Type the URL of the repository that you are linking to.
	* `root_url` - (Optional, String) e.g. https://github.example.com.
	* `source_repo_url` - (Optional, String) Type the URL of the repository that you are forking or cloning.
	* `title` - (Optional, String) e.g. My GitHub Enterprise Server.
	* `token_url` - (Optional, String) Integration token URL.
	* `type` - (Optional, String)
	  * Constraints: Allowable values are: `new`, `fork`, `clone`, `link`.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind tool to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cd_toolchain_tool_githubpublic.
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

You can import the `ibm_cd_toolchain_tool_githubpublic` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `tool_id` in the following format:

```
<toolchain_id>/<tool_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind tool to.
* `tool_id`: A string. ID of the tool bound to the toolchain.

# Syntax
```
$ terraform import ibm_cd_toolchain_tool_githubpublic.cd_toolchain_tool_githubpublic <toolchain_id>/<tool_id>
```
