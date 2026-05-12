---
layout: "ibm"
page_title: "IBM : ibm_logs_router_routes"
description: |-
  Get information about logs_router_routes
subcategory: "Logs Routing API Version 3"
---

# ibm_logs_router_routes

Provides a read-only data source to retrieve information about logs_router_routes. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_router_routes" "logs_router_routes" {
	name = ibm_logs_router_route.logs_router_route_instance.name
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `name` - (Optional, String) The name of the route.
  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_router_routes.
* `routes` - (List) A list of route resources.
  * Constraints: The maximum length is `30` items. The minimum length is `0` items.
Nested schema for **routes**:
	* `created_at` - (String) The timestamp of the route creation time.
	* `crn` - (String) The crn of the route resource.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character.
	* `id` - (String) The UUID of the route resource.
	  * Constraints: The maximum length is `1028` characters. The minimum length is `24` characters.
	* `managed_by` - (String) Present when the route is enterprise-managed (`managed_by: enterprise`).
	  * Constraints: Allowable values are: `enterprise`, `account`.
	* `name` - (String) The name of the route.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character.
	* `rules` - (List) The routing rules that will be evaluated in their order of the array. Once a rule is matched, the remaining rules in the route definition will be skipped.
	  * Constraints: The maximum length is `10` items. The minimum length is `1` item.
	Nested schema for **rules**:
		* `action` - (String) The action if the inclusion_filters matches, default is `send` action.
		  * Constraints: Allowable values are: `send`, `drop`.
		* `inclusion_filters` - (List) A list of conditions to be satisfied for routing platform logs to pre-defined target.
		  * Constraints: The maximum length is `7` items. The minimum length is `0` items.
		Nested schema for **inclusion_filters**:
			* `operand` - (String) Part of CRN that can be compared with values. Currently only location is supported.
			  * Constraints: Allowable values are: `location`.
			* `operator` - (String) The operation to be performed between operand and the provided values. 'is' to be used with one value and 'in' can support up to 20 values in the array.
			  * Constraints: Allowable values are: `is`, `in`.
			* `values` - (List) The provided string values of the operand to be compared with.
			  * Constraints: The maximum length is `20` items. The minimum length is `1` item.
		* `targets` - (List) The target ID List. Platform logs will be sent to all targets listed in the rule. You can include targets from other regions.
		  * Constraints: The maximum length is `3` items. The minimum length is `0` items.
		Nested schema for **targets**:
			* `crn` - (String) The CRN of a pre-defined logs-router target.
			  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 \\-._:\/]+$/`.
			* `id` - (String) The target uuid for a pre-defined platform logs router target.
			  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 \\-._:]+$/`.
			* `name` - (String) The name of a pre-defined logs-router target.
			  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character.
			* `target_type` - (String) The type of the target.
			  * Constraints: Allowable values are: `cloud_logs`.
	* `updated_at` - (String) The timestamp of the route last updated time.

