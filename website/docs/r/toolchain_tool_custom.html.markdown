---
layout: "ibm"
page_title: "IBM : ibm_toolchain_tool_custom"
description: |-
  Manages toolchain_tool_custom.
subcategory: "IBM Toolchain API"
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

* `container` - (Optional, Forces new resource, List) 
Nested scheme for **container**:
	* `guid` - (Required, String)
	* `type` - (Required, String)
	  * Constraints: Allowable values are: `organization_guid`, `resource_group_id`.
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
* `parameters_references` - (Optional, Map) Decoded values used on provision in the broker that reference fields in the parameters.
* `toolchain_id` - (Required, Forces new resource, String) 

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_tool_custom.
* `dashboard_url` - (Required, String) The URL of a user-facing user interface for this instance of a service.

## Import

You can import the `ibm_toolchain_tool_custom` resource by using `instance_id`. The id of the created or updated service instance.

# Syntax
```
$ terraform import ibm_toolchain_tool_custom.toolchain_tool_custom <instance_id>
```

# Example
```
$ terraform import ibm_toolchain_tool_custom.toolchain_tool_custom 4f107490-3820-400b-a008-f7f38d4163ed
```
