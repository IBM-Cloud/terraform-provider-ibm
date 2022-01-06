---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : regions"
description: |-
  Manages IBM Cloud regions.
---

# ibm_is_regions
Retrieve information about VPC Generation 2 list of regions as a read only data source. For more information, about managing IBM Cloud region, see [creating a VPC in a different region](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-a-vpc-in-a-different-region).

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

data "ibm_is_regions" "example" {
}

```



## Attribute reference
Following attribute references can be accessed after your data source is created.

- `regions` - (List) List of all regions in the IBM Cloud infrastructure.

  Nested scheme for `regions`:
    - `endpoint` - (String) The endpoint of the region.
    - `href` - (String) The url for this region.
    - `name` - (String) The name of the region.
    - `status` - (String) The status of the region.
