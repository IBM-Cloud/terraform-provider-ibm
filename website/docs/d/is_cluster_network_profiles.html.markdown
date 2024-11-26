---
layout: "ibm"
page_title: "IBM : ibm_is_cluster_network_profiles"
description: |-
  Get information about ClusterNetworkProfileCollection
subcategory: "VPC infrastructure"
---

# ibm_is_cluster_network_profiles

Provides a read-only data source to retrieve information about a ClusterNetworkProfileCollection. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_cluster_network_profiles" "is_cluster_network_profiles" {
}
```


## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the ClusterNetworkProfileCollection.
- `profiles` - (List) A page of cluster network profiles.
	
	Nested schema for **profiles**:
	- `family` - (String) The product family this cluster network profile belongs to.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	- `href` - (String) The URL for this cluster network profile.
	- `name` - (String) The globally unique name for this cluster network profile.
	- `resource_type` - (String) The resource type.
	- `supported_instance_profiles` - (List) The instance profiles that support this cluster network profile.
	Nested schema for **supported_instance_profiles**:
		- `href` - (String) The URL for this virtual server instance profile.
		- `name` - (String) The globally unique name for this virtual server instance profile.
		- `resource_type` - (String) The resource type.
	- `zones` - (List) Zones in this region that support this cluster network profile.
		
		Nested schema for **zones**:
		- `href` - (String) The URL for this zone.
		- `name` - (String) The globally unique name for this zone.

