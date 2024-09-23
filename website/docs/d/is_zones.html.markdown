---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : zones"
description: |-
  Manages IBM Cloud zones.
---

# ibm_is_zones
Retrieve information of an existing IBM Cloud zones in a particular region as a read-only data source. For more information, about IBM Cloud zones, see [creating a VPC in a different region](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-a-vpc-in-a-different-region).

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
data "ibm_is_zones" "example" {
  region = "us-south"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `region` - (Required, String) The name of the region.
- `status` - (Optional, String) Filter the list by status of zones.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `zones` - (String) The list of zones in an IBM Cloud region.  For example, **us-south-1**,**us-south-2**.
- `zone_info` - (List) Collection of zones.
  Nested schema for **zone_info**:
	- `data_center` - (String) The physical data center assigned to this logical zone. If absent, no physical data center has been assigned.
	- `name` - (String) The name of the zone.
	- `status` - (String) The status of the zone.
	- `universal_name` - (String) The universal name for this zone. Will be absent if this zone has a status of unassigned.
