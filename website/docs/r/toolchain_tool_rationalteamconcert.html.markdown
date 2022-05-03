---
layout: "ibm"
page_title: "IBM : ibm_toolchain_tool_rationalteamconcert"
description: |-
  Manages toolchain_tool_rationalteamconcert.
subcategory: "Toolchain"
---

# ibm_toolchain_tool_rationalteamconcert

Provides a resource for toolchain_tool_rationalteamconcert. This allows toolchain_tool_rationalteamconcert to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_toolchain_tool_rationalteamconcert" "toolchain_tool_rationalteamconcert" {
  parameters {
		server_url = "server_url"
		user_id = "user_id"
		password = "password"
		type = "new"
		project_area = "project_area"
		process_template = "process_template"
		enable_traceability = true
  }
  toolchain_id = "toolchain_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `name` - (Optional, String) Name of tool integration.
* `parameters` - (Optional, List) Tool integration parameters.
Nested scheme for **parameters**:
	* `enable_traceability` - (Optional, Boolean) Select this check box to track the deployment of code changes by creating tags, comments on work items.
	  * Constraints: The default value is `false`.
	* `password` - (Required, String) Type your password for Rational Team Concert (Jazz) server access.
	* `process_template` - (Optional, String) Type the Rational Team Concert process template to use to create the project.
	* `project_area` - (Required, String) Type the name of the Rational Team Concert project area to add to the toolchain.
	* `server_url` - (Required, String) Type the server URL for your Rational Team Concert instance.
	* `type` - (Required, String)
	  * Constraints: Allowable values are: `new`, `existing`.
	* `user_id` - (Required, String) Type your user id for Rational Team Concert (Jazz) server access.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind integration to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_tool_rationalteamconcert.
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

You can import the `ibm_toolchain_tool_rationalteamconcert` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `integration_id` in the following format:

```
<toolchain_id>/<integration_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind integration to.
* `integration_id`: A string. ID of the tool integration to be deleted.

# Syntax
```
$ terraform import ibm_toolchain_tool_rationalteamconcert.toolchain_tool_rationalteamconcert <toolchain_id>/<integration_id>
```
