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
  receive_global_events = false
  rules {
    target_ids = [ ibm_atracker_target.atracker_target.id ]
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

* `name` - (Required, String) The name of the route. The name must be 1000 characters or less and cannot include any special characters other than `(space) - . _ :`.
  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 -._:]+$/`
* `receive_global_events` - (Required, Boolean) Indicates whether or not all global events should be forwarded to this region.
* `rules` - (Required, List) Routing rules that will be evaluated in their order of the array.
Nested scheme for **rules**:
	* `target_ids` - (Required, List) The target ID List. Only 1 target id is supported.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the Activity Tracker Route.
* `created` - (Optional, String) The timestamp of the route creation time.
* `crn` - (Required, String) The crn of the route resource.
* `updated` - (Optional, String) The timestamp of the route last updated time.
* `version` - (Optional, Integer) The version of the route.

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
