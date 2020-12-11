---
layout: "ibm"
page_title: "IBM : atracker_routes"
sidebar_current: "docs-ibm-datasource-atracker-routes"
description: |-
  Get information about A list of route resources.
---

# ibm\_atracker_routes

Provides a read-only data source for A list of route resources.. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "atracker_routes" "atracker_routes" {
	name = "my-route"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) The name of this route.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the A list of route resources..
* `routes` - A list of route resources. Nested `routes` blocks have the following structure:
	* `id` - The uuid of this route resource.
	* `name` - The name of this route.
	* `instance_id` - The uuid of ATracker services in this region.
	* `crn` - The crn of this route type resource.
	* `version` - The version of this route.
	* `receive_global_events` - Whether or not all global events should be forwarded to this region.
	* `rules` - The routing rules that will be evaluated in their order of the array. Nested `rules` blocks have the following structure:
		* `target_ids` - The target ID List. Only one target id is supported. For regional route, the id must be V4 uuid of a target in the same region. For global route, it will be region-code and target-id separated by colon.

