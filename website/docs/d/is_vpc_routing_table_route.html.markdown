---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_vpc_routing_table_route"
description: |-
  Get information about VPC routing table route.
---

# ibm_is_vpc_routing_table_route

Provides a read-only data source for Route. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.For more information, about VPC default routing table, see [about routing tables and routes](https://cloud.ibm.com/docs/vpc?topic=vpc-about-custom-routes).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage (using route id)

```terraform
data "ibm_is_vpc_routing_table_route" "example_route" {
  vpc 			= ibm_is_vpc.example_vpc.id
  routing_table = ibm_is_vpc_routing_table.example_rt.routing_table
  route_id 		= ibm_is_vpc_routing_table_route.example_route.route_id
}
```

## Example Usage (using route name)
```terraform		
data "ibm_is_vpc_routing_table_route" "example_route_name" {
  vpc 			= ibm_is_vpc.example_vpc.id
  routing_table = ibm_is_vpc_routing_table.example_rt.routing_table
  name 			= ibm_is_vpc_routing_table_route.example_route.name
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `name` - (Optional, String) The VPC routing table name. Mutually exclusive with `routing_table`, one of them is required
- `route_id` - (Required, String) The VPC routing table route identifier.
- `routing_table` - (Optional, String) The VPC routing table identifier. Mutually exclusive with `name`, one of them is required
- `vpc` - (Required, String) The VPC identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `action` - (String) The action to perform with a packet matching the route, allowable values are: `delegate`, `delegate_vpc`, `deliver`, `drop`.
- `advertise` - (Boolean) Indicates whether this route will be advertised to the ingress sources specified by the `advertise_routes_to` routing table property.
	- `delegate`: delegate to the system's built-in routes
	- `delegate_vpc`: delegate to the system's built-in routes, ignoring Internet-bound  routes
	- `deliver`: deliver the packet to the specified `next_hop`
	- `drop`: drop the packet.
- `created_at` - (String) The date and time that the route was created.
- `creator` - (List) If present, the resource that created the route. Routes with this property present cannot bedirectly deleted. All routes with an `origin` of `learned` or `service` will have thisproperty set, and future `origin` values may also have this property set.
Nested scheme for **creator**:
    - `crn` - (String) The VPN gateway's CRN.
      - Constraints: The maximum length is `512` characters. The minimum length is `9` characters.
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
    Nested scheme for **deleted**:
        - `more_info` - (Required, String) Link to documentation about deleted resources.
          - Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
    - `href` - (String) The VPN gateway's canonical URL.
      - Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
    - `id` - (String) The unique identifier for this VPN gateway.
      - Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
    - `name` - (String) The user-defined name for this VPN gateway.
      - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
    - `resource_type` - (String) The resource type.
      - Constraints: Allowable values are: `vpn_gateway`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
- `destination` - (String) The destination of the route.
- `href` - (String) The URL for this route.
- `id` - (String) The unique identifier of the Route.
- `lifecycle_state` - (String) The lifecycle state of the route.
  - Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`.
- `name` - (String) The user-defined name for this route.
- `next_hop` - (List) If `action` is `deliver`, the next hop that packets will be delivered to.  For other `action` values, its `address` will be `0.0.0.0`.
	Nested scheme for **next_hop**:
	- `address` - (String) The IP address.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
		Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The VPN connection's canonical URL.
	- `id` - (String) The unique identifier for this VPN gateway connection.
	- `name` - (String) The user-defined name for this VPN connection.
	- `resource_type` - (String) The resource type.

- `origin` - (String) The origin of this route:- `service`: route was directly created by a service- `user`: route was directly created by a userThe enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the route on which the unexpected property value was encountered.
  - Constraints: Allowable values are: `learned`, `service`, `user`.
- `priority` - (Integer) The route's priority. Smaller values have higher priority. If a routing table contains routes with the same destination, the route with the highest priority (smallest value) is selected. For Example (2), supports values from 0 to 4. Default is 2.
- `zone` - (List) The zone the route applies to. (Traffic from subnets in this zone will be subject to this route).
	Nested scheme for **zone**:
	- `href` - (String) The URL for this zone.
	- `name` - (String) The globally unique name for this zone.
