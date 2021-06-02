---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_vpc_address_prefix"
description: |-
  Get information about VPC Address Prefixes
---

# ibm\_is_vpc_address_prefix

Provides a read-only data source for AddressPrefixCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_vpc_address_prefix" "is_vpc_address_prefix_name" {
  vpc  = "r134-b5938d43-cb2f-4666-bc99-9410863ed305"
  name = "outsider-sense-motor-chomp"
}
```

## Argument Reference

The following arguments are supported:

* `vpc` - (Required, string) The VPC identifier.
* `name` - (Optional, string) The user-defined name for this address prefix. Names must be unique within the VPC the address prefix resides in.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the AddressPrefixCollection.
* `address_prefixes` - Collection of address prefixes. Nested `address_prefixes` blocks have the following structure:
	* `cidr` - The CIDR block for this prefix.
	* `created_at` - The date and time that the prefix was created.
	* `has_subnets` - Indicates whether subnets exist with addresses from this prefix.
	* `href` - The URL for this address prefix.
	* `id` - The unique identifier for this address prefix.
	* `is_default` - Indicates whether this is the default prefix for this zone in this VPC. If a default prefix was automatically created when the VPC was created, the prefix is automatically named using a hyphenated list of randomly-selected words, but may be updated with a user-specified name.
	* `name` - The user-defined name for this address prefix. Names must be unique within the VPC the address prefix resides in.
	* `zone` - The zone this address prefix resides in. Nested `zone` blocks have the following structure:
		* `href` - The URL for this zone.
		* `name` - The globally unique name for this zone.
* `total_count` - The total number of resources across all pages.

