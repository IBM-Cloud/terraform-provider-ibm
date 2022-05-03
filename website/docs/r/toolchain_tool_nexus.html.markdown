---
layout: "ibm"
page_title: "IBM : ibm_toolchain_tool_nexus"
description: |-
  Manages toolchain_tool_nexus.
subcategory: "Toolchain"
---

# ibm_toolchain_tool_nexus

Provides a resource for toolchain_tool_nexus. This allows toolchain_tool_nexus to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_toolchain_tool_nexus" "toolchain_tool_nexus" {
  parameters {
		name = "name"
		dashboard_url = "dashboard_url"
		type = "npm"
		user_id = "user_id"
		token = "token"
		release_url = "release_url"
		mirror_url = "mirror_url"
		snapshot_url = "snapshot_url"
  }
  toolchain_id = "toolchain_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `name` - (Optional, String) Name of tool integration.
* `parameters` - (Optional, List) Tool integration parameters.
Nested scheme for **parameters**:
	* `dashboard_url` - (Optional, String) Type the URL that you want to navigate to when you click the Nexus integration tile.
	* `mirror_url` - (Optional, String) Type the URL for your Nexus virtual repository, which is a repository that can see your private repositories and a cache of the public repositories.
	* `name` - (Required, String) Type a name for this tool integration, for example: my-nexus. This name displays on your toolchain.
	* `release_url` - (Optional, String) Type the URL for your Nexus release repository.
	* `snapshot_url` - (Optional, String) Type the URL for your Nexus snapshot repository.
	* `token` - (Optional, String) Type the password or authentication token for your Nexus repository.
	* `type` - (Required, String) Choose the type of repository for your Nexus integration.
	  * Constraints: Allowable values are: `npm`, `maven`.
	* `user_id` - (Optional, String) Type the User ID or email for your Nexus repository.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind integration to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_tool_nexus.
* `crn` - (Required, String) 
* `get_integration_by_id_response_id` - (Required, String) 
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

## Import

You can import the `ibm_toolchain_tool_nexus` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `integration_id` in the following format:

```
<toolchain_id>/<integration_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind integration to.
* `integration_id`: A string. ID of the tool integration to be deleted.

# Syntax
```
$ terraform import ibm_toolchain_tool_nexus.toolchain_tool_nexus <toolchain_id>/<integration_id>
```
