---
layout: "ibm"
page_title: "IBM : vpc"
sidebar_current: "docs-ibm-datasources-is-vpc"
description: |-
  Manages IBM virtual private cloud.
---

# ibm\_is_vpc

Import the details of an existing IBM Virtual Private cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
resource "ibm_is_vpc" "testacc_vpc" {
    name = "test"
}

data "ibm_is_vpc" "ds_vpc" {
    name = "test"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the VPC.

## Attribute Reference

The following attributes are exported:

* `status` - The status of VPC.
* `default_network_acl` - ID of the default network ACL.
* `classic_access` - Indicates whether this VPC is connected to Classic Infrastructure.
* `resource_group` - The resource group where the VPC created.
* `tags` - Tags associated with the instance.