---
layout: "ibm"
page_title: "IBM : is_share_profiles"
description: |-
  Get information about ShareProfileCollection
subcategory: "VPC infrastructure"
---

# ibm\_is_share_profiles

Provides a read-only data source for ShareProfileCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_share_profiles" "example" {
}
```

## Attribute Reference

The following attributes are exported:

- `profiles` - (List) Collection of share profiles. Nested `profiles` blocks have the following structure:
	- `family` - (String) The product family this share profile belongs to.
	- `href` - (String) The URL for this share profile.
	- `name` - (String) The globally unique name for this share profile.
	- `resource_type` - (String) The resource type.
	- `allowed_access_protocols` - (List) The possible allowed access protocols for a share with this profile.. Nested `allowed_access_protocols` blocks have the following structure:
		- `default` - (List of String) The default allowed access protocol modes for shares with this profile.
		- `type` - (String) The type for this profile field
		- `values` - (List of String) The possible allowed access protocols for shares with this profile.
	- `allowed_transit_encryption_modes` - (List) The possible allowed transit encryption modes for a share with this profile.. Nested `allowed_transit_encryption_modes` blocks have the following structure:
		- `default` - (List of String) The default allowed transit encryption modes for shares with this profile.
		- `type` - (String) The type for this profile field.
		- `values` - (List of String) The allowed transit encryption modes for a share with this profile.
	- `availability_modes` - (List) The data availability mode of a share with this profile. Nested `availability_modes` blocks have the following structure:
		- `default` - (List of String) The default allowed transit encryption modes for shares with this profile.
		- `type` - (String) The type for this profile field.
		- `value` - (String) The value for this profile field
		- `values` - (List of String) The allowed transit encryption modes for a share with this profile.
	- `bandwidth` - (List) The permitted bandwidth (in megabits per second) for a share with this profile. Nested `capacity` blocks have the following structure:
		- `default` - (Integer) The default capacity for this share profile
		- `max` - (Integer) The max capacity for this share profile
		- `min` - (Integer) The min capacity for this share profile
		- `step` - (Integer) The increment step value for this profile field
		- `type` - (String) The type for this profile field
		- `value` - (Integer) The value for this profile field
		- `values` - (List of Integers) The permitted values for this profile field
	- `capacity` - (List) The permitted capacity range (in gigabytes) for a share with this profile. Nested `capacity` blocks have the following structure:
		- `default` - (Integer) The default capacity for this share profile
		- `max` - (Integer) The max capacity for this share profile
		- `min` - (Integer) The min capacity for this share profile
		- `step` - (Integer) The increment step value for this profile field
		- `type` - (String) The type for this profile field
		- `value` - (Integer) The value for this profile field
		- `values` - (List of Integers) The permitted values for this profile field
	- `iops` - (List) The permitted IOPS range for a share with this profile. Nested `iops` blocks have the following structure:
		- `default` - (Integer) The default iops for this share profile
		- `max` - (Integer) The max iops for this share profile
		- `min` - (Integer) The min iops for this share profile
		- `step` - (Integer) The increment step value for this profile field
		- `type` - (String) The type for this profile field
		- `value` - (Integer) The value for this profile field
		- `values` - (List of Integers) The permitted values for this profile field

- `total_count` - The total number of resources across all pages.

