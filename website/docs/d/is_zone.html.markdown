---
layout: "ibm"
page_title: "IBM : zone"
sidebar_current: "docs-ibm-datasources-is-zone"
description: |-
  Manages IBM Cloud Zone.
---

# ibm\_is_zone

Import the details of an existing IBM Cloud zone in a particular region as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_zone" "ds_zone" {
    name = "us-south-1"
    region = "us-south"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the zone.
* `region` - (Required, string) The name of the region.

## Attribute Reference

The following attributes are exported:

* `status` - The status of zone.