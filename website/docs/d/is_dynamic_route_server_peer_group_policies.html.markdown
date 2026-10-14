---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dynamic_route_server_peer_group_policies"
description: |-
  Get information about the policies of an IBM dynamic route server peer group.
---

# ibm_is_dynamic_route_server_peer_group_policies

Provides a read-only data source to retrieve the policies of a dynamic route server peer group. For more information, see [about dynamic route server](https://cloud.ibm.com/docs/vpc?topic=vpc-about-dynamic-route-server).

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
data "ibm_is_dynamic_route_server_peer_group_policies" "example" {
  dynamic_route_server_id            = ibm_is_dynamic_route_server.example.id
  dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group.example.dynamic_route_server_peer_group_id
  sort                               = "name"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `dynamic_route_server_id` - (Required, String) The dynamic route server identifier.
* `dynamic_route_server_peer_group_id` - (Required, String) The dynamic route server peer group identifier.
* `sort` - (Optional, String) Sorts the returned collection by the specified property name in ascending order. A `-` may be prepended to the name to sort in descending order. Allowable values are: `created_at`, `name`. The default value is `-created_at`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the data source.
* `policies` - (List) The policies of the dynamic route server peer group. All pages of the collection are retrieved.

  Nested schema for **policies**:
  * `id` - (String) The unique identifier for this policy.
  * `created_at` - (String) The date and time that the policy was created or last updated.
  * `href` - (String) The URL for this policy.
  * `lifecycle_state` - (String) The lifecycle state of the policy.
  * `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).

    Nested schema for **lifecycle_reasons**:
    * `code` - (String) A reason code for this lifecycle state.
    * `message` - (String) An explanation of the reason for this lifecycle state.
    * `more_info` - (String) A link to documentation about the reason for this lifecycle state.
  * `name` - (String) The name for this policy.
  * `resource_type` - (String) The resource type.
  * `state` - (String) The state of the policy. One of `enabled` or `disabled`.
  * `type` - (String) The type of policy. One of `custom_routes`, `learned_routes`, `vpc_address_prefixes`, `vpc_routing_tables`.
  * `custom_routes` - (List) The custom routes to advertise (type `custom_routes`).

    Nested schema for **custom_routes**:
    * `destination` - (String) The destination CIDR of the route.
  * `excluded_prefixes` - (List) The prefixes excluded from monitoring when learning routes from connected peer groups.

    Nested schema for **excluded_prefixes**:
    * `ge` - (Integer) The minimum prefix length to match. Zero disables `ge` filtering.
    * `le` - (Integer) The maximum prefix length to match. Zero disables `le` filtering.
    * `prefix` - (String) The prefix to be excluded.
  * `peer_groups` - (List) The peer groups from which routes are learned (type `learned_routes`).

    Nested schema for **peer_groups**:
    * `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.

      Nested schema for **deleted**:
      * `more_info` - (String) A link to documentation about deleted resources.
    * `href` - (String) The URL for this peer group.
    * `id` - (String) The unique identifier for this peer group.
    * `name` - (String) The name for this peer group.
    * `resource_type` - (String) The resource type.
  * `next_hops` - (List) The next hop peer groups (type `vpc_routing_tables`).

    Nested schema for **next_hops**:
    * `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.

      Nested schema for **deleted**:
      * `more_info` - (String) A link to documentation about deleted resources.
    * `href` - (String) The URL for this peer group.
    * `id` - (String) The unique identifier for this peer group.
    * `name` - (String) The name for this peer group.
    * `resource_type` - (String) The resource type.
  * `route_delete_delay` - (Integer) The number of seconds to wait before deleting a route from `routing_tables` when the status of the route's associated peer is `down`, or after it is no longer advertised.
  * `routing_tables` - (List) The routing tables to update with the learned routes (type `vpc_routing_tables`).

    Nested schema for **routing_tables**:
    * `advertise` - (Boolean) Indicates whether this route will be advertised to the ingress sources specified by the `advertise_routes_to` routing table property.
    * `vpc_routing_table` - (List) The VPC routing table.

      Nested schema for **vpc_routing_table**:
      * `crn` - (String) The CRN for this routing table.
      * `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.

        Nested schema for **deleted**:
        * `more_info` - (String) A link to documentation about deleted resources.
      * `href` - (String) The URL for this routing table.
      * `id` - (String) The unique identifier for this routing table.
      * `name` - (String) The name for this routing table.
      * `resource_type` - (String) The resource type.
