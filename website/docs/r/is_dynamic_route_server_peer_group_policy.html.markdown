---
layout: "ibm"
page_title: "IBM : ibm_is_dynamic_route_server_peer_group_policy"
description: |-
  Manages is_dynamic_route_server_peer_group_policy.
subcategory: "Virtual Private Cloud API"
---

# ibm_is_dynamic_route_server_peer_group_policy

Create, update, and delete is_dynamic_route_server_peer_group_policys with this resource.

## Example Usage

```hcl
resource "ibm_is_dynamic_route_server_peer_group_policy" "is_dynamic_route_server_peer_group_policy_instance" {
  custom_routes {
		destination = "192.168.3.0/24"
  }
  dynamic_route_server_id = "dynamic_route_server_id"
  dynamic_route_server_peer_group_id = "dynamic_route_server_peer_group_id"
  excluded_prefixes {
		ge = 0
		le = 32
		prefix = "prefix"
  }
  name = "my-dynamic-route-server-peer-group-policy"
  next_hops {
		deleted {
			more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
		}
		href = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"
		id = "r006-8228af60-189d-11ed-861d-0242ac120004"
		name = "my-dynamic-route-server-peer-group"
		resource_type = "dynamic_route_server_peer_group"
  }
  peer_groups {
		deleted {
			more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
		}
		href = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002"
		id = "r006-8228af60-189d-11ed-861d-0242ac120004"
		name = "my-dynamic-route-server-peer-group"
		resource_type = "dynamic_route_server_peer_group"
  }
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
  type = "custom_routes"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `custom_routes` - (Optional, List) The custom routes to advertise.
  * Constraints: The maximum length is `10` items. The minimum length is `1` item.
Nested schema for **custom_routes**:
	* `destination` - (Required, String) The destination CIDR of the route. The host identifier in the CIDR must be zero.
	  * Constraints: The maximum length is `18` characters. The minimum length is `9` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$/`.
* `dynamic_route_server_id` - (Required, Forces new resource, String) The dynamic route server identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `dynamic_route_server_peer_group_id` - (Required, Forces new resource, String) The dynamic route server peer group identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `excluded_prefixes` - (Optional, List) A list of prefixes that are excluded from monitoring by the dynamic route server when learning routes from connected peer groups. They are applied in addition to any `excluded_prefixes` defined on the peer group. Each excluded prefix must be a subset of the `monitored_prefixes` configured in the peer group.
Nested schema for **excluded_prefixes**:
	* `ge` - (Optional, Integer) The minimum prefix length to match. If non-zero, the prefix watcher will only exclude routes within the `prefix` that have a prefix length greater than or equal to this value.If zero, `ge` match filtering will not be applied.If non-zero, `ge` must be:- Greater than or equal to the network `prefix` length.- Less than or equal to 32.If both `ge` and `le` are non-zero, `ge` must be less than or equal to `le`.
	  * Constraints: The default value is `0`. The maximum value is `32`. The minimum value is `0`.
	* `le` - (Optional, Integer) The maximum prefix length to match. If non-zero, the prefix watcher will only match routes within the `prefix` that have a prefix length less than or equal to this value.If zero, `le` match filtering will not be applied.If non-zero, `le` value must be:- Less than or equal to the network `prefix` length.- Greater than `ge`.
	  * Constraints: The default value is `0`. The maximum value is `32`. The minimum value is `0`.
	* `prefix` - (Required, String) The prefix to be excluded.
* `name` - (Optional, String) The name for this dynamic route server peer group policy. The name must not be used by another peer group policy. Names starting with ibm- are reserved for ibm managed policies, and are not allowed. If unspecified, the name will be a hyphenated list of randomly-selected words.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
* `next_hops` - (Optional, List) The next hop resources.
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
* `peer_groups` - (Optional, List) The peer groups from which the routes are being learned and then advertised back to the peer group on which this peer group policy has been applied.
Nested schema for **peer_groups**:
	* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		* `more_info` - (Required, String) A link to documentation about deleted resources.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `href` - (Required, String) The URL for this dynamic route server peer group.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (Required, String) The ID for this dynamic route server peer group peer.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `name` - (Required, String) The name for this dynamic route server peer group.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	* `resource_type` - (Required, String) The resource type.
	  * Constraints: Allowable values are: `dynamic_route_server_peer_group`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `route_delete_delay` - (Optional, Integer) The number of seconds to wait before deleting a route from `routing_tables` when the`status` of the route's associated peer is `down`, or after it is no longer advertised by peers in the peer group.
  * Constraints: The maximum value is `60`. The minimum value is `0`.
* `routing_tables` - (Optional, List) The routing tables to update with the learned routes.When `next_hops` specifies multiple next hop addresses, individual routes are added for each address according to the `priority` value of the peer with that next hop `endpoint.address`:- For ingress routing tables, the `priority` value of the peer is used.- For egress routing tables, the route `priority` is determined by  [Zone priority rules](https://cloud.ibm.com/docs/__TBD__).
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
* `state` - (Optional, String) The state used for this dynamic route server peer group policy:- `disabled`: The peer group policy is disabled, and dynamic route server   will not apply the configured policy rules.- `enabled`: The peer group policy is enabled, and dynamic route server will  apply the configured policy rules.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
  * Constraints: Allowable values are: `disabled`, `enabled`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `type` - (Required, String) The type of dynamic route server peer group policy:- `custom_routes`: A policy used for custom routes.  The custom routes are advertised to other peer groups based on the policies  configured.- `learned_routes`:  A policy used for learned routes.  Learned routes are advertised to other peer groups based on the policies  configured.- `vpc_address_prefixes`: A policy used for advertising VPC address prefixes.  The VPC address prefixes are advertised to other peer groups based on the policies  configured.- `vpc_routing_tables`:  A policy used for updating VPC routing tables.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
  * Constraints: Allowable values are: `custom_routes`, `learned_routes`, `vpc_address_prefixes`, `vpc_routing_tables`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the is_dynamic_route_server_peer_group_policy.
* `created_at` - (String) The date and time that the peer group policy was created or last updated.
* `href` - (String) The URL for this dynamic route server peer group policy.
  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `is_dynamic_route_server_peer_group_policy_id` - (String) The unique identifier for this dynamic route server peer group policy.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `resource_type` - (String) The resource type.
  * Constraints: Allowable values are: `dynamic_route_server_peer_group_policy`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

* `etag` - ETag identifier for is_dynamic_route_server_peer_group_policy.

## Import

You can import the `ibm_is_dynamic_route_server_peer_group_policy` resource by using `id`.
The `id` property can be formed from `dynamic_route_server_id`, `dynamic_route_server_peer_group_id`, and `is_dynamic_route_server_peer_group_policy_id` in the following format:

<pre>
&lt;dynamic_route_server_id&gt;/&lt;dynamic_route_server_peer_group_id&gt;/&lt;is_dynamic_route_server_peer_group_policy_id&gt;
</pre>
* `dynamic_route_server_id`: A string. The dynamic route server identifier.
* `dynamic_route_server_peer_group_id`: A string. The dynamic route server peer group identifier.
* `is_dynamic_route_server_peer_group_policy_id`: A string in the format `r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5`. The unique identifier for this dynamic route server peer group policy.

# Syntax
<pre>
$ terraform import ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy &lt;dynamic_route_server_id&gt;/&lt;dynamic_route_server_peer_group_id&gt;/&lt;is_dynamic_route_server_peer_group_policy_id&gt;
</pre>
