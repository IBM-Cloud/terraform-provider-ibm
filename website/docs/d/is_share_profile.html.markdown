---
layout: "ibm"
page_title: "IBM : is_share_profile"
description: |-
  Get information about ShareProfile
subcategory: "VPC infrastructure"
---

# ibm\_is_share_profile

Provides a read-only data source for ShareProfile. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_is_share_profile" "example" {
	name = "dp2"
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required, string) The file share profile name.

## Attribute Reference

The following attributes are exported:

- `family` - (String) The product family this share profile belongs to.
- `href` - (String) The URL for this share profile.
- `resource_type` - (String) The resource type.
- `capacity` - (List) - The permitted capacity range (in gigabytes) for a share with this profile. Nested `capacity` blocks have the following structure:
		- `default` - (Integer) The default capacity for this share profile
		- `max` - (Integer) The max capacity for this share profile
		- `min` - (Integer) The min capacity for this share profile
		- `step` - (Integer) The increment step value for this profile field
		- `type` - (String) The type for this profile field
		- `value` - (Integer) The value for this profile field
		- `values` - (List of Integers) The permitted values for this profile field
- `iops` - (List) - The permitted IOPS range for a share with this profile. Nested `iops` blocks have the following structure:
  - `default` - (Integer) The default iops for this share profile
  - `max` - (Integer) The max iops for this share profile
  - `min` - (Integer) The min iops for this share profile
  - `step` - (Integer) The increment step value for this profile field
  - `type` - (String) The type for this profile field
  - `value` - (Integer) The value for this profile field
  - `values` - (List of Integers) The permitted values for this profile field
