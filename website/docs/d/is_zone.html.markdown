---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : zone"
description: |-
  Manages IBM Cloud zone.
---

# ibm_is_zone
Retrieve information of an existing IBM Cloud zone in a particular region as a read-only data source. For more information, about IBM Cloud zone, see [creating a VPC in a different region](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-a-vpc-in-a-different-region).

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

data "ibm_is_zone" "example" {
  name   = "us-south-1"
  region = "us-south"
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Required, String) The name of the zone.
- `region` - (Required, String) The name of the region.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `status` - (String) The status of the zone.
- `data_center` - (String) The physical data center assigned to this logical zone. If absent, no physical data center has been assigned.
- `universal_name` - (String) The universal name for this zone. Will be absent if this zone has a status of unassigned.
