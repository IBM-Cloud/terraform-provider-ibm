---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_vpc_address_prefixes"
description: |-
  Get information about VPC address prefixes
---

# ibm_is_vpc_address_prefixes

Retrieve information of an existing IBM Cloud address prefix collection. For more information, about VPC address prefix, see [address prefixes](https://cloud.ibm.com/docs/vpc?topic=vpc-vpc-behind-the-curtain#address-prefixes).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform
data "ibm_is_vpc_address_prefixes" "example" {
  vpc  = ibm_is_vpc.example.id
  name = "example-address-prefix"
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `name` - (Optional, String) The unique user-defined name within the VPC the address prefix.
- `vpc`  - (Required, String) The VPC identifier.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `address_prefixes` - (List) Collection of the address prefixes.

  Nested `address_prefixes` blocks have the following structure:
  - `created_at` - (Timestamp) The date and time that the prefix was created.
  - `cidr` - (String) The CIDR block for this prefix.
  - `has_subnets` - (String) Indicates whether subnets exist with addresses from this prefix.
  - `href` - (String) The URL for this address prefix.
  - `id` - (String) The unique identifier for this address prefix.
  - `is_default` - (String) Indicates whether this is the default prefix for this zone in this VPC. If a default prefix was automatically created when the VPC was created, the prefix is automatically named using a hyphenated list of randomly-selected words, but may be updated with a user-specified name.
  - `name` - (String) The user-defined name for this address prefix. Names must be unique within the VPC the address prefix resides in.
  - `zone` - (List) The zone this address prefix resides in.
  
      Nested `zone` blocks have the following structure:
      - `href` - (String) The URL for this zone.
      - `name` - (String) The globally unique name for this zone.
  
  - `id` - (String) The unique identifier of the AddressPrefixCollection.
  - `total_count` - (String) The total number of resources across all pages.