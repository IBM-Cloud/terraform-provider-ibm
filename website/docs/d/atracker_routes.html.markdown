---
layout: "ibm"
page_title: "IBM : ibm_atracker_routes"
description: |-
  Get information about atracker_routes
subcategory: "Activity Tracking API"
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

The following arguments are supported:

* `name` - (Optional, string) The name of the route.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the atracker_routes.
* `routes` - A list of route resources. Nested `routes` blocks have the following structure:
	* `id` - The uuid of the route resource.
	* `name` - The name of the route.
	* `crn` - The crn of the route resource.
	* `version` - The version of the route.
	* `receive_global_events` - Indicates whether or not all global events should be forwarded to this region.
	* `rules` - The routing rules that will be evaluated in their order of the array. Nested `rules` blocks have the following structure:
		* `target_ids` - The target ID List. Only 1 target id is supported.
	* `created` - The timestamp of the route creation time.
	* `updated` - The timestamp of the route last updated time.

