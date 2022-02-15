---
layout: "ibm"
page_title: "IBM : ibm_toolchain_integration"
description: |-
  Manages toolchain_integration.
subcategory: "IBM Toolchain API"
---

# ibm_toolchain_integration

Provides a resource for toolchain_integration. This allows toolchain_integration to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_toolchain_integration" "toolchain_integration" {
  service_id = "todolist"
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
* `parameters` - (Optional, String) 
* `parameters_references` - (Optional, Map) Decoded values used on provision in the broker that reference fields in the parameters.
* `service_id` - (Required, Forces new resource, String) The unique short name of the service that should be provisioned.
  * Constraints: The maximum length is `100` characters.
* `toolchain_id` - (Required, Forces new resource, String) 

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_integration.
* `dashboard_url` - (Required, String) The URL of a user-facing user interface for this instance of a service.

## Import

You can import the `ibm_toolchain_integration` resource by using `instance_id`. The id of the created or updated service instance.

# Syntax
```
$ terraform import ibm_toolchain_integration.toolchain_integration <instance_id>
```

# Example
```
$ terraform import ibm_toolchain_integration.toolchain_integration 4f107490-3820-400b-a008-f7f38d4163ed
```
