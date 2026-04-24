---
layout: "ibm"
page_title: "IBM : ibm_logs_router_route"
description: |-
  Manages logs_router_route.
subcategory: "Logs Routing API Version 3"
---

# ibm_logs_router_route

Create, update, and delete logs_router_routes with this resource.

## Example Usage

```hcl
resource "ibm_logs_router_route" "logs_router_route_instance" {
  managed_by = "enterprise"
  name = "my-route"
  rules {
		action = "send"
		targets {
			id = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
		}
		inclusion_filters {
			operand = "location"
			operator = "is"
			values = [ "us-south" ]
		}
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `managed_by` - (Optional, String) Present when the route is enterprise-managed (`managed_by: enterprise`).
  * Constraints: Allowable values are: `enterprise`, `account`.
* `name` - (Required, String) The name of the route.
  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character.
* `rules` - (Required, List) The routing rules that will be evaluated in their order of the array. Once a rule is matched, the remaining rules in the route definition will be skipped.
  * Constraints: The maximum length is `10` items. The minimum length is `1` item.
Nested schema for **rules**:
	* `action` - (Optional, String) The action if the inclusion_filters matches, default is `send` action.
	  * Constraints: Allowable values are: `send`, `drop`.
	* `inclusion_filters` - (Required, List) A list of conditions to be satisfied for routing platform logs to pre-defined target.
	  * Constraints: The maximum length is `7` items. The minimum length is `0` items.
	Nested schema for **inclusion_filters**:
		* `operand` - (Required, String) Part of CRN that can be compared with values. Currently only location is supported.
		  * Constraints: Allowable values are: `location`.
		* `operator` - (Required, String) The operation to be performed between operand and the provided values. 'is' to be used with one value and 'in' can support up to 20 values in the array.
		  * Constraints: Allowable values are: `is`, `in`.
		* `values` - (Required, List) The provided string values of the operand to be compared with.
		  * Constraints: The maximum length is `20` items. The minimum length is `1` item.
	* `targets` - (Required, List) The target ID List. Platform logs will be sent to all targets listed in the rule. You can include targets from other regions.
	  * Constraints: The maximum length is `3` items. The minimum length is `0` items.
	Nested schema for **targets**:
		* `crn` - (Required, String) The CRN of a pre-defined logs-router target.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 \\-._:\/]+$/`.
		* `id` - (Required, String) The target uuid for a pre-defined platform logs router target.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 \\-._:]+$/`.
		* `name` - (Required, String) The name of a pre-defined logs-router target.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character.
		* `target_type` - (Required, String) The type of the target.
		  * Constraints: Allowable values are: `cloud_logs`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_router_route.
* `created_at` - (String) The timestamp of the route creation time.
* `crn` - (String) The crn of the route resource.
  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character.
* `updated_at` - (String) The timestamp of the route last updated time.


## Import

You can import the `ibm_logs_router_route` resource by using `id`. The UUID of the route resource.

# Syntax
<pre>
$ terraform import ibm_logs_router_route.logs_router_route &lt;id&gt;
</pre>

# Example
```
$ terraform import ibm_logs_router_route.logs_router_route c3af557f-fb0e-4476-85c3-0889e7fe7bc4
```
