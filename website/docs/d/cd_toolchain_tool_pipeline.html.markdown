---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_pipeline"
description: |-
  Get information about cd_toolchain_tool_pipeline
subcategory: "Toolchain"
---

# ibm_cd_toolchain_tool_pipeline

Provides a read-only data source for cd_toolchain_tool_pipeline. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cd_toolchain_tool_pipeline" "cd_toolchain_tool_pipeline" {
	integration_id = "integration_id"
	toolchain_id = ibm_cd_toolchain_tool_pipeline.cd_toolchain_tool_pipeline.toolchain_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `integration_id` - (Required, Forces new resource, String) ID of the tool integration bound to the toolchain.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the cd_toolchain_tool_pipeline.
* `crn` - (Required, String) Tool integration CRN.

* `get_integration_by_id_response_id` - (Required, String) Tool integration ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

* `href` - (Required, String) URI representing the tool integration.

* `name` - (Optional, String) Tool integration name.

* `parameters` - (Required, List) Parameters to be used to create the integration.
Nested scheme for **parameters**:
	* `name` - (Optional, String)
	* `type` - (Optional, String)
	  * Constraints: Allowable values are: `classic`, `tekton`.
	* `ui_pipeline` - (Optional, Boolean) When this check box is selected, the applications that this pipeline deploys are shown in the View app menu on the toolchain page. This setting is best for UI apps that can be accessed from a browser.
	  * Constraints: The default value is `false`.

* `referent` - (Required, List) Information on URIs to access this resource through the UI or API.
Nested scheme for **referent**:
	* `api_href` - (Optional, String) URI representing the this resource through an API.
	* `ui_href` - (Optional, String) URI representing the this resource through the UI.

* `resource_group_id` - (Required, String) Resource group where tool integration can be found.

* `state` - (Required, String) Current configuration state of the tool integration.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.

* `toolchain_crn` - (Required, String) CRN of toolchain which the integration is bound to.

* `updated_at` - (Required, String) Latest tool integration update timestamp.

