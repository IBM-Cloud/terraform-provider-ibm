---
layout: "ibm"
page_title: "IBM : ibm_metrics_router_routes"
description: |-
  Get information about metrics_router_routes
subcategory: "Metrics Router"
---

# ibm_metrics_router_routes

Provides a read-only data source for metrics_router_routes. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_metrics_router_routes" "metrics_router_routes" {
	name = ibm_metrics_router_route.metrics_router_route.name
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `name` - (Optional, String) The name of the route.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the metrics_router_routes.
* `routes` - (List) A list of route resources.
  * Constraints: The maximum length is `4` items. The minimum length is `0` items.
Nested scheme for **routes**:
	* `api_version` - (Integer) The API version of the route.
	* `created_at` - (String) The timestamp of the route creation time.
	* `crn` - (String) The crn of the route resource.
	* `id` - (String) The UUID of the route resource.
	* `message` - (String) An optional message containing information about the route.
	* `name` - (String) The name of the route.
	* `rules` - (List) The routing rules that will be evaluated in their order of the array. Once a rule is matched, the remaining rules in the route definition will be skipped.
	  * Constraints: The maximum length is `4` items. The minimum length is `0` items.
	Nested scheme for **rules**:
		* `inclusion_filters` - (List) A list of conditions to be satisfied for routing events to pre-defined target.
		  * Constraints: The maximum length is `7` items. The minimum length is `0` items.
		Nested scheme for **inclusion_filters**:
			* `operand` - (String) Part of CRN that can be compared with values.
			  * Constraints: Allowable values are: `location`, `service_name`, `service_instance`, `resource_type`, `resource`.
			* `operator` - (String) The operation to be performed between operand and the provided values. 'is' to be used with one value and 'in' can support upto 20 values in the array.
			  * Constraints: Allowable values are: `is`, `in`.
			* `value` - (List) The provided values of the operand to be compared with. With 'is' operator, a single string is also supported with minLength 2 and maxLength 100.
			  * Constraints: The maximum length is `20` items. The minimum length is `1` item.
		* `target_ids` - (List) The target ID List. All the events will be sent to all targets listed in the rule. You can include targets from other regions.
		  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9 -._:]+$/`. The maximum length is `3` items. The minimum length is `0` items.
	* `updated_at` - (String) The timestamp of the route last updated time.
	* `version` - (Integer) The version of the route.

