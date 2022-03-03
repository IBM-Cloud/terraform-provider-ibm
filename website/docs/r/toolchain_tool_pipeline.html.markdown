---
layout: "ibm"
page_title: "IBM : ibm_toolchain_tool_pipeline"
description: |-
  Manages toolchain_tool_pipeline.
subcategory: "IBM Toolchain API"
---

# ibm_toolchain_tool_pipeline

Provides a resource for toolchain_tool_pipeline. This allows toolchain_tool_pipeline to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_toolchain_tool_pipeline" "toolchain_tool_pipeline" {
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
	* `name` - (Required, String)
	* `type` - (Required, String)
	  * Constraints: Allowable values are: `classic`, `tekton`.
	* `ui_pipeline` - (Optional, Boolean) When this check box is selected, the applications that this pipeline deploys are shown in the View app menu on the toolchain page. This setting is best for UI apps that can be accessed from a browser.
	  * Constraints: The default value is `false`.
* `parameters_references` - (Optional, Map) Decoded values used on provision in the broker that reference fields in the parameters.
* `toolchain_id` - (Required, Forces new resource, String) 

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_tool_pipeline.
* `dashboard_url` - (Required, String) The URL of a user-facing user interface for this instance of a service.

## Import

You can import the `ibm_toolchain_tool_pipeline` resource by using `instance_id`. The id of the created or updated service instance.

# Syntax
```
$ terraform import ibm_toolchain_tool_pipeline.toolchain_tool_pipeline <instance_id>
```

# Example
```
$ terraform import ibm_toolchain_tool_pipeline.toolchain_tool_pipeline 4f107490-3820-400b-a008-f7f38d4163ed
```
