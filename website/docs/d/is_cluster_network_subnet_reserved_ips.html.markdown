---
layout: "ibm"
page_title: "IBM : ibm_is_cluster_network_subnet_reserved_ips"
description: |-
  Get information about ClusterNetworkSubnetReservedIPCollection
subcategory: "VPC infrastructure"
---

# ibm_is_cluster_network_subnet_reserved_ips

Provides a read-only data source to retrieve information about a ClusterNetworkSubnetReservedIPCollection. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_cluster_network_subnet_reserved_ips" "is_cluster_network_subnet_reserved_ips_instance" {
  cluster_network_id = ibm_is_cluster_network.is_cluster_network_instance.id
  cluster_network_subnet_id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `cluster_network_id` - (Required, Forces new resource, String) The cluster network identifier.
- `cluster_network_subnet_id` - (Required, Forces new resource, String) The cluster network subnet identifier.
- `name` - (Optional, String) Filters the collection to resources with a `name` property matching the exact specified name.
- `sort` - (Optional, String) Sorts the returned collection by the specified property name in ascending order. A `-` may be prepended to the name to sort in descending order. For example, the value `-created_at` sorts the collection by the `created_at` property in descending order, and the value `name` sorts it by the `name` property in ascending order.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the ClusterNetworkSubnetReservedIPCollection.
- `reserved_ips` - (List) A page of reserved IPs for the cluster network subnet.
	Nested schema for **reserved_ips**:
	- `address` - (String) The IP address.If the address is pending allocation, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
	- `auto_delete` - (Boolean) Indicates whether this cluster network subnet reserved IP member will be automatically deleted when either `target` is deleted, or the cluster network subnet reserved IP is unbound.
	- `created_at` - (String) The date and time that the cluster network subnet reserved IP was created.
	- `href` - (String) The URL for this cluster network subnet reserved IP.
	- `id` - (String) The unique identifier for this cluster network subnet reserved IP.
	- `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
		Nested schema for **lifecycle_reasons**:
		- `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
		- `message` - (String) An explanation of the reason for this lifecycle state.
		- `more_info` - (String) Link to documentation about the reason for this lifecycle state.
	- `lifecycle_state` - (String) The lifecycle state of the cluster network subnet reserved IP.
	- `name` - (String) The name for this cluster network subnet reserved IP. The name is unique across all reserved IPs in a cluster network subnet.
	- `owner` - (String) The owner of the cluster network subnet reserved IPThe enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	- `resource_type` - (String) The resource type.
	- `target` - (List) The target this cluster network subnet reserved IP is bound to.If absent, this cluster network subnet reserved IP is provider-owned or unbound.
		Nested schema for **target**:
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this cluster network interface.
		- `id` - (String) The unique identifier for this cluster network interface.
		- `name` - (String) The name for this cluster network interface. The name is unique across all interfaces in the cluster network.
		- `resource_type` - (String) The resource type.

