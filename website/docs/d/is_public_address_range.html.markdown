---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : public_address_range"
description: |-
  Manages IBM Cloud public address rnage.
---
# ibm_is_public_address_range
Retrieve information of an existing public address range data source as a read only data source. For more information, about public address range , see [creating public address range](https://cloud.ibm.com/docs/vpc?topic=vpc-par-creating&interface=ui).

**Note:** 
The Public Address Range feature is currently available only with `Select Availability`.

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

```terraform
data "ibm_is_public_address_range" "example" {
	identifier = ibm_is_public_address_range.example.id
}
```

## Argument Reference
Review the argument reference that you can specify for your data source.

- `identifier` - (Optional, String) The ID of the VPN server.
- `name` - (Optional, String) The name of the VPN server.

  ~> **NOTE**
    `identifier` and `name` are mutually exclusive.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the PublicAddressRange.
- `cidr` - (String) The public IPv4 range, expressed in CIDR format.
- `created_at` - (String) The date and time that the public address range was created.
- `crn` - (String) The CRN for this public address range.
- `href` - (String) The URL for this public address range.
- `ipv4_address_count` - (Integer) The number of IPv4 addresses in this public address range.
- `lifecycle_state` - (String) The lifecycle state of the public address range.
- `name` - (String) The name for this public address range. The name is unique across all public address ranges in the region.
- `resource_group` - (List) The resource group for this public address range.
	
	Nested schema for `resource_group`:
	- `href` - (String) The URL for this resource group.
	- `id` - (String) The unique identifier for this resource group.
	- `name` - (String) The name for this resource group.
- `resource_type` - (String) The resource type.
- `target` - (List) The target this public address range is bound to.If absent, this public address range is not bound to a target.
	
	Nested schema for `target`:
	- `vpc` - (List) The VPC this public address range is bound to.
		
		Nested schema for `vpc`:
		- `crn` - (String) The CRN for this VPC.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			
			Nested schema for `deleted`:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this VPC.
		- `id` - (String) The unique identifier for this VPC.
		- `name` - (String) The name for this VPC. The name is unique across all VPCs in the region.
		- `resource_type` - (String) The resource type.
	- `zone` - (List) The zone this public address range resides in.
		
		Nested schema for `zone`:
		- `href` - (String) The URL for this zone.
		- `name` - (String) The globally unique name for this zone.
