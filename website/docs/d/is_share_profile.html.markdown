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
	name = "tier-3iops"
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required, string) The file share profile name.

## Attribute Reference

The following attributes are exported:

- `family` - The product family this share profile belongs to.
- `href` - The URL for this share profile.
- `resource_type` - The resource type.

