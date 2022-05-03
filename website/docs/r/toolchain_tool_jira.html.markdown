---
layout: "ibm"
page_title: "IBM : ibm_toolchain_tool_jira"
description: |-
  Manages toolchain_tool_jira.
subcategory: "Toolchain"
---

# ibm_toolchain_tool_jira

Provides a resource for toolchain_tool_jira. This allows toolchain_tool_jira to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_toolchain_tool_jira" "toolchain_tool_jira" {
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

* `name` - (Optional, String) Name of tool integration.
* `parameters` - (Optional, List) Tool integration parameters.
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
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind integration to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_tool_jira.
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

You can import the `ibm_toolchain_tool_jira` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `integration_id` in the following format:

```
<toolchain_id>/<integration_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind integration to.
* `integration_id`: A string. ID of the tool integration to be deleted.

# Syntax
```
$ terraform import ibm_toolchain_tool_jira.toolchain_tool_jira <toolchain_id>/<integration_id>
```
