---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_hashicorpvault"
description: |-
  Get information about cd_toolchain_tool_hashicorpvault
subcategory: "CD Toolchain"
---

# ibm_cd_toolchain_tool_hashicorpvault

~> **Beta:** This data source is in Beta, and is subject to change.

Provides a read-only data source for cd_toolchain_tool_hashicorpvault. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cd_toolchain_tool_hashicorpvault" "cd_toolchain_tool_hashicorpvault" {
	tool_id = "tool_id"
	toolchain_id = ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault.toolchain_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `tool_id` - (Required, Forces new resource, String) ID of the tool bound to the toolchain.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the cd_toolchain_tool_hashicorpvault.
* `crn` - (String) Tool CRN.

* `href` - (String) URI representing the tool.

* `name` - (String) Tool name.

* `parameters` - (List) Unique key-value pairs representing parameters to be used to create the tool.
Nested scheme for **parameters**:
	* `authentication_method` - (String) Choose the authentication method for your HashiCorp Vault instance.
	  * Constraints: Allowable values are: `token`, `approle`, `userpass`, `github`.
	* `dashboard_url` - (String) Type the URL that you want to navigate to when you click the HashiCorp Vault integration tile.
	* `default_secret` - (String) Type a default secret name that will be selected or used if no list of secret names are returned from your HashiCorp Vault instance.
	* `name` - (String) Enter a name for this tool integration. This name is displayed on your toolchain.
	* `password` - (String) Type or select the authentication password for your HashiCorp Vault instance.
	* `path` - (String) Type the mount path where your secrets are stored in your HashiCorp Vault instance.
	* `role_id` - (String) Type or select the authentication role ID for your HashiCorp Vault instance.
	* `secret_filter` - (String) Type a regular expression to filter the list of secret names returned from your HashiCorp Vault instance.
	* `secret_id` - (String) Type or select the authentication secret ID for your HashiCorp Vault instance.
	* `server_url` - (String) Type the server URL for your HashiCorp Vault instance.
	* `token` - (String) Type or select the authentication token for your HashiCorp Vault instance.
	* `username` - (String) Type or select the authentication username for your HashiCorp Vault instance.

* `referent` - (List) Information on URIs to access this resource through the UI or API.
Nested scheme for **referent**:
	* `api_href` - (String) URI representing the this resource through an API.
	* `ui_href` - (String) URI representing the this resource through the UI.

* `resource_group_id` - (String) Resource group where tool can be found.

* `state` - (String) Current configuration state of the tool.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.

* `toolchain_crn` - (String) CRN of toolchain which the tool is bound to.


* `updated_at` - (String) Latest tool update timestamp.

