---
layout: "ibm"
page_title: "IBM : ibm_toolchain_tool_insights"
description: |-
  Manages toolchain_tool_insights.
subcategory: "IBM Toolchain API"
---

# ibm_toolchain_tool_insights

Provides a resource for toolchain_tool_insights. This allows toolchain_tool_insights to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_toolchain_tool_insights" "toolchain_tool_insights" {
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
* `parameters_references` - (Optional, Map) Decoded values used on provision in the broker that reference fields in the parameters.
* `toolchain_id` - (Required, Forces new resource, String) 

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_tool_insights.
* `dashboard_url` - (Required, String) The URL of a user-facing user interface for this instance of a service.

## Import

You can import the `ibm_toolchain_tool_insights` resource by using `instance_id`. The id of the created or updated service instance.

# Syntax
```
$ terraform import ibm_toolchain_tool_insights.toolchain_tool_insights <instance_id>
```

# Example
```
$ terraform import ibm_toolchain_tool_insights.toolchain_tool_insights 4f107490-3820-400b-a008-f7f38d4163ed
```
