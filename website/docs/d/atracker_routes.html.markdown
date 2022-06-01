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

```terraform
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
* `routes` - (List) A list of route resources.
Nested scheme for **routes**:
	* `api_version` - (Integer) The API version of the route.
	* `crn` - (String) The crn of the route resource.
	* `id` - (String) The uuid of the route resource.
	* `name` - (String) The name of the route.
	* `rules` - (List) The routing rules that will be evaluated in their order of the array. Once a rule is matched, the remaining rules in the route definition will be skipped.
	Nested scheme for **rules**:
		* `target_ids` - (List) The target ID List. All the events will be send to all targets listed in the rule. You can include targets from other regions.
		  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9 -._:]+$/`.
		* `locations` - (List) Logs from these locations will be sent to the targets specified. Locations is a superset of regions including global and *.
		  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9 -._:]+$/`.
	* `created_at` - (String) The timestamp of the route creation time.
	* `updated_at` - (String) The timestamp of the route last updated time.
	* `version` - (Integer) The version of the route.
	* `created` - **DEPRECATED** (String) The timestamp of the route creation time.
	* `updated` - **DEPRECATED** (String) The timestamp of the route last updated time.
	* `receive_global_events` - **DEPRECATED** (Boolean) Indicates whether or not all global events should be forwarded to this region.
