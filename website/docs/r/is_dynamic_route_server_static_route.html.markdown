---
layout: "ibm"
page_title: "IBM : ibm_is_dynamic_route_server_static_route"
description: |-
  Manages DynamicRouteServerStaticRoute.
subcategory: "Virtual Private Cloud API"
---

# ibm_is_dynamic_route_server_static_route

Create, update, and delete DynamicRouteServerStaticRoutes with this resource.

## Example Usage

```hcl
resource "ibm_is_dynamic_route_server_static_route" "is_dynamic_route_server_static_route_instance" {
  destination = "192.168.3.0/24"
  dynamic_route_server_id = "dynamic_route_server_id"
  name = "my-dynamic-route-server-static-route"
  next_hops {
		deleted {
			more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
		}
		href = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"
		id = "r006-8228af60-189d-11ed-861d-0242ac120004"
		name = "my-dynamic-route-server-peer-group"
		resource_type = "dynamic_route_server_peer_group"
  }
  route_delete_delay = 1
  routing_tables {
		advertise = true
		vpc_routing_table {
			crn = "crn:v1:bluemix:public:is:us-south:a/aa2432b1fa4d4ace891e9b80fc104e34::vpc-routing-table:r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/r006-6885e83f-03b2-4603-8a86-db2a0f55c840"
			deleted {
				more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
			}
			href = "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-4727d842-f94f-4a2d-824a-9bc9b02c523b/routing_tables/r006-6885e83f-03b2-4603-8a86-db2a0f55c840"
			id = "r006-6885e83f-03b2-4603-8a86-db2a0f55c840"
			name = "my-routing-table-1"
			resource_type = "routing_table"
		}
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `destination` - (Required, String) The destination CIDR of the route. The host identifier in the CIDR must be zero.At most two routes per `zone` in a VPC routing table can have the same `destination` and `priority`, and only if both routes have an `action` of `deliver` and the`next_hop` is an IP address.
  * Constraints: The maximum length is `18` characters. The minimum length is `9` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$/`.
* `dynamic_route_server_id` - (Required, Forces new resource, String) The dynamic route server identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `name` - (Optional, String) The name for this static route. The name must not be used by another static route for this dynamic route server. Names starting with ibm- are reserved for ibm managed policies, and are not allowed. If unspecified, the name will be a hyphenated list of randomly-selected words.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
* `next_hops` - (Required, List) The next hop resources.
  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
Nested schema for **next_hops**:
	* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		* `more_info` - (Required, String) A link to documentation about deleted resources.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `href` - (Optional, String) The URL for this dynamic route server peer group.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (Optional, String) The ID for this dynamic route server peer group peer.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `name` - (Optional, String) The name for this dynamic route server peer group.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	* `resource_type` - (Optional, String) The resource type.
	  * Constraints: Allowable values are: `dynamic_route_server_peer_group`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `route_delete_delay` - (Optional, Integer) The number of seconds to wait before deleting a route from `routing_tables` when the `status` of the route's associated peer is `down`.
  * Constraints: The maximum value is `60`. The minimum value is `0`.
* `routing_tables` - (Required, List) The routing tables to update.A route is added to each routing table for each peer in the `next_hop` resource. Each route uses its associated peer's `endpoint.address` for the route's `next_hop`, and:- For ingress routing tables, the route's `priority` uses the peer's `priority`.- For egress routing tables, the route's `priority` is determined by  [Zone priority rules](https://cloud.ibm.com/docs/__TBD__).
Nested schema for **routing_tables**:
	* `advertise` - (Required, Boolean) Indicates whether this route will be advertised to the ingress sources specified by the `advertise_routes_to` routing table property.
	* `vpc_routing_table` - (Required, List)
	Nested schema for **vpc_routing_table**:
		* `crn` - (Required, String) The CRN for this VPC routing table.
		  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
		* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested schema for **deleted**:
			* `more_info` - (Required, String) A link to documentation about deleted resources.
			  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `href` - (Required, String) The URL for this routing table.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `id` - (Required, String) The unique identifier for this routing table.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
		* `name` - (Required, String) The name for this routing table. The name is unique across all routing tables for the VPC.
		  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
		* `resource_type` - (Required, String) The resource type.
		  * Constraints: Allowable values are: `routing_table`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the DynamicRouteServerStaticRoute.
* `added_routes` - (List) The routes added to the VPC routing tables for this dynamic route server static route.
  * Constraints: The minimum length is `0` items.
Nested schema for **added_routes**:
	* `vpc_routing_table_route` - (List) The route reference for the applied route, If absent the routecould not be added to the routing table.For more information, see[Dynamic Route Server Static Route Failure](https://cloud.ibm.com/docs/__TBD__).
	Nested schema for **vpc_routing_table_route**:
		* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested schema for **deleted**:
			* `more_info` - (String) A link to documentation about deleted resources.
			  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `href` - (String) The URL for this route.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `id` - (String) The unique identifier for this route.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
		* `name` - (String) The name for this route. The name is unique across all routes in the routing table.
		  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
* `href` - (String) The URL for this dynamic route server static routes.
  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `is_dynamic_route_server_static_route_id` - (String) The unique identifier for this dynamic route server static routes.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
  * Constraints: The minimum length is `0` items.
Nested schema for **lifecycle_reasons**:
	* `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `internal_error`, `resource_suspended_by_provider`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `message` - (String) An explanation of the reason for this lifecycle state.
	* `more_info` - (String) A link to documentation about the reason for this lifecycle state.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `lifecycle_state` - (String) The lifecycle state of the static route.
  * Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `resource_type` - (String) The resource type.
  * Constraints: Allowable values are: `dynamic_route_server_static_route`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

* `etag` - ETag identifier for DynamicRouteServerStaticRoute.

## Import

You can import the `ibm_is_dynamic_route_server_static_route` resource by using `id`.
The `id` property can be formed from `dynamic_route_server_id`, and `is_dynamic_route_server_static_route_id` in the following format:

<pre>
&lt;dynamic_route_server_id&gt;/&lt;is_dynamic_route_server_static_route_id&gt;
</pre>
* `dynamic_route_server_id`: A string. The dynamic route server identifier.
* `is_dynamic_route_server_static_route_id`: A string in the format `r006-8228af60-189d-11ed-861d-0242ac120002`. The unique identifier for this dynamic route server static routes.

# Syntax
<pre>
$ terraform import ibm_is_dynamic_route_server_static_route.is_dynamic_route_server_static_route &lt;dynamic_route_server_id&gt;/&lt;is_dynamic_route_server_static_route_id&gt;
</pre>
