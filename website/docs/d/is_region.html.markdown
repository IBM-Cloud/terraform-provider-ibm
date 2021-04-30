---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : region"
description: |-
  Manages IBM Cloud Region.
---

# ibm\_is_region

Import the details of an existing IBM Cloud region as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_region" "ds_region" {
  name = "us-south"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the region.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `status` - The status of region.
* `endpoint` - The endpoint of the region.