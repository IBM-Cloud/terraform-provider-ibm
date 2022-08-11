---
layout: "ibm"
page_title: "IBM : ibm_atracker_route"
description: |-
  Manages Activity Tracker Route.
subcategory: "Activity Tracker"
---

# ibm_atracker_route

Provides a resource for Activity Tracker Route. This allows Activity Tracker Route to be created, updated and deleted.

## Example usage

```terraform
resource "ibm_atracker_route" "atracker_route" {
  name = "my-route"
  rules {
    target_ids = [ ibm_atracker_target.atracker_target.id ]
    locations = [ "us-south", "global" ]
  }
  lifecycle {
    # Recommended to ensure that if a target ID is removed here and destroyed in a plan, this is updated first
    create_before_destroy = true
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

* `name` - (Required, String) The name of the route. The name must be 1000 characters or less and cannot include any special characters other than `(space) - . _ :`.
  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 -._:]+$/`
* `receive_global_events` - **DEPRECATED** (Optional, Boolean) Indicates whether or not all global events should be forwarded to this region.  Use rules.locations instead with `global` included.
* `rules` - (Required, List) Routing rules that will be evaluated in their order of the array.
Nested scheme for **rules**:
	* `target_ids` - (Required, List) The target ID List. All the events will be send to all targets listed in the rule. You can include targets from other regions.
	* `locations` - (Optional, List) Logs from these locations will be sent to the targets specified. Locations is a superset of regions including global and *.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the atracker_route.
* `api_version` - (Integer) The API version of the route.
* `created_at` - (String) The timestamp of the route creation time.
* `crn` - (Required, String) The crn of the route resource.
* `updated_at` - (String) The timestamp of the route last updated time.
* `version` - (Optional, Integer) The version of the route.
* `updated` - **DEPRECATED** (Optional, String) The timestamp of the route last updated time.
* `created` - **DEPRECATED** (Optional, String) The timestamp of the route creation time.

## Import

You can import the `ibm_atracker_route` resource by using `id`. The uuid of the route resource.

# Syntax
```
$ terraform import ibm_atracker_route.atracker_route <id>
```

# Example
```
$ terraform import ibm_atracker_route.atracker_route c3af557f-fb0e-4476-85c3-0889e7fe7bc4
```
