---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_keyprotect"
description: |-
  Get information about cd_toolchain_tool_keyprotect
subcategory: "CD Toolchain"
---

# ibm_cd_toolchain_tool_keyprotect

Provides a read-only data source for cd_toolchain_tool_keyprotect. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cd_toolchain_tool_keyprotect" "cd_toolchain_tool_keyprotect" {
	tool_id = "tool_id"
	toolchain_id = ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect.toolchain_id
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

* `id` - The unique identifier of the cd_toolchain_tool_keyprotect.
* `crn` - (Required, String) Tool CRN.

* `get_tool_by_id_response_id` - (Required, String) Tool ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

* `href` - (Required, String) URI representing the tool.

* `name` - (Optional, String) Tool name.

* `parameters` - (Required, List) Parameters to be used to create the tool.
Nested scheme for **parameters**:
	* `instance_name` - (Required, String) The name of your Key Protect instance. You should choose an entry from the list provided based on the selected region and resource group. e.g: Key Protect-01.
	  * Constraints: The value must match regular expression `/\\S/`.
	* `name` - (Required, String) Enter a name for this tool integration. This name is displayed on your toolchain.
	* `region` - (Required, String) Region.
	* `resource_group` - (Required, String) Resource group.

* `referent` - (Required, List) Information on URIs to access this resource through the UI or API.
Nested scheme for **referent**:
	* `api_href` - (Optional, String) URI representing the this resource through an API.
	* `ui_href` - (Optional, String) URI representing the this resource through the UI.

* `resource_group_id` - (Required, String) Resource group where tool can be found.

* `state` - (Required, String) Current configuration state of the tool.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.

* `toolchain_crn` - (Required, String) CRN of toolchain which the tool is bound to.

* `updated_at` - (Required, String) Latest tool update timestamp.

