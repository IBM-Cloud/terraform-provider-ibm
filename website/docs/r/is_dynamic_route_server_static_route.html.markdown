---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dynamic_route_server_static_route"
description: |-
  Manages IBM dynamic route server static route.
---

# ibm_is_dynamic_route_server_static_route

Create, update, and delete a dynamic route server static route. A static route adds a route for `destination` to each of the given VPC routing tables, with the peers of the `next_hops` peer groups as next hop. For more information, see [about dynamic route server](https://cloud.ibm.com/docs/vpc?topic=vpc-about-dynamic-route-server).

**Note:**
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

```terraform
resource "ibm_is_dynamic_route_server_static_route" "example" {
  dynamic_route_server_id = ibm_is_dynamic_route_server.example.id
  name                    = "example-static-route"
  destination             = "192.168.3.0/24"
  route_delete_delay      = 30
  next_hops {
    id = ibm_is_dynamic_route_server_peer_group.example.dynamic_route_server_peer_group_id
  }
  routing_tables {
    advertise = true
    vpc_routing_table {
      id = ibm_is_vpc_routing_table.example.routing_table
    }
  }
}
```

## Timeouts

The `ibm_is_dynamic_route_server_static_route` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

* `create` - (Default 10 minutes) Used for creating the static route and waiting for it to become `stable`.
* `update` - (Default 10 minutes) Used for updating the static route and waiting for it to become `stable`.
* `delete` - (Default 10 minutes) Used for deleting the static route and waiting for it to be removed.

## Argument Reference

You can specify the following arguments for this resource.

* `dynamic_route_server_id` - (Required, Forces new resource, String) The dynamic route server identifier.
* `destination` - (Required, Forces new resource, String) The destination CIDR of the route. The host identifier in the CIDR must be zero.
* `name` - (Optional, String) The name for this static route. The name must not be used by another static route for this dynamic route server. Names starting with `ibm-` are reserved. If unspecified, the name will be a hyphenated list of randomly-selected words.
* `next_hops` - (Required, List) The next hop peer groups. At least one item is required.

  Nested schema for **next_hops**:
  * `id` - (Optional, String) The unique identifier for the peer group. Exactly one of `id` or `href` must be set.
  * `href` - (Optional, String) The URL for the peer group.
* `route_delete_delay` - (Optional, Integer) The number of seconds to wait before deleting a route from `routing_tables` when the status of the route's associated peer is `down`. Must be between `0` and `60`. The default value is `0`.
* `routing_tables` - (Required, List) The routing tables to update. At least one item is required. A route is added to each routing table for each peer in the `next_hops` peer groups.

  Nested schema for **routing_tables**:
  * `advertise` - (Required, Boolean) Indicates whether this route will be advertised to the ingress sources specified by the `advertise_routes_to` routing table property.
  * `vpc_routing_table` - (Required, List) The VPC routing table. Exactly one item is allowed.

    Nested schema for **vpc_routing_table**:
    * `id` - (Optional, String) The unique identifier for the routing table. Exactly one of `id`, `href` or `crn` must be set.
    * `href` - (Optional, String) The URL for the routing table.
    * `crn` - (Optional, String) The CRN for the routing table.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the static route resource, in the format `<dynamic_route_server_id>/<static_route_id>`.
* `dynamic_route_server_static_route_id` - (String) The unique identifier for this static route.
* `added_routes` - (List) The routes added to the VPC routing tables for this static route.

  Nested schema for **added_routes**:
  * `vpc_routing_table_route` - (List) The route reference for the applied route. If absent, the route could not be added to the routing table.

    Nested schema for **vpc_routing_table_route**:
    * `id` - (String) The unique identifier for this route.
    * `href` - (String) The URL for this route.
    * `name` - (String) The name for this route.
    * `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.

      Nested schema for **deleted**:
      * `more_info` - (String) A link to documentation about deleted resources.
* `href` - (String) The URL for this static route.
* `lifecycle_state` - (String) The lifecycle state of the static route.
* `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).

  Nested schema for **lifecycle_reasons**:
  * `code` - (String) A reason code for this lifecycle state.
  * `message` - (String) An explanation of the reason for this lifecycle state.
  * `more_info` - (String) A link to documentation about the reason for this lifecycle state.
* `resource_type` - (String) The resource type.
* `etag` - (String) The ETag of the static route, used for conditional updates.
* `next_hops` items additionally expose `name`, `resource_type` and `deleted` (a list with `more_info`) as read from the service.
* `routing_tables.vpc_routing_table` items additionally expose `name`, `resource_type` and `deleted` (a list with `more_info`) as read from the service.

## Import

You can import the `ibm_is_dynamic_route_server_static_route` resource by using `id`.
The `id` property can be formed from `dynamic_route_server_id` and `dynamic_route_server_static_route_id` in the following format:

```
<dynamic_route_server_id>/<dynamic_route_server_static_route_id>
```

* `dynamic_route_server_id`: A string. The dynamic route server identifier.
* `dynamic_route_server_static_route_id`: A string. The static route identifier.

**Syntax**

```
$ terraform import ibm_is_dynamic_route_server_static_route.example <dynamic_route_server_id>/<dynamic_route_server_static_route_id>
```
