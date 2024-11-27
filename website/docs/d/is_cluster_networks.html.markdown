---
layout: "ibm"
page_title: "IBM : ibm_is_cluster_networks"
description: |-
  Get information about ClusterNetworkCollection
subcategory: "VPC infrastructure"
---

# ibm_is_cluster_networks

Provides a read-only data source to retrieve information about a ClusterNetworkCollection. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_cluster_networks" "is_cluster_networks_instance" {
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `name` - (Optional, String) Filters the collection to resources with a `name` property matching the exact specified name.
- `resource_group_id` - (Optional, String) Filters the collection to resources with a `resource_group.id` property matching the specified identifier.
- `sort` - (Optional, String) Sorts the returned collection by the specified property name in ascending order. A `-` may be prepended to the name to sort in descending order. For example, the value `-created_at` sorts the collection by the `created_at` property in descending order, and the value `name` sorts it by the `name` property in ascending order.
- `vpc_crn` - (Optional, String) Filters the collection to cluster networks with a `vpc.crn` property matching the specified CRN.
- `vpc_id` - (Optional, String) Filters the collection to cluster networks with a `vpc.id` property matching the specified id.
- `vpc_name` - (Optional, String) Filters the collection to cluster networks with a `vpc.name` property matching the specified name.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the ClusterNetworkCollection.
- `cluster_networks` - (List) A page of cluster networks.

	Nested schema for **cluster_networks**:
	- `created_at` - (String) The date and time that the cluster network was created.
	- `crn` - (String) The CRN for this cluster network.
	- `href` - (String) The URL for this cluster network.
	- `id` - (String) The unique identifier for this cluster network.
	- `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
		Nested schema for **lifecycle_reasons**:
		- `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
		- `message` - (String) An explanation of the reason for this lifecycle state.
		- `more_info` - (String) Link to documentation about the reason for this lifecycle state.
	- `lifecycle_state` - (String) The lifecycle state of the cluster network.
	- `name` - (String) The name for this cluster network. The name must not be used by another cluster network in the region.
	- `profile` - (List) The profile for this cluster network.
		Nested schema for **profile**:
		- `href` - (String) The URL for this cluster network profile.
		- `name` - (String) The globally unique name for this cluster network profile.
		- `resource_type` - (String) The resource type.
	- `resource_group` - (List) The resource group for this cluster network.
		Nested schema for **resource_group**:
		- `href` - (String) The URL for this resource group.
		- `id` - (String) The unique identifier for this resource group.
		- `name` - (String) The name for this resource group.
	- `resource_type` - (String) The resource type.
	- `subnet_prefixes` - (List) The IP address ranges available for subnets for this cluster network.
		Nested schema for **subnet_prefixes**:
		- `allocation_policy` - (String) The allocation policy for this subnet prefix:- `auto`: Subnets created by total count in this cluster network can use this prefix.
		- `cidr` - (String) The CIDR block for this prefix.
	- `vpc` - (List) The VPC this cluster network resides in.
		Nested schema for **vpc**:
		- `crn` - (String) The CRN for this VPC.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this VPC.
		- `id` - (String) The unique identifier for this VPC.
		- `name` - (String) The name for this VPC. The name is unique across all VPCs in the region.
		- `resource_type` - (String) The resource type.
	- `zone` - (List) The zone this cluster network resides in.
		Nested schema for **zone**:
		- `href` - (String) The URL for this zone.
		- `name` - (String) The globally unique name for this zone.

