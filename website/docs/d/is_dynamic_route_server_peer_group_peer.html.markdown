---
layout: "ibm"
page_title: "IBM : ibm_is_dynamic_route_server_peer_group_peer"
description: |-
  Get information about is_dynamic_route_server_peer_group_peer
subcategory: "Virtual Private Cloud API"
---

# ibm_is_dynamic_route_server_peer_group_peer

Provides a read-only data source to retrieve information about an is_dynamic_route_server_peer_group_peer. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_dynamic_route_server_peer_group_peer" "is_dynamic_route_server_peer_group_peer" {
	dynamic_route_server_id = ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance.dynamic_route_server_id
	dynamic_route_server_peer_group_id = ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance.dynamic_route_server_peer_group_id
	is_dynamic_route_server_peer_group_peer_id = ibm_is_dynamic_route_server_peer_group_peer.is_dynamic_route_server_peer_group_peer_instance.is_dynamic_route_server_peer_group_peer_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `dynamic_route_server_id` - (Required, Forces new resource, String) The dynamic route server identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `dynamic_route_server_peer_group_id` - (Required, Forces new resource, String) The dynamic route server peer group identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `is_dynamic_route_server_peer_group_peer_id` - (Required, Forces new resource, String) The dynamic route server peer group peer identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the is_dynamic_route_server_peer_group_peer.
* `asn` - (Integer) The autonomous system number (ASN) for this dynamic route server peer group address peer.
* `bidirectional_forwarding_detection` - (List) The bidirectional forwarding detection (BFD) configuration for this dynamicroute server peer group peer.
Nested schema for **bidirectional_forwarding_detection**:
	* `detect_multiplier` - (Integer) The desired detection time multiplier for bidirectional forwarding detection control packets on this dynamic route server for this peer.
	  * Constraints: The maximum value is `255`. The minimum value is `2`.
	* `enabled` - (Boolean) Indicates whether bidirectional forwarding detection (BFD) is enabled on this dynamic route server peer group peer.
	* `mode` - (String) The bidirectional forwarding detection mode of this peer:- `asynchronous`: Each peer sends BFD control packets independently. Session failure is detected when expected packets are not received within the detection interval as defined in[RFC 5880](https://www.rfc-editor.org/rfc/rfc5880.html?#section-3.2)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `asynchronous`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `receive_interval` - (Integer) The minimum interval, in microseconds, between received bidirectional forwarding detection control packets that this dynamic route server is capable of supporting. The actual interval is negotiated between this dynamic route server and the peer.
	  * Constraints: The maximum value is `60000`. The minimum value is `20`.
	* `role` - (String) The bidirectional forwarding detection role used in session initialization:  - `active`: Actively initiates BFD control packets to bring up the session.  - `passive`: Waits for BFD control packets from the peer and does not initiate    the session.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `active`, `passive`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `sessions` - (List) The bidirectional forwarding detection sessions for this peer.
	  * Constraints: The maximum length is `3` items. The minimum length is `1` item.
	Nested schema for **sessions**:
		* `local` - (List) The local peer for this bidirectional forwarding detection session.
		Nested schema for **local**:
			* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			Nested schema for **deleted**:
				* `more_info` - (String) A link to documentation about deleted resources.
				  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
			* `href` - (String) The URL for this dynamic route server member.
			  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
			* `id` - (String) The unique identifier for this dynamic route server member.
			  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
			* `name` - (String) The name for this dynamic route server member. The name is unique across all members in the dynamic route server.
			  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
			* `resource_type` - (String) The resource type.
			  * Constraints: Allowable values are: `dynamic_route_server_member`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
			* `virtual_network_interfaces` - (List) The virtual network interfaces for this dynamic route server member.
			  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
			Nested schema for **virtual_network_interfaces**:
				* `crn` - (String) The CRN for this virtual network interface.
				  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
				* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
				Nested schema for **deleted**:
					* `more_info` - (String) A link to documentation about deleted resources.
					  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
				* `href` - (String) The URL for this virtual network interface.
				  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^https:\/\/([^\/?#]*)([^?#]*)\/virtual_network_interfaces\/[-0-9a-z_]+$/`.
				* `id` - (String) The unique identifier for this virtual network interface.
				  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
				* `name` - (String) The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.
				  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
				* `primary_ip` - (List) The primary IP for this virtual network interface.
				Nested schema for **primary_ip**:
					* `address` - (String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
					  * Constraints: The maximum length is `15` characters. The minimum length is `7` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`.
					* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
					Nested schema for **deleted**:
						* `more_info` - (String) A link to documentation about deleted resources.
						  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
					* `href` - (String) The URL for this reserved IP.
					  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
					* `id` - (String) The unique identifier for this reserved IP.
					  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
					* `name` - (String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
					  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
					* `resource_type` - (String) The resource type.
					  * Constraints: Allowable values are: `subnet_reserved_ip`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
				* `resource_type` - (String) The resource type.
				  * Constraints: Allowable values are: `virtual_network_interface`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
				* `subnet` - (List) The associated subnet.
				Nested schema for **subnet**:
					* `crn` - (String) The CRN for this subnet.
					  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
					* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
					Nested schema for **deleted**:
						* `more_info` - (String) A link to documentation about deleted resources.
						  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
					* `href` - (String) The URL for this subnet.
					  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
					* `id` - (String) The unique identifier for this subnet.
					  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
					* `name` - (String) The name for this subnet. The name is unique across all subnets in the VPC.
					  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
					* `resource_type` - (String) The resource type.
					  * Constraints: Allowable values are: `subnet`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
		* `remote` - (List) The remote peer for this bidirectional forwarding detection session.
		Nested schema for **remote**:
			* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			Nested schema for **deleted**:
				* `more_info` - (String) A link to documentation about deleted resources.
				  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
			* `href` - (String) The URL for this dynamic route server peer group peer.
			  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
			* `id` - (String) The ID for this dynamic route server peer group peer.
			  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
			* `name` - (String) The name for this dynamic route server peer group peer. The name must not be used by another peer in the peer group.
			  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
			* `resource_type` - (String) The resource type.
			  * Constraints: Allowable values are: `dynamic_route_server_peer_group_peer`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
		* `state` - (String) The current state of this bidirectional forwarding detection session as observed by the dynamic route server. The states are defined in [RFC 5880](https://www.rfc-editor.org/rfc/rfc5880.html#section-4.1).
		  * Constraints: Allowable values are: `admin_down`, `down`, `init`, `up`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `transmit_interval` - (Integer) The minimum interval, in microseconds that this dynamic route server prefer to use when transmitting bidirectional forwarding detection control to this peer. The actual interval is negotiated between this dynamic route server and the peer.
	  * Constraints: The maximum value is `60000`. The minimum value is `20`.
* `created_at` - (String) The date and time that this dynamic route server peer group peer was created.
* `creator` - (String) The type of resource that created this peer:  - `dynamic_route_server`: The peer was created by a dynamic router server.  - `transit_gateway`:  The peer was created by a transit gateway.The enumerated values for this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
  * Constraints: Allowable values are: `dynamic_route_server`, `transit_gateway`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `endpoint` - (List) The endpoint for a dynamic route server peer group address peer.
Nested schema for **endpoint**:
	* `address` - (String) The IP address.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
	  * Constraints: The maximum length is `15` characters. The minimum length is `7` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`.
	* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		* `more_info` - (String) A link to documentation about deleted resources.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `gateway` - (List) The gateway IP address of the dynamic route server peer group address peer endpoint.
	Nested schema for **gateway**:
		* `address` - (String) The IP address.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
		  * Constraints: The maximum length is `15` characters. The minimum length is `7` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`.
	* `href` - (String) The URL for this reserved IP.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (String) The unique identifier for this reserved IP.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `name` - (String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	* `resource_type` - (String) The resource type.
	  * Constraints: Allowable values are: `subnet_reserved_ip`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `href` - (String) The URL for this dynamic route server peer group peer.
  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
  * Constraints: The minimum length is `0` items.
Nested schema for **lifecycle_reasons**:
	* `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `internal_error`, `resource_suspended_by_provider`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `message` - (String) An explanation of the reason for this lifecycle state.
	* `more_info` - (String) A link to documentation about the reason for this lifecycle state.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `lifecycle_state` - (String) The lifecycle state of the dynamic route server peer group peer.
  * Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `name` - (String) The name for this dynamic route server peer group peer. The name must not be used by another peer in the peer group.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
* `priority` - (Integer) The priority of the peer group peer. The priority is used to determine the preferred path for routing. A lower value indicates a higher priority.
  * Constraints: The maximum value is `4`. The minimum value is `0`.
* `resource_type` - (String) The resource type.
  * Constraints: Allowable values are: `dynamic_route_server_peer_group_peer`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `sessions` - (List) The BGP sessions for this peer group peer.Empty if `health_monitor.mode` is `none`.
  * Constraints: The maximum length is `10` items. The minimum length is `0` items.
Nested schema for **sessions**:
	* `established_at` - (String) The date and time that the BGP session was established. This property will be present only when the session `state` is `established`.
	* `local` - (List) The local peer for this BGP session.
	Nested schema for **local**:
		* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested schema for **deleted**:
			* `more_info` - (String) A link to documentation about deleted resources.
			  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `href` - (String) The URL for this dynamic route server member.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `id` - (String) The unique identifier for this dynamic route server member.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
		* `name` - (String) The name for this dynamic route server member. The name is unique across all members in the dynamic route server.
		  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
		* `resource_type` - (String) The resource type.
		  * Constraints: Allowable values are: `dynamic_route_server_member`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
		* `virtual_network_interfaces` - (List) The virtual network interfaces for this dynamic route server member.
		  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
		Nested schema for **virtual_network_interfaces**:
			* `crn` - (String) The CRN for this virtual network interface.
			  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
			* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			Nested schema for **deleted**:
				* `more_info` - (String) A link to documentation about deleted resources.
				  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
			* `href` - (String) The URL for this virtual network interface.
			  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^https:\/\/([^\/?#]*)([^?#]*)\/virtual_network_interfaces\/[-0-9a-z_]+$/`.
			* `id` - (String) The unique identifier for this virtual network interface.
			  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
			* `name` - (String) The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.
			  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
			* `primary_ip` - (List) The primary IP for this virtual network interface.
			Nested schema for **primary_ip**:
				* `address` - (String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
				  * Constraints: The maximum length is `15` characters. The minimum length is `7` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`.
				* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
				Nested schema for **deleted**:
					* `more_info` - (String) A link to documentation about deleted resources.
					  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
				* `href` - (String) The URL for this reserved IP.
				  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
				* `id` - (String) The unique identifier for this reserved IP.
				  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
				* `name` - (String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
				  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
				* `resource_type` - (String) The resource type.
				  * Constraints: Allowable values are: `subnet_reserved_ip`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
			* `resource_type` - (String) The resource type.
			  * Constraints: Allowable values are: `virtual_network_interface`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
			* `subnet` - (List) The associated subnet.
			Nested schema for **subnet**:
				* `crn` - (String) The CRN for this subnet.
				  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
				* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
				Nested schema for **deleted**:
					* `more_info` - (String) A link to documentation about deleted resources.
					  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
				* `href` - (String) The URL for this subnet.
				  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
				* `id` - (String) The unique identifier for this subnet.
				  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
				* `name` - (String) The name for this subnet. The name is unique across all subnets in the VPC.
				  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
				* `resource_type` - (String) The resource type.
				  * Constraints: Allowable values are: `subnet`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `protocol_state` - (String) The state of the routing protocol with this dynamic route server peer group peer. The states follow the conventions defined in [RFC 4274](https://datatracker.ietf.org/doc/html/rfc4274#section-2.3).- `initializing`: The BGP session is being initialized.
	  * Constraints: Allowable values are: `active`, `connect`, `established`, `idle`, `initializing`, `open_confirm`, `open_sent`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `remote` - (List) The remote peer for this BGP session.
	Nested schema for **remote**:
		* `name` - (String) The name for this dynamic route server peer group peer. The name must not be used by another peer in the peer group.
		  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
* `state` - (String) The administrative state used for this peer group peer:- `disabled`: The peer group peer is disabled, and the dynamic route server members will  not establish a BGP session with this peer group peer.- `enabled`: The peer group peer is enabled, and the dynamic route server members will  try to establish a BGP session with this peer group peer.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
  * Constraints: Allowable values are: `disabled`, `enabled`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `status` - (String) The status of this dynamic route server peer group peer connection:- `down`: not operational.- `up`: operating normally.
  * Constraints: Allowable values are: `down`, `up`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `status_reasons` - (List) The status reasons for the dynamic route server peer group peer.
  * Constraints: The minimum length is `0` items.
Nested schema for **status_reasons**:
	* `code` - (String) The reasons for the current dynamic route server service connection status (if any).- `internal_error`- `peer_not_responding`- __TBD__.
	  * Constraints: Allowable values are: `internal_error`, `peer_not_responding`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `message` - (String) An explanation of the reason for this dynamic route server service connection's status.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~\\n\\r\\t]*$/`.
	* `more_info` - (String) A link to documentation about this status reason.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

