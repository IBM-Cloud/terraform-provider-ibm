---
layout: "ibm"
page_title: "IBM : ibm_toolchain_tool_custom"
description: |-
  Manages toolchain_tool_custom.
subcategory: "Toolchain"
---

# ibm_toolchain_tool_custom

Provides a resource for toolchain_tool_custom. This allows toolchain_tool_custom to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_toolchain_tool_custom" "toolchain_tool_custom" {
  toolchain_id = "toolchain_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `name` - (Optional, String) Name of tool integration.
* `parameters` - (Optional, List) 
Nested scheme for **parameters**:
	* `additional_properties` - (Optional, String) (Advanced) Type any information that is needed to integrate with other tools in your toolchain.
	* `dashboard_url` - (Required, String) Type the URL that you want to navigate to when you click the tool integration card.
	* `description` - (Optional, String) Type a description for the tool instance.
	* `documentation_url` - (Optional, String) Type the URL for your tool's documentation.
	* `image_url` - (Optional, String) Type the URL of the icon to show on your tool integration's card.
	* `lifecycle_phase` - (Required, String) Select the lifecycle phase of the IBM Cloud Garage Method that is the most closely associated with this tool.
	  * Constraints: Allowable values are: `THINK`, `CODE`, `DELIVER`, `RUN`, `MANAGE`, `LEARN`, `CULTURE`.
	* `name` - (Required, String) Type a name for this specific tool integration; for example: My Build and Deploy Pipeline.
	* `type` - (Required, String) Type the name of the tool that you are integrating; for example: Delivery Pipeline.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind integration to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_tool_custom.
* `crn` - (Required, String) 
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

You can import the `ibm_toolchain_tool_custom` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `integration_id` in the following format:

```
<toolchain_id>/<integration_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind integration to.
* `integration_id`: A string. ID of the tool integration to be deleted.

# Syntax
```
$ terraform import ibm_toolchain_tool_custom.toolchain_tool_custom <toolchain_id>/<integration_id>
```
