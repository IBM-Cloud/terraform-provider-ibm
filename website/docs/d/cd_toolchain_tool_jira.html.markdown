---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_jira"
description: |-
  Get information about cd_toolchain_tool_jira
subcategory: "CD Toolchain"
---

# ibm_cd_toolchain_tool_jira

Provides a read-only data source for cd_toolchain_tool_jira. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cd_toolchain_tool_jira" "cd_toolchain_tool_jira" {
	tool_id = "tool_id"
	toolchain_id = ibm_cd_toolchain_tool_jira.cd_toolchain_tool_jira.toolchain_id
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

* `id` - The unique identifier of the cd_toolchain_tool_jira.
* `crn` - (Required, String) Tool CRN.

* `get_tool_by_id_response_id` - (Required, String) Tool ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

* `href` - (Required, String) URI representing the tool.

* `name` - (Optional, String) Tool name.

* `parameters` - (Required, List) Parameters to be used to create the tool.
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

* `referent` - (Required, List) Information on URIs to access this resource through the UI or API.
Nested scheme for **referent**:
	* `api_href` - (Optional, String) URI representing the this resource through an API.
	* `ui_href` - (Optional, String) URI representing the this resource through the UI.

* `resource_group_id` - (Required, String) Resource group where tool can be found.

* `state` - (Required, String) Current configuration state of the tool.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.

* `toolchain_crn` - (Required, String) CRN of toolchain which the tool is bound to.

* `updated_at` - (Required, String) Latest tool update timestamp.

