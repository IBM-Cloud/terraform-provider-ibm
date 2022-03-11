---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_vpc_address_prefix"
description: |-
  Get information about VPC Address Prefix
---

# ibm_is_vpc_address_prefix

Provides a read-only data source for VPC Address Prefix. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_vpc_address_prefix" "example" {
  vpc = ibm_is_vpc.example.id
  address_prefix = ibm_is_vpc_address_prefix.example.address_prefix
}
data "ibm_is_vpc_address_prefix" "example-1" {
  vpc_name = ibm_is_vpc.example.name
  address_prefix = ibm_is_vpc_address_prefix.example.address_prefix
}
data "ibm_is_vpc_address_prefix" "example-2" {
  vpc = ibm_is_vpc.example.id
  address_prefix_name = ibm_is_vpc_address_prefix.example.name
}
data "ibm_is_vpc_address_prefix" "example-3" {
  vpc_name = ibm_is_vpc.example.name
  address_prefix_name = ibm_is_vpc_address_prefix.example.name
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `address_prefix` - (String) The address prefix identifier.
- `address_prefix_name` - (String) The address prefix name.

~> **Note:**
  Provide exactly one of `address_prefix`, `address_prefix_name`

- `vpc` - (String) The VPC identifier
- `vpc_name` - (String) Name of the VPC
  
  ~> **Note:**
  Provide exactly one of `vpc`, `vpc_name`

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier of the AddressPrefix.
- `cidr` - (String) The CIDR block for this prefix.

- `created_at` - (String) The date and time that the prefix was created.

- `has_subnets` - (Boolean) Indicates whether subnets exist with addresses from this prefix.

- `href` - (String) The URL for this address prefix.

- `is_default` - (Boolean) Indicates whether this is the default prefix for this zone in this VPC. If a default prefix was automatically created when the VPC was created, the prefix is automatically named using a hyphenated list of randomly-selected words, but may be updated with a user-specified name.

- `name` - (String) The user-defined name for this address prefix. Names must be unique within the VPC the address prefix resides in.

- `zone` - (List) The zone this address prefix resides in.
  Nested scheme for **zone**:
	- `href` - (String) The URL for this zone.
	- `name` - (String) The globally unique name for this zone.

