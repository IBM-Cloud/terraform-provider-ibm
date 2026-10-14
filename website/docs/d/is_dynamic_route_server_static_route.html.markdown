---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dynamic_route_server_static_route"
description: |-
  Get information about an IBM dynamic route server static route.
---

# ibm_is_dynamic_route_server_static_route

Provides a read-only data source to retrieve information about a dynamic route server static route. For more information, see [about dynamic route server](https://cloud.ibm.com/docs/vpc?topic=vpc-about-dynamic-route-server).

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
data "ibm_is_dynamic_route_server_static_route" "example" {
  dynamic_route_server_id                 = ibm_is_dynamic_route_server.example.id
  dynamic_route_server_static_route_id = ibm_is_dynamic_route_server_static_route.example.dynamic_route_server_static_route_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `dynamic_route_server_id` - (Required, String) The dynamic route server identifier.
* `dynamic_route_server_static_route_id` - (Required, String) The dynamic route server static route identifier.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the data source, in the format `<dynamic_route_server_id>/<static_route_id>`.
* `added_routes` - (List) The routes added to the VPC routing tables for this static route.

  Nested schema for **added_routes**:
  * `vpc_routing_table_route` - (List) The route reference for the applied route. If absent, the route could not be added to the routing table.

    Nested schema for **vpc_routing_table_route**:
    * `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.

      Nested schema for **deleted**:
      * `more_info` - (String) A link to documentation about deleted resources.
    * `href` - (String) The URL for this route.
    * `id` - (String) The unique identifier for this route.
    * `name` - (String) The name for this route.
* `destination` - (String) The destination CIDR of the route.
* `href` - (String) The URL for this static route.
* `lifecycle_state` - (String) The lifecycle state of the static route.
* `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).

  Nested schema for **lifecycle_reasons**:
  * `code` - (String) A reason code for this lifecycle state.
  * `message` - (String) An explanation of the reason for this lifecycle state.
  * `more_info` - (String) A link to documentation about the reason for this lifecycle state.
* `name` - (String) The name for this static route.
* `next_hops` - (List) The next hop peer groups.

  Nested schema for **next_hops**:
  * `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.

    Nested schema for **deleted**:
    * `more_info` - (String) A link to documentation about deleted resources.
  * `href` - (String) The URL for this peer group.
  * `id` - (String) The unique identifier for this peer group.
  * `name` - (String) The name for this peer group.
  * `resource_type` - (String) The resource type.
* `resource_type` - (String) The resource type.
* `route_delete_delay` - (Integer) The number of seconds to wait before deleting a route from `routing_tables` when the status of the route's associated peer is `down`.
* `routing_tables` - (List) The routing tables to update.

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
