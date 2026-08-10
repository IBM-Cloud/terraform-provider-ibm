---
layout: "ibm"
page_title: "IBM : ibm_is_dynamic_route_server_peer_group_policy"
description: |-
  Get information about is_dynamic_route_server_peer_group_policy
subcategory: "Virtual Private Cloud API"
---

# ibm_is_dynamic_route_server_peer_group_policy

Provides a read-only data source to retrieve information about an is_dynamic_route_server_peer_group_policy. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_dynamic_route_server_peer_group_policy" "is_dynamic_route_server_peer_group_policy" {
	dynamic_route_server_id = ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance.dynamic_route_server_id
	dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance.dynamic_route_server_peer_group_id
	is_dynamic_route_server_peer_group_policy_id = ibm_is_dynamic_route_server_peer_group_policy.is_dynamic_route_server_peer_group_policy_instance.is_dynamic_route_server_peer_group_policy_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `dynamic_route_server_id` - (Required, Forces new resource, String) The dynamic route server identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `dynamic_route_server_peer_group_id` - (Required, Forces new resource, String) The dynamic route server peer group identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `is_dynamic_route_server_peer_group_policy_id` - (Required, Forces new resource, String) The dynamic route server peer group policy identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the is_dynamic_route_server_peer_group_policy.
* `created_at` - (String) The date and time that the peer group policy was created or last updated.
* `custom_routes` - (List) The custom routes to advertise.
  * Constraints: The maximum length is `10` items. The minimum length is `1` item.
Nested schema for **custom_routes**:
	* `destination` - (String) The destination CIDR of the route. The host identifier in the CIDR must be zero.
	  * Constraints: The maximum length is `18` characters. The minimum length is `9` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$/`.
* `excluded_prefixes` - (List) A list of prefixes that are excluded from monitoring by the dynamic route server when learning routes from connected peer groups. They are applied in addition to any `excluded_prefixes` defined on the peer group. Each excluded prefix must be a subset of the `monitored_prefixes` configured in the peer group.
Nested schema for **excluded_prefixes**:
	* `ge` - (Integer) The minimum prefix length to match. If non-zero, the prefix watcher will only exclude routes within the `prefix` that have a prefix length greater than or equal to this value.If zero, `ge` match filtering will not be applied.If non-zero, `ge` must be:- Greater than or equal to the network `prefix` length.- Less than or equal to 32.If both `ge` and `le` are non-zero, `ge` must be less than or equal to `le`.
	  * Constraints: The default value is `0`. The maximum value is `32`. The minimum value is `0`.
	* `le` - (Integer) The maximum prefix length to match. If non-zero, the prefix watcher will only match routes within the `prefix` that have a prefix length less than or equal to this value.If zero, `le` match filtering will not be applied.If non-zero, `le` value must be:- Less than or equal to the network `prefix` length.- Greater than `ge`.
	  * Constraints: The default value is `0`. The maximum value is `32`. The minimum value is `0`.
	* `prefix` - (String) The prefix to be excluded.
* `href` - (String) The URL for this dynamic route server peer group policy.
  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `name` - (String) The name for this dynamic route server peer group policy. The name must not be used by another peer group policy. Names starting with ibm- are reserved for ibm managed policies, and are not allowed. If unspecified, the name will be a hyphenated list of randomly-selected words.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
* `next_hops` - (List) The next hop resources.
Nested schema for **next_hops**:
	* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		* `more_info` - (String) A link to documentation about deleted resources.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `href` - (String) The URL for this dynamic route server peer group.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (String) The ID for this dynamic route server peer group peer.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `name` - (String) The name for this dynamic route server peer group.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	* `resource_type` - (String) The resource type.
	  * Constraints: Allowable values are: `dynamic_route_server_peer_group`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `peer_groups` - (List) The peer groups from which the routes are being learned and then advertised back to the peer group on which this peer group policy has been applied.
Nested schema for **peer_groups**:
	* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		* `more_info` - (String) A link to documentation about deleted resources.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `href` - (String) The URL for this dynamic route server peer group.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (String) The ID for this dynamic route server peer group peer.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `name` - (String) The name for this dynamic route server peer group.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	* `resource_type` - (String) The resource type.
	  * Constraints: Allowable values are: `dynamic_route_server_peer_group`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `resource_type` - (String) The resource type.
  * Constraints: Allowable values are: `dynamic_route_server_peer_group_policy`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `route_delete_delay` - (Integer) The number of seconds to wait before deleting a route from `routing_tables` when the`status` of the route's associated peer is `down`, or after it is no longer advertised by peers in the peer group.
  * Constraints: The maximum value is `60`. The minimum value is `0`.
* `routing_tables` - (List) The routing tables to update with the learned routes.When `next_hops` specifies multiple next hop addresses, individual routes are added for each address according to the `priority` value of the peer with that next hop `endpoint.address`:- For ingress routing tables, the `priority` value of the peer is used.- For egress routing tables, the route `priority` is determined by  [Zone priority rules](https://cloud.ibm.com/docs/__TBD__).
Nested schema for **routing_tables**:
	* `advertise` - (Boolean) Indicates whether this route will be advertised to the ingress sources specified by the `advertise_routes_to` routing table property.
	* `vpc_routing_table` - (List)
	Nested schema for **vpc_routing_table**:
		* `crn` - (String) The CRN for this VPC routing table.
		  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
		* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested schema for **deleted**:
			* `more_info` - (String) A link to documentation about deleted resources.
			  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `href` - (String) The URL for this routing table.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `id` - (String) The unique identifier for this routing table.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
		* `name` - (String) The name for this routing table. The name is unique across all routing tables for the VPC.
		  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
		* `resource_type` - (String) The resource type.
		  * Constraints: Allowable values are: `routing_table`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `state` - (String) The state used for this dynamic route server peer group policy:- `disabled`: The peer group policy is disabled, and dynamic route server   will not apply the configured policy rules.- `enabled`: The peer group policy is enabled, and dynamic route server will  apply the configured policy rules.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
  * Constraints: Allowable values are: `disabled`, `enabled`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `type` - (String) The type of dynamic route server peer group policy:- `custom_routes`: A policy used for custom routes.  The custom routes are advertised to other peer groups based on the policies  configured.- `learned_routes`:  A policy used for learned routes.  Learned routes are advertised to other peer groups based on the policies  configured.- `vpc_address_prefixes`: A policy used for advertising VPC address prefixes.  The VPC address prefixes are advertised to other peer groups based on the policies  configured.- `vpc_routing_tables`:  A policy used for updating VPC routing tables.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
  * Constraints: Allowable values are: `custom_routes`, `learned_routes`, `vpc_address_prefixes`, `vpc_routing_tables`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

