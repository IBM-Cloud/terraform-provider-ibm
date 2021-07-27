---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : zones"
description: |-
  Manages IBM Cloud zones.
---

# ibm_is_zones
Retrieve information of an existing IBM Cloud zones in a particular region as a read-only data source. For more information, about IBM Cloud zones, see [creating a VPC in a different region](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-a-vpc-in-a-different-region).

## Example usage

```terraform
data "ibm_is_zones" "ds_zones" {
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
