---
layout: "ibm"
page_title: "IBM : ibm_atracker_route"
description: |-
  Manages atracker_route.
subcategory: "Activity Tracker API Version 2"
---

# ibm_atracker_route

Create, update, and delete atracker_routes with this resource.

## Example Usage

```hcl
resource "ibm_atracker_route" "atracker_route_instance" {
  managed_by = "enterprise"
  name = "my-route"
  rules {
		target_ids = [ "c3af557f-fb0e-4476-85c3-0889e7fe7bc4" ]
		locations = [ "us-south" ]
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `managed_by` - (Optional, String) Present when the route is enterprise-managed (`managed_by: enterprise`).
  * Constraints: Allowable values are: `enterprise`, `account`.
* `name` - (Required, String) The name of the route.
* `rules` - (Required, List) The routing rules that will be evaluated in their order of the array. Once a rule is matched, the remaining rules in the route definition will be skipped.
  * Constraints: The maximum length is `10` items. The minimum length is `1` item.
Nested schema for **rules**:
	* `locations` - (Required, List) Logs from these locations will be sent to the targets specified. Locations is a superset of regions including global and *.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9 -._:]+$/`.
	* `target_ids` - (Required, List) The target ID List. All the events will be send to all targets listed in the rule. You can include targets from other regions.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9 -._:]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the atracker_route.
* `api_version` - (Integer) The API version of the route.
  * Constraints: The maximum value is `2`. The minimum value is `2`.
* `created_at` - (String) The timestamp of the route creation time.
* `crn` - (String) The crn of the route resource.
* `message` - (String) An optional message containing information about the route.
* `updated_at` - (String) The timestamp of the route last updated time.
* `version` - (Integer) The version of the route.
  * Constraints: The maximum value is `99999`. The minimum value is `0`.


## Import

You can import the `ibm_atracker_route` resource by using `id`. The uuid of the route resource.

# Syntax
<pre>
$ terraform import ibm_atracker_route.atracker_route &lt;id&gt;
</pre>

# Example
```
$ terraform import ibm_atracker_route.atracker_route c3af557f-fb0e-4476-85c3-0889e7fe7bc4
```
