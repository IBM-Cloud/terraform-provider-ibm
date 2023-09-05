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

- `profiles` - Collection of share profiles. Nested `profiles` blocks have the following structure:
	- `family` - The product family this share profile belongs to.
	- `href` - The URL for this share profile.
	- `name` - The globally unique name for this share profile.
	- `resource_type` - The resource type.
- `total_count` - The total number of resources across all pages.

