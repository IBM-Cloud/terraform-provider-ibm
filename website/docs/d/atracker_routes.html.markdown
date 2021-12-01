---
layout: "ibm"
page_title: "IBM : ibm_atracker_routes"
description: |-
  Get information about atracker_routes
subcategory: "Activity Tracker"
---

# ibm_atracker_routes

Provides a read-only data source for atracker_routes. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_atracker_routes" "atracker_routes" {
	name = "my-route"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `name` - (Optional, String) The name of the route.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the atracker_routes.
* `routes` - (Required, List) A list of route resources.
Nested scheme for **routes**:
	* `id` - (Required, String) The uuid of the route resource.
	* `name` - (Required, String) The name of the route.
	* `crn` - (Required, String) The crn of the route resource.
	* `version` - (Optional, Integer) The version of the route.
	* `receive_global_events` - (Required, Boolean) Indicates whether or not all global events should be forwarded to this region.
	* `rules` - (Required, List) The routing rules that will be evaluated in their order of the array.
	Nested scheme for **rules**:
		* `target_ids` - (Required, List) The target ID List. Only 1 target id is supported.
	* `created` - (Optional, String) The timestamp of the route creation time.
	* `updated` - (Optional, String) The timestamp of the route last updated time.

