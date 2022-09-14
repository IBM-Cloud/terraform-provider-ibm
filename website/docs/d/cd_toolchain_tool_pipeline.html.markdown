---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_pipeline"
description: |-
  Get information about cd_toolchain_tool_pipeline
subcategory: "CD Toolchain"
---

# ibm_cd_toolchain_tool_pipeline

~> **Beta:** This data source is in Beta, and is subject to change.

Provides a read-only data source for cd_toolchain_tool_pipeline. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cd_toolchain_tool_pipeline" "cd_toolchain_tool_pipeline" {
	tool_id = "tool_id"
	toolchain_id = ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline.toolchain_id
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

* `id` - The unique identifier of the cd_toolchain_tool_pipeline.
* `crn` - (String) Tool CRN.

* `href` - (String) URI representing the tool.

* `name` - (String) Tool name.

* `parameters` - (List) Unique key-value pairs representing parameters to be used to create the tool.
Nested scheme for **parameters**:
	* `name` - (String)
	* `type` - (String)
	  * Constraints: Allowable values are: `classic`, `tekton`.
	* `ui_pipeline` - (Boolean) When this check box is selected, the applications that this pipeline deploys are shown in the View app menu on the toolchain page. This setting is best for UI apps that can be accessed from a browser.
	  * Constraints: The default value is `false`.

* `referent` - (List) Information on URIs to access this resource through the UI or API.
Nested scheme for **referent**:
	* `api_href` - (String) URI representing the this resource through an API.
	* `ui_href` - (String) URI representing the this resource through the UI.

* `resource_group_id` - (String) Resource group where tool can be found.

* `state` - (String) Current configuration state of the tool.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.

* `toolchain_crn` - (String) CRN of toolchain which the tool is bound to.


* `updated_at` - (String) Latest tool update timestamp.

