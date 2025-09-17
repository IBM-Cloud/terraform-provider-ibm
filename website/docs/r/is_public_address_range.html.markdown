---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : public_address_range"
description: |-
  Manages IBM public address range.
---

# ibm_is_public_address_range

Create, update, and delete a public address range. For more information, see [creating public address range](https://cloud.ibm.com/docs/vpc?topic=vpc-par-creating&interface=ui).

**Note:** 
The Public Address Range feature is currently available only with the `Select Availability`.

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage
The following example shows how you can create a public address range for a vpc that are located in a specific zone.

```terraform
resource "ibm_is_public_address_range" "public_address_range_instance" {
  ipv4_address_count = "16"
  name               = "example-public-address-range"
  resource_group {
    id = "11caaa983d9c4beb82690daab18717e9"
  }
  target {
    vpc {
      id = ibm_is_vpc.testacc_vpc.id
    }
    zone {
      name = "us-south-3"
    }
  }
}
```

An example shows how you can create public address range not attached to vpc and zone

```terraform
resource "ibm_is_public_address_range" "public_address_range_instance" {
  ipv4_address_count = "16"
  name               = "example-public-address-range"
  resource_group {
    id = "11caaa983d9c4beb82690daab08717e9"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

- `ipv4_address_count` - (Required, Integer) The number of IPv4 addresses in this public address range.
- `name` - (Optional, String) The name for this public address range. The name is unique across all public address ranges in the region.
- `resource_group` - (Optional, List) The resource group for this public address range.
    
	Nested schema for `resource_group`:
	- `id` - (Required, String) The unique identifier for this resource group.
- `target` - (Optional, List) The target this public address range is bound to.If absent, this public address range is not bound to a target.
    
	Nested schema for `target`:
	- `vpc` - (Required, List) The VPC this public address range is bound to. If present, any of the below value must be specified.
	    
		Nested schema for `vpc`:
		- `crn` - (Optional, String) The CRN for this VPC.
		- `href` - (Optional, String) The URL for this VPC.
		- `id` - (Optional, String) The unique identifier for this VPC.
	- `zone` - (Required, List) The zone this public address range resides in. If present, any of the below value must be specified.
	    
		Nested schema for `zone`:
		- `href` - (Optional, String) The URL for this zone.
		- `name` - (Optional, String) The globally unique name for this zone.

## Attribute Reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - The unique identifier of the PublicAddressRange.
- `cidr` - (String) The public IPv4 range, expressed in CIDR format.
- `created_at` - (String) The date and time that the public address range was created.
- `crn` - (String) The CRN for this public address range.
- `href` - (String) The URL for this public address range.
- `lifecycle_state` - (String) The lifecycle state of the public address range.
- `resource_type` - (String) The resource type.
- `resource_group` - (List) The resource group for this public address range.
    
	Nested schema for `resource_group`:
	- `href` - (String) The URL for this resource group.
	- `id` - (String) The unique identifier for this resource group.
	- `name` - (String) The name for this resource group.
- `target` - (List) The target this public address range is bound to.If absent, this public address range is not bound to a target.
    
	Nested schema for `target`:
	- `vpc` - (List) The VPC this public address range is bound to.
	    
		Nested schema for `vpc`:
		- `crn` - (String) The CRN for this VPC.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			
			Nested schema for `deleted`:
			- `more_info` - (Computed, String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this VPC.
		- `id` - (String) The unique identifier for this VPC.
		- `name` - (Computed, String) The name for this VPC. The name is unique across all VPCs in the region.
		- `resource_type` - (Computed, String) The resource type.
	- `zone` - (List) The zone this public address range resides in.
	    
		Nested schema for `zone`:
		- `href` - (String) The URL for this zone.
		- `name` - (String) The globally unique name for this zone.

## Import

You can import the `ibm_is_public_address_range` resource by using `id`. The unique identifier for this public address range.

# Syntax
```
$ terraform import ibm_is_public_address_range.is_public_address_range <id>
```

# Example
```
$ terraform import ibm_is_public_address_range.is_public_address_range r006-a4841334-b584-4293-938e-3bc63b4a5b6a
```
