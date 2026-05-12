---
layout: "ibm"
page_title: "IBM : ibm_is_cluster_network_interfaces"
description: |-
  Get information about ClusterNetworkInterfaceCollection
subcategory: "VPC infrastructure"
---

# ibm_is_cluster_network_interfaces

Provides a read-only data source to retrieve information about a ClusterNetworkInterfaceCollection. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_cluster_network_interfaces" "is_cluster_network_interfaces_instance" {
  cluster_network_id = ibm_is_cluster_network.is_cluster_network_instance.id
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `cluster_network_id` - (Required, Forces new resource, String) The cluster network identifier.
- `name` - (Optional, String) Filters the collection to resources with a `name` property matching the exact specified name.
- `sort` - (Optional, String) Sorts the returned collection by the specified property name in ascending order. A `-` may be prepended to the name to sort in descending order. For example, the value `-created_at` sorts the collection by the `created_at` property in descending order, and the value `name` sorts it by the `name` property in ascending order.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the ClusterNetworkInterfaceCollection.
- `interfaces` - (List) A page of cluster network interfaces.
	Nested schema for **interfaces**:
	- `allow_ip_spoofing` - (Boolean) Indicates whether source IP spoofing is allowed on this cluster network interface. If `false`, source IP spoofing is prevented on this cluster network interface. If `true`, source IP spoofing is allowed on this cluster network interface.
	- `auto_delete` - (Boolean) Indicates whether this cluster network interface will be automatically deleted when `target` is deleted.
	- `created_at` - (String) The date and time that the cluster network interface was created.
	- `enable_infrastructure_nat` - (Boolean) If `true`:- The VPC infrastructure performs any needed NAT operations.- `floating_ips` must not have more than one floating IP.If `false`:- Packets are passed unchanged to/from the virtual network interface,  allowing the workload to perform any needed NAT operations.
	- `href` - (String) The URL for this cluster network interface.
	- `id` - (String) The unique identifier for this cluster network interface.
	- `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
		Nested schema for **lifecycle_reasons**:
		- `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
		- `message` - (String) An explanation of the reason for this lifecycle state.
		- `more_info` - (String) Link to documentation about the reason for this lifecycle state.
	- `lifecycle_state` - (String) The lifecycle state of the cluster network interface.
	- `mac_address` - (String) The MAC address of the cluster network interface. May be absent if`lifecycle_state` is `pending`.
	- `name` - (String) The name for this cluster network interface. The name is unique across all interfaces in the cluster network.
	- `primary_ip` - (List) The cluster network subnet reserved IP for this cluster network interface.
		Nested schema for **primary_ip**:
		- `address` - (String) The IP address.If the address is pending allocation, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this cluster network subnet reserved IP.
		- `id` - (String) The unique identifier for this cluster network subnet reserved IP.
		- `name` - (String) The name for this cluster network subnet reserved IP. The name is unique across all reserved IPs in a cluster network subnet.
		- `resource_type` - (String) The resource type.
	- `resource_type` - (String) The resource type.
	- `subnet` - (List)
		Nested schema for **subnet**:
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this cluster network subnet.
		- `id` - (String) The unique identifier for this cluster network subnet.
		- `name` - (String) The name for this cluster network subnet. The name is unique across all cluster network subnets in the cluster network.
		- `resource_type` - (String) The resource type.
	- `target` - (List) The target of this cluster network interface.If absent, this cluster network interface is not attached to a target.The resources supported by this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
		Nested schema for **target**:
		- `href` - (String) The URL for this instance cluster network attachment.
		- `id` - (String) The unique identifier for this instance cluster network attachment.
		- `name` - (String) The name for this instance cluster network attachment. The name is unique across all network attachments for the instance.
		- `resource_type` - (String) The resource type.
	- `vpc` - (List) The VPC this cluster network interface resides in.
		Nested schema for **vpc**:
		- `crn` - (String) The CRN for this VPC.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this VPC.
		- `id` - (String) The unique identifier for this VPC.
		- `name` - (String) The name for this VPC. The name is unique across all VPCs in the region.
		- `resource_type` - (String) The resource type.
	- `zone` - (List) The zone this cluster network interface resides in.
		Nested schema for **zone**:
		- `href` - (String) The URL for this zone.
		- `name` - (String) The globally unique name for this zone.

