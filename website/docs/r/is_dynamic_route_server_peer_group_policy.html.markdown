---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dynamic_route_server_peer_group_policy"
description: |-
  Manages IBM dynamic route server peer group policy.
---

# ibm_is_dynamic_route_server_peer_group_policy

Create, update, and delete a dynamic route server peer group policy. A policy controls which routes a dynamic route server advertises to, or learns from, the peers in a peer group. For more information, see [about dynamic route server](https://cloud.ibm.com/docs/vpc?topic=vpc-about-dynamic-route-server).

**Note:**
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

The policy `type` selects which other arguments are allowed:

| `type` | Allowed arguments |
| --- | --- |
| `custom_routes` | `custom_routes` |
| `learned_routes` | `peer_groups`, `excluded_prefixes` |
| `vpc_address_prefixes` | none |
| `vpc_routing_tables` | `next_hops`, `routing_tables`, `route_delete_delay`, `excluded_prefixes` |

Setting an argument that is not allowed for the selected `type` is rejected at plan time.

A policy that advertises custom routes:

```terraform
resource "ibm_is_dynamic_route_server_peer_group_policy" "example_custom_routes" {
  dynamic_route_server_id            = ibm_is_dynamic_route_server.example.id
  dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group.example.dynamic_route_server_peer_group_id
  name                               = "example-custom-routes-policy"
  type                               = "custom_routes"
  custom_routes {
    destination = "192.168.3.0/24"
  }
}
```

A policy that advertises routes learned from another peer group:

```terraform
resource "ibm_is_dynamic_route_server_peer_group_policy" "example_learned_routes" {
  dynamic_route_server_id            = ibm_is_dynamic_route_server.example.id
  dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group.example.dynamic_route_server_peer_group_id
  name                               = "example-learned-routes-policy"
  type                               = "learned_routes"
  peer_groups {
    id = ibm_is_dynamic_route_server_peer_group.source.dynamic_route_server_peer_group_id
  }
  excluded_prefixes {
    prefix = "10.10.0.0/16"
    ge     = 20
    le     = 24
  }
}
```

A policy that updates a VPC routing table with learned routes:

```terraform
resource "ibm_is_dynamic_route_server_peer_group_policy" "example_vpc_routing_tables" {
  dynamic_route_server_id            = ibm_is_dynamic_route_server.example.id
  dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group.example.dynamic_route_server_peer_group_id
  name                               = "example-vpc-routing-tables-policy"
  type                               = "vpc_routing_tables"
  route_delete_delay                 = 30
  next_hops {
    id = ibm_is_dynamic_route_server_peer_group.next_hop.dynamic_route_server_peer_group_id
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

The `ibm_is_dynamic_route_server_peer_group_policy` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

* `create` - (Default 10 minutes) Used for creating the policy and waiting for it to become `stable`.
* `update` - (Default 10 minutes) Used for updating the policy and waiting for it to become `stable`.
* `delete` - (Default 10 minutes) Used for deleting the policy and waiting for it to be removed.

## Argument Reference

You can specify the following arguments for this resource.

* `dynamic_route_server_id` - (Required, Forces new resource, String) The dynamic route server identifier.
* `dynamic_route_server_peer_group_id` - (Required, Forces new resource, String) The dynamic route server peer group identifier.
* `type` - (Required, Forces new resource, String) The type of policy. Allowable values are: `custom_routes`, `learned_routes`, `vpc_address_prefixes`, `vpc_routing_tables`.
* `name` - (Optional, String) The name for this policy. The name must not be used by another policy of the peer group. Names starting with `ibm-` are reserved. If unspecified, the name will be a hyphenated list of randomly-selected words.
* `state` - (Optional, String) The state of the policy. Allowable values are: `enabled`, `disabled`. The default value is `enabled`.
* `custom_routes` - (Optional, List) The custom routes to advertise. Only allowed when `type` is `custom_routes`.

  Nested schema for **custom_routes**:
  * `destination` - (Required, String) The destination CIDR of the route. The host identifier in the CIDR must be zero.
* `excluded_prefixes` - (Optional, List) The prefixes excluded from monitoring when learning routes from connected peer groups. They are applied in addition to any `excluded_prefixes` defined on the peer group. Only allowed when `type` is `learned_routes` or `vpc_routing_tables`.

  Nested schema for **excluded_prefixes**:
  * `prefix` - (Required, String) The prefix to be excluded.
  * `ge` - (Optional, Integer) The minimum prefix length to match. Zero (the default) disables `ge` filtering.
  * `le` - (Optional, Integer) The maximum prefix length to match. Zero (the default) disables `le` filtering.
* `peer_groups` - (Optional, List) The peer groups from which routes are learned and then advertised to the peer group this policy is applied to. Only allowed when `type` is `learned_routes`.

  Nested schema for **peer_groups**:
  * `id` - (Optional, String) The unique identifier for the peer group. Exactly one of `id` or `href` must be set.
  * `href` - (Optional, String) The URL for the peer group.
* `next_hops` - (Optional, List) The next hop peer groups. Only allowed when `type` is `vpc_routing_tables`.

  Nested schema for **next_hops**:
  * `id` - (Optional, String) The unique identifier for the peer group. Exactly one of `id` or `href` must be set.
  * `href` - (Optional, String) The URL for the peer group.
* `route_delete_delay` - (Optional, Integer) The number of seconds to wait before deleting a route from `routing_tables` when the status of the route's associated peer is `down`, or after it is no longer advertised by peers in the peer group. Must be between `0` and `60`. The default value is `0`. Only allowed when `type` is `vpc_routing_tables`.
* `routing_tables` - (Optional, List) The routing tables to update with the learned routes. Only allowed when `type` is `vpc_routing_tables`.

  Nested schema for **routing_tables**:
  * `advertise` - (Required, Boolean) Indicates whether this route will be advertised to the ingress sources specified by the `advertise_routes_to` routing table property.
  * `vpc_routing_table` - (Required, List) The VPC routing table. Exactly one item is allowed.

    Nested schema for **vpc_routing_table**:
    * `id` - (Optional, String) The unique identifier for the routing table. Exactly one of `id`, `href` or `crn` must be set.
    * `href` - (Optional, String) The URL for the routing table.
    * `crn` - (Optional, String) The CRN for the routing table.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the policy resource, in the format `<dynamic_route_server_id>/<dynamic_route_server_peer_group_id>/<policy_id>`.
* `dynamic_route_server_peer_group_policy_id` - (String) The unique identifier for this policy.
* `created_at` - (String) The date and time that the policy was created or last updated.
* `href` - (String) The URL for this policy.
* `lifecycle_state` - (String) The lifecycle state of the policy.
* `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).

  Nested schema for **lifecycle_reasons**:
  * `code` - (String) A reason code for this lifecycle state.
  * `message` - (String) An explanation of the reason for this lifecycle state.
  * `more_info` - (String) A link to documentation about the reason for this lifecycle state.
* `resource_type` - (String) The resource type.
* `etag` - (String) The ETag of the policy, used for conditional updates.
* `peer_groups` and `next_hops` items additionally expose `name`, `resource_type` and `deleted` (a list with `more_info`) as read from the service.
* `routing_tables.vpc_routing_table` items additionally expose `name`, `resource_type` and `deleted` (a list with `more_info`) as read from the service.

## Import

You can import the `ibm_is_dynamic_route_server_peer_group_policy` resource by using `id`.
The `id` property can be formed from `dynamic_route_server_id`, `dynamic_route_server_peer_group_id`, and `dynamic_route_server_peer_group_policy_id` in the following format:

```
<dynamic_route_server_id>/<dynamic_route_server_peer_group_id>/<dynamic_route_server_peer_group_policy_id>
```

* `dynamic_route_server_id`: A string. The dynamic route server identifier.
* `dynamic_route_server_peer_group_id`: A string. The dynamic route server peer group identifier.
* `dynamic_route_server_peer_group_policy_id`: A string. The policy identifier.

**Syntax**

```
$ terraform import ibm_is_dynamic_route_server_peer_group_policy.example <dynamic_route_server_id>/<dynamic_route_server_peer_group_id>/<dynamic_route_server_peer_group_policy_id>
```
