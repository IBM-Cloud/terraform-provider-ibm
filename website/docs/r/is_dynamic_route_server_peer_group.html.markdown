---
layout: "ibm"
page_title: "IBM : ibm_is_dynamic_route_server_peer_group"
description: |-
  Manages DynamicRouteServerPeerGroup.
subcategory: "Virtual Private Cloud API"
---

# ibm_is_dynamic_route_server_peer_group

Create, update, and delete DynamicRouteServerPeerGroups with this resource.

## Example Usage

```hcl
resource "ibm_is_dynamic_route_server_peer_group" "is_dynamic_route_server_peer_group_instance" {
  dynamic_route_server_id = "dynamic_route_server_id"
  health_monitor {
		mode = "bfd"
		roles = [ "next_hop" ]
  }
  name = "my-dynamic-route-server-peer-group"
  peers {
		created_at = "2026-01-02T03:04:05.006Z"
		href = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002/peers/r006-8228af60-189d-11ed-861d-0242ac120004"
		id = "r006-8228af60-189d-11ed-861d-0242ac120004"
		lifecycle_reasons {
			code = "resource_suspended_by_provider"
			message = "The resource has been suspended. Contact IBM support with the CRN for next steps."
			more_info = "https://cloud.ibm.com/apidocs/vpc#resource-suspension"
		}
		lifecycle_state = "stable"
		name = "my-dynamic-route-server-peer-group-peer"
		resource_type = "dynamic_route_server_peer_group_peer"
		status = "up"
		status_reasons {
			code = "peer_not_responding"
			message = "The connection is down because the peer is not responding."
			more_info = "https://cloud.ibm.com/docs/drs?topic=vpc-drs-health#__TBD__"
		}
		asn = 64520
		bidirectional_forwarding_detection {
			detect_multiplier = 3
			enabled = true
			mode = "asynchronous"
			receive_interval = 300
			role = "active"
			sessions {
				local {
					deleted {
						more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
					}
					href = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/members/r006-d7cc5196-9924-48d4-82d8-3f30da41fdc"
					id = "0717-9dce0ab3-a719-4582-ad71-00abfaae43ed"
					name = "my-dynamic-route-server-member1"
					resource_type = "dynamic_route_server_member"
					virtual_network_interfaces {
						crn = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
						deleted {
							more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
						}
						href = "https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
						id = "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
						name = "my-virtual-network-interface"
						primary_ip {
							address = "192.168.3.4"
							deleted {
								more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
							}
							href = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
							id = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
							name = "my-reserved-ip"
							resource_type = "subnet_reserved_ip"
						}
						resource_type = "virtual_network_interface"
						subnet {
							crn = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
							deleted {
								more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
							}
							href = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
							id = "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
							name = "my-subnet"
							resource_type = "subnet"
						}
					}
				}
				remote {
					deleted {
						more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
					}
					href = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/peer_groups/r006-8228af60-189d-11ed-861d-0242ac120002/peers/r006-8228af60-189d-11ed-861d-0242ac120004"
					id = "r006-8228af60-189d-11ed-861d-0242ac120004"
					name = "my-dynamic-route-server-peer-group-peer"
					resource_type = "dynamic_route_server_peer_group_peer"
				}
				state = "admin_down"
			}
			transmit_interval = 300
		}
		creator = "dynamic_route_server"
		endpoint {
			address = "192.168.3.4"
			gateway {
				address = "192.168.3.4"
			}
		}
		priority = 1
		sessions {
			established_at = "2026-01-02T03:04:05.006Z"
			local {
				deleted {
					more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
				}
				href = "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/members/r006-d7cc5196-9924-48d4-82d8-3f30da41fdc"
				id = "0717-9dce0ab3-a719-4582-ad71-00abfaae43ed"
				name = "my-dynamic-route-server-member1"
				resource_type = "dynamic_route_server_member"
				virtual_network_interfaces {
					crn = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
					deleted {
						more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
					}
					href = "https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
					id = "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
					name = "my-virtual-network-interface"
					primary_ip {
						address = "192.168.3.4"
						deleted {
							more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
						}
						href = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
						id = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
						name = "my-reserved-ip"
						resource_type = "subnet_reserved_ip"
					}
					resource_type = "virtual_network_interface"
					subnet {
						crn = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
						deleted {
							more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
						}
						href = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
						id = "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
						name = "my-subnet"
						resource_type = "subnet"
					}
				}
			}
			protocol_state = "active"
			remote {
				name = "my-dynamic-route-server-peer-group-peer"
			}
		}
		state = "enabled"
  }
  prefix_watcher {
		excluded_prefixes {
			ge = 0
			le = 32
			prefix = "prefix"
		}
		monitored_prefixes {
			ge = 0
			le = 32
			prefix = "prefix"
		}
  }
  state = "disabled"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `dynamic_route_server_id` - (Required, Forces new resource, String) The dynamic route server identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `health_monitor` - (Required, List) The health monitoring configuration used for this dynamic route server peer group.
Nested schema for **health_monitor**:
	* `mode` - (Optional, String) The mode used for this health monitor:- `bfd`: Both BGP monitoring and Bidirectional Forwarding Detection (BFD)  are enabled. When a peer in this peer group becomes unreachable, routes  whose `next_hop` is this peer are automatically withdrawn. BFD provides  faster failure detection.- `bgp`: Only BGP monitoring is enabled. When a peer in this peer group  becomes unreachable, routes whose `next_hop` is this peer are automatically  withdrawn.- `none`: Monitoring is disabled. When a peer in this peer group becomes  unreachable, routes whose `next_hop` is this peer are not automatically  withdrawn.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `bfd`, `bgp`, `none`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `roles` - (Optional, List) The `roles` used for this dynamic route server peer group health monitor configuration:- `next_hop`: Health monitor configuration used for `next_hop` role peers  of this dynamic route server peer group.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable list items are: `next_hop`. The maximum length is `1` item. The minimum length is `1` item.
* `name` - (Optional, String) The name for this dynamic route server peer group.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
* `peers` - (Required, List) A collection of peers of the same type having a similar network topology grouped together to apply identical policies and maintain consistent behavior.
  * Constraints: The maximum length is `6` items. The minimum length is `1` item.
Nested schema for **peers**:
	* `asn` - (Optional, Integer) The autonomous system number (ASN) for this dynamic route server peer group address peer.
	* `bidirectional_forwarding_detection` - (Optional, List) The bidirectional forwarding detection (BFD) configuration for this dynamicroute server peer group peer.
	Nested schema for **bidirectional_forwarding_detection**:
		* `detect_multiplier` - (Computed, Integer) The desired detection time multiplier for bidirectional forwarding detection control packets on this dynamic route server for this peer.
		  * Constraints: The maximum value is `255`. The minimum value is `2`.
		* `enabled` - (Computed, Boolean) Indicates whether bidirectional forwarding detection (BFD) is enabled on this dynamic route server peer group peer.
		* `mode` - (Computed, String) The bidirectional forwarding detection mode of this peer:- `asynchronous`: Each peer sends BFD control packets independently. Session failure is detected when expected packets are not received within the detection interval as defined in[RFC 5880](https://www.rfc-editor.org/rfc/rfc5880.html?#section-3.2)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
		  * Constraints: Allowable values are: `asynchronous`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
		* `receive_interval` - (Computed, Integer) The minimum interval, in microseconds, between received bidirectional forwarding detection control packets that this dynamic route server is capable of supporting. The actual interval is negotiated between this dynamic route server and the peer.
		  * Constraints: The maximum value is `60000`. The minimum value is `20`.
		* `role` - (Computed, String) The bidirectional forwarding detection role used in session initialization:  - `active`: Actively initiates BFD control packets to bring up the session.  - `passive`: Waits for BFD control packets from the peer and does not initiate    the session.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
		  * Constraints: Allowable values are: `active`, `passive`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
		* `sessions` - (Required, List) The bidirectional forwarding detection sessions for this peer.
		  * Constraints: The maximum length is `3` items. The minimum length is `1` item.
		Nested schema for **sessions**:
			* `local` - (Required, List) The local peer for this bidirectional forwarding detection session.
			Nested schema for **local**:
				* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
				Nested schema for **deleted**:
					* `more_info` - (Computed, String) A link to documentation about deleted resources.
					  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
				* `href` - (Computed, String) The URL for this dynamic route server member.
				  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
				* `id` - (Computed, String) The unique identifier for this dynamic route server member.
				  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
				* `name` - (Computed, String) The name for this dynamic route server member. The name is unique across all members in the dynamic route server.
				  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
				* `resource_type` - (Computed, String) The resource type.
				  * Constraints: Allowable values are: `dynamic_route_server_member`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
				* `virtual_network_interfaces` - (Required, List) The virtual network interfaces for this dynamic route server member.
				  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
				Nested schema for **virtual_network_interfaces**:
					* `crn` - (Computed, String) The CRN for this virtual network interface.
					  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
					* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
					Nested schema for **deleted**:
						* `more_info` - (Computed, String) A link to documentation about deleted resources.
						  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
					* `href` - (Computed, String) The URL for this virtual network interface.
					  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^https:\/\/([^\/?#]*)([^?#]*)\/virtual_network_interfaces\/[-0-9a-z_]+$/`.
					* `id` - (Computed, String) The unique identifier for this virtual network interface.
					  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
					* `name` - (Computed, String) The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.
					  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
					* `primary_ip` - (Required, List) The primary IP for this virtual network interface.
					Nested schema for **primary_ip**:
						* `address` - (Computed, String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
						  * Constraints: The maximum length is `15` characters. The minimum length is `7` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`.
						* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
						Nested schema for **deleted**:
							* `more_info` - (Computed, String) A link to documentation about deleted resources.
							  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
						* `href` - (Computed, String) The URL for this reserved IP.
						  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
						* `id` - (Computed, String) The unique identifier for this reserved IP.
						  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
						* `name` - (Computed, String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
						  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
						* `resource_type` - (Computed, String) The resource type.
						  * Constraints: Allowable values are: `subnet_reserved_ip`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
					* `resource_type` - (Computed, String) The resource type.
					  * Constraints: Allowable values are: `virtual_network_interface`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
					* `subnet` - (Required, List) The associated subnet.
					Nested schema for **subnet**:
						* `crn` - (Computed, String) The CRN for this subnet.
						  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
						* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
						Nested schema for **deleted**:
							* `more_info` - (Computed, String) A link to documentation about deleted resources.
							  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
						* `href` - (Computed, String) The URL for this subnet.
						  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
						* `id` - (Computed, String) The unique identifier for this subnet.
						  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
						* `name` - (Computed, String) The name for this subnet. The name is unique across all subnets in the VPC.
						  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
						* `resource_type` - (Computed, String) The resource type.
						  * Constraints: Allowable values are: `subnet`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
			* `remote` - (Required, List) The remote peer for this bidirectional forwarding detection session.
			Nested schema for **remote**:
				* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
				Nested schema for **deleted**:
					* `more_info` - (Computed, String) A link to documentation about deleted resources.
					  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
				* `href` - (Computed, String) The URL for this dynamic route server peer group peer.
				  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
				* `id` - (Computed, String) The ID for this dynamic route server peer group peer.
				  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
				* `name` - (Computed, String) The name for this dynamic route server peer group peer. The name must not be used by another peer in the peer group.
				  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
				* `resource_type` - (Computed, String) The resource type.
				  * Constraints: Allowable values are: `dynamic_route_server_peer_group_peer`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
			* `state` - (Computed, String) The current state of this bidirectional forwarding detection session as observed by the dynamic route server. The states are defined in [RFC 5880](https://www.rfc-editor.org/rfc/rfc5880.html#section-4.1).
			  * Constraints: Allowable values are: `admin_down`, `down`, `init`, `up`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
		* `transmit_interval` - (Computed, Integer) The minimum interval, in microseconds that this dynamic route server prefer to use when transmitting bidirectional forwarding detection control to this peer. The actual interval is negotiated between this dynamic route server and the peer.
		  * Constraints: The maximum value is `60000`. The minimum value is `20`.
	* `created_at` - (Computed, String) The date and time that this dynamic route server peer group peer was created.
	* `creator` - (Computed, String) The type of resource that created this peer:  - `dynamic_route_server`: The peer was created by a dynamic router server.  - `transit_gateway`:  The peer was created by a transit gateway.The enumerated values for this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `dynamic_route_server`, `transit_gateway`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `endpoint` - (Optional, List) The endpoint for a dynamic route server peer group address peer.
	Nested schema for **endpoint**:
		* `address` - (Optional, String) The IP address.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
		  * Constraints: The maximum length is `15` characters. The minimum length is `7` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`.
		* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested schema for **deleted**:
			* `more_info` - (Computed, String) A link to documentation about deleted resources.
			  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `gateway` - (Optional, List) The gateway IP address of the dynamic route server peer group address peer endpoint.
		Nested schema for **gateway**:
			* `address` - (Required, String) The IP address.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
			  * Constraints: The maximum length is `15` characters. The minimum length is `7` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`.
		* `href` - (Optional, String) The URL for this reserved IP.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `id` - (Optional, String) The unique identifier for this reserved IP.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
		* `name` - (Computed, String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
		  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
		* `resource_type` - (Computed, String) The resource type.
		  * Constraints: Allowable values are: `subnet_reserved_ip`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `href` - (Computed, String) The URL for this dynamic route server peer group peer.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (Computed, String) The ID for this dynamic route server peer group peer.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `lifecycle_reasons` - (Required, List) The reasons for the current `lifecycle_state` (if any).
	  * Constraints: The minimum length is `0` items.
	Nested schema for **lifecycle_reasons**:
		* `code` - (Computed, String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
		  * Constraints: Allowable values are: `internal_error`, `resource_suspended_by_provider`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
		* `message` - (Computed, String) An explanation of the reason for this lifecycle state.
		* `more_info` - (Computed, String) A link to documentation about the reason for this lifecycle state.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `lifecycle_state` - (Computed, String) The lifecycle state of the dynamic route server peer group peer.
	  * Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `name` - (Required, String) The name for this dynamic route server peer group peer. The name must not be used by another peer in the peer group.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	* `priority` - (Optional, Integer) The priority of the peer group peer. The priority is used to determine the preferred path for routing. A lower value indicates a higher priority.
	  * Constraints: The maximum value is `4`. The minimum value is `0`.
	* `resource_type` - (Computed, String) The resource type.
	  * Constraints: Allowable values are: `dynamic_route_server_peer_group_peer`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `sessions` - (Optional, List) The BGP sessions for this peer group peer.Empty if `health_monitor.mode` is `none`.
	  * Constraints: The maximum length is `10` items. The minimum length is `0` items.
	Nested schema for **sessions**:
		* `established_at` - (Computed, String) The date and time that the BGP session was established. This property will be present only when the session `state` is `established`.
		* `local` - (Required, List) The local peer for this BGP session.
		Nested schema for **local**:
			* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			Nested schema for **deleted**:
				* `more_info` - (Computed, String) A link to documentation about deleted resources.
				  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
			* `href` - (Computed, String) The URL for this dynamic route server member.
			  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
			* `id` - (Computed, String) The unique identifier for this dynamic route server member.
			  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
			* `name` - (Computed, String) The name for this dynamic route server member. The name is unique across all members in the dynamic route server.
			  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
			* `resource_type` - (Computed, String) The resource type.
			  * Constraints: Allowable values are: `dynamic_route_server_member`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
			* `virtual_network_interfaces` - (Required, List) The virtual network interfaces for this dynamic route server member.
			  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
			Nested schema for **virtual_network_interfaces**:
				* `crn` - (Computed, String) The CRN for this virtual network interface.
				  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
				* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
				Nested schema for **deleted**:
					* `more_info` - (Computed, String) A link to documentation about deleted resources.
					  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
				* `href` - (Computed, String) The URL for this virtual network interface.
				  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^https:\/\/([^\/?#]*)([^?#]*)\/virtual_network_interfaces\/[-0-9a-z_]+$/`.
				* `id` - (Computed, String) The unique identifier for this virtual network interface.
				  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
				* `name` - (Computed, String) The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.
				  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
				* `primary_ip` - (Required, List) The primary IP for this virtual network interface.
				Nested schema for **primary_ip**:
					* `address` - (Computed, String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
					  * Constraints: The maximum length is `15` characters. The minimum length is `7` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`.
					* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
					Nested schema for **deleted**:
						* `more_info` - (Computed, String) A link to documentation about deleted resources.
						  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
					* `href` - (Computed, String) The URL for this reserved IP.
					  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
					* `id` - (Computed, String) The unique identifier for this reserved IP.
					  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
					* `name` - (Computed, String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
					  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
					* `resource_type` - (Computed, String) The resource type.
					  * Constraints: Allowable values are: `subnet_reserved_ip`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
				* `resource_type` - (Computed, String) The resource type.
				  * Constraints: Allowable values are: `virtual_network_interface`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
				* `subnet` - (Required, List) The associated subnet.
				Nested schema for **subnet**:
					* `crn` - (Computed, String) The CRN for this subnet.
					  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
					* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
					Nested schema for **deleted**:
						* `more_info` - (Computed, String) A link to documentation about deleted resources.
						  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
					* `href` - (Computed, String) The URL for this subnet.
					  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
					* `id` - (Computed, String) The unique identifier for this subnet.
					  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
					* `name` - (Computed, String) The name for this subnet. The name is unique across all subnets in the VPC.
					  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
					* `resource_type` - (Computed, String) The resource type.
					  * Constraints: Allowable values are: `subnet`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
		* `protocol_state` - (Computed, String) The state of the routing protocol with this dynamic route server peer group peer. The states follow the conventions defined in [RFC 4274](https://datatracker.ietf.org/doc/html/rfc4274#section-2.3).- `initializing`: The BGP session is being initialized.
		  * Constraints: Allowable values are: `active`, `connect`, `established`, `idle`, `initializing`, `open_confirm`, `open_sent`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
		* `remote` - (Required, List) The remote peer for this BGP session.
		Nested schema for **remote**:
			* `name` - (Computed, String) The name for this dynamic route server peer group peer. The name must not be used by another peer in the peer group.
			  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	* `state` - (Optional, String) The administrative state used for this peer group peer:- `disabled`: The peer group peer is disabled, and the dynamic route server members will  not establish a BGP session with this peer group peer.- `enabled`: The peer group peer is enabled, and the dynamic route server members will  try to establish a BGP session with this peer group peer.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `disabled`, `enabled`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `status` - (Computed, String) The status of this dynamic route server peer group peer connection:- `down`: not operational.- `up`: operating normally.
	  * Constraints: Allowable values are: `down`, `up`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `status_reasons` - (Optional, List) The status reasons for the dynamic route server peer group peer.
	  * Constraints: The minimum length is `0` items.
	Nested schema for **status_reasons**:
		* `code` - (Computed, String) The reasons for the current dynamic route server service connection status (if any).- `internal_error`- `peer_not_responding`- __TBD__.
		  * Constraints: Allowable values are: `internal_error`, `peer_not_responding`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
		* `message` - (Computed, String) An explanation of the reason for this dynamic route server service connection's status.
		  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~\\n\\r\\t]*$/`.
		* `more_info` - (Computed, String) A link to documentation about this status reason.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `prefix_watcher` - (Optional, List) The prefix watcher configuration used for this dynamic route server peer group.
Nested schema for **prefix_watcher**:
	* `excluded_prefixes` - (Required, List) The prefixes that are excluded from `monitored_prefixes` for this prefix watcher.If empty, no prefixes are excluded from `monitored_prefixes`.
	  * Constraints: The maximum length is `10` items. The minimum length is `1` item.
	Nested schema for **excluded_prefixes**:
		* `ge` - (Required, Integer) The minimum matched prefix length. If non-zero, only routes within the `prefix` that have a prefix length greater or equal to this value are excluded.If zero, `ge` matching is not applied.
		  * Constraints: The maximum value is `32`. The minimum value is `0`.
		* `le` - (Required, Integer) The maximum matched prefix length. If non-zero, only routes within the `prefix` that have a prefix length less than or equal to this value are excluded.If zero, `le` matching is not applied.
		  * Constraints: The maximum value is `32`. The minimum value is `0`.
		* `prefix` - (Required, String) The prefix excluded from `monitored_prefixes`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 address prefixes in the future.
	* `monitored_prefixes` - (Required, List) The prefixes that are monitored by this prefix watcher.
	  * Constraints: The maximum length is `10` items. The minimum length is `1` item.
	Nested schema for **monitored_prefixes**:
		* `ge` - (Required, Integer) The minimum prefix length to match. If non-zero, the prefix watcher matches only routes within the prefix that have a prefix length greater or equal to this value.If zero, `ge` matching is not applied.
		  * Constraints: The maximum value is `32`. The minimum value is `0`.
		* `le` - (Required, Integer) The maximum prefix length to match. If non-zero, the prefix watcher matches only routes within the prefix that have a prefix length less than or equal to this value.If zero, `le` matching is not applied.
		  * Constraints: The maximum value is `32`. The minimum value is `0`.
		* `prefix` - (Required, String) The IP address prefix matched by this prefix watcher.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 address prefixes in the future.
* `roles` - (Required, List) The roles assigned to peers in this dynamic route server peer group: - `next_hop`: The peer can serve as a next hop for routing traffic. - `router`: The peer can participate in route exchange and route filtering. The enumerated values for this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
  * Constraints: Allowable list items are: `next_hop`, `router`. The list items must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`. The maximum length is `2` items. The minimum length is `1` item.
* `state` - (Optional, String) The administrative state used for the dynamic route server peer group:- `disabled`: The peer group is disabled, and the dynamic route server members will  not establish a BGP session with peers in this peer group.- `enabled`: The peer group is enabled, and the dynamic route server members will  try to establish a BGP session with peers in this peer group.
  * Constraints: Allowable values are: `disabled`, `enabled`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the DynamicRouteServerPeerGroup.
* `created_at` - (String) The date and time that this dynamic route server peer group was created.
* `creator` - (List) The type of resource that created this dynamic route server peer group:  - `dynamic_route_server`: dynamic router server created peer group.  - `transit_gateway`:  transit gateway created peer group.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
  * Constraints: Allowable list items are: `dynamic_route_server`, `transit_gateway`.
* `href` - (String) The URL for this dynamic route server peer group.
  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `is_dynamic_route_server_peer_group_id` - (String) The ID for this dynamic route server peer group peer.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
  * Constraints: The minimum length is `0` items.
Nested schema for **lifecycle_reasons**:
	* `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `internal_error`, `resource_suspended_by_provider`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `message` - (String) An explanation of the reason for this lifecycle state.
	* `more_info` - (String) A link to documentation about the reason for this lifecycle state.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `lifecycle_state` - (String) The lifecycle state of the dynamic route server peer group.
  * Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `resource_type` - (String) The resource type.
  * Constraints: Allowable values are: `dynamic_route_server_peer_group`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `status` - (String) The aggregate session status of the peers in this dynamic route server peer group:  - `degraded`: One or more dynamic route server peer group peers    were unable to establish a session.  - `down`: All dynamic route server peer group peers were unable    to establish a session.  - `initializing`: The dynamic route server peer group is in the process of establishing     sessions with its peers.  - `unknown`: The aggregate session status of the dynamic route server peer group     peers could not be determined.  - `up`: All dynamic route server peer group peers have established    sessions and are operating normally.
  * Constraints: Allowable values are: `degraded`, `down`, `initializing`, `unknown`, `up`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `status_reasons` - (List) The reasons for the current dynamic route server peer group connection status (if any).
  * Constraints: The minimum length is `0` items.
Nested schema for **status_reasons**:
	* `code` - (String) A snake case string succinctly identifying the status reason.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `cannot_connect_peer_group_peer`, `cannot_start_dynamic_route_server_member`, `degraded`, `initializing`, `maintenance`, `stopped_by_provider`, `stopped_by_user`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `message` - (String) An explanation of the status reason.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~\\n\\r\\t]*$/`.
	* `more_info` - (String) A link to documentation about this status reason.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `etag` - ETag identifier for DynamicRouteServerPeerGroup.

## Import

You can import the `ibm_is_dynamic_route_server_peer_group` resource by using `id`.
The `id` property can be formed from `dynamic_route_server_id`, and `is_dynamic_route_server_peer_group_id` in the following format:

<pre>
&lt;dynamic_route_server_id&gt;/&lt;is_dynamic_route_server_peer_group_id&gt;
</pre>
* `dynamic_route_server_id`: A string. The dynamic route server identifier.
* `is_dynamic_route_server_peer_group_id`: A string in the format `r006-8228af60-189d-11ed-861d-0242ac120004`. The ID for this dynamic route server peer group peer.

# Syntax
<pre>
$ terraform import ibm_is_dynamic_route_server_peer_group.is_dynamic_route_server_peer_group &lt;dynamic_route_server_id&gt;/&lt;is_dynamic_route_server_peer_group_id&gt;
</pre>
