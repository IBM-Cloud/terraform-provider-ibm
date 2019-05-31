---
layout: "ibm"
page_title: "IBM : vpc-address-prefix"
sidebar_current: "docs-ibm-resource-is-vpc-address-prefix"
description: |-
  Manages IBM IS VPC Address prefix.
---

# ibm\_is_vpc_address_prefix

Provides a vpc address prefix resource. This allows vpc address prefix to be created, updated, and cancelled.


## Example Usage

```hcl
resource "ibm_is_vpc" "testacc_vpc" {
  name = "testvpc"
}

resource "ibm_is_vpc_address_prefix" "testacc_vpc_address_prefix" {
  name = "test"
  zone   = "us-south-1"
  vpc         = "${ibm_is_vpc.testacc_vpc.id}"
  cidr        = "10.240.0.0/24"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The address prefix name.
* `vpc` - (Required, string) The vpc id. 
* `zone` - (Required, string) Name of the zone. 
* `cidr` - (Required, string) The CIDR block for the address prefix. 

## Attribute Reference

The following attributes are exported:

* `id` - The id of the address prefix.
* `has_subnets` - Indicates whether subnets exist with addresses from this prefix.

## Import

ibm_is_vpc_address_prefix can be imported using ID, eg

```
$ terraform import ibm_is_vpc_address_prefix.example d7bec597-4726-451f-8a63-e62e6f19c32c
```