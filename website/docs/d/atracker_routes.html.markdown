---
layout: "ibm"
page_title: "IBM : ibm_atracker_routes"
description: |-
  Get information about atracker_routes
subcategory: "Activity Tracker Event Routing"
---

# ibm_atracker_routes

Provides a read-only data source to retrieve information about atracker_routes. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```terraform
data "ibm_atracker_routes" "atracker_routes" {
	name = "my-route"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `name` - (Optional, String) The name of the route.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the atracker_routes.
* `routes` - (List) A list of route resources.
  * Constraints: The maximum length is `30` items. The minimum length is `0` items.
Nested schema for **routes**:
	* `api_version` - (Integer) The API version of the route.
	  * Constraints: The maximum value is `2`. The minimum value is `2`.
	* `created_at` - (String) The timestamp of the route creation time.
	* `crn` - (String) The crn of the route resource.
	* `id` - (String) The uuid of the route resource.
	* `message` - (String) An optional message containing information about the route.
	* `name` - (String) The name of the route.
	* `rules` - (List) The routing rules that will be evaluated in their order of the array. Once a rule is matched, the remaining rules in the route definition will be skipped.
	  * Constraints: The maximum length is `10` items. The minimum length is `1` item.
	Nested schema for **rules**:
		* `target_ids` - (List) The target ID List. All the events will be send to all targets listed in the rule. You can include targets from other regions.
		  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9 -._:]+$/`.
		* `locations` - (List) Logs from these locations will be sent to the targets specified. Locations is a superset of regions including global and *.
		  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9 -._:]+$/`.
	* `created_at` - (String) The timestamp of the route creation time.
	* `updated_at` - (String) The timestamp of the route last updated time.
	* `version` - (Integer) The version of the route.
