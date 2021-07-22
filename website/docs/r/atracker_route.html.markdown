---
layout: "ibm"
page_title: "IBM : ibm_atracker_route"
description: |-
  Manages Activity Tracking Route.
subcategory: "Activity Tracking API"
---

# ibm_atracker_route

Provides a resource for Activity Tracking Route. This allows Activity Tracking Route to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_atracker_route" "atracker_route" {
  name = "my-route"
  receive_global_events = false
  rules = { "target_ids" : [ "target_ids" ] }
}
```

## Timeouts

atracker_route provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for creating a Activity Tracking Route.
* `update` - (Default 20 minutes) Used for updating a Activity Tracking Route.
* `delete` - (Default 10 minutes) Used for deleting a Activity Tracking Route.

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the route. The name must be 1000 characters or less and cannot include any special characters other than `(space) - . _ :`.
  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 -._:]+$/`
* `receive_global_events` - (Required, bool) Indicates whether or not all global events should be forwarded to this region.
* `rules` - (Required, List) Routing rules that will be evaluated in their order of the array.
  * `target_ids` - (Required, []interface{}) The target ID List. Only 1 target id is supported.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the Activity Tracking Route.
* `created` - The timestamp of the route creation time.
* `crn` - The crn of the route resource.
* `updated` - The timestamp of the route last updated time.
* `version` - The version of the route.

## Import

You can import the `ibm_atracker_route` resource by using `id`. The uuid of the route resource.

```
$ terraform import ibm_atracker_route.atracker_route c3af557f-fb0e-4476-85c3-0889e7fe7bc4
```
