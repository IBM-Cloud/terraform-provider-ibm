---
layout: "ibm"
page_title: "IBM : ibm_is_cluster_network_subnets"
description: |-
  Get information about ClusterNetworkSubnetCollection
subcategory: "VPC infrastructure"
---

# ibm_is_cluster_network_subnets

Provides a read-only data source to retrieve information about a ClusterNetworkSubnetCollection. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_cluster_network_subnets" "is_cluster_network_subnets_instance" {
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

- `id` - The unique identifier of the ClusterNetworkSubnetCollection.
- `subnets` - (List) A page of subnets for the cluster network.
	
	Nested schema for **subnets**:
	- `available_ipv4_address_count` - (Integer) The number of IPv4 addresses in this cluster network subnet that are not in use, and have not been reserved by the user or the provider.
	- `created_at` - (String) The date and time that the cluster network subnet was created.
	- `href` - (String) The URL for this cluster network subnet.
	- `id` - (String) The unique identifier for this cluster network subnet.
	- `ip_version` - (String) The IP version for this cluster network subnet.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	- `ipv4_cidr_block` - (String) The IPv4 range of this cluster network subnet, expressed in CIDR format.
	- `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
		Nested schema for **lifecycle_reasons**:
		- `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
		- `message` - (String) An explanation of the reason for this lifecycle state.
		- `more_info` - (String) Link to documentation about the reason for this lifecycle state.
	- `lifecycle_state` - (String) The lifecycle state of the cluster network subnet.
	- `name` - (String) The name for this cluster network subnet. The name is unique across all cluster network subnets in the cluster network.
	- `resource_type` - (String) The resource type.
	- `total_ipv4_address_count` - (Integer) The total number of IPv4 addresses in this cluster network subnet.Note: This is calculated as 2<sup>(32 - prefix length)</sup>. For example, the prefix length `/24` gives:<br> 2<sup>(32 - 24)</sup> = 2<sup>8</sup> = 256 addresses.

