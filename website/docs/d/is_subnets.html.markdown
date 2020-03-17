---
layout: "ibm"
page_title: "IBM : Subnets"
sidebar_current: "docs-ibm-datasources-is-subnets"
description: |-
  Manages IBM Cloud Infrastructure Subnets.
---

# ibm\_is_subnets

Import the details of an existing IBM Cloud Infrastructure subnets as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_subnets" "ds_subnets" {
}

```

## Attribute Reference

The following attributes are exported:

* `subnets` - List of all subnets in the IBM Cloud Infrastructure.
  * `name` - The name for this subnet.
  * `id` - The unique identifier for this subnet.
  * `ipv4_cidr_block` - The IPv4 CIDR block for this subnet.
  * `status` - The status of this subnet.



