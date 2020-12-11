---
layout: "ibm"
page_title: "IBM : atracker_route"
sidebar_current: "docs-ibm-resource-atracker-route"
description: |-
  Manages ATracker Route.
---

# ibm\_atracker_route

Provides a resource for ATracker Route. This allows ATracker Route to be created, updated and deleted.

## Example Usage

```hcl
resource "atracker_route" "atracker_route" {
  name = "my-route"
  receive_global_events = false
  rules = { example: "object" }
}
```

## Timeouts

atracker_route provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for creating a ATracker Route.
* `update` - (Default 20 minutes) Used for updating a ATracker Route.
* `delete` - (Default 10 minutes) Used for deleting a ATracker Route.

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the route. Must be 180 characters or less and cannot include any special characters other than `(space) - . _ :`.
* `receive_global_events` - (Required, bool) Whether or not all global events should be forwarded to this region.
* `rules` - (Required, List) Routing rules that will be evaluated in their order of the array.
  * `target_ids` - (Required, []interface{}) The target ID List. Only one target id is supported. For regional route, the id must be V4 uuid of a target in the same region. For global route, it will be region-code and target-id separated by colon.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the ATracker Route.
* `instance_id` - The uuid of ATracker services in this region.
* `crn` - The crn of this route type resource.
* `version` - The version of this route.
